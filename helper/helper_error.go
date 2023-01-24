package helper

func PanicIferror(err error) {
	if err != nil {
		panic(err)
	}
}
