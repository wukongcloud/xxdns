## Bind9 配置段



### named.conf

| 配置段       | 描述                                                         |
| ------------ | ------------------------------------------------------------ |
| acl          | 定义访问控制列表                                             |
| controls     | 定义 rndc 命令使用的控制通道，若省略此句，则只允许经过 rndc.key 认证的 127.0.0.1 的 rndc 控制，参考 rndc |
| include      | 将其他文件包含到本配置文件当中                               |
| key          | 定义用于 TSIG 的授权密钥                                     |
| logging      | 定义日志的记录规范                                           |
| lwres        | 将 named 同时配置成一个轻量级的解析器                        |
| options      | 定义全局配置选项                                             |
| trusted-keys | 为服务器定义 DNSSEC 加密密钥                                 |
| server       | 设置每个服务器的特有的选项                                   |
| view         | 定义域名空间的一个视图                                       |
| zone         | 定义一个区声明                                               |



#### 资源记录
标准资源记录的基本格式是：
[name] [ttl] [class] type data

- class 字段: IN、CH 、 HS
- type 字段： [RecordType.md](RecordType.md) 

