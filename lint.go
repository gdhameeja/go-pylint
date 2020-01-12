package main

import (
	"fmt"
)

// master class that handles all the checkers
// also controls the report generation
type PyLinter struct {
	checkers []Checker

	// Our current design to store reports against checker
	// is going to be with a checker_id (unique to each checker)
	// against a slice of strings. This is because golang doesn't
	// let us define comparable types (which may be used as
	// keys in maps
	rm ReportManager
}

// regsiter a new checker with the linter
// also registers the reports associated with the checker
func (p PyLinter) RegisterChecker(checker Checker) {
	p.checkers = append(p.checkers, checker)
	rId, rTitle := checker.report[0], checker.report[1]
	p.rm.RegisterReport(rId, rTitle, checker)
}

// main entry point; check files or modules by their
// names. filesOrModules is a list of strings presenting the
// modules to check; although if config.jobs (number of workers)
// is specified as 1 (default) we use internal method called _do_check()
func (p PyLinter) Check(filesOrModules []string) {}

func (p PyLinter) setReporter(reporter Reporter) {}



// use the given linter to lint the files with given amount of workers (jobs)
// for now we're going to do everything on the main thread, then maybe
// later split it into number of jobs
func checkParallel(linter PyLinter, jobs int, files []string) {
	// use the jobs arguement later
	fmt.Println(jobs)
}
