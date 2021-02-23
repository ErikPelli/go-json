package json_test

import (
	"bytes"
	"testing"

	"github.com/goccy/go-json"
)

func TestCoverInt(t *testing.T) {
	type structInt struct {
		A int `json:"a"`
	}
	type structIntOmitEmpty struct {
		A int `json:"a,omitempty"`
	}
	type structIntString struct {
		A int `json:"a,string"`
	}

	type structIntPtr struct {
		A *int `json:"a"`
	}
	type structIntPtrOmitEmpty struct {
		A *int `json:"a,omitempty"`
	}
	type structIntPtrString struct {
		A *int `json:"a,string"`
	}

	type structIntTriplePtr struct {
		A ***int `json:"a"`
	}
	type structIntTriplePtrOmitEmpty struct {
		A ***int `json:"a,omitempty"`
	}
	type structIntTriplePtrString struct {
		A ***int `json:"a,string"`
	}

	tests := []struct {
		name string
		data interface{}
	}{
		{
			name: "Int",
			data: 10,
		},
		{
			name: "IntPtr",
			data: intptr(10),
		},
		{
			name: "IntTriplePtr",
			data: intptr3(10),
		},
		{
			name: "IntSlice",
			data: []int{1, 2, 3, 4, 5},
		},
		{
			name: "IntPtrSlice",
			data: []*int{nil, nil, intptr(1), nil, nil},
		},
		{
			name: "StructIntSliceNil",
			data: ([]structInt)(nil),
		},
		{
			name: "PtrStructIntSliceNil",
			data: ([]*structInt)(nil),
		},
		{
			name: "StructIntSliceZero",
			data: []structInt{},
		},
		{
			name: "PtrStructIntSliceZero",
			data: []*structInt{},
		},
		{
			name: "PtrStructIntNilSlice",
			data: []*structInt{nil, nil},
		},
		{
			name: "StructIntOmitEmptySliceNil",
			data: ([]structIntOmitEmpty)(nil),
		},
		{
			name: "PtrStructIntOmitEmptySliceNil",
			data: ([]*structIntOmitEmpty)(nil),
		},
		{
			name: "StructIntOmitEmptySliceZero",
			data: []structIntOmitEmpty{},
		},
		{
			name: "PtrStructIntOmitEmptySliceZero",
			data: []*structIntOmitEmpty{},
		},
		{
			name: "PtrStructIntOmitEmptyNilSlice",
			data: []*structIntOmitEmpty{nil, nil},
		},

		// HeadIntZero
		{
			name: "HeadIntZero",
			data: struct {
				A int `json:"a"`
			}{},
		},
		{
			name: "HeadIntZeroOmitEmpty",
			data: struct {
				A int `json:"a,omitempty"`
			}{},
		},
		{
			name: "HeadIntZeroString",
			data: struct {
				A int `json:"a,string"`
			}{},
		},

		// HeadInt
		{
			name: "HeadInt",
			data: struct {
				A int `json:"a"`
			}{A: -1},
		},
		{
			name: "HeadIntOmitEmpty",
			data: struct {
				A int `json:"a,omitempty"`
			}{A: -1},
		},
		{
			name: "HeadIntString",
			data: struct {
				A int `json:"a,string"`
			}{A: -1},
		},

		// HeadIntPtr
		{
			name: "HeadIntPtr",
			data: struct {
				A *int `json:"a"`
			}{A: intptr(-1)},
		},
		{
			name: "HeadIntPtrOmitEmpty",
			data: struct {
				A *int `json:"a,omitempty"`
			}{A: intptr(-1)},
		},
		{
			name: "HeadIntPtrString",
			data: struct {
				A *int `json:"a,string"`
			}{A: intptr(-1)},
		},

		// HeadIntPtrNil
		{
			name: "HeadIntPtrNil",
			data: struct {
				A *int `json:"a"`
			}{A: nil},
		},
		{
			name: "HeadIntPtrNilOmitEmpty",
			data: struct {
				A *int `json:"a,omitempty"`
			}{A: nil},
		},
		{
			name: "HeadIntPtrNilString",
			data: struct {
				A *int `json:"a,string"`
			}{A: nil},
		},

		// PtrHeadIntZero
		{
			name: "PtrHeadIntZero",
			data: &struct {
				A int `json:"a"`
			}{},
		},
		{
			name: "PtrHeadIntZeroOmitEmpty",
			data: &struct {
				A int `json:"a,omitempty"`
			}{},
		},
		{
			name: "PtrHeadIntZeroString",
			data: &struct {
				A int `json:"a,string"`
			}{},
		},

		// PtrHeadInt
		{
			name: "PtrHeadInt",
			data: &struct {
				A int `json:"a"`
			}{A: -1},
		},
		{
			name: "PtrHeadIntOmitEmpty",
			data: &struct {
				A int `json:"a,omitempty"`
			}{A: -1},
		},
		{
			name: "PtrHeadIntString",
			data: &struct {
				A int `json:"a,string"`
			}{A: -1},
		},

		// PtrHeadIntPtr
		{
			name: "PtrHeadIntPtr",
			data: &struct {
				A *int `json:"a"`
			}{A: intptr(-1)},
		},
		{
			name: "PtrHeadIntPtrOmitEmpty",
			data: &struct {
				A *int `json:"a,omitempty"`
			}{A: intptr(-1)},
		},
		{
			name: "PtrHeadIntPtrString",
			data: &struct {
				A *int `json:"a,string"`
			}{A: intptr(-1)},
		},

		// PtrHeadIntPtrNil
		{
			name: "PtrHeadIntPtrNil",
			data: &struct {
				A *int `json:"a"`
			}{A: nil},
		},
		{
			name: "PtrHeadIntPtrNilOmitEmpty",
			data: &struct {
				A *int `json:"a,omitempty"`
			}{A: nil},
		},
		{
			name: "PtrHeadIntPtrNilString",
			data: &struct {
				A *int `json:"a,string"`
			}{A: nil},
		},

		// PtrHeadIntNil
		{
			name: "PtrHeadIntNil",
			data: (*struct {
				A *int `json:"a"`
			})(nil),
		},
		{
			name: "PtrHeadIntNilOmitEmpty",
			data: (*struct {
				A *int `json:"a,omitempty"`
			})(nil),
		},
		{
			name: "PtrHeadIntNilString",
			data: (*struct {
				A *int `json:"a,string"`
			})(nil),
		},

		// HeadIntZeroMultiFields
		{
			name: "HeadIntZeroMultiFields",
			data: struct {
				A int `json:"a"`
				B int `json:"b"`
				C int `json:"c"`
			}{},
		},
		{
			name: "HeadIntZeroMultiFieldsOmitEmpty",
			data: struct {
				A int `json:"a,omitempty"`
				B int `json:"b,omitempty"`
				C int `json:"c,omitempty"`
			}{},
		},
		{
			name: "HeadIntZeroMultiFields",
			data: struct {
				A int `json:"a,string"`
				B int `json:"b,string"`
				C int `json:"c,string"`
			}{},
		},

		// HeadIntMultiFields
		{
			name: "HeadIntMultiFields",
			data: struct {
				A int `json:"a"`
				B int `json:"b"`
				C int `json:"c"`
			}{A: -1, B: 2, C: 3},
		},
		{
			name: "HeadIntMultiFieldsOmitEmpty",
			data: struct {
				A int `json:"a,omitempty"`
				B int `json:"b,omitempty"`
				C int `json:"c,omitempty"`
			}{A: -1, B: 2, C: 3},
		},
		{
			name: "HeadIntMultiFieldsString",
			data: struct {
				A int `json:"a,string"`
				B int `json:"b,string"`
				C int `json:"c,string"`
			}{A: -1, B: 2, C: 3},
		},

		// HeadIntPtrMultiFields
		{
			name: "HeadIntPtrMultiFields",
			data: struct {
				A *int `json:"a"`
				B *int `json:"b"`
				C *int `json:"c"`
			}{A: intptr(-1), B: intptr(2), C: intptr(3)},
		},
		{
			name: "HeadIntPtrMultiFieldsOmitEmpty",
			data: struct {
				A *int `json:"a,omitempty"`
				B *int `json:"b,omitempty"`
				C *int `json:"c,omitempty"`
			}{A: intptr(-1), B: intptr(2), C: intptr(3)},
		},
		{
			name: "HeadIntPtrMultiFieldsString",
			data: struct {
				A *int `json:"a,string"`
				B *int `json:"b,string"`
				C *int `json:"c,string"`
			}{A: intptr(-1), B: intptr(2), C: intptr(3)},
		},

		// HeadIntPtrNilMultiFields
		{
			name: "HeadIntPtrNilMultiFields",
			data: struct {
				A *int `json:"a"`
				B *int `json:"b"`
				C *int `json:"c"`
			}{A: nil, B: nil, C: nil},
		},
		{
			name: "HeadIntPtrNilMultiFieldsOmitEmpty",
			data: struct {
				A *int `json:"a,omitempty"`
				B *int `json:"b,omitempty"`
				C *int `json:"c,omitempty"`
			}{A: nil, B: nil, C: nil},
		},
		{
			name: "HeadIntPtrNilMultiFieldsString",
			data: struct {
				A *int `json:"a,string"`
				B *int `json:"b,string"`
				C *int `json:"c,string"`
			}{A: nil, B: nil, C: nil},
		},

		// PtrHeadIntZeroMultiFields
		{
			name: "PtrHeadIntZeroMultiFields",
			data: &struct {
				A int `json:"a"`
				B int `json:"b"`
			}{},
		},
		{
			name: "PtrHeadIntZeroMultiFieldsOmitEmpty",
			data: &struct {
				A int `json:"a,omitempty"`
				B int `json:"b,omitempty"`
			}{},
		},
		{
			name: "PtrHeadIntZeroMultiFieldsString",
			data: &struct {
				A int `json:"a,string"`
				B int `json:"b,string"`
			}{},
		},

		// PtrHeadIntMultiFields
		{
			name: "PtrHeadIntMultiFields",
			data: &struct {
				A int `json:"a"`
				B int `json:"b"`
			}{A: -1, B: 2},
		},
		{
			name: "PtrHeadIntMultiFieldsOmitEmpty",
			data: &struct {
				A int `json:"a,omitempty"`
				B int `json:"b,omitempty"`
			}{A: -1, B: 2},
		},
		{
			name: "PtrHeadIntMultiFieldsString",
			data: &struct {
				A int `json:"a,string"`
				B int `json:"b,string"`
			}{A: -1, B: 2},
		},

		// PtrHeadIntPtrMultiFields
		{
			name: "PtrHeadIntPtrMultiFields",
			data: &struct {
				A *int `json:"a"`
				B *int `json:"b"`
			}{A: intptr(-1), B: intptr(2)},
		},
		{
			name: "PtrHeadIntPtrMultiFieldsOmitEmpty",
			data: &struct {
				A *int `json:"a,omitempty"`
				B *int `json:"b,omitempty"`
			}{A: intptr(-1), B: intptr(2)},
		},
		{
			name: "PtrHeadIntPtrMultiFieldsString",
			data: &struct {
				A *int `json:"a,string"`
				B *int `json:"b,string"`
			}{A: intptr(-1), B: intptr(2)},
		},

		// PtrHeadIntPtrNilMultiFields
		{
			name: "PtrHeadIntPtrNilMultiFields",
			data: &struct {
				A *int `json:"a"`
				B *int `json:"b"`
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadIntPtrNilMultiFieldsOmitEmpty",
			data: &struct {
				A *int `json:"a,omitempty"`
				B *int `json:"b,omitempty"`
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadIntPtrNilMultiFieldsString",
			data: &struct {
				A *int `json:"a,string"`
				B *int `json:"b,string"`
			}{A: nil, B: nil},
		},

		// PtrHeadIntNilMultiFields
		{
			name: "PtrHeadIntNilMultiFields",
			data: (*struct {
				A *int `json:"a"`
				B *int `json:"b"`
			})(nil),
		},
		{
			name: "PtrHeadIntNilMultiFieldsOmitEmpty",
			data: (*struct {
				A *int `json:"a,omitempty"`
				B *int `json:"b,omitempty"`
			})(nil),
		},
		{
			name: "PtrHeadIntNilMultiFieldsString",
			data: (*struct {
				A *int `json:"a,string"`
				B *int `json:"b,string"`
			})(nil),
		},

		// HeadIntZeroNotRoot
		{
			name: "HeadIntZeroNotRoot",
			data: struct {
				A struct {
					A int `json:"a"`
				}
			}{},
		},
		{
			name: "HeadIntZeroNotRootOmitEmpty",
			data: struct {
				A struct {
					A int `json:"a,omitempty"`
				}
			}{},
		},
		{
			name: "HeadIntZeroNotRootString",
			data: struct {
				A struct {
					A int `json:"a,string"`
				}
			}{},
		},

		// HeadIntNotRoot
		{
			name: "HeadIntNotRoot",
			data: struct {
				A struct {
					A int `json:"a"`
				}
			}{A: struct {
				A int `json:"a"`
			}{A: -1}},
		},
		{
			name: "HeadIntNotRootOmitEmpty",
			data: struct {
				A struct {
					A int `json:"a,omitempty"`
				}
			}{A: struct {
				A int `json:"a,omitempty"`
			}{A: -1}},
		},
		{
			name: "HeadIntNotRootString",
			data: struct {
				A struct {
					A int `json:"a,string"`
				}
			}{A: struct {
				A int `json:"a,string"`
			}{A: -1}},
		},

		// HeadIntPtrNotRoot
		{
			name: "HeadIntPtrNotRoot",
			data: struct {
				A struct {
					A *int `json:"a"`
				}
			}{A: struct {
				A *int `json:"a"`
			}{intptr(-1)}},
		},
		{
			name: "HeadIntPtrNotRootOmitEmpty",
			data: struct {
				A struct {
					A *int `json:"a,omitempty"`
				}
			}{A: struct {
				A *int `json:"a,omitempty"`
			}{intptr(-1)}},
		},
		{
			name: "HeadIntPtrNotRootString",
			data: struct {
				A struct {
					A *int `json:"a,string"`
				}
			}{A: struct {
				A *int `json:"a,string"`
			}{intptr(-1)}},
		},

		// HeadIntPtrNilNotRoot
		{
			name: "HeadIntPtrNilNotRoot",
			data: struct {
				A struct {
					A *int `json:"a"`
				}
			}{},
		},
		{
			name: "HeadIntPtrNilNotRootOmitEmpty",
			data: struct {
				A struct {
					A *int `json:"a,omitempty"`
				}
			}{},
		},
		{
			name: "HeadIntPtrNilNotRootString",
			data: struct {
				A struct {
					A *int `json:"a,string"`
				}
			}{},
		},

		// PtrHeadIntZeroNotRoot
		{
			name: "PtrHeadIntZeroNotRoot",
			data: struct {
				A *struct {
					A int `json:"a"`
				}
			}{A: new(struct {
				A int `json:"a"`
			})},
		},
		{
			name: "PtrHeadIntZeroNotRootOmitEmpty",
			data: struct {
				A *struct {
					A int `json:"a,omitempty"`
				}
			}{A: new(struct {
				A int `json:"a,omitempty"`
			})},
		},
		{
			name: "PtrHeadIntZeroNotRootString",
			data: struct {
				A *struct {
					A int `json:"a,string"`
				}
			}{A: new(struct {
				A int `json:"a,string"`
			})},
		},

		// PtrHeadIntNotRoot
		{
			name: "PtrHeadIntNotRoot",
			data: struct {
				A *struct {
					A int `json:"a"`
				}
			}{A: &(struct {
				A int `json:"a"`
			}{A: -1})},
		},
		{
			name: "PtrHeadIntNotRootOmitEmpty",
			data: struct {
				A *struct {
					A int `json:"a,omitempty"`
				}
			}{A: &(struct {
				A int `json:"a,omitempty"`
			}{A: -1})},
		},
		{
			name: "PtrHeadIntNotRootString",
			data: struct {
				A *struct {
					A int `json:"a,string"`
				}
			}{A: &(struct {
				A int `json:"a,string"`
			}{A: -1})},
		},

		// PtrHeadIntPtrNotRoot
		{
			name: "PtrHeadIntPtrNotRoot",
			data: struct {
				A *struct {
					A *int `json:"a"`
				}
			}{A: &(struct {
				A *int `json:"a"`
			}{A: intptr(-1)})},
		},
		{
			name: "PtrHeadIntPtrNotRootOmitEmpty",
			data: struct {
				A *struct {
					A *int `json:"a,omitempty"`
				}
			}{A: &(struct {
				A *int `json:"a,omitempty"`
			}{A: intptr(-1)})},
		},
		{
			name: "PtrHeadIntPtrNotRootString",
			data: struct {
				A *struct {
					A *int `json:"a,string"`
				}
			}{A: &(struct {
				A *int `json:"a,string"`
			}{A: intptr(-1)})},
		},

		// PtrHeadIntPtrNilNotRoot
		{
			name: "PtrHeadIntPtrNilNotRoot",
			data: struct {
				A *struct {
					A *int `json:"a"`
				}
			}{A: &(struct {
				A *int `json:"a"`
			}{A: nil})},
		},
		{
			name: "PtrHeadIntPtrNilNotRootOmitEmpty",
			data: struct {
				A *struct {
					A *int `json:"a,omitempty"`
				}
			}{A: &(struct {
				A *int `json:"a,omitempty"`
			}{A: nil})},
		},
		{
			name: "PtrHeadIntPtrNilNotRootString",
			data: struct {
				A *struct {
					A *int `json:"a,string"`
				}
			}{A: &(struct {
				A *int `json:"a,string"`
			}{A: nil})},
		},

		// PtrHeadIntNilNotRoot
		{
			name: "PtrHeadIntNilNotRoot",
			data: struct {
				A *struct {
					A *int `json:"a"`
				}
			}{A: nil},
		},
		{
			name: "PtrHeadIntNilNotRootOmitEmpty",
			data: struct {
				A *struct {
					A *int `json:"a,omitempty"`
				} `json:",omitempty"`
			}{A: nil},
		},
		{
			name: "PtrHeadIntNilNotRootString",
			data: struct {
				A *struct {
					A *int `json:"a,string"`
				} `json:",string"`
			}{A: nil},
		},

		// HeadIntZeroMultiFieldsNotRoot
		{
			name: "HeadIntZeroMultiFieldsNotRoot",
			data: struct {
				A struct {
					A int `json:"a"`
				}
				B struct {
					B int `json:"b"`
				}
			}{},
		},
		{
			name: "HeadIntZeroMultiFieldsNotRootOmitEmpty",
			data: struct {
				A struct {
					A int `json:"a,omitempty"`
				}
				B struct {
					B int `json:"b,omitempty"`
				}
			}{},
		},
		{
			name: "HeadIntZeroMultiFieldsNotRootString",
			data: struct {
				A struct {
					A int `json:"a,string"`
				}
				B struct {
					B int `json:"b,string"`
				}
			}{},
		},

		// HeadIntMultiFieldsNotRoot
		{
			name: "HeadIntMultiFieldsNotRoot",
			data: struct {
				A struct {
					A int `json:"a"`
				}
				B struct {
					B int `json:"b"`
				}
			}{A: struct {
				A int `json:"a"`
			}{A: -1}, B: struct {
				B int `json:"b"`
			}{B: 2}},
		},
		{
			name: "HeadIntMultiFieldsNotRootOmitEmpty",
			data: struct {
				A struct {
					A int `json:"a,omitempty"`
				}
				B struct {
					B int `json:"b,omitempty"`
				}
			}{A: struct {
				A int `json:"a,omitempty"`
			}{A: -1}, B: struct {
				B int `json:"b,omitempty"`
			}{B: 2}},
		},
		{
			name: "HeadIntMultiFieldsNotRootString",
			data: struct {
				A struct {
					A int `json:"a,string"`
				}
				B struct {
					B int `json:"b,string"`
				}
			}{A: struct {
				A int `json:"a,string"`
			}{A: -1}, B: struct {
				B int `json:"b,string"`
			}{B: 2}},
		},

		// HeadIntPtrMultiFieldsNotRoot
		{
			name: "HeadIntPtrMultiFieldsNotRoot",
			data: struct {
				A struct {
					A *int `json:"a"`
				}
				B struct {
					B *int `json:"b"`
				}
			}{A: struct {
				A *int `json:"a"`
			}{A: intptr(-1)}, B: struct {
				B *int `json:"b"`
			}{B: intptr(2)}},
		},
		{
			name: "HeadIntPtrMultiFieldsNotRootOmitEmpty",
			data: struct {
				A struct {
					A *int `json:"a,omitempty"`
				}
				B struct {
					B *int `json:"b,omitempty"`
				}
			}{A: struct {
				A *int `json:"a,omitempty"`
			}{A: intptr(-1)}, B: struct {
				B *int `json:"b,omitempty"`
			}{B: intptr(2)}},
		},
		{
			name: "HeadIntPtrMultiFieldsNotRootString",
			data: struct {
				A struct {
					A *int `json:"a,string"`
				}
				B struct {
					B *int `json:"b,string"`
				}
			}{A: struct {
				A *int `json:"a,string"`
			}{A: intptr(-1)}, B: struct {
				B *int `json:"b,string"`
			}{B: intptr(2)}},
		},

		// HeadIntPtrNilMultiFieldsNotRoot
		{
			name: "HeadIntPtrNilMultiFieldsNotRoot",
			data: struct {
				A struct {
					A *int `json:"a"`
				}
				B struct {
					B *int `json:"b"`
				}
			}{A: struct {
				A *int `json:"a"`
			}{A: nil}, B: struct {
				B *int `json:"b"`
			}{B: nil}},
		},
		{
			name: "HeadIntPtrNilMultiFieldsNotRootOmitEmpty",
			data: struct {
				A struct {
					A *int `json:"a,omitempty"`
				}
				B struct {
					B *int `json:"b,omitempty"`
				}
			}{A: struct {
				A *int `json:"a,omitempty"`
			}{A: nil}, B: struct {
				B *int `json:"b,omitempty"`
			}{B: nil}},
		},
		{
			name: "HeadIntPtrNilMultiFieldsNotRootString",
			data: struct {
				A struct {
					A *int `json:"a,string"`
				}
				B struct {
					B *int `json:"b,string"`
				}
			}{A: struct {
				A *int `json:"a,string"`
			}{A: nil}, B: struct {
				B *int `json:"b,string"`
			}{B: nil}},
		},

		// PtrHeadIntZeroMultiFieldsNotRoot
		{
			name: "PtrHeadIntZeroMultiFieldsNotRoot",
			data: &struct {
				A struct {
					A int `json:"a"`
				}
				B struct {
					B int `json:"b"`
				}
			}{},
		},
		{
			name: "PtrHeadIntZeroMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A struct {
					A int `json:"a,omitempty"`
				}
				B struct {
					B int `json:"b,omitempty"`
				}
			}{},
		},
		{
			name: "PtrHeadIntZeroMultiFieldsNotRootString",
			data: &struct {
				A struct {
					A int `json:"a,string"`
				}
				B struct {
					B int `json:"b,string"`
				}
			}{},
		},

		// PtrHeadIntMultiFieldsNotRoot
		{
			name: "PtrHeadIntMultiFieldsNotRoot",
			data: &struct {
				A struct {
					A int `json:"a"`
				}
				B struct {
					B int `json:"b"`
				}
			}{A: struct {
				A int `json:"a"`
			}{A: -1}, B: struct {
				B int `json:"b"`
			}{B: 2}},
		},
		{
			name: "PtrHeadIntMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A struct {
					A int `json:"a,omitempty"`
				}
				B struct {
					B int `json:"b,omitempty"`
				}
			}{A: struct {
				A int `json:"a,omitempty"`
			}{A: -1}, B: struct {
				B int `json:"b,omitempty"`
			}{B: 2}},
		},
		{
			name: "PtrHeadIntMultiFieldsNotRootString",
			data: &struct {
				A struct {
					A int `json:"a,string"`
				}
				B struct {
					B int `json:"b,string"`
				}
			}{A: struct {
				A int `json:"a,string"`
			}{A: -1}, B: struct {
				B int `json:"b,string"`
			}{B: 2}},
		},

		// PtrHeadIntPtrMultiFieldsNotRoot
		{
			name: "PtrHeadIntPtrMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A *int `json:"a"`
				}
				B *struct {
					B *int `json:"b"`
				}
			}{A: &(struct {
				A *int `json:"a"`
			}{A: intptr(-1)}), B: &(struct {
				B *int `json:"b"`
			}{B: intptr(2)})},
		},
		{
			name: "PtrHeadIntPtrMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A *int `json:"a,omitempty"`
				}
				B *struct {
					B *int `json:"b,omitempty"`
				}
			}{A: &(struct {
				A *int `json:"a,omitempty"`
			}{A: intptr(-1)}), B: &(struct {
				B *int `json:"b,omitempty"`
			}{B: intptr(2)})},
		},
		{
			name: "PtrHeadIntPtrMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A *int `json:"a,string"`
				}
				B *struct {
					B *int `json:"b,string"`
				}
			}{A: &(struct {
				A *int `json:"a,string"`
			}{A: intptr(-1)}), B: &(struct {
				B *int `json:"b,string"`
			}{B: intptr(2)})},
		},

		// PtrHeadIntPtrNilMultiFieldsNotRoot
		{
			name: "PtrHeadIntPtrNilMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A *int `json:"a"`
				}
				B *struct {
					B *int `json:"b"`
				}
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadIntPtrNilMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A *int `json:"a,omitempty"`
				} `json:",omitempty"`
				B *struct {
					B *int `json:"b,omitempty"`
				} `json:",omitempty"`
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadIntPtrNilMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A *int `json:"a,string"`
				} `json:",string"`
				B *struct {
					B *int `json:"b,string"`
				} `json:",string"`
			}{A: nil, B: nil},
		},

		// PtrHeadIntNilMultiFieldsNotRoot
		{
			name: "PtrHeadIntNilMultiFieldsNotRoot",
			data: (*struct {
				A *struct {
					A *int `json:"a"`
				}
				B *struct {
					B *int `json:"b"`
				}
			})(nil),
		},
		{
			name: "PtrHeadIntNilMultiFieldsNotRootOmitEmpty",
			data: (*struct {
				A *struct {
					A *int `json:"a,omitempty"`
				}
				B *struct {
					B *int `json:"b,omitempty"`
				}
			})(nil),
		},
		{
			name: "PtrHeadIntNilMultiFieldsNotRootString",
			data: (*struct {
				A *struct {
					A *int `json:"a,string"`
				}
				B *struct {
					B *int `json:"b,string"`
				}
			})(nil),
		},

		// PtrHeadIntDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadIntDoubleMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A int `json:"a"`
					B int `json:"b"`
				}
				B *struct {
					A int `json:"a"`
					B int `json:"b"`
				}
			}{A: &(struct {
				A int `json:"a"`
				B int `json:"b"`
			}{A: -1, B: 2}), B: &(struct {
				A int `json:"a"`
				B int `json:"b"`
			}{A: 3, B: 4})},
		},
		{
			name: "PtrHeadIntDoubleMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A int `json:"a,omitempty"`
					B int `json:"b,omitempty"`
				}
				B *struct {
					A int `json:"a,omitempty"`
					B int `json:"b,omitempty"`
				}
			}{A: &(struct {
				A int `json:"a,omitempty"`
				B int `json:"b,omitempty"`
			}{A: -1, B: 2}), B: &(struct {
				A int `json:"a,omitempty"`
				B int `json:"b,omitempty"`
			}{A: 3, B: 4})},
		},
		{
			name: "PtrHeadIntDoubleMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A int `json:"a,string"`
					B int `json:"b,string"`
				}
				B *struct {
					A int `json:"a,string"`
					B int `json:"b,string"`
				}
			}{A: &(struct {
				A int `json:"a,string"`
				B int `json:"b,string"`
			}{A: -1, B: 2}), B: &(struct {
				A int `json:"a,string"`
				B int `json:"b,string"`
			}{A: 3, B: 4})},
		},

		// PtrHeadIntNilDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadIntNilDoubleMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A int `json:"a"`
					B int `json:"b"`
				}
				B *struct {
					A int `json:"a"`
					B int `json:"b"`
				}
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadIntNilDoubleMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A int `json:"a,omitempty"`
					B int `json:"b,omitempty"`
				} `json:",omitempty"`
				B *struct {
					A int `json:"a,omitempty"`
					B int `json:"b,omitempty"`
				} `json:",omitempty"`
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadIntNilDoubleMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A int `json:"a,string"`
					B int `json:"b,string"`
				}
				B *struct {
					A int `json:"a,string"`
					B int `json:"b,string"`
				}
			}{A: nil, B: nil},
		},

		// PtrHeadIntNilDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadIntNilDoubleMultiFieldsNotRoot",
			data: (*struct {
				A *struct {
					A int `json:"a"`
					B int `json:"b"`
				}
				B *struct {
					A int `json:"a"`
					B int `json:"b"`
				}
			})(nil),
		},
		{
			name: "PtrHeadIntNilDoubleMultiFieldsNotRootOmitEmpty",
			data: (*struct {
				A *struct {
					A int `json:"a,omitempty"`
					B int `json:"b,omitempty"`
				}
				B *struct {
					A int `json:"a,omitempty"`
					B int `json:"b,omitempty"`
				}
			})(nil),
		},
		{
			name: "PtrHeadIntNilDoubleMultiFieldsNotRootString",
			data: (*struct {
				A *struct {
					A int `json:"a,string"`
					B int `json:"b,string"`
				}
				B *struct {
					A int `json:"a,string"`
					B int `json:"b,string"`
				}
			})(nil),
		},

		// PtrHeadIntPtrDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadIntPtrDoubleMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A *int `json:"a"`
					B *int `json:"b"`
				}
				B *struct {
					A *int `json:"a"`
					B *int `json:"b"`
				}
			}{A: &(struct {
				A *int `json:"a"`
				B *int `json:"b"`
			}{A: intptr(-1), B: intptr(2)}), B: &(struct {
				A *int `json:"a"`
				B *int `json:"b"`
			}{A: intptr(3), B: intptr(4)})},
		},
		{
			name: "PtrHeadIntPtrDoubleMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A *int `json:"a,omitempty"`
					B *int `json:"b,omitempty"`
				}
				B *struct {
					A *int `json:"a,omitempty"`
					B *int `json:"b,omitempty"`
				}
			}{A: &(struct {
				A *int `json:"a,omitempty"`
				B *int `json:"b,omitempty"`
			}{A: intptr(-1), B: intptr(2)}), B: &(struct {
				A *int `json:"a,omitempty"`
				B *int `json:"b,omitempty"`
			}{A: intptr(3), B: intptr(4)})},
		},
		{
			name: "PtrHeadIntPtrDoubleMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A *int `json:"a,string"`
					B *int `json:"b,string"`
				}
				B *struct {
					A *int `json:"a,string"`
					B *int `json:"b,string"`
				}
			}{A: &(struct {
				A *int `json:"a,string"`
				B *int `json:"b,string"`
			}{A: intptr(-1), B: intptr(2)}), B: &(struct {
				A *int `json:"a,string"`
				B *int `json:"b,string"`
			}{A: intptr(3), B: intptr(4)})},
		},

		// PtrHeadIntPtrNilDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadIntPtrNilDoubleMultiFieldsNotRoot",
			data: &struct {
				A *struct {
					A *int `json:"a"`
					B *int `json:"b"`
				}
				B *struct {
					A *int `json:"a"`
					B *int `json:"b"`
				}
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadIntPtrNilDoubleMultiFieldsNotRootOmitEmpty",
			data: &struct {
				A *struct {
					A *int `json:"a,omitempty"`
					B *int `json:"b,omitempty"`
				} `json:",omitempty"`
				B *struct {
					A *int `json:"a,omitempty"`
					B *int `json:"b,omitempty"`
				} `json:",omitempty"`
			}{A: nil, B: nil},
		},
		{
			name: "PtrHeadIntPtrNilDoubleMultiFieldsNotRootString",
			data: &struct {
				A *struct {
					A *int `json:"a,string"`
					B *int `json:"b,string"`
				}
				B *struct {
					A *int `json:"a,string"`
					B *int `json:"b,string"`
				}
			}{A: nil, B: nil},
		},

		// PtrHeadIntPtrNilDoubleMultiFieldsNotRoot
		{
			name: "PtrHeadIntPtrNilDoubleMultiFieldsNotRoot",
			data: (*struct {
				A *struct {
					A *int `json:"a"`
					B *int `json:"b"`
				}
				B *struct {
					A *int `json:"a"`
					B *int `json:"b"`
				}
			})(nil),
		},
		{
			name: "PtrHeadIntPtrNilDoubleMultiFieldsNotRootOmitEmpty",
			data: (*struct {
				A *struct {
					A *int `json:"a,omitempty"`
					B *int `json:"b,omitempty"`
				}
				B *struct {
					A *int `json:"a,omitempty"`
					B *int `json:"b,omitempty"`
				}
			})(nil),
		},
		{
			name: "PtrHeadIntPtrNilDoubleMultiFieldsNotRootString",
			data: (*struct {
				A *struct {
					A *int `json:"a,string"`
					B *int `json:"b,string"`
				}
				B *struct {
					A *int `json:"a,string"`
					B *int `json:"b,string"`
				}
			})(nil),
		},

		// AnonymousHeadInt
		{
			name: "AnonymousHeadInt",
			data: struct {
				structInt
				B int `json:"b"`
			}{
				structInt: structInt{A: -1},
				B:         2,
			},
		},
		{
			name: "AnonymousHeadIntOmitEmpty",
			data: struct {
				structIntOmitEmpty
				B int `json:"b,omitempty"`
			}{
				structIntOmitEmpty: structIntOmitEmpty{A: -1},
				B:                  2,
			},
		},
		{
			name: "AnonymousHeadIntString",
			data: struct {
				structIntString
				B int `json:"b,string"`
			}{
				structIntString: structIntString{A: -1},
				B:               2,
			},
		},

		// PtrAnonymousHeadInt
		{
			name: "PtrAnonymousHeadInt",
			data: struct {
				*structInt
				B int `json:"b"`
			}{
				structInt: &structInt{A: -1},
				B:         2,
			},
		},
		{
			name: "PtrAnonymousHeadIntOmitEmpty",
			data: struct {
				*structIntOmitEmpty
				B int `json:"b,omitempty"`
			}{
				structIntOmitEmpty: &structIntOmitEmpty{A: -1},
				B:                  2,
			},
		},
		{
			name: "PtrAnonymousHeadIntString",
			data: struct {
				*structIntString
				B int `json:"b,string"`
			}{
				structIntString: &structIntString{A: -1},
				B:               2,
			},
		},

		// NilPtrAnonymousHeadInt
		{
			name: "NilPtrAnonymousHeadInt",
			data: struct {
				*structInt
				B int `json:"b"`
			}{
				structInt: nil,
				B:         2,
			},
		},
		{
			name: "NilPtrAnonymousHeadIntOmitEmpty",
			data: struct {
				*structIntOmitEmpty
				B int `json:"b,omitempty"`
			}{
				structIntOmitEmpty: nil,
				B:                  2,
			},
		},
		{
			name: "NilPtrAnonymousHeadIntString",
			data: struct {
				*structIntString
				B int `json:"b,string"`
			}{
				structIntString: nil,
				B:               2,
			},
		},

		// AnonymousHeadIntPtr
		{
			name: "AnonymousHeadIntPtr",
			data: struct {
				structIntPtr
				B *int `json:"b"`
			}{
				structIntPtr: structIntPtr{A: intptr(-1)},
				B:            intptr(2),
			},
		},
		{
			name: "AnonymousHeadIntPtrOmitEmpty",
			data: struct {
				structIntPtrOmitEmpty
				B *int `json:"b,omitempty"`
			}{
				structIntPtrOmitEmpty: structIntPtrOmitEmpty{A: intptr(-1)},
				B:                     intptr(2),
			},
		},
		{
			name: "AnonymousHeadIntPtrString",
			data: struct {
				structIntPtrString
				B *int `json:"b,string"`
			}{
				structIntPtrString: structIntPtrString{A: intptr(-1)},
				B:                  intptr(2),
			},
		},

		// AnonymousHeadIntPtrNil
		{
			name: "AnonymousHeadIntPtrNil",
			data: struct {
				structIntPtr
				B *int `json:"b"`
			}{
				structIntPtr: structIntPtr{A: nil},
				B:            intptr(2),
			},
		},
		{
			name: "AnonymousHeadIntPtrNilOmitEmpty",
			data: struct {
				structIntPtrOmitEmpty
				B *int `json:"b,omitempty"`
			}{
				structIntPtrOmitEmpty: structIntPtrOmitEmpty{A: nil},
				B:                     intptr(2),
			},
		},
		{
			name: "AnonymousHeadIntPtrNilString",
			data: struct {
				structIntPtrString
				B *int `json:"b,string"`
			}{
				structIntPtrString: structIntPtrString{A: nil},
				B:                  intptr(2),
			},
		},

		// PtrAnonymousHeadIntPtr
		{
			name: "PtrAnonymousHeadIntPtr",
			data: struct {
				*structIntPtr
				B *int `json:"b"`
			}{
				structIntPtr: &structIntPtr{A: intptr(-1)},
				B:            intptr(2),
			},
		},
		{
			name: "PtrAnonymousHeadIntPtrOmitEmpty",
			data: struct {
				*structIntPtrOmitEmpty
				B *int `json:"b,omitempty"`
			}{
				structIntPtrOmitEmpty: &structIntPtrOmitEmpty{A: intptr(-1)},
				B:                     intptr(2),
			},
		},
		{
			name: "PtrAnonymousHeadIntPtrString",
			data: struct {
				*structIntPtrString
				B *int `json:"b,string"`
			}{
				structIntPtrString: &structIntPtrString{A: intptr(-1)},
				B:                  intptr(2),
			},
		},

		// NilPtrAnonymousHeadIntPtr
		{
			name: "NilPtrAnonymousHeadIntPtr",
			data: struct {
				*structIntPtr
				B *int `json:"b"`
			}{
				structIntPtr: nil,
				B:            intptr(2),
			},
		},
		{
			name: "NilPtrAnonymousHeadIntPtrOmitEmpty",
			data: struct {
				*structIntPtrOmitEmpty
				B *int `json:"b,omitempty"`
			}{
				structIntPtrOmitEmpty: nil,
				B:                     intptr(2),
			},
		},
		{
			name: "NilPtrAnonymousHeadIntPtrString",
			data: struct {
				*structIntPtrString
				B *int `json:"b,string"`
			}{
				structIntPtrString: nil,
				B:                  intptr(2),
			},
		},

		// AnonymousHeadIntOnly
		{
			name: "AnonymousHeadIntOnly",
			data: struct {
				structInt
			}{
				structInt: structInt{A: -1},
			},
		},
		{
			name: "AnonymousHeadIntOnlyOmitEmpty",
			data: struct {
				structIntOmitEmpty
			}{
				structIntOmitEmpty: structIntOmitEmpty{A: -1},
			},
		},
		{
			name: "AnonymousHeadIntOnlyString",
			data: struct {
				structIntString
			}{
				structIntString: structIntString{A: -1},
			},
		},

		// PtrAnonymousHeadIntOnly
		{
			name: "PtrAnonymousHeadIntOnly",
			data: struct {
				*structInt
			}{
				structInt: &structInt{A: -1},
			},
		},
		{
			name: "PtrAnonymousHeadIntOnlyOmitEmpty",
			data: struct {
				*structIntOmitEmpty
			}{
				structIntOmitEmpty: &structIntOmitEmpty{A: -1},
			},
		},
		{
			name: "PtrAnonymousHeadIntOnlyString",
			data: struct {
				*structIntString
			}{
				structIntString: &structIntString{A: -1},
			},
		},

		// NilPtrAnonymousHeadIntOnly
		{
			name: "NilPtrAnonymousHeadIntOnly",
			data: struct {
				*structInt
			}{
				structInt: nil,
			},
		},
		{
			name: "NilPtrAnonymousHeadIntOnlyOmitEmpty",
			data: struct {
				*structIntOmitEmpty
			}{
				structIntOmitEmpty: nil,
			},
		},
		{
			name: "NilPtrAnonymousHeadIntOnlyString",
			data: struct {
				*structIntString
			}{
				structIntString: nil,
			},
		},

		// AnonymousHeadIntPtrOnly
		{
			name: "AnonymousHeadIntPtrOnly",
			data: struct {
				structIntPtr
			}{
				structIntPtr: structIntPtr{A: intptr(-1)},
			},
		},
		{
			name: "AnonymousHeadIntPtrOnlyOmitEmpty",
			data: struct {
				structIntPtrOmitEmpty
			}{
				structIntPtrOmitEmpty: structIntPtrOmitEmpty{A: intptr(-1)},
			},
		},
		{
			name: "AnonymousHeadIntPtrOnlyString",
			data: struct {
				structIntPtrString
			}{
				structIntPtrString: structIntPtrString{A: intptr(-1)},
			},
		},

		// AnonymousHeadIntPtrNilOnly
		{
			name: "AnonymousHeadIntPtrNilOnly",
			data: struct {
				structIntPtr
			}{
				structIntPtr: structIntPtr{A: nil},
			},
		},
		{
			name: "AnonymousHeadIntPtrNilOnlyOmitEmpty",
			data: struct {
				structIntPtrOmitEmpty
			}{
				structIntPtrOmitEmpty: structIntPtrOmitEmpty{A: nil},
			},
		},
		{
			name: "AnonymousHeadIntPtrNilOnlyString",
			data: struct {
				structIntPtrString
			}{
				structIntPtrString: structIntPtrString{A: nil},
			},
		},

		// PtrAnonymousHeadIntPtrOnly
		{
			name: "PtrAnonymousHeadIntPtrOnly",
			data: struct {
				*structIntPtr
			}{
				structIntPtr: &structIntPtr{A: intptr(-1)},
			},
		},
		{
			name: "PtrAnonymousHeadIntPtrOnlyOmitEmpty",
			data: struct {
				*structIntPtrOmitEmpty
			}{
				structIntPtrOmitEmpty: &structIntPtrOmitEmpty{A: intptr(-1)},
			},
		},
		{
			name: "PtrAnonymousHeadIntPtrOnlyString",
			data: struct {
				*structIntPtrString
			}{
				structIntPtrString: &structIntPtrString{A: intptr(-1)},
			},
		},

		// NilPtrAnonymousHeadIntPtrOnly
		{
			name: "NilPtrAnonymousHeadIntPtrOnly",
			data: struct {
				*structIntPtr
			}{
				structIntPtr: nil,
			},
		},
		{
			name: "NilPtrAnonymousHeadIntPtrOnlyOmitEmpty",
			data: struct {
				*structIntPtrOmitEmpty
			}{
				structIntPtrOmitEmpty: nil,
			},
		},
		{
			name: "NilPtrAnonymousHeadIntPtrOnlyString",
			data: struct {
				*structIntPtrString
			}{
				structIntPtrString: nil,
			},
		},
	}
	for _, test := range tests {
		for _, indent := range []bool{true, false} {
			for _, htmlEscape := range []bool{true, false} {
				var buf bytes.Buffer
				enc := json.NewEncoder(&buf)
				enc.SetEscapeHTML(htmlEscape)
				if indent {
					enc.SetIndent("", "  ")
				}
				if err := enc.Encode(test.data); err != nil {
					t.Fatalf("%s(htmlEscape:%T): %+v: %s", test.name, htmlEscape, test.data, err)
				}
				stdresult := encodeByEncodingJSON(test.data, indent, htmlEscape)
				if buf.String() != stdresult {
					t.Errorf("%s(htmlEscape:%T): doesn't compatible with encoding/json. expected %q but got %q", test.name, htmlEscape, stdresult, buf.String())
				}
			}
		}
	}
}
