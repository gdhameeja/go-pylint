package main


// The AST type representing the python ast
// We will let python handle the parsing part and consume info through IPC
// and store all the details here
type AST struct {
	// all child nodes
	_fields []Node

	// line number
	lineNo int
}

// the ast node
type Node struct {}



// recursively yield all descendant nodes in the tree starting at *node*
// in no specified order. Useful if only want to modify nodes in place
// and don't care about the context.


// call python code to get all the ast details and return an AST
func BuildAST() *AST {
	ast := new(AST)
	return ast
}
