# go-ls
A recreation of a `ls` command written in go.

The command takes zero or one arguments. If given no arguments, the output will be based on the current directory. Otherwise, a path can be provided to print the contents of the directory with provided path. 


## Install 

To install the binary to use in your terminal, run the following commands from the root directory of this repository.

#### Reqirements

Go must be installed.


#### Steps 
1. Compile and install binary into go/bin directory.
```
cd go-ls && go install .
```

2. Identify location where go binaries are installed on your machine and copy it!
```
go env GOPATH
```

3. To use `go-ls` on the command line, add the following to your  `~/.zshrc`, replacing the path with the path outputte from the previous command.
```
export PATH=$PATH:/path/to/your/go/bin
```
5. Then open a new terminal window and `go-ls` will be available for you to use. 