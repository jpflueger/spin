package postgres

type DbDataTypeKind int

const (
	DbDataTypeKindBoolean DbDataTypeKind = iota
	DbDataTypeKindInt8
	DbDataTypeKindInt16
	DbDataTypeKindInt32
	DbDataTypeKindInt64
	DbDataTypeKindUint8
	DbDataTypeKindUint16
	DbDataTypeKindUint32
	DbDataTypeKindUint64
	DbDataTypeKindFloating32
	DbDataTypeKindFloating64
	DbDataTypeKindStr
	DbDataTypeKindBinary
	DbDataTypeKindOther
)

type DbDataType struct {
	kind DbDataTypeKind
}

func (n DbDataType) Kind() DbDataTypeKind {
	return n.kind
}

func DbDataTypeBoolean() DbDataType {
	return DbDataType{kind: DbDataTypeKindBoolean}
}

func DbDataTypeInt8() DbDataType {
	return DbDataType{kind: DbDataTypeKindInt8}
}

func DbDataTypeInt16() DbDataType {
	return DbDataType{kind: DbDataTypeKindInt16}
}

func DbDataTypeInt32() DbDataType {
	return DbDataType{kind: DbDataTypeKindInt32}
}

func DbDataTypeInt64() DbDataType {
	return DbDataType{kind: DbDataTypeKindInt64}
}

func DbDataTypeUint8() DbDataType {
	return DbDataType{kind: DbDataTypeKindUint8}
}

func DbDataTypeUint16() DbDataType {
	return DbDataType{kind: DbDataTypeKindUint16}
}

func DbDataTypeUint32() DbDataType {
	return DbDataType{kind: DbDataTypeKindUint32}
}

func DbDataTypeUint64() DbDataType {
	return DbDataType{kind: DbDataTypeKindUint64}
}

func DbDataTypeFloating32() DbDataType {
	return DbDataType{kind: DbDataTypeKindFloating32}
}

func DbDataTypeFloating64() DbDataType {
	return DbDataType{kind: DbDataTypeKindFloating64}
}

func DbDataTypeStr() DbDataType {
	return DbDataType{kind: DbDataTypeKindStr}
}

func DbDataTypeBinary() DbDataType {
	return DbDataType{kind: DbDataTypeKindBinary}
}

func DbDataTypeOther() DbDataType {
	return DbDataType{kind: DbDataTypeKindOther}
}