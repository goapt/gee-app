# Gee Application

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