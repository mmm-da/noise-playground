package main

func panic_if_error(err error) {
	if err != nil {
		panic(err)
	}
}
