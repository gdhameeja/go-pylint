package main

//TODO(gaurav): looks like Checker will have to be an interface

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
