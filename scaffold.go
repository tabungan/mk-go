package mk

// Scaffold scaffolds entry on the top of basePath.
//
// An error will be returned if the scaffolding process fails, otherwise
// nil will be returned.
func Scaffold(basePath string, entry Entry) error {
	return entry.Scaffold(basePath)
}

// ScaffoldOrPanic scaffolds entry on the top of basePath, or
// panic if the scaffolding process fails.
func ScaffoldOrPanic(basePath string, entry Entry) {
	if err := Scaffold(basePath, entry); err != nil {
		panic(err)
	}
}
