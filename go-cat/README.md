# go-cat
A recreation of a `cat` command written in go.

The command takes one argument. The command takes a path to a file as an argument. 

## Install 

To install the binary to use in your terminal, run the following commands from the root directory of this repository.

#### Reqirements

Go must be installed.


#### Steps 
1. Compile and install binary into go/bin directory.
```
cd go-cat && go install .
```

2. Identify location where go binaries are installed on your machine and copy it!
```
go env GOPATH
```

3. To use `go-cat` on the command line, add the following to your  `~/.zshrc`, replacing the path with the path outputte from the previous command.
```
export PATH=$PATH:/path/to/your/go/bin
```
5. Then open a new terminal window and `go-cat` will be available for you to use. 