# pathinfo

<br/>

## Usage
```go
pathParts := pathinfo.NewPath("/www/htdocs/inc/lib.inc.php")
fmt.Println(pathParts.ExtractDirName())
fmt.Println(pathParts.ExtractBaseName())
fmt.Println(pathParts.ExtractExtension())
fmt.Println(pathParts.ExtractFileName())
```
