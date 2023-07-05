# qijiangAPI

## 1	环境变量

### 本地环境
| 参数名 | 字段值 |
| ------ | ------ |
|baseUrl|http://localhost:8080|
### 测试环境
| 参数名 | 字段值                       |
| ------ |---------------------------|
|baseUrl| http://192.168.0.162:8080 |


## 2	CPU

> POST: http://localhost:8080/api/CPU/:ip
### 接口说明
> 根据ip地址返回相对应服务器的当前CPU使用率
### 请求体(Request Body)
| 参数名称 | 数据类型 | 默认值 | 不为空 | 描述 |
| ------ | ------ |-----| ------ | ------ |
| pwd|string|     |true|服务器ssh密码|

### 响应体

● 200: OK 响应数据格式：JSON

| 参数名称 | 类型 | 默认值 | 不为空 | 描述 |
| ------ | ------ |-----| ------ | ------ |
| IpAddress:|string|     |true|192.168.0.162|
| Utilization|number|     |true|0.7|

###### **失败返回**
<table border="1" style="border: 1px solid #a9adb1;margin: 0;border-collapse: collapse;border-spacing: 0;"><tr><td colspan="1" rowspan="1">code：-999<br></td><td colspan="2" rowspan="1">msg：连接ssh服务器失败<br></td></tr><tr><td colspan="1" rowspan="1">code：-9<br></td><td colspan="2" rowspan="1">msg：创建ssh会话失败<br></td></tr><tr><td colspan="1" rowspan="1">code：-99<br></td><td colspan="2" rowspan="1">msg：CPU利用率信息获取失败<br></td></tr></table>

## 3	内存

> POST  http://localhost:8080/api/Mem/:ip
### 接口说明
> 根据ip地址返回内存利用率
### 请求体(Request Body)
| 参数名称 | 数据类型 | 默认值 | 不为空 | 描述 |
| ------ | ------ |-----| ------ | ------ |
| pwd|string|     |true|对应服务器的ssh密码|
### 响应体
● 200: OK 响应数据格式：JSON

| 参数名称 | 类型 | 默认值 | 不为空 | 描述 |
| ------ | ------ |-----| ------ | ------ |
| IpAddress:|string|     |true|192.168.0.162|
| Utilization|string|     |true|11%|
失败返回
<table border="1" style="border: 1px solid #a9adb1;margin: 0;border-collapse: collapse;border-spacing: 0;"><tr><td colspan="1" rowspan="1">code：-999<br></td><td colspan="2" rowspan="1">msg：连接ssh服务器失败<br></td></tr><tr><td colspan="1" rowspan="1">code：-9<br></td><td colspan="2" rowspan="1">msg：创建ssh会话失败<br></td></tr><tr><td colspan="1" rowspan="1">code：-99<br></td><td colspan="2" rowspan="1">msg：内存利用率信息获取失败<br></td></tr></table>

## 4	硬盘

> POST  http://localhost:8080/api/Disk/:ip
### 接口说明
> 根据ip地址返回对应服务器的硬盘数量，每一块硬盘的大小，利用率
### 请求体(Request Body)
| 参数名称 | 数据类型 | 默认值 | 不为空 | 描述 |
| ------ | ------ | ------ | ------ | ------ |
| pwd|string||true|对应服务器的ssh密码|
### 响应体
● 200: OK 响应数据格式：JSON

| 参数名称 | 类型 | 默认值 | 不为空 | 描述 |
| ------ | ------ |-----| ------ | ------ |
| IpAddress|string|     |true|192.168.0.162|
| diskMsg|object|     |true||
|⇥ disks|array[object]|     |true||
|⇥⇥ 使用率|string|     |true|85%|
|⇥⇥ 可用空间|string|     |true|8.0G|
|⇥⇥ 已用空间|string|     |true|43G|
|⇥⇥ 挂载地址|string|     |true|/var/lib/docker/overlay2/b4037659753f13fb30a880e1f397b6168762529031b9cdba318b183ca770287f/merged|
|⇥⇥ 文件系统总大小|string|     |true|50G|
失败返回
<table border="1" style="border: 1px solid #a9adb1;margin: 0;border-collapse: collapse;border-spacing: 0;"><tr><td colspan="1" rowspan="1">code：-999<br></td><td colspan="2" rowspan="1">msg：连接ssh服务器失败<br></td></tr><tr><td colspan="1" rowspan="1">code：-9<br></td><td colspan="2" rowspan="1">msg：创建ssh会话失败<br></td></tr><tr><td colspan="1" rowspan="1">code：-99<br></td><td colspan="2" rowspan="1">msg：硬盘利用率信息获取失败<br></td></tr></table>

## 5	端口

> POST  http://localhost:8080/api/ip/:ip/:port
### 接口说明
> 根据ip和端口号查询该端口是否正常工作
### 请求体(Request Body)
| 参数名称 | 数据类型 | 默认值 | 不为空 | 描述 |
| ------ | ------ |-----| ------ | ------ |
| pwd|string|     |true|对应服务器的ssh密码|
### 响应体
● 200: OK 响应数据格式：JSON

| 参数名称 | 类型 | 默认值 | 不为空 | 描述 |
| ------ | ------ |-----| ------ | ------ |
| IpAddress|string|     |true|192.168.0.162|
| PortStatus|string|     |true|Listening|
| port|string|     |true|22|

## 部署
### 环境依赖
go version >= 1.19.6
### 访问方式
项目部署在192.168.0.162服务器
1. 将 http://localhost:8080替换为 http://192.168.0.162:8080
2. 后面拼接欲访问的接口，并将待查服务器的ip附在最后,如：
 > http://192.168.0.162:8080/api/Mem/192.168.15.16
> 
3. 打开Postman，向该链接附带body中的参数发送请求。
![img.png](img.png)
