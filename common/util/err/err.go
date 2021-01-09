package err

// Check will exit 1 when error is not nil
func Check(err error) {
	if err == nil {
		return
	}
	panic(err)
}

