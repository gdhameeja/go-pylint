// contains the ASTWalker type and the constructor to build it

package main


type ASTWalker struct {
	nbStatements int

	// subject to change
	visitEvents map[string][]string
	leaveEvents map[string][]string
}

func InitWalker(linter PyLinter) *ASTWalker {
	walker := new(ASTWalker)
	walker.linter = linter
	return &walker
}
