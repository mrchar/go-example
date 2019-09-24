# 关于使用 go modules 的几个实验

## example1 
top 包中声明module为github.com/mrchar/go-example/gomod/example1/top，引用sub包。在sub目录中创建go.mod文件之前就可以在top目录下使用go build直接获取到sub包，所以在使用go modules时，依赖的包不必使用go modules。
然后在sub中创建go.mod文件，声明module为github.com/mrchar/go-example/gomod/example1/sub，在top目录下依然可以顺利获得sub包并成功编译。
不过使用 go get 命令获取 github.com/mrchar/go-example/gomod/example1/sub go.mod文件中记录的将是sub包的版本，而不是go-example的仓库版本。

## example2
使用一个module同时管理多个包

## example3
在sub目录中的go.mod声明错误的module，module名和路径没有对应的情况下，在top包中引用，会发生错误。

## example4
在top包中引用sub包的时候不使用完整路径，而是直接使用sub作为包名，执行go build，提示找不到包。然后在top目录下执行go get github.com/mrchar/go-example/gomod/example4/sub，go.mod文件中会添加github.com/mrchar/go-example/gomod/example4/sub，编辑go.mod文件，使用replace将sub映射到github.com/mrchar/go-example/gomod/example4/sub，再次执行go build可以通过。