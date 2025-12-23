package ptrutils

import "testing"

func TestDerefOr_String(t *testing.T) {
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
			actual := DerefOr(test.pointer, test.defaultValue)
			if actual != test.expected {
				t.Errorf("expected: %s, got: %s", test.expected, actual)
			}
		})
	}
}

func TestDerefOrL_String(t *testing.T) {
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
			actual := DerefOrL(test.pointer, lazyIdentity(test.defaultValue))
			if actual != test.expected {
				t.Errorf("expected: %s, got: %s", test.expected, actual)
			}
		})
	}
}

func TestPtr_String(t *testing.T) {
	example := "example"
	emptyString := ""

	tests := map[string]struct {
		value func() *string
	}{
		"string with value": {
			value: lazyIdentity(&example),
		},
		"empty string": {
			value: lazyIdentity(&emptyString),
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
