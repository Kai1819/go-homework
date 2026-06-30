# Go Homework

## 项目地址
https://github.com/Kai1819/go-homework

## 使用说明

1. 每个作业有对应目录，分别写实现和测试用例。
2. 编辑 `homework.go`，完成各个函数的实现。
3. 在本地运行 `go test -v` 验证代码。
4. 提交代码到你的分支。GitHub Actions 会自动运行测试。

## 清空缓存

1. go clean -cache -testcache -modcache
2. 重启cursor
3. 每次跑测试代码，清空缓存，cursor配置："go.testFlags": ["-v", "-count=1"]

## 拓展学习
#### 体系学习
刘丹冰：https://github.com/aceld/golang
案例：https://gobyexample-cn.github.io/
菜鸟：https://www.runoob.com/go/go-tutorial.html
博客：https://github.com/0voice

#### 专题学习
协程：https://github.com/0voice/Introduction-to-Golang/blob/main/%E6%96%87%E7%AB%A0/Go%E8%AF%AD%E8%A8%80%E4%B9%8Bgoroutine%E5%8D%8F%E7%A8%8B%E8%AF%A6%E8%A7%A3.md
并发：https://geektutu.com/post/hpg-mutex.html
并发：https://www.topgoer.com/%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B/%E5%B9%B6%E5%8F%91%E4%BB%8B%E7%BB%8D.html
并发：https://go.cyub.vip/concurrency/


#### 线上练习
https://exercism.org/tracks/go/exercises/annalyns-infiltration/edit