购物车系统

需求：
1.实现登录，登出（只有最新一次登录有效）
2.实现多用户登录
3.实现用户余额，账单以及商品列表的查询
4.实现充值功能，可给任意用户充值
5.实现下单功能

使用方式:
1.环境准备：
环境已安装mysql和go

2.创建用户并赋权：
mysql -u root
use mysql;
CREATE USER 'admin'@'localhost' IDENTIFIED BY 'admin123';
GRANT ALL PRIVILEGES ON *.*  TO 'admin'@'localhost' IDENTIFIED BY 'admin123'  WITH GRANT OPTION;

3.插入数据:
运行/shop/mysql/sqlFile下所有的sql文件

4.开启服务端
go run runServer.go

5.客户端登录输入密码：
go run runClient.go
输入 admin/admin123

6.查询余额:
按照指引，operate输入check,ID输入用户id，如(oUT385ZLmRr6R_a9xKSfSW9SekYI),type输入balance即可返回对应用户余额

7.查询账单
按照指引，operate输入check,ID输入用户id，如(oUT385ZLmRr6R_a9xKSfSW9SekYI),type输入orderList即可返回对应用户余额

8.查询商品列表
按照指引，operate输入check,ID输入用户id，如(oUT385ZLmRr6R_a9xKSfSW9SekYI),type输入shopList即可返回对应用户余额

9.充值
按照指引，operate输入recharge,ID输入用户id，如(oUT385ZLmRr6R_a9xKSfSW9SekYI),value输入充值数目，如(5000)，即可完成充值

10.下单
按照指引，operate输入check,ID输入用户id，如(oUT385ZLmRr6R_a9xKSfSW9SekYI),shopID输入购买的商品ID，如(9)，即可完成下单

代码目录：
common
    -http.go 自己封装的所有项目使用的网络共用库
shop
    client
        -client:开启客户端
        -clientProcess:客户端数据处理流程
    common
        -common.go:主要记录网络传输过程中的参数结构体
    config
        -config.go:主要用户解析shop.yaml配置文件，里面记录主要端口与host
    dataRecord
        -userInfo:用户账户密码文件
    mysql
        -mysql.go:自己封装的对mysql的相关操作
        sqlFile:保存预先运行的sql文件
    -runClient.go:客户端启动脚本
    -runServer.go:服务端启动脚本
    server
        -server:开启服务端
        -serverProcess:服务端数据处理流程
        -serverProcess_test.go:数据处理单测文件
    shop.yaml:配置文件
