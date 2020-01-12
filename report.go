package main

type ReportManager struct {

	// Our current design to store reports against checker
	// is going to be with a checker_id (unique to each checker)
	// against a slice of strings. This is because golang doesn't
	// let us define comparable types (which may be used as
	// keys in maps
	reports map[int][]string
}

// rId is the unique identifier for the report
// rTitle is the report's title
// checker is the checker defining the report
// all checker's will implement a common method
// to generate the report
func (rm ReportManager) RegisterReport(rId string, rTitle string, checker Checker) {
}
