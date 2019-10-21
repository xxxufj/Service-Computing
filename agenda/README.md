## 任务目标：
1. 安装cobra
2. cobra的简单使用
3. agenda项目开发

## 安装 cobra
安装时在工作区使用命令 `go get -v github.com/spf13/cobra/cobra`

其实也可以直接用 `go get -d github.com/spf13/cobra/cobra` 指令进行只下载不安装的操作，因为有些包国内下载不了


执行指令后会出现类似下面的错误
```
Fetching https://golang.org/x/sys/unix?go-get=1
https fetch failed: Get https://golang.org/x/sys/unix?go-get=1: dial tcp 216.239.37.1:443: i/o timeout
```
通过报错信息了解到下载失败的包是 sys 和 text

解决方法是通过下面的指令从git上下载

```
git clone https://github.com/golang/sys
git clone https://github.com/golang/text
```
下载完成后重新进行cobra的安装

`go install github.com/spf13/cobra/cobra`

![img](https://github.com/xxxufj/Service-Computing/blob/master/agenda/pictures/1/1.PNG)

此时安装就不再出现报错信息了


## cobra 的简单使用
简单使用的目标是创建一个处理命令 register -uTestUser 的程序

首先创建程序 testcobra
`cobra init test --pkg-name=github.com/homework/test/testcobra`

添加指令 register

`cobra add register`

此时需要的文件就产生了

因为 `register` 有参数 `username`， 所以需要对 `register` 的 `init` 函数和 `Run` 匿名回调函数分别做修改

在 `init` 中添加
```go
  registerCmd.Flags().StringP("user", "u", "Anonymous", "Help message for username")
```

在 `run` 中添加
```go
username, _ := cmd.Flags().GetString("user")
fmt.Println("register called by " + username)
```
对 `register` 函数进行测试

```go
go run main.go register --user=TestUser
```

测试结果如下

![img](https://github.com/xxxufj/Service-Computing/tree/blob/master/agenda/pictures/2/1.PNG)




## agenda 项目开发
### register 
#### register 需求
* 注册新用户时，用户需设置一个唯一的用户名和一个密码。另外，还需登记邮箱及电话信息；
* 如果注册时提供的用户名已由其他用户使用，应反馈一个适当的出错信息
* 成功注册后，亦应反馈一个成功注册的信息；
* 使用方法：register -u username -p password -e email -t phoneNumber

#### 测试结果：

新用户注册

![img](https://github.com/xxxufj/Service-Computing/tree/blob/master/agenda/pictures/3/1.PNG)


错误范例：用户名重复注册

![img](https://github.com/xxxufj/Service-Computing/tree/blob/master/agenda/pictures/3/2.PNG)



#### login 需求
* 用户使用用户名和密码登录 Agenda 系统。
* 用户名和密码同时正确则登录成功并反馈一个成功登录的信息；
* 登录失败反馈一个失败登录的信息。
* 使用方法：login -u username -p password

#### 测试结果：

用户正确登陆 与 用户账号与密码不匹配

![img](https://github.com/xxxufj/Service-Computing/tree/blob/master/agenda/pictures/3/3.PNG)





