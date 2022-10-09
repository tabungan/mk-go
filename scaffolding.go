package mk

func Scaffolding(basePath string, entry Entry) error {
	return entry.Scaffold(basePath)
}
