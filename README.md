## Porting pylint to golang.

Initial strategy is to call python code to build ast and provide details then use checkers written in go to run over them and produce reports
