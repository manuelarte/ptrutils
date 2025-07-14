package ptrutils

// DerefOr Dereference the pointer unless nil, then it returns defaultValue.
func DerefOr[T any](ptr *T, defaultValue T) T {
	return DerefOrL(ptr, LazyIdentity(defaultValue))
}

// DerefOrL Dereference the pointer unless nil, then it runs a function that returns a default value.
func DerefOrL[T any](ptr *T, f func() T) T {
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
