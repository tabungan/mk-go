package mk

func Scaffolding(basePath string, entry Entry) error {
	return entry.Scaffold(basePath)
}

func ScaffoldingOrPanic(basePath string, entry Entry) {
	Scaffolding(basePath, entry)
}