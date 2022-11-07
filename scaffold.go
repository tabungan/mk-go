package mk

func Scaffold(basePath string, entry Entry) error {
	return entry.Scaffold(basePath)
}

func ScaffoldOrPanic(basePath string, entry Entry) {
	if err := Scaffold(basePath, entry); err != nil {
		panic(err)
	}
}
