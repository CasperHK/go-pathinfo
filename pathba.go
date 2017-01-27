package main

import (
    "fmt"
    "path/filepath"
    "path"
    "strings"
)

func main() {
    fmt.Println("Hello World")
    a := NewPath("/home/asb.doc")
    a.Print()
    fmt.Println(a.Base())
    fmt.Println(a.Dir())
    fmt.Println(a.Ext())
    fmt.Println(a.Name())
    fmt.Println(a.HasExt())
    fmt.Println(a.RemoveExt())
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

func (p *Path) Dir() string {
    dirn := path.Dir(p.path)
    return dirn
}

func (p *Path) Base() string {
    bn := path.Base(p.path)
    return bn
}

func (p *Path) Ext() string {
    ext := filepath.Ext(p.path)
    if ext != "" {
        ext = strings.TrimPrefix(ext, ".")
    }
    return strings.ToLower(ext)
}

func (p *Path) Name() string {
    var base = path.Base(p.path)
    var ext = path.Ext(p.path)
    var name = base[0:len(base)-len(ext)]
    return name
}

func (p *Path) HasExt() bool {
	if filepath.Ext(p.path) == "" {
		return false
	} else {
		return true
	}
}

func (p *Path) RemoveExt() string {
    ext := path.Ext(p.path)
    var pathWithoutExt string
    if ext != "" {
        pathWithoutExt = strings.TrimSuffix(p.path, ext)
    }
    return pathWithoutExt
}

// func (p *Path) MoveUp(numOfLevel int) Path {
// 	return Path{path: ""}
// }

// func (p *Path) ValidateExt(validExts []string) string {

// }

// func (p *Path) ValidateNotExt(invalidExts []string) string {

// }

// func (p *Path) IsFromUnixLike() bool {

// }

// func (p *Path) IsFromWindow() bool {

// }

// func (p *Path) IsBaseOnly() bool {

// }

// func (p *Path) IsAbs() bool {} {

// }

// func (p *Path) IsRel() bool {
	
// }