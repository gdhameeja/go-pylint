package src

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
// for now we will only do one one worker
// NOTE: Add whether message which are registered in the message store may
//       be emitted or not as in `PyLinter.check`
func (p PyLinter) Check(filesOrModules []string) {
	p.doCheck(filesOrModules)
}

// walks the `filesOrModules` and builds the ast
func (p PyLinter) doCheck(filesOrModules []string) {
	walker := InitWalker(p)
}

func (p PyLinter) setReporter(reporter Reporter) {}


// PrepareChcecker returns array of all the needed checkers
// for now it returns all the available checkers.
func (p PyLinter) PrepareCheckers() []Checker {
	return []Checker {p}
}

// communicates with python code to get the ast
func (p PyLinter) GetAst(filepath string, modname string) {
}



// use the given linter to lint the files with given amount of workers (jobs)
// for now we're going to do everything on the main thread, then maybe
// later split it into number of jobs
func checkParallel(linter PyLinter, jobs int, files []string) {
	// use the jobs arguement later
	fmt.Println(jobs)
}
