package postgres

import "fmt"

type ParameterValueKind int

const (
	ParameterValueKindBoolean ParameterValueKind = iota
	ParameterValueKindInt8
	ParameterValueKindInt16
	ParameterValueKindInt32
	ParameterValueKindInt64
	ParameterValueKindUint8
	ParameterValueKindUint16
	ParameterValueKindUint32
	ParameterValueKindUint64
	ParameterValueKindFloating32
	ParameterValueKindFloating64
	ParameterValueKindStr
	ParameterValueKindBinary
	ParameterValueKindDbNull
)

type ParameterValue struct {
	kind ParameterValueKind
	val  interface{}
}

func (n ParameterValue) Kind() ParameterValueKind {
	return n.kind
}

func ParameterValueBoolean(v bool) ParameterValue {
	return ParameterValue{kind: ParameterValueKindBoolean, val: v}
}

func (n ParameterValue) GetBoolean() bool {
	if g, w := n.Kind(), ParameterValueKindBoolean; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(bool)
}

func (n *ParameterValue) SetBoolean(v bool) {
	n.val = v
	n.kind = ParameterValueKindBoolean
}

func ParameterValueInt8(v int8) ParameterValue {
	return ParameterValue{kind: ParameterValueKindInt8, val: v}
}

func (n ParameterValue) GetInt8() int8 {
	if g, w := n.Kind(), ParameterValueKindInt8; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(int8)
}

func (n *ParameterValue) SetInt8(v int8) {
	n.val = v
	n.kind = ParameterValueKindInt8
}

func ParameterValueInt16(v int16) ParameterValue {
	return ParameterValue{kind: ParameterValueKindInt16, val: v}
}

func (n ParameterValue) GetInt16() int16 {
	if g, w := n.Kind(), ParameterValueKindInt16; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(int16)
}

func (n *ParameterValue) SetInt16(v int16) {
	n.val = v
	n.kind = ParameterValueKindInt16
}

func ParameterValueInt32(v int32) ParameterValue {
	return ParameterValue{kind: ParameterValueKindInt32, val: v}
}

func (n ParameterValue) GetInt32() int32 {
	if g, w := n.Kind(), ParameterValueKindInt32; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(int32)
}

func (n *ParameterValue) SetInt32(v int32) {
	n.val = v
	n.kind = ParameterValueKindInt32
}

func ParameterValueInt64(v int64) ParameterValue {
	return ParameterValue{kind: ParameterValueKindInt64, val: v}
}

func (n ParameterValue) GetInt64() int64 {
	if g, w := n.Kind(), ParameterValueKindInt64; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(int64)
}

func (n *ParameterValue) SetInt64(v int64) {
	n.val = v
	n.kind = ParameterValueKindInt64
}

func ParameterValueUint8(v uint8) ParameterValue {
	return ParameterValue{kind: ParameterValueKindUint8, val: v}
}

func (n ParameterValue) GetUint8() uint8 {
	if g, w := n.Kind(), ParameterValueKindUint8; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(uint8)
}

func (n *ParameterValue) SetUint8(v uint8) {
	n.val = v
	n.kind = ParameterValueKindUint8
}

func ParameterValueUint16(v uint16) ParameterValue {
	return ParameterValue{kind: ParameterValueKindUint16, val: v}
}

func (n ParameterValue) GetUint16() uint16 {
	if g, w := n.Kind(), ParameterValueKindUint16; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(uint16)
}

func (n *ParameterValue) SetUint16(v uint16) {
	n.val = v
	n.kind = ParameterValueKindUint16
}

func ParameterValueUint32(v uint32) ParameterValue {
	return ParameterValue{kind: ParameterValueKindUint32, val: v}
}

func (n ParameterValue) GetUint32() uint32 {
	if g, w := n.Kind(), ParameterValueKindUint32; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(uint32)
}

func (n *ParameterValue) SetUint32(v uint32) {
	n.val = v
	n.kind = ParameterValueKindUint32
}

func ParameterValueUint64(v uint64) ParameterValue {
	return ParameterValue{kind: ParameterValueKindUint64, val: v}
}

func (n ParameterValue) GetUint64() uint64 {
	if g, w := n.Kind(), ParameterValueKindUint64; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(uint64)
}

func (n *ParameterValue) SetUint64(v uint64) {
	n.val = v
	n.kind = ParameterValueKindUint64
}

func ParameterValueFloating32(v float32) ParameterValue {
	return ParameterValue{kind: ParameterValueKindFloating32, val: v}
}

func (n ParameterValue) GetFloating32() float32 {
	if g, w := n.Kind(), ParameterValueKindFloating32; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(float32)
}

func (n *ParameterValue) SetFloating32(v float32) {
	n.val = v
	n.kind = ParameterValueKindFloating32
}

func ParameterValueFloating64(v float64) ParameterValue {
	return ParameterValue{kind: ParameterValueKindFloating64, val: v}
}

func (n ParameterValue) GetFloating64() float64 {
	if g, w := n.Kind(), ParameterValueKindFloating64; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(float64)
}

func (n *ParameterValue) SetFloating64(v float64) {
	n.val = v
	n.kind = ParameterValueKindFloating64
}

func ParameterValueStr(v string) ParameterValue {
	return ParameterValue{kind: ParameterValueKindStr, val: v}
}

func (n ParameterValue) GetStr() string {
	if g, w := n.Kind(), ParameterValueKindStr; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.(string)
}

func (n *ParameterValue) SetStr(v string) {
	n.val = v
	n.kind = ParameterValueKindStr
}

func ParameterValueBinary(v []uint8) ParameterValue {
	return ParameterValue{kind: ParameterValueKindBinary, val: v}
}

func (n ParameterValue) GetBinary() []uint8 {
	if g, w := n.Kind(), ParameterValueKindBinary; g != w {
		panic(fmt.Sprintf("Attr kind is %v, not %v", g, w))
	}
	return n.val.([]uint8)
}

func (n *ParameterValue) SetBinary(v []uint8) {
	n.val = v
	n.kind = ParameterValueKindBinary
}

func ParameterValueDbNull() ParameterValue {
	return ParameterValue{kind: ParameterValueKindDbNull}
}