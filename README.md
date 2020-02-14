## Porting pylint to golang.

Initial strategy is to call python code to build ast and provide details then use checkers written in go to run over them and produce reports

### Notes
- We build ast in python, and we tokenize the code in python
- Remember, tokenize uses ast node from built by astroid, which underneath is just
bufferedstream:
this is what utils.tokenize_module does with ast_node in PyLint
```
ipdb> ast_node.stream                                                                                                                                                           
<bound method Module.stream of <Module.foo l.0 at 0x7f70b9493080>>
ipdb> ast_node.stream()                                                                                                                                                         
<_io.BufferedReader name='/home/melvault/Documents/programs/read_pylint/foo.py'>
ipdb> ast_node.stream().readline                                                                                                                                                
<built-in method readline of _io.BufferedReader object at 0x7f70b9e32728>
ipdb> tokens   
```

- we also need tokenize to detect encoding of the file we are trying to read for some reason
```
astroid.builder.py
def open_source_file(filename):
    with open(filename, "rb") as byte_stream:
        encoding = detect_encoding(byte_stream.readline)[0]
    stream = open(filename, "r", newline=None, encoding=encoding)
    data = stream.read()
    return stream, encoding, data
```
