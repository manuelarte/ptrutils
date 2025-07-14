package ptrutils

import "testing"

func TestPointerUtils_DefOr_String(t *testing.T) {
	example := "example"

	tests := map[string]struct {
		pointer      *string
		defaultValue string
		expected     string
	}{
		"pointer is null": {
			pointer:      nil,
			defaultValue: "defaultValue",
			expected:     "defaultValue",
		},
		"pointer is not null": {
			pointer:      &example,
			defaultValue: "defaultValue",
			expected:     example,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := DefOr(test.pointer, test.defaultValue)
			if actual != test.expected {
				t.Errorf("expected: %s, got: %s", test.expected, actual)
			}
		})
	}
}

func TestPointerUtils_DefOrL_String(t *testing.T) {
	example := "example"

	tests := map[string]struct {
		pointer      *string
		defaultValue string
		expected     string
	}{
		"pointer is null": {
			pointer:      nil,
			defaultValue: "defaultValue",
			expected:     "defaultValue",
		},
		"pointer is not null": {
			pointer:      &example,
			defaultValue: "defaultValue",
			expected:     example,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := DefOrL(test.pointer, LazyIdentity(test.defaultValue))
			if actual != test.expected {
				t.Errorf("expected: %s, got: %s", test.expected, actual)
			}
		})
	}
}

func TestPointerUtils_PointerTo_String(t *testing.T) {
	example := "example"
	emptyString := ""

	tests := map[string]struct {
		value func() *string
	}{
		"string with value": {
			value: LazyIdentity(&example),
		},
		"empty string": {
			value: LazyIdentity(&emptyString),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			expected := test.value()

			actual := Ptr(*expected)
			if *actual != *test.value() {
				t.Errorf("expected: %v, got: %v", expected, actual)
			}
		})
	}
}
