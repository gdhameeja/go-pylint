package src

import "github.com/go-pylint/src/checkers"
import "github.com/go-pylint/src/ast"

type ASTWalker struct {
	nbStatements int

	// subject to change
	// visitEvents and leaveEvents are maps of nodes and the methods
	// (func visit.. leave..) on them
	visitEvents map[ast.Node][]string
	leaveEvents map[ast.Node][]string
	linter      PyLinter
}


// walk the checker's methods and collect all func with visit.. or leave..
// for example: visitFunctionDef, visitSimpleStatement, leaveFor, leaveIf etc.
func (w ASTWalker) AddChecker(checker checkers.Checker) {
	// TODO(gaurav): string slices for now
	// check the golang reflection api
	var visitFuncs []string
	var leaveFuncs []string
}


// param astroidNode of type Node is a single node in the ast
func (walker ASTWalker) Walk(astroidNode ast.Node) {
	visitEvents := walker.visitEvents[astroidNode]
	leaveEvents := walker.leaveEvents[astroidNode]
	if astroidNode.IsStatement() {
		walker.nbStatements++
	}

	for vistmethod := range visitEvents {
		visitMethod(astroidNode)
	}

	for leaveMethod := range leaveEvents {
		leaveMethod(astroidNode)
	}

}

func InitWalker(linter PyLinter) *ASTWalker {
	walker := new(ASTWalker)
	walker.linter = linter
	return walker
}
