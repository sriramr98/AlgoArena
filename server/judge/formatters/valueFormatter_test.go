package formatters

import (
	"reflect"
	"testing"
)

func TestValueFormatter(t *testing.T) {
	t.Run("format integer array", func(t *testing.T) {
		formatter := ValueFormatter[[]int]{}
		input := "[1, 2, 3, 4, 5]"
		expected := []int{1, 2, 3, 4, 5}

		result, err := formatter.Format(input)

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got: %v", expected, result)
		}
	})

	t.Run("format string array", func(t *testing.T) {
		formatter := ValueFormatter[[]string]{}
		input := `["apple", "banana", "cherry"]`
		expected := []string{"apple", "banana", "cherry"}

		result, err := formatter.Format(input)

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got: %v", expected, result)
		}
	})

	t.Run("format boolean array", func(t *testing.T) {
		formatter := ValueFormatter[[]bool]{}
		input := "[true, false, true]"
		expected := []bool{true, false, true}

		result, err := formatter.Format(input)

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got: %v", expected, result)
		}
	})

	t.Run("format float array", func(t *testing.T) {
		formatter := ValueFormatter[[]float64]{}
		input := "[1.1, 2.2, 3.3]"
		expected := []float64{1.1, 2.2, 3.3}

		result, err := formatter.Format(input)

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got: %v", expected, result)
		}
	})

	t.Run("format empty array", func(t *testing.T) {
		formatter := ValueFormatter[[]int]{}
		input := "[]"
		expected := []int{}

		result, err := formatter.Format(input)

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got: %v", expected, result)
		}
	})

	t.Run("format nested array", func(t *testing.T) {
		formatter := ValueFormatter[[][]int]{}
		input := "[[1, 2], [3, 4], [5, 6]]"
		expected := [][]int{{1, 2}, {3, 4}, {5, 6}}

		result, err := formatter.Format(input)

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got: %v", expected, result)
		}
	})

	t.Run("format array of objects", func(t *testing.T) {
		type Person struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		formatter := ValueFormatter[[]Person]{}
		input := `[{"name":"John","age":30},{"name":"Jane","age":25}]`
		expected := []Person{
			{Name: "John", Age: 30},
			{Name: "Jane", Age: 25},
		}

		result, err := formatter.Format(input)

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got: %v", expected, result)
		}
	})

	t.Run("format simple object", func(t *testing.T) {
		type Person struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		formatter := ValueFormatter[Person]{}
		input := `{"name":"John","age":30}`
		expected := Person{Name: "John", Age: 30}

		result, err := formatter.Format(input)

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got: %v", expected, result)
		}
	})

	t.Run("format primitive types", func(t *testing.T) {
		// Test with integer
		intFormatter := ValueFormatter[int]{}
		intInput := "42"
		intExpected := 42

		intResult, err := intFormatter.Format(intInput)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
		if intResult != intExpected {
			t.Errorf("Expected %v, got: %v", intExpected, intResult)
		}

		// Test with string
		strFormatter := ValueFormatter[string]{}
		strInput := `"hello world"`
		strExpected := "hello world"

		strResult, err := strFormatter.Format(strInput)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
		if strResult != strExpected {
			t.Errorf("Expected %v, got: %v", strExpected, strResult)
		}

		// Test with boolean
		boolFormatter := ValueFormatter[bool]{}
		boolInput := "true"
		boolExpected := true

		boolResult, err := boolFormatter.Format(boolInput)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
		if boolResult != boolExpected {
			t.Errorf("Expected %v, got: %v", boolExpected, boolResult)
		}
	})

	t.Run("handle invalid JSON", func(t *testing.T) {
		formatter := ValueFormatter[[]int]{}
		input := "[1, 2, 3," // Invalid JSON missing closing bracket

		_, err := formatter.Format(input)

		if err == nil {
			t.Error("Expected an error for invalid JSON, got nil")
		}
	})

	t.Run("handle type mismatch", func(t *testing.T) {
		formatter := ValueFormatter[[]int]{}
		input := `["a", "b", "c"]` // String values, not integers

		_, err := formatter.Format(input)

		if err == nil {
			t.Error("Expected an error for type mismatch, got nil")
		}
	})

	t.Run("handle incompatible types", func(t *testing.T) {
		// Try to parse an object as an array
		formatter := ValueFormatter[[]int]{}
		input := `{"key": "value"}` // Object, not array

		_, err := formatter.Format(input)

		if err == nil {
			t.Error("Expected an error for incompatible types, got nil")
		}
	})
}
