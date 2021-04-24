# Package

In the most basic terms, A package is nothing but a directory inside your Go workspace containing one or more Go source files, or other Go package. Every Go source file belongs to a package. To declare a source file to be part of a package, we use the following syntax:
```
package <packagename>
```
The above package declaration must be the first command of code in Go source file. All the functions, types, and variables defined in the source file become part of the declared package. 
In simple words package is in Go like as namespace in other programming languages. 

## Package usage

Every package which will be used in Go program must be imported via `import` keyword. Go compiler has a strict rule about importing packages, that is each imported package **must** be used in code. That means if you declare some package in `import` block and don't use any variable, function or interface, compiler will report "panic" error. 

There are two types of packages or libraries *standard* and *custom*. 

**Standard libraries** are well tested, documented, and provided by go standard language installation. List of all available *standard* libraries you can find [here](https://golang.org/pkg/#stdlib).

**Custom libraries** are open source packages provided by golang community or your private package. Usually custom packages are stored on github repository and we can download them via following command from CLI `go get <repository_url>`. Also they can be installed on local storage via the following command set in CLI (UNIX OS) :
```
mkdir ~/go/src/<packageName>
cp <packageName>.go ~/go/src/<packageName>
go install <packageName>
```

    *It is good practice to make package on "Single responsibility" principe because it is easier to testing, documenting and users don't have to install/import functionality which they don't need.*



