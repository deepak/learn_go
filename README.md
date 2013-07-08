# Notes

Installed using homebrew

brew install go --cross-compile-common

Did not need to set $GOROOT
It is only a convention 
http://golang.org/doc/install#gopath

Code installed with homebrew have their bins in PATH 
so that was also not needed

## GOPATH
http://golang.org/doc/code.html#GOPATH

GOPATH is where downloaded code is installed

If GOPATH is not set:

> go get code.google.com/p/go-tour/gotour
package code.google.com/p/go-tour/gotour: cannot download, $GOPATH not set. For more details see: go help gopath

Can set GOPATH to

export GOPATH="/usr/local/var/go"

CHECK: It can be colon seperated if multiple paths are to be specified
CHECK: ig hg is needed for `go get`. I had hg installed previously

## Folder structure for packages

TODO: http://golang.org/doc/code.html#PackagePaths

## Running Go Programs

go run hello.go

## package manager aka rubygems ?

# list of community code 

# installing code (go get)

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



 