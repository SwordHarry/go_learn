# Go
## Goroutine
## tcmalloc
## 垃圾回收：gc 堪用，离好用还有一段距离

## 用 go 开发的项目
- Docker
- Kubernetes
- etcd
- beego
- martini
- codis
- delve

## go 环境搭建

gocode: 代码补全
godef: 跳转到代码对应定义
guru: guru原来叫做oracle，能获得全部代码对应引用
gorename: 重命名Go源码文件
goimports: 自动帮你格式化import语句的，它来帮你决定import顺序

golangci-lint: 旧的用法是gometalinter，现在应该选择快5倍的golangci-lint，这个工具用来做静态检查工具，它会反馈代码风格并给出建议，这样就可以在源码保存是运行它能告诉你现有代码中有什么问题
go-outline: 当前文件做符号(Symbol)搜索，相当于文件大纲
go-symbols: 工作区符号搜索
gogetdoc: 鼠标悬停时可以显示文档(方法和类的签名等帮助信息)
gomodifytags: 提供tags管理，可以对struct的tag增删改
gotests: 用例测试
impl: 自动生成接口实现代码(stubs)
gopkgs: 列出Go包。

 【待下载】
fillstruct。根据结构体定义在使用时自动填充默认值


## go 命令
### go build
如果不存在 main 包，则只对当前目录下的程序源码进行语法检查，不会生成可执行文件。
### go run



***注意：Go 没有三目运算符，所以不支持* ?: *形式的条件判断。***

