# Notes

Installed using homebrew

brew install go --cross-compile-common

Did not need to set $GOROOT
It is only a convention 
http://golang.org/doc/install#gopath

Code installed with homebrew have their bins in PATH 
so that was also not needed

mercurial (hg) is needed for `go get`
might use svn, hg, git or bzr
where is this config set in packages ?
https://code.google.com/p/go-wiki/wiki/GoGetTools

> go get code.google.com/p/go-tour/gotour
go: missing Mercurial command. See http://golang.org/s/gogetcmd
package code.google.com/p/go-tour/gotour: exec: "hg": executable file not found in $PATH

## GOPATH
http://golang.org/doc/code.html#GOPATH

GOPATH is where downloaded code is installed

If GOPATH is not set:

> go get code.google.com/p/go-tour/gotour
package code.google.com/p/go-tour/gotour: cannot download, $GOPATH not set. For more details see: go help gopath

Can set GOPATH to

export GOPATH="/usr/local/var/go"

CHECK: It can be colon seperated if multiple paths are to be specified


## terminology

- packages
- functions take parameters
- 

## Folder structure for packages

TODO: http://golang.org/doc/code.html#PackagePaths

## Running Go Programs

go run hello.go

## package manager aka rubygems ?

# warnings

importing a package but not using it will throw a warning
TODO: automatically collect a list of such warnings in a csv. maybe driven via tests

# list of community code 

# installing code (go get)


## comparisons with Ruby

no implicit returns like ruby
semicolons can be used as terminators. but is not necessary

## Questions

Programs start running in package main.

import (
    "fmt"
    "math"
)
fmt.Println("Happy", math.Pi, "Day")
like a ruby module function. not like include ie. still has seperate namespace

By convention, the package name is the same as the last element of the import path.

List names exported by a package

After importing a package, you can refer to the names it exports.
In Go, a name is exported if it begins with a capital letter.
Foo is an exported name, as is FOO. The name foo is not exported.

func add(x int, y int) int {
}
to define a function. why not keep it like c ?
TODO: read http://golang.org/doc/articles/gos_declaration_syntax.html

emacs mode for go ?

import (
    "fmt"
    "math"
)
what internal data-structure is that ?

TDD in go. what is the culture regarding tests in the community ?

how to find type of a variable ? and list of types in golang
reflection in go. variables/types in current scope ?
any type for "result parameter" in http://tour.golang.org/#9

what is `:=`
  a, b := swap("hello", "world")

return_all.go: do not care about unused variables
otherwise program does not run

