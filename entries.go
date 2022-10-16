package mk

type Entries []Entry

func (entries Entries) Scaffold(basePath string) error {
	for _, entry := range entries {
		if err := entry.Scaffold(basePath); err != nil {
			return err
		}
	}
	return nil
}
