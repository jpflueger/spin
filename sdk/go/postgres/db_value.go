package postgres

import "fmt"

type DbValueKind int

const (
	DbValueKindBoolean DbValueKind = iota
	DbValueKindInt8
	DbValueKindInt16
	DbValueKindInt32
	DbValueKindInt64
	DbValueKindUint8
	DbValueKindUint16
	DbValueKindUint32
	DbValueKindUint64
	DbValueKindFloating32
	DbValueKindFloating64
	DbValueKindStr
	DbValueKindBinary
	DbValueKindDbNull
	DbValueKindUnsupported
)

type DbValue struct {
	kind DbValueKind
	val  interface{}
}

func (n DbValue) Kind() DbValueKind {
	return n.kind
}

func DbValueBoolean(v bool) DbValue {
	return DbValue{kind: DbValueKindBoolean, val: v}
}

func (n DbValue) GetBoolean() bool {
	if g, w := n.Kind(), DbValueKindBoolean; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(bool)
}

func (n *DbValue) SetBoolean(v bool) {
	n.val = v
	n.kind = DbValueKindBoolean
}

func DbValueInt8(v int8) DbValue {
	return DbValue{kind: DbValueKindInt8, val: v}
}

func (n DbValue) GetInt8() int8 {
	if g, w := n.Kind(), DbValueKindInt8; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(int8)
}

func (n *DbValue) SetInt8(v int8) {
	n.val = v
	n.kind = DbValueKindInt8
}

func DbValueInt16(v int16) DbValue {
	return DbValue{kind: DbValueKindInt16, val: v}
}

func (n DbValue) GetInt16() int16 {
	if g, w := n.Kind(), DbValueKindInt16; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(int16)
}

func (n *DbValue) SetInt16(v int16) {
	n.val = v
	n.kind = DbValueKindInt16
}

func DbValueInt32(v int32) DbValue {
	return DbValue{kind: DbValueKindInt32, val: v}
}

func (n DbValue) GetInt32() int32 {
	if g, w := n.Kind(), DbValueKindInt32; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(int32)
}

func (n *DbValue) SetInt32(v int32) {
	n.val = v
	n.kind = DbValueKindInt32
}

func DbValueInt64(v int64) DbValue {
	return DbValue{kind: DbValueKindInt64, val: v}
}

func (n DbValue) GetInt64() int64 {
	if g, w := n.Kind(), DbValueKindInt64; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(int64)
}

func (n *DbValue) SetInt64(v int64) {
	n.val = v
	n.kind = DbValueKindInt64
}

func DbValueUint8(v uint8) DbValue {
	return DbValue{kind: DbValueKindUint8, val: v}
}

func (n DbValue) GetUint8() uint8 {
	if g, w := n.Kind(), DbValueKindUint8; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(uint8)
}

func (n *DbValue) SetUint8(v uint8) {
	n.val = v
	n.kind = DbValueKindUint8
}

func DbValueUint16(v uint16) DbValue {
	return DbValue{kind: DbValueKindUint16, val: v}
}

func (n DbValue) GetUint16() uint16 {
	if g, w := n.Kind(), DbValueKindUint16; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(uint16)
}

func (n *DbValue) SetUint16(v uint16) {
	n.val = v
	n.kind = DbValueKindUint16
}

func DbValueUint32(v uint32) DbValue {
	return DbValue{kind: DbValueKindUint32, val: v}
}

func (n DbValue) GetUint32() uint32 {
	if g, w := n.Kind(), DbValueKindUint32; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(uint32)
}

func (n *DbValue) SetUint32(v uint32) {
	n.val = v
	n.kind = DbValueKindUint32
}

func DbValueUint64(v uint64) DbValue {
	return DbValue{kind: DbValueKindUint64, val: v}
}

func (n DbValue) GetUint64() uint64 {
	if g, w := n.Kind(), DbValueKindUint64; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(uint64)
}

func (n *DbValue) SetUint64(v uint64) {
	n.val = v
	n.kind = DbValueKindUint64
}

func DbValueFloating32(v float32) DbValue {
	return DbValue{kind: DbValueKindFloating32, val: v}
}

func (n DbValue) GetFloating32() float32 {
	if g, w := n.Kind(), DbValueKindFloating32; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(float32)
}

func (n *DbValue) SetFloating32(v float32) {
	n.val = v
	n.kind = DbValueKindFloating32
}

func DbValueFloating64(v float64) DbValue {
	return DbValue{kind: DbValueKindFloating64, val: v}
}

func (n DbValue) GetFloating64() float64 {
	if g, w := n.Kind(), DbValueKindFloating64; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(float64)
}

func (n *DbValue) SetFloating64(v float64) {
	n.val = v
	n.kind = DbValueKindFloating64
}

func DbValueStr(v string) DbValue {
	return DbValue{kind: DbValueKindStr, val: v}
}

func (n DbValue) GetStr() string {
	if g, w := n.Kind(), DbValueKindStr; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(string)
}

func (n *DbValue) SetStr(v string) {
	n.val = v
	n.kind = DbValueKindStr
}

func DbValueBinary(v []uint8) DbValue {
	return DbValue{kind: DbValueKindBinary, val: v}
}

func (n DbValue) GetBinary() []uint8 {
	if g, w := n.Kind(), DbValueKindBinary; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.([]uint8)
}

func (n *DbValue) SetBinary(v []uint8) {
	n.val = v
	n.kind = DbValueKindBinary
}

func DbValueDbNull() DbValue {
	return DbValue{kind: DbValueKindDbNull}
}

func DbValueUnsupported() DbValue {
	return DbValue{kind: DbValueKindUnsupported}
}
