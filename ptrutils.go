package ptrutils

// DefOr Dereference the pointer unless nil, then it returns defaultValue.
func DefOr[T any](ptr *T, defaultValue T) T {
	return DefOrL(ptr, LazyIdentity(defaultValue))
}

// DefOrL Dereference the pointer unless nil, then it runs a function that returns a default value.
func DefOrL[T any](ptr *T, f func() T) T {
	if ptr == nil {
		return f()
	} else {
		return *ptr
	}
}

// LazyIdentity Defer the identity operation until the function is called.
func LazyIdentity[T any](value T) func() T {
	return func() T {
		return value
	}
}

// Ptr converts a value to a pointer.
func Ptr[T any](value T) *T {
	return &value
}
