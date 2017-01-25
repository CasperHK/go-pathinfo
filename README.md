# pathinfo

<br/>
## Installation
```bash

```

## Usage
```go
pathParts := pathinfo.NewPath("/www/htdocs/inc/lib.inc.php")

fmt.Println(pathParts.ExtrDirName())
fmt.Println(pathParts.ExtrBaseName())
fmt.Println(pathParts.ExtrExtension())
fmt.Println(pathParts.ExtrFileName())
```
