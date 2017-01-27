# PathBa

<br/>
## Installation
```bash
go get github.com/CasperHK/pathinfo
```

## Usage
```go
import {
    "github.com/CasperHK/pathinfo"
}
```
```go
path := pathinfo.NewPath("/www/htdocs/inc/lib.inc.php")

fmt.Println(path.Dir())
fmt.Println(path.Base())
fmt.Println(path.Ext())
fmt.Println(path.Name())

fmt.Println(path.HasExtension())
fmt.Println(path.MoveUp(2))
fmt.Println(path.IsRelative())
fmt.Println(path.IsAbsolute())
fmt.Println(path.IsExt([]string{"php", "doc"}))
fmt.Println(path.IsNotExt([]string{"php", "doc"}))
fmt.Println(path.IsFromUnixLike())
fmt.Println(path.IsFromWindow())
fmt.Println(path.IsOnlyBaseName())
fmt.Println(path.RemoveExtension())
```
