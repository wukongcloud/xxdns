## DNS的区域类型

| 类型    | 描述     |
| ------- | -------- |
| Master  | 主服务器 |
| Slave   | 从服务器 |
| Forward | 转发域   |
| Hit     | 根域     |





### 例子

#### master

```
zone "xxdns.org" IN {
      type master;
      file "db.xxdns.org";
      allow-transfer { 192.168.53.54; }; //acl also works
};
```



#### slave

```
zone "xxdns.org" IN {
      type slave;
      file "slaves/db.xxdns.org";
      masters {192.168.53.53; };
};
```



#### forward

```
zone "xxdns.org" IN {
    type forward;
    forwarders { 8.8.8.8; };
    forward only;
};
```



#### hit

```
zone "." IN {
        type hint;
        file "named.ca";
};
```





