package ptrutils

import "testing"

func TestDerefOr_String(t *testing.T) {
	example := "example"

	tests := map[string]struct {
		pointer      *string
		defaultValue string
		want         string
	}{
		"pointer is null": {
			pointer:      nil,
			defaultValue: "defaultValue",
			want:         "defaultValue",
		},
		"pointer is not null": {
			pointer:      &example,
			defaultValue: "defaultValue",
			want:         example,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := DerefOr(test.pointer, test.defaultValue)
			if got != test.want {
				t.Errorf("DerefOr = %s, want: %s", got, test.want)
			}
		})
	}
}

func TestDerefOrL_String(t *testing.T) {
	example := "example"

	tests := map[string]struct {
		pointer      *string
		defaultValue string
		want         string
	}{
		"pointer is null": {
			pointer:      nil,
			defaultValue: "defaultValue",
			want:         "defaultValue",
		},
		"pointer is not null": {
			pointer:      &example,
			defaultValue: "defaultValue",
			want:         example,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := DerefOrL(test.pointer, lazyIdentity(test.defaultValue))
			if got != test.want {
				t.Errorf("DerefOrL = %s, want: %s", got, test.want)
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
			want := test.value()

			got := Ptr(*want)
			if *got != *test.value() {
				t.Errorf("Ptr = %v, want: %v", got, want)
			}
		})
	}
}
