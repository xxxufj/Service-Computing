# 使用 GO 开发第一个包
## 一、选择包路径
按照文档说明，我的第一个包放在 `$ mkdir $GOPATH/src/github.com/user/hello` 下
在该目录中通过 `vim hello.go` 指令创建第一个文件
添加代码如下：
```go
package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
}
```

