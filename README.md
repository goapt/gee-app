# Gee Application

<a href="https://github.com/goapt/gee-app/actions"><img src="https://github.com/goapt/gee-app/workflows/build/badge.svg" alt="Build Status"></a>
<a href="https://codecov.io/gh/goapt/gee-app"><img src="https://codecov.io/gh/goapt/gee-app/branch/master/graph/badge.svg" alt="codecov"></a>
<a href="https://goreportcard.com/report/github.com/goapt/gee-app"><img src="https://goreportcard.com/badge/github.com/goapt/gee-app" alt="Go Report Card
"></a>
<a href="https://pkg.go.dev/github.com/goapt/gee-app"><img src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square" alt="GoDoc"></a>
<a href="https://opensource.org/licenses/mit-license.php" rel="nofollow"><img src="https://badges.frapsoft.com/os/mit/mit.svg?v=103"></a>


## 使用方法
1. `git clone https://github.com/goapt/gee-app.git app`
2. `go mod tidy`
3. 修改 `config.toml` 配置信息，如 `app_name` `storage_path` 等
4. `make` 编译

## Http Server
如果项目是一个http的接口项目，路由在`router`包中管理，`handler` 写到 `handler` 包，使用如下命令行启动

```
./app http --addr=:8081
```


## Cli Server
如果项目是一个cli的命令行程序，命令行程序写在 `cmd` 包中，运行时指定subcommand名称，如
```
./app test --id=4
```

## 编译运行

```
make
./app subcommend
```

## 版本信息

```
./app -v
```

## HTTP调试

详见 `app/app.http` 文件