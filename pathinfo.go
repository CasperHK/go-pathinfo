package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    fmt.Println("Hello World")
    a := NewPath("/home/asb.doc")
    a.Print()
    fmt.Println(a.HasExtension())
}

type Path struct {
    path string
}

func NewPath(p string) Path {
    return Path{path: p}
}

func (p *Path) Print() {
	fmt.Println(p.path)
}

func (p *Path) ExtrDirName() string {
	return filepath.Ext(p.path)
}

// func (p *Path) ExtrBaseName() string {

// }

// func (p *Path) ExtrExtension() string {

// }

// func (p *Path) ExtrFileName() string {

// }

func (p *Path) HasExtension() bool {
	if filepath.Ext(p.path) == "" {
		return false
	} else {
		return true
	}
}

// func (p *Path) MoveUp(numOfLevel int) Path {
// 	return Path{path: ""}
// }

