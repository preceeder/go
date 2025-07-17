### nsqlookupd
```shell
# -broadcast-address 广播地址(nsqd 注册的时候要用到这个ip地址) 一般0.0.0.0 就可以
# -http-address 接收nsqd http 注册的 端口
# -tcp-address 接收nsqdtcp注册的端口
#启动： 
./nsqlookupd -broadcast-address=0.0.0.0 -http-address=0.0.0.0:4131 -tcp-address=0.0.0.0:4130 -inactive-producer-timeout=5m0s

```


### nsqd

```shell
# 这里我是在本地运行的所以一般ip都是 127.0.0.1   分布式的情况下  就是 nsqlookupd 的ip地址
# -broadcast-address nsqd注册到nsqlookupd的IP地址  一般本地可以是0.0.0.0
#-broadcast-http-port   nsqd注册到nsqlookupd的 http 端口
#-broadcast-tcp-port nsqd注册到nsqlookupd的 tcp 端口
# -data-path  数据备份到磁盘的路径
#-tcp-address nsqd tcp 链接地址 <host>:<port>
#-http-address nsqd http 链接地址 <host>:<port>     // port 和 -broadcast-http-port  保持一致 这两个参数保留一个就可以
#-https-address nsqd https 链接地址 <host>:<port>   // port 和 -broadcast-tcp-port  保持一致 这两个参数保留一个就可以
#-lookupd-tcp-address  lookupd的tcp 链接地址

./nsqd -broadcast-address=0.0.0.0 -broadcast-http-port=4161 -broadcast-tcp-port=4160 -tcp-address=0.0.0.0:4160  -http-address=0.0.0.0:4161 -https-address=0.0.0.0:4162  -lookupd-tcp-address=127.0.0.1:4130 -data-path=./
```



### admin
```shell
# -http-address <addr>:<port>  // default "0.0.0.0:4171"   访问的地址
# -lookupd-http-address   
./nsqadmin --lookupd-http-address=127.0.0.1:4131

# 本地浏览器打开
http://127.0.0.1:4171
```