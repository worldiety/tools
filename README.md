# tools  [![Travis-CI](https://travis-ci.com/worldiety/tools.svg?branch=master)](https://travis-ci.com/worldiety/tools) [![Go Report Card](https://goreportcard.com/badge/github.com/worldiety/tools)](https://goreportcard.com/report/github.com/worldiety/tools) [![GoDoc](https://godoc.org/github.com/worldiety/tools?status.svg)](http://godoc.org/github.com/worldiety/tools)

tools is a go library with wrappings and tools around the language, which are usually private or hidden.

- `go list -json` is wrapped in a type safe way by 
    ```go
    package example
    import "github.com/worldiety/tools"
    
    package main(){
        pkg, err := tools.GoList(".")
        //...
    }
    ```

