"""
this module uses the astroid library and builds up an
ast of the python code we are trying to lint. The output
of this module will be read by the go code which will then
try to lint the code, since we can't do ast of python in go
we call this from go code and read the ast in go
"""
import astroid


MANAGER = astroid.MANAGER

class Walker:

    def __init__(self, files, modname=None):
        # for now modname is none because we just wanna be able to
        # make it work for one file, also I don't know what modename does
        self.file = file
        self.modname = modename
        

    def build_ast(self):
        try:
            return MANAGER.ast_from_file(self.file)
