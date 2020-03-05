// contains the ASTWalker type and the constructor to build it

package src

import "github.com/go-pylint/src/checkers"

type ASTWalker struct {
	nbStatements int

	// subject to change
	visitEvents map[string][]string
	leaveEvents map[string][]string
	linter      PyLinter
}

func (w ASTWalker) AddChecker(checker checkers.Checker) {
}

func InitWalker(linter PyLinter) *ASTWalker {
	walker := new(ASTWalker)
	walker.linter = linter
	return walker
}
