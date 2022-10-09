package mk

type Entry interface {
	Scaffold(basePath string) error
}
