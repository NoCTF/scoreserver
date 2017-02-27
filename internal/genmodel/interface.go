package main

import (
	"go/ast"
	"reflect"
)

type scanner interface {
	Scan(...interface{}) error
}

type processFunc func(*context, *ast.Package) error
type inspector func(ast.Node) bool

type context struct {
	Dir     string
	PkgName string
	Models  []model
}

type field struct {
	Name       string
	Tag        reflect.StructTag
	Type       string
	JSONName   string
	IsNullable bool
}

type model struct {
	Fields  []field
	Name    string
	PkgName string
}

type processor struct {
	Dir string
}
