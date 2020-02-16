package checkers


// IAstroidChecker
// Interface for checker which prefers receive events according to statement type
type Checker interface {
	getMessages() map[string]MessageDefinition
}

// Interface for checker which need to parse the raw file
type IRawChecker interface {
	ProcessModule()
}

// Interface for checkers that need access to the token list
type ITokenChecker interface {}


// Where actual checking happens
// AST based
type BaseChecker struct {
	checkerId int
	checkerName string

	// reportId, reportTitle
	Report [2]string
}

// Identifies with the property 'messages' that checkers have
func (c BaseChecker) getMessages() map[string]MessageDefinition {
	msgMap := make(map[string]MessageDefinition)
	return msgMap
}

// Initialize visit variables and stats
func (c BaseChecker) Open() {
}

func (c BaseChecker) Close() {
}
