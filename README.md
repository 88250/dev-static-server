# Dev Static Server

开发环境用静态资源 HTTP 伺服器。

## 背景

* 在 Windows 上使用 goexec 实时生成伺服器的话每次使用都需要设置防火墙，长久下来防火墙规则会非常多
* 使用 NGINX 等的话配置有点麻烦，我只需要一个简单的 HTTP 服务在开发时引入静态资源用

## 使用

1. 安装 golang，然后获取并编译 `go get -u github.com/b3log/dev-static-server`
2. 编译成功后将在 ${GOPATH}/bin 下生成名为 `dev-static-server` 的可执行文件
3. 将 `dev-static-server` 移到需要伺服静态资源的目录下启动即可，端口 9090

## 授权

Dev Static Server 使用 [木兰宽松许可证, 第1版](http://license.coscl.org.cn/MulanPSL) 开源协议。
