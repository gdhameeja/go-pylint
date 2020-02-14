package src

// IAstroidChecker
// Interface for checker which prefers receive events according to statement type
type IAstroidChecker interface {}

// Interface for checker which need to parse the raw file
type IRawChecker interface {
	ProcessModule()
}

// Interface for checkers that need access to the token list
type ITokenChecker interface {}


// Where actual checking happens
// AST based
type Checker struct {
	checkerId int
	checkerName string

	// reportId, reportTitle
	report [2]string
}

// Identifies with the property 'messages' that checkers have
func (c Checker) getMessages() map[string]MessageDefinition {
	msgMap := make(map[string]MessageDefinition)
	return msgMap
}

// Initialize visit variables and stats
func (c Checker) Open() {
}

func (c Checker) Close() {
}
