//! Functionality to get a prepared Spin application from Bindle.

#![deny(missing_docs)]

/// Module to prepare the assets for the components of an application.
mod assets;
/// Configuration representation for a Spin application in Bindle.
pub mod config;
/// Bindle helper functions.
mod utils;

use crate::bindle::{
    config::{RawAppManifest, RawComponentManifest},
    utils::{find_manifest, BindleReader},
};
use anyhow::{anyhow, Context, Result};
use bindle::{
    client::{tokens::NoToken, Client},
    Invoice,
};
use futures::future;
use spin_config::{
    ApplicationInformation, ApplicationOrigin, Configuration, CoreComponent, ModuleSource,
    SpinVersion, WasmConfig,
};
use std::path::Path;
use tracing::log;
pub use utils::{BindleTokenManager, SPIN_MANIFEST_MEDIA_TYPE};

/// Given a Bindle server URL and reference, pull it, expand its assets locally, and get a
/// prepared application configuration consumable by a Spin execution context.
/// If a directory is provided, use it as the base directory to expand the assets,
/// otherwise create a new temporary directory.
pub async fn from_bindle(
    id: &str,
    url: &str,
    base_dst: impl AsRef<Path>,
) -> Result<Configuration<CoreComponent>> {
    // TODO
    // Handle Bindle authentication.
    let manager = BindleTokenManager::NoToken(NoToken);
    let client = Client::new(url, manager)?;
    let reader = BindleReader::remote(&client, &id.parse()?);

    prepare(id, url, &reader, base_dst).await
}

/// Converts a Bindle invoice into Spin configuration.
async fn prepare(
    id: &str,
    url: &str,
    reader: &BindleReader,
    base_dst: impl AsRef<Path>,
) -> Result<Configuration<CoreComponent>> {
    // First, get the invoice from the Bindle server.
    let invoice = reader
        .get_invoice()
        .await
        .with_context(|| anyhow!("Failed to load invoice '{}' from '{}'", id, url))?;

    // Then, reconstruct the application manifest from the parcels.
    let raw: RawAppManifest =
        toml::from_slice(&reader.get_parcel(&find_manifest(&invoice)?).await?)?;
    log::trace!("Recreated manifest from bindle: {:?}", raw);

    let info = info(&raw, &invoice, url);
    log::trace!("Application information from bindle: {:?}", info);
    let components = future::join_all(
        raw.components
            .into_iter()
            .map(|c| async { core(c, &invoice, reader, &base_dst).await })
            .collect::<Vec<_>>(),
    )
    .await
    .into_iter()
    .map(|x| x.expect("Cannot prepare component"))
    .collect::<Vec<_>>();

    Ok(Configuration { info, components })
}

/// Given a raw component manifest, prepare its assets and return a fully formed core component.
async fn core(
    raw: RawComponentManifest,
    invoice: &Invoice,
    reader: &BindleReader,
    base_dst: impl AsRef<Path>,
) -> Result<CoreComponent> {
    let bytes = reader
        .get_parcel(&raw.source)
        .await
        .with_context(|| anyhow!("Cannot get module source from bindle"))?;

    let source = ModuleSource::Buffer(bytes, format!("parcel {}", raw.source));
    let id = raw.id;
    let parcels = invoice
        .parcel
        .as_ref()
        .unwrap_or(&vec![])
        .iter()
        .map(|p| p.label.clone())
        .collect::<Vec<_>>();
    let mounts = match raw.wasm.files {
        Some(_) => vec![
            assets::prepare_component(reader, &invoice.bindle.id, &parcels, base_dst, &id).await?,
        ],
        None => vec![],
    };
    let environment = raw.wasm.environment.unwrap_or_default();
    let allowed_http_hosts = raw.wasm.allowed_http_hosts.unwrap_or_default();
    let wasm = WasmConfig {
        environment,
        mounts,
        allowed_http_hosts,
    };
    let trigger = raw.trigger;

    Ok(CoreComponent {
        source,
        id,
        wasm,
        trigger,
    })
}

/// Converts the raw application manifest from the bindle invoice into the
/// standard application configuration.
fn info(raw: &RawAppManifest, invoice: &Invoice, url: &str) -> ApplicationInformation {
    ApplicationInformation {
        // TODO
        // Handle API version and namespace.
        spin_version: SpinVersion::V1,
        name: invoice.bindle.id.name().to_string(),
        version: invoice.bindle.id.version_string(),
        description: invoice.bindle.description.clone(),
        authors: invoice.bindle.authors.clone().unwrap_or_default(),
        trigger: raw.trigger.clone(),
        namespace: None,
        origin: ApplicationOrigin::Bindle {
            id: invoice.bindle.id.to_string(),
            server: url.to_string(),
        },
    }
}