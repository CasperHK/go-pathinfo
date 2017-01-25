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
```
