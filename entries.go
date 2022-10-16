package mk

type Entries []Entry

func (entries Entries) Scaffold(basePath string) error {
	entries[0].Scaffold(basePath)
	return nil
}
