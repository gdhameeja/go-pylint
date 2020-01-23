package src

import (
	"fmt"
	"reflect"
)

// Actual message omitted by checkers. These are used in reports??
// NOTE: can be configured to not be emitted on `PyLinter`
type Message struct {
	path string
	description string
	lineNumber int
}


// Definition of a message
// maintains the link to checkers
// Note: the original impl of MessageDefiniton also has a msgId
// but there is already a mapping from msgId -> MessageDefiniton in
// MessageDefinitionStore. So we're not going to store msgId at instance level
type MessageDefinition struct {
	msg          Message
	description  string
	checkerName  string
}



// Stores information about every possible message definition.
type MessageDefinitionStore struct {
	msgDefinitions map[string]MessageDefinition
}

func (store MessageDefinitionStore) registerMessage(msgId string, message MessageDefinition) {
	store.msgDefinitions[msgId] = message
}

func (store MessageDefinitionStore) registerMessagesFromChecker(checker Checker) {
	var messages map[string]MessageDefinition
	messages = checker.getMessages()
	for msgId, message := range messages {
		store.registerMessage(msgId, message)
	}
}

// get a msgId by MessageDefinition
// original pylint stores msgId as an attribute
// on the MessageDefinition class
func (store MessageDefinitionStore) getMessageId(message MessageDefinition) (string, error) {
	for key, value := range store.msgDefinitions {
		if reflect.DeepEqual(message, value) {
			return key, nil
		}
	}
	return "", fmt.Errorf("No key found with MessageDefinition value %s", message)
}
