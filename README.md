# HTTP文件服务器

一个简单的Go语言HTTP文件服务器，用于快速启动本地文件服务。

## 功能特性

- 🚀 快速启动本地HTTP文件服务器
- 📁 提供当前目录下文件的HTTP访问
- 📊 自动记录访问日志
- ⚙️ 可自定义监听地址和端口
- 🔧 简单易用的命令行参数

## 安装方式

### 方式一：直接下载可执行文件

从 [Releases](https://github.com/55gY/http/releases) 页面下载最新版本的可执行文件。

### 方式二：从源码编译

```bash
git clone https://github.com/55gY/http.git
cd http
go mod init http
go get github.com/gorilla/handlers
go build -o http.exe http.go
```

## 使用说明

### 基本用法

```bash
# 使用默认配置启动（127.0.0.1:8080）
./http.exe

```

### 自定义端口和地址

```bash
# 自定义端口
./http.exe -port 127.0.0.1:3000

# 监听所有网络接口
./http.exe -port 0.0.0.0:8080

# 只监听本地回环
./http.exe -port localhost:9000
```

### 命令行参数

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `-port` | `127.0.0.1:8080` | 监听地址和端口 |

## 使用场景

- 快速共享本地文件
- 前端开发时提供静态文件服务
- 临时搭建文件下载服务
- 本地测试HTTP服务

## 示例

启动服务器后，在浏览器中访问对应地址即可浏览当前目录下的文件：

```
http://127.0.0.1:8080
```

## 依赖

- Go 1.16+
- [gorilla/handlers](https://github.com/gorilla/handlers) - HTTP中间件

## 许可证

MIT License

## 贡献

欢迎提交Issue和Pull Request！

---

如果这个项目对您有帮助，请给个⭐️星标支持！