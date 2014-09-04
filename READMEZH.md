关于
===

基于[yujiod's wiki](https://github.com/yujiod/wiki)系统实现，移植到Beego框架下，主要方便于部署。 revel似乎不支持直接打包成一个binary文件。

从源码生成
---

1. 下载压缩包，解压缩至 $GOPATH\src\miniwiki

2. 安装依赖：

    ```
    go get github.com/pmezard/go-difflib/difflib
    go get github.com/astaxie/beego
    go get github.com/jinzhu/gorm
    go get github.com/mattn/go-sqlite3
    go get github.com/oal/beego-pongo2
    go get github.com/shurcooL/go/github_flavored_markdown
    ```
3. 生成二进制文件
4. 执行可执行文件

授权协议
---
MIT协议开源。