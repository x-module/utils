# Fileutil
fileutil包支持文件基本操作。

<div STYLE="page-break-after: always;"></div>

## 源码:

- [https://github.com/duke-git/lancet/blob/main/fileutil/file.go](https://github.com/duke-git/lancet/blob/main/fileutil/file.go)

<div STYLE="page-break-after: always;"></div>

## 用法:
```go
import (
    "github.com/x-module/utils/fileutil"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录
- [ClearFile](#ClearFile)
- [CreateFile](#CreateFile)
- [CreateDir](#CreateDir)
- [CopyFile](#CopyFile)
- [FileMode](#FileMode)
- [MiMeType](#MiMeType)
- [IsExist](#IsExist)
- [IsLink](#IsLink)
- [IsDir](#IsDir)
- [ListFileNames](#ListFileNames)
- [RemoveFile](#RemoveFile)
- [ReadFileToString](#ReadFileToString)
- [ReadFileByLine](#ReadFileByLine)
- [Zip](#Zip)
- [UnZip](#UnZip)

<div STYLE="page-break-after: always;"></div>

## 文档



### <span id="ClearFile">ClearFile</span>
<p>清空文件内容</p>

<b>函数签名:</b>

```go
func ClearFile(path string) error
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/fileutil"
)

func main() {
    err := fileutil.ClearFile("./test.txt")
    if err != nil {
        fmt.Println(err)
    }
}
```

### <span id="CreateFile">CreateFile</span>
<p>创建文件，创建成功返回true, 否则返回false</p>

<b>函数签名:</b>

```go
func CreateFile(path string) bool
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/fileutil"
)

func main() {
    isCreatedSucceed := fileutil.CreateFile("./test.txt")
    fmt.Println(isCreatedSucceed)
}
```


### <span id="CreateDir">CreateDir</span>
<p>使用绝对路径创建嵌套目录，例如/a/, /a/b/</p>

<b>函数签名:</b>

```go
func CreateDir(absPath string) error
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/fileutil"
)

func main() {
    err := fileutil.CreateDir("/a/")
    fmt.Println(err)
}
```



### <span id="CopyFile">CopyFile</span>
<p>拷贝文件，会覆盖原有的文件</p>

<b>函数签名:</b>

```go
func CopyFile(srcFilePath string, dstFilePath string) error
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/fileutil"
)

func main() {
    err := fileutil.CopyFile("./test.txt", "./test_copy.txt")
    if err != nil {
        fmt.Println(err)
    }
}
```



### <span id="FileMode">FileMode</span>
<p>获取文件mode信息</p>

<b>函数签名:</b>

```go
func FileMode(path string) (fs.FileMode, error)
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/fileutil"
)

func main() {
    mode, err := fileutil.FileMode("./test.txt")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(mode)
}
```



### <span id="MiMeType">MiMeType</span>
<p>获取文件mime类型, 'file'参数的类型必须是string或者*os.File</p>

<b>函数签名:</b>

```go
func MiMeType(file any) string
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "os"
    "github.com/x-module/utils/fileutil"
)

func main() {
    type1 := fileutil.MiMeType("./test.txt")
    fmt.Println(type1) //text/plain; charset=utf-8

    f, _ := os.Open("./file.go")
    type2 := fileutil.MiMeType(f)
    fmt.Println(type2) //text/plain; charset=utf-8
}
```




### <span id="IsExist">IsExist</span>
<p>判断文件或目录是否存在</p>

<b>函数签名:</b>

```go
func IsExist(path string) bool
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/fileutil"
)

func main() {
    fileutil.CreateFile("./test.txt")
    isFileExist := fileutil.IsExist("./test.txt")
    fmt.Println(isFileExist) //true
}
```



### <span id="IsLink">IsLink</span>
<p>判断文件是否是符号链接</p>

<b>函数签名:</b>

```go
func IsLink(path string) bool
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/fileutil"
)

func main() {
    isLinkFile := fileutil.IsLink("./test.txt")
    fmt.Println(isLinkFile) //false
}
```



### <span id="IsDir">IsDir</span>
<p>判断参数是否是目录</p>

<b>函数签名:</b>

```go
func IsDir(path string) bool 
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/fileutil"
)

func main() {
    isDir := fileutil.IsDir("./")
    fmt.Println(isDir) //true

    isDir = fileutil.IsDir("./test.txt")
    fmt.Println(isDir) //false
}
```



### <span id="ListFileNames">ListFileNames</span>
<p>返回目录下所有文件名</p>

<b>函数签名:</b>

```go
func ListFileNames(path string) ([]string, error)
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/fileutil"
)

func main() {
    fileNames, _ := fileutil.ListFileNames("./")
    fmt.Println(fileNames)
}
```



### <span id="RemoveFile">RemoveFile</span>
<p>删除文件</p>

<b>函数签名:</b>

```go
func RemoveFile(path string) error
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/fileutil"
)

func main() {
    err := fileutil.RemoveFile("./test.txt")
    if err != nil {
        fmt.Println(err)
    }
}
```


### <span id="ReadFileToString">ReadFileToString</span>
<p>读取文件内容并返回字符串</p>

<b>函数签名:</b>

```go
func ReadFileToString(path string) (string, error)
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "os"
    "github.com/x-module/utils/fileutil"
)

func main() {
    path := "./test.txt"
    fileutil.CreateFile(path)

    f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
    f.WriteString("hello world")

    content, _ := fileutil.ReadFileToString(path)
    fmt.Println(content) //hello world
}
```



### <span id="ReadFileByLine">ReadFileByLine</span>
<p>按行读取文件内容，返回字符串切片包含每一行</p>

<b>函数签名:</b>

```go
func ReadFileByLine(path string)([]string, error)
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "os"
    "github.com/x-module/utils/fileutil"
)

func main() {
    path := "./text.txt"
    fileutil.CreateFile(path)

    f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
    defer f.Close()
    f.WriteString("hello\nworld")

    contents, _ := fileutil.ReadFileByLine(path)
    fmt.Println(contents) //[]string{"hello", "world"}
}
```



### <span id="Zip">Zip</span>
<p>zip压缩文件, fpath参数可以是文件或目录</p>

<b>函数签名:</b>

```go
func Zip(fpath string, destPath string) error
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/fileutil"
)

func main() {
    err := fileutil.Zip("./test.txt", "./test.zip")
    if err != nil {
        fmt.Println(err)
    }
}
```




### <span id="UnZip">UnZip</span>
<p>zip解压缩文件并保存在目录中</p>

<b>Signature:</b>

```go
func UnZip(zipFile string, destPath string) error
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/x-module/utils/fileutil"
)

func main() {
    err := fileutil.Zip("./test.zip", "./unzip/test.txt")
    if err != nil {
        fmt.Println(err)
    }
}
```





