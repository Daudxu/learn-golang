# Learn Golang

> https://github.com/goproxy/goproxy.cn


```
 go env  // 输出 go 环境有关的信息

// 代理配置
 go env -w GO111MODULE=on
 go env -w GOPROXY=https://goproxy.cn,direct

 // goimports
go get -v -u golang.org/x/tools/cmd/goimports
// go mod init youer progect
```