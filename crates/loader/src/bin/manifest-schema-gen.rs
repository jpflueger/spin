use schemars::schema_for;
use spin_loader::local::config::RawAppManifestAnyVersion;
use std::fs;

fn main() {
    let schema = schema_for!(RawAppManifestAnyVersion);
    let data = serde_json::to_string_pretty(&schema).unwrap();
    fs::write("./schema.json", data).expect("Unable to write file");
}