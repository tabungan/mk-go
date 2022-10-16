package mk

type Entries []Entry

func (entries Entries) Scaffold(basePath string) error {
	for _, entry := range entries {
		entry.Scaffold(basePath)
	}
	return nil
}
