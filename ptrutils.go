package ptrutils

// DerefOr Dereference the pointer unless nil, then it returns defaultValue.
func DerefOr[T any](ptr *T, defaultValue T) T {
	return DerefOrL(ptr, lazyIdentity(defaultValue))
}

// DerefOrL Dereference the pointer unless nil, then it runs a function that returns a default value.
func DerefOrL[T any](ptr *T, f func() T) T {
	if ptr == nil {
		return f()
	}

	return *ptr
}

// Ptr converts a value to a pointer.
func Ptr[T any](value T) *T {
	return &value
}

// lazyIdentity Defer the identity operation until the function is called.
func lazyIdentity[T any](value T) func() T {
	return func() T {
		return value
	}
}
