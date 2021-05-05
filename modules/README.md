# Modules 

A **module** is first presented in version 1.11. Module is a collection of packages that are released. versioned, and distributed together. Module can be downloaded directly from version control repositories or from proxy servers via the following command: 
```
go get <repository path>
```
A module is defined by a *module path*, which is declared in a `go.mod` file, together with information about the module's dependencies. The *module root directory* is the directory that contains the `go.mod` file. The *main module* is the module containing the directory where the `go` command is invoked. Go modules use **Semantic Versioning** as version control.  


## Custom modules

If we wont to create custom go module, first we should make new repository on some version control service or on proxy server. After cloning the repository on local storage we must initialize go module by `go mod init` command, it will create `go.mod` file. After that we can create go module execution file. Module file and package should be named as module. 