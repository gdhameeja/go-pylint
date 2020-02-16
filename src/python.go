// this is the module where we will write all interfacing code
// from golang to python

package src

import "github.com/go-pylint/src/ast"

// struct type to do all the ipc (inter process communication)
// and hold all the details from the IPC
type IPC struct {
	purePython bool
	tokenInfo []map[string]string
}

func GetAst(filepath, modname string) (*ast.AST, error) {
	return new(ast.AST), nil
}

