package deepcopy

import (
	"reflect"
	"testing"
)

type testSrcModifyMap struct {
	I   int
	I8  int8
	I16 int16
	I32 int32
	I64 int64

	U   uint
	U8  uint8
	U16 uint16
	U32 uint32
	U64 uint64

	F32 float32
	F64 float64

	S string

	B bool
}

type testDstModifyMap struct {
	I   int
	I8  int8
	I16 int16
	I32 int32
	I64 int64

	U   uint
	U8  uint8
	U16 uint16
	U32 uint32
	U64 uint64

	F32 float32
	F64 float64

	S string

	B bool
}

func Test_ModifyMap(t *testing.T) {
	t.Run("测试 src没有值时调用的回调函数", func(t *testing.T) {
		var src = testSrcModifyMap{}

		var need = testDstModifyMap{
			I:   1,
			I8:  1,
			I16: 1,
			I32: 1,
			I64: 1,

			U:   1,
			U8:  1,
			U16: 1,
			U32: 1,
			U64: 1,

			F32: 1,
			F64: 1,

			S: "hello",
			B: true,
		}

		var dst testDstModifyMap

		err := CopyEx(&dst, &src, WithModifyDstField(map[string]ModifyDstFieldFunc{
			"I": func(v interface{}) interface{} {
				return 1
			},
			"I8": func(v interface{}) interface{} {
				return int8(1)
			},
			"I16": func(v interface{}) interface{} {
				return int16(1)
			},
			"I32": func(v interface{}) interface{} {
				return int32(1)
			},
			"I64": func(v interface{}) interface{} {
				return int64(1)
			},

			"U": func(v interface{}) interface{} {
				return uint(1)
			},
			"U8": func(v interface{}) interface{} {
				return uint8(1)
			},
			"U16": func(v interface{}) interface{} {
				return uint16(1)
			},
			"U32": func(v interface{}) interface{} {
				return uint32(1)
			},
			"U64": func(v interface{}) interface{} {
				return uint64(1)
			},
			"F32": func(v interface{}) interface{} {
				return float32(1)
			},
			"F64": func(v interface{}) interface{} {
				return float64(1)
			},
			"S": func(srcArg interface{}) (newDst interface{}) {
				return "hello"
			},
			"B": func(srcArg interface{}) (newDst interface{}) {
				return true
			},
		}))

		if err != nil {
			t.Errorf("test faild, err:%v", err)
		}
		if !reflect.DeepEqual(dst, need) {
			t.Errorf("test faild, got:%v", dst)
		}
	})
	t.Run("测试 src有值时调用的回调函数", func(t *testing.T) {
		var src = testSrcModifyMap{
			I:   2,
			I8:  2,
			I16: 2,
			I32: 2,
			I64: 2,
			U:   2,
			U8:  2,
			U16: 2,
			U32: 2,
			U64: 2,
			F32: 2,
			F64: 2,
			S:   "hello 123",
			B:   true,
		}

		var need = testDstModifyMap{
			I:   1,
			I8:  1,
			I16: 1,
			I32: 1,
			I64: 1,

			U:   1,
			U8:  1,
			U16: 1,
			U32: 1,
			U64: 1,

			F32: 1,
			F64: 1,

			S: "hello",
			B: true,
		}

		var dst testDstModifyMap

		err := CopyEx(&dst, &src, WithModifySrcField(map[string]ModifySrcValue{
			"I": {
				"I", "I", func(v interface{}) interface{} {
					return 1
				},
			},
			"I8": {
				"I8", "I8", func(v interface{}) interface{} {
					return int8(1)
				},
			},
			"I16": {
				"I16", "I16", func(v interface{}) interface{} {
					return int16(1)
				},
			},
			"I32": {
				"I32", "I32", func(v interface{}) interface{} {
					return int32(1)
				},
			},
			"I64": {
				"I64", "I64", func(v interface{}) interface{} {
					return int64(1)
				},
			},
			"U": {
				"U", "U", func(v interface{}) interface{} {
					return uint(1)
				},
			},
			"U8": {
				"U8", "U8", func(v interface{}) interface{} {
					return uint8(1)
				},
			},
			"U16": {
				"U16", "U16", func(v interface{}) interface{} {
					return uint16(1)
				},
			},
			"U32": {
				"U32", "U32", func(v interface{}) interface{} {
					return uint32(1)
				},
			},
			"U64": {
				"U64", "U64", func(v interface{}) interface{} {
					return uint64(1)
				},
			},
			"F32": {
				"F32", "F32", func(v interface{}) interface{} {
					return float32(1)
				},
			},
			"F64": {
				"F64", "F64", func(v interface{}) interface{} {
					return float64(1)
				},
			},
			"S": {
				"S", "S", func(srcArg interface{}) interface{} {
					return "hello"
				},
			},
			"B": {
				"B", "B", func(srcArg interface{}) interface{} {
					return true
				},
			},
		}))

		if err != nil {
			t.Errorf("test faild, err:%v", err)
		}
		if !reflect.DeepEqual(dst, need) {
			t.Errorf("test faild, got:%v", dst)
		}
	})

	t.Run("测试 src有值时调用的回调函数2", func(t *testing.T) {
		var src = testSrcModifyMap{
			I: 2,
		}

		var need = testDstModifyMap{

			I8: 3,
		}

		var dst testDstModifyMap
		err := CopyEx(&dst, &src, WithModifySrcField(map[string]ModifySrcValue{
			"I": {
				DstFieldName: "I8", SrcFieldName: "I", Callback: func(v interface{}) interface{} {
					return int8(v.(int) + 1)
				},
			},
		}))

		if err != nil {
			t.Errorf("test faild, err:%v", err)
		}
		if !reflect.DeepEqual(dst, need) {
			t.Errorf("test faild, got:%v", dst)
		}
	})
}
