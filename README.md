# groupcache demo

## 部署

三个终端分别启动：

```shell
go run src/server/server.go -addr=:8080 -pool=http://127.0.0.1:8080,http://127.0.0.1:8081,http://127.0.0.1:8082
go run src/server/server.go -addr=:8081 -pool=http://127.0.0.1:8081,http://127.0.0.1:8080,http://127.0.0.1:8082
go run src/server/server.go -addr=:8082 -pool=http://127.0.0.1:8082,http://127.0.0.1:8080,http://127.0.0.1:8081
```

浏览器打开：  `localhost:8080/color?name=red`, `localhost:8081/color?name=red`

日志：

```shell
➜  groupcache-demo git:(master) ✗ go run src/server/server.go -addr=:8080 -pool=http://127.0.0.1:8080,http://127.0.0.1:8081,http://127.0.0.1:8082

2023/01/04 11:31:09 获取 name=red
2023/01/04 11:31:20 获取 name=red


# ================================================================================

➜  groupcache-demo git:(master) ✗ go run src/server/server.go -addr=:8081 -pool=http://127.0.0.1:8081,http://127.0.0.1:8080,http://127.0.0.1:8082

2023/01/04 11:31:09 读DB red
2023/01/04 11:31:43 获取 name=red

# ================================================================================
➜  groupcache-demo git:(master) ✗ go run src/server/server.go -addr=:8082 -pool=http://127.0.0.1:8082,http://127.0.0.1:8080,http://127.0.0.1:8081



```

## ref

* https://gist.github.com/fiorix/816117cfc7573319b72d

