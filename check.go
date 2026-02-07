package main

func mustNoError(err error) {
	if err != nil {
		panic(err)
	}
}
