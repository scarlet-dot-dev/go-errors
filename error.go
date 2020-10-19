package handle

// Handle will call fn on *e if it is not nil, or if the function is recovering
// from a panic.
func Handle(e *error, fn func(err error) error) {
	// first check for a panic
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			panic(err)
		}
		*e = fn(err)
		return
	}

	// not a panic / recover
	if err := *e; err != nil {
		*e = fn(err)
	}
}

// Check error will panic with err if the error is not nil.
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
