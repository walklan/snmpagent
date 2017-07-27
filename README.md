采集机http代理
====

采集机http代理服务端, 接收http的get或post请求, 实现snmp采集和批量ping测试, 返回json格式数据

启动方式
----

```
./service-start.sh
```

使用方法
----

* snmp采集get方式
```
http://127.0.0.1:1216/snmpagent?seq=1111&ip=127.0.0.1&version=v2c&community=public&oids=get:.1.3.6.1.2.1.1.2.0!table:.1.3.6.1.2.1.31.1.1.1.1,.1.3.6.1.2.1.31.1.1.1.10
```

* ping测试get方式

支持同时ping100个地址

```
http://127.0.0.1:1216/pingagent?seq=1111&ip=192.168.1.1,192.168.1.2,192.168.1.3
```
