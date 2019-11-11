# 程序效果

![](https://github.com/xxxufj/Service-Computing/blob/master/cloudgo/img/2.PNG)

# 概述
开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理

# 任务目标
* 熟悉 go 服务器工作原理
* 基于现有 web 库，编写一个简单 web 应用类似 cloudgo。
* 使用 curl 工具访问 web 程序
* 对 web 执行压力测试

# 要求
* 编程 web 服务程序 类似 cloudgo 应用。
  * 要求有详细的注释
  * 是否使用框架、选哪个框架自己决定 请在 README.md 说明你决策的依据
* 使用 curl 测试，将测试结果写入 README.md
* 使用 ab 测试，将测试结果写入 README.md。并解释重要参数

# 编程 web 服务程序
## 框架选择
* Martini是一个强大为了编写模块化Web应用而生的GO语言框架
* 与revel、beego等其他框架相比，martini 是一个新锐的微型框架，只带有简单的核心，包括路由功能和依赖注入容器inject
* 由于这一特征，使用martini做本次作业这样轻量级的程序时很有优势，因此本次作业中我选择了用 martini
* 换个角度说，martini营造的不是一个大而全的框架，而是一种组件生态martini-contrib，而且他的DI实现，让第三方库很容易改造为martini规范的中间件。

## martini 实现第一个框架
```go
package main

import "github.com/go-martini/martini"

func main() {
  // 创建一个典型的martini实例
  m := martini.Classic()
  
  // 接收对'\'的GET方法请求，第二个参数是对一请求的处理方法
  m.Get("/", func() string {
    return "Hello world!"
  })
  
  // 运行服务器
  m.Run()
}
```

* 执行 `go run main.go`

![](https://github.com/xxxufj/Service-Computing/blob/master/cloudgo/img/1.PNG)

<br>

## web 页面渲染
* 页面渲染使用martini的中间件 `render`
* 基本用法
```go
ackage main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)

func main() {
  m := martini.Classic()
  // render html templates from templates directory
  m.Use(render.Renderer())

  m.Get("/", func(r render.Render) {
    r.HTML(200, "hello", "jeremy")
  })

  m.Run()
}
```

## 实现一个简单的用户注册网页服务程序
### 用户注册信息定义
```go
type User struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
```

### 设置 render 的使用
```go
	// render html templates from templates directory
	m.Use(render.Renderer())

	// use pictures in assets directory
	m.Use(martini.Static("assets"))
```

### 用户发起请求 
```go
// This will set the Content-Type header to "application/html; charset=ISO-8859-1"
	m.Get("/", func(r render.Render) {
		//render 'login.tmpl'
		r.HTML(200, "login", map[string]interface{}{})
	})

```

![](https://github.com/xxxufj/Service-Computing/blob/master/cloudgo/img/2.PNG)

## 用户提交表单
```go
m.Post("/", binding.Bind(User{}), func(u User, r render.Render) {
		p := User{Username: u.Username, Password: u.Password}
		// render 'info.tmpl', and add information of user to the webpage
		r.HTML(200, "info", map[string]interface{}{"user": p})
	})
```

![](https://github.com/xxxufj/Service-Computing/blob/master/cloudgo/img/3.PNG)


# 使用 curl 工具访问 web 程序

![](https://github.com/xxxufj/Service-Computing/blob/master/cloudgo/img/4.PNG)

# 对 web 执行压力测试
* 执行命令 `ab -n 10000 -c 1000 http://localhost:3000/` 来进行压力测试
* 参数说明： -n表示请求数，-c表示并发数
* 可以看到，在1000个请求，100个并发的情况下，总共使用了1.2秒左右时间来完成处理

![](https://github.com/xxxufj/Service-Computing/blob/master/cloudgo/img/5.PNG)

* ab 命令参数
  * -n 执行的请求数量
  * -c 并发请求个数
  * -t 测试所进行的最大秒数
  * -p 包含了需要POST的数据的文件
  * -T POST数据所使用的Content-type头信息
  * -k 启用HTTP KeepAlive功能，即在一个HTTP会话中执行多个请求，默认时，不启用KeepAlive功能
  
* 结果说明
  * Document Path: 请求的资源
  * Document Length: 文档返回的长度，不包括相应头
  * Concurrency Level：并发数
  * Time taken for tests：完成所有请求总共花费的时间
  * Complete requests：成功请求的次数
  * Failed requests：失败请求的次数
  * Total transferred：总共传输的字节数
  * HTML transferred：实际页面传输的字节数
  * Requests per second：每秒请求数
  * Time per request: [ms] (mean)： 平均每个请求消耗的时间
  * Time per request: [ms] (mean, across all concurrent requests) ：上面的请求除以并发数
  * Transfer rate：传输速率

