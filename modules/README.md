# Modules 

A **module** is first presented in version 1.11. Module is a collection of packages that are released. versioned, and distributed together. Module can be downloaded directly from version control repositories or from proxy servers via the following command: 
```
go get <repository path>
```
A module is defined by a *module path*, which is declared in a `go.mod` file, together with information about the module's dependencies. The *module root directory* is the directory that contains the `go.mod` file. The *main module* is the module containing the directory where the `go` command is invoked. Go modules use **Semantic Versioning**
