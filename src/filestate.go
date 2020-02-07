package main

// Hold the internal state specific to the currently analyzed file
// basically hold info about file being checked
type FileState struct {
	baseName string
	moduleMessagesState map[string]string
	rawModuleMessagesState map[string]string
	ignoredMessages map[string]string
	effectiveMaxLineNumber int
}

func InitFileState(modName string) *FileState {
	fileState := new(FileState)
	fileState.baseName = modName
	return &fileState
}
