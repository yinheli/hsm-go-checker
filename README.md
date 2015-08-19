# SJL22 加密机健康状态检查

应运维部要求写的一个小工具, 用于检查 `SJL22` 型号加密机的健康状态.

## 工作原理

> 向加密机(HSM) 发送 `NC` 指令, 并检测收到的响应
> 当网络不可达, 网络IO错误, 加密机宕机等, print `error` 以及相关错误原因, 否则 print `OK`

## 编译

> 需要先安装 `golang` 环境, 程序没有三方依赖, 直接编译即可

```
go build hsm_checker.go
```

## 查看帮助

```
./hsm_checker --help
```

```
Usage of ./hsm_checker:
  -h="192.168.1.200": hsm host
  -p=8000: hsm port
```

## 执行
```
./hsm_checker
```

> 不带参数, 默认主机: `192.168.1.200`, 端口: `8000`

