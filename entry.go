// Package mk provides utilities to scaffold files and directories.
//
// Inspired by mk* utilities,
// package entities are named in such a way to resemble them, e.g.
// mk.Dir represents the mkdir command.
// This package is initially intended for testing, but
// can be used for general scaffolding purposes.
package mk

// Type Entry defines the trait of scaffoldable entities.
//
// Scaffold scaffolds the entity on the top of basePath.
// An error should be returned if the scaffolding process fails, otherwise
// nil should be returned.
type Entry interface {
	Scaffold(basePath string) error
}
