package lib

func Assert(b bool, message string) {
	if !b {
		panic(message)
	}
}
