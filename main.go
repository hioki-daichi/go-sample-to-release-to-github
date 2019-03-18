//go:generate statik -src=./public

package main

import (
	"fmt"
	"io/ioutil"

	_ "github.com/hioki-daichi/go-sample-to-release-to-github/statik"
	"github.com/rakyll/statik/fs"
)

func main() {
	fmt.Println(foo())
}

func foo() string {
	statikFS, _ := fs.New()
	f, _ := statikFS.Open("/hello.txt")
	b, _ := ioutil.ReadAll(f)
	return string(b)
}
