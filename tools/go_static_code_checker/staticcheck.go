//go:build tools
// +build tools

// Package tools is a place holder for all golang binary toolings
// needed to maintain the repository health but is not a compilation dependency.
//
// The package is kept away from standard compilations by using a specific
// 'tools' build tag. Imported packages therefore could be 'main' package of
// other modules as this tools package will never be compiled.
package tools

import (
	_ "github.com/sluongng/staticcheck-codegen"
)
