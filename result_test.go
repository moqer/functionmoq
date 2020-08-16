package funcmoq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type CustomStruct1 struct {
	y1 interface{}
	y2 string
}

func newCustomStruct() CustomStruct {
	return CustomStruct{
		x2: make([]int, 4),
		x3: [5]int{1, 2, 3, 4, 5},
		x4: "test",
		x5: CustomStruct1{},
		x6: make([]CustomStruct1, 0),
		x8: make(map[string]CustomStruct1),
	}
}

type CustomStruct struct {
	x1 int
	x2 []int
	x3 [5]int
	x4 string
	x5 CustomStruct1
	x6 []CustomStruct1
	x7 interface{}
	x8 map[string]CustomStruct1
}

func (s CustomStruct) Test() string {
	return "test"
}

type Tester interface {
	Test() string
}

func TestSlice_Int(t *testing.T) {
	expected := []int{1, 2, 3, 4}
	result := Result{}
	result.Returning(nil, expected)
	var actual []int
	var err error
	assert.Nil(t, result.Retrieve(&err, &actual))
	for i := range expected {
		assert.Equal(t, actual[i], expected[i])
	}
}

func TestSlice_CustomStruct(t *testing.T) {
	expected := []CustomStruct{newCustomStruct()}
	result := Result{}
	result.Returning(nil, expected)
	var actual []CustomStruct
	var err error
	assert.Nil(t, result.Retrieve(&err, &actual))
	for i := range expected {
		assert.Equal(t, actual[i], expected[i])
	}
}

func TestArray_Int(t *testing.T) {
	expected := [4]int{1, 2, 3, 4}
	result := Result{}
	result.Returning(nil, expected)
	var actual [4]int
	var err error
	assert.Nil(t, result.Retrieve(&err, &actual))
	for i := range expected {
		assert.Equal(t, actual[i], expected[i])
	}
}

func TestArray_CustomStruct(t *testing.T) {
	expected := [2]CustomStruct{newCustomStruct(), newCustomStruct()}
	result := Result{}
	result.Returning(nil, expected)
	var actual [2]CustomStruct
	var err error
	assert.Nil(t, result.Retrieve(&err, &actual))
	for i := range expected {
		assert.Equal(t, actual[i], expected[i])
	}
}

func TestArray_CustomStructPointers(t *testing.T) {
	c1 := newCustomStruct()
	c2 := newCustomStruct()
	expected := [2]*CustomStruct{&c1, &c2}
	result := Result{}
	result.Returning(nil, expected)
	var actual [2]*CustomStruct
	var err error
	assert.Nil(t, result.Retrieve(&err, &actual))
	for i := range expected {
		assert.Equal(t, actual[i], expected[i])
	}
}

func TestElement_Int(t *testing.T) {
	expected := 1
	result := Result{}
	result.Returning(nil, expected)
	var actual int
	var err error
	assert.Nil(t, result.Retrieve(&err, &actual))
	assert.Equal(t, expected, actual)
}

func TestElement_string(t *testing.T) {
	expected := "test"
	result := Result{}
	result.Returning(nil, expected)
	var actual string
	var err error
	assert.Nil(t, result.Retrieve(&err, &actual))
	assert.Equal(t, expected, actual)
}

func TestElement_CustomStruct(t *testing.T) {
	expected := newCustomStruct()
	result := Result{}
	result.Returning(nil, expected)
	var actual CustomStruct
	var err error
	assert.Nil(t, result.Retrieve(&err, &actual))
	assert.Equal(t, expected, actual)
}

func TestElement_IntPointer(t *testing.T) {
	tmp := 1
	expected := &tmp
	result := Result{}
	result.Returning(nil, expected)
	var actual *int
	var err error
	assert.Nil(t, result.Retrieve(&err, &actual))
	assert.Equal(t, expected, actual)
}
func TestElement_StringPointer(t *testing.T) {
	tmp := "test"
	expected := &tmp
	result := Result{}
	result.Returning(nil, expected)
	var actual *string
	var err error
	assert.Nil(t, result.Retrieve(&err, &actual))
	assert.Equal(t, expected, actual)
}
func TestElement_CustomStringPointer(t *testing.T) {
	tmp := newCustomStruct()
	expected := &tmp
	result := Result{}
	result.Returning(nil, expected)
	var actual *CustomStruct
	var err error
	assert.Nil(t, result.Retrieve(&err, &actual))
	assert.Equal(t, expected, actual)
}
func TestElement_EmptyInterface(t *testing.T) {
	var expected interface{} = 1
	result := Result{}
	result.Returning(nil, expected)
	var actual interface{}
	var err error
	assert.Nil(t, result.Retrieve(&err, &actual))
	assert.Equal(t, expected, actual)
}

func TestElement_TesterInterface(t *testing.T) {
	var expected Tester = newCustomStruct()
	result := Result{}
	result.Returning(nil, expected)
	var actual CustomStruct
	var err error
	assert.Nil(t, result.Retrieve(&err, &actual))
	assert.Equal(t, expected, actual)
}

func TestElement_NilErrorInterface(t *testing.T) {
	var expected error
	result := Result{}
	result.Returning(nil, expected)
	var actual error
	var err error
	assert.Nil(t, result.Retrieve(&err, &actual))
	assert.Equal(t, expected, actual)
}
