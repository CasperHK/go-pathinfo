# pathinfo

<br/>
## Installation
```bash

```

## Usage
```go
path := pathinfo.NewPath("/www/htdocs/inc/lib.inc.php")

fmt.Println(path.ExtrDirName())
fmt.Println(path.ExtrBaseName())
fmt.Println(path.ExtrExtension())
fmt.Println(path.ExtrFileName())

fmt.Println(path.HasExtension())
fmt.Println(path.MoveUp(2))
fmt.Println(path.IsRelative())
fmt.Println(path.IsAbsolute())
fmt.Println(path.IsExt([]string{"php", "doc"}))
fmt.Println(path.IsNotExt([]string{"php", "doc"}))
fmt.Println(path.IsFromUnixLike())
fmt.Println(path.IsFromWindow())
```
