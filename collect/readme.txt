系统负载监控系统

Agent端：

1）收集系统指标
①主机信息：主机名、操作系统信息
②网络信息：IP地址列表、网络读写字节/包个数
③CPU信息：CPU型号、CPU逻辑核数、CPU物理核心、CPU使用率、
④内存使用情况：总内存大小、实际内存大小、交换区大小
⑤硬盘使用情况：总硬盘大小、已使用硬盘大小
⑥参考库： github.com/shirou/gopsutil

定时上报
①接口鉴权：将固定AppCode放在Header中，请求Header中添加的Authorization字段；配置Authorization字段的值为“APPCODE ＋ 半角空格 ＋APPCODE值”。例如 Authorization:APPCODE AppCode值
②每分钟上报一次系统指标结构化数据

Server端
（1）支持Agent上报数据接口
①接口鉴权：校验http header中的AppCode值是否一致
②上报数据持久化存储

（2）支持多用户登录
①用户信息需要做持久化，用户密码不能明文存放
②用户账户可以提前创建好，支持admin和guest
③除登录接口外，其它接口需要鉴权

（3）提供查询上报数据接口
①查询维度：主机名；操作系统类型 Linux、Windows、Mac；上报时间段


使用方法：

1.go run storeServer.go 开启服务端
2.go run login.go 执行登录功能(不执行此脚本，执行其他的，会提示先登录，服务端没有用户的信息) 账户 admin 密码 admin123
3.go run collectClient.go 开启采集与定时上报功能，提供持久化存储至dataRecord目录下
4.go run getData.go -hostName xxx -osName xxx -startTime xxx -endTime xxx 其中参数为过滤条件

自我测试，可成功完成可持续化存储，并调用getData.go成功获取过滤的信息


代码目录介绍
common 多个项目公用的公共库
    - http.go 创建http连接的，相关公函封装
collect 监控项目主目录
    - collectClient.go 定时采集与上报客户端
    collectMethod 采集方法目录
        - collectMethod.go 采集方法
        - collectMethod_test.go 单测脚本
    common collect项目公共使用目录
        - common.go 主要记录接口鉴权与检验登录的方法
    dataRecord 可持续化数据记录路径
        - collectData 监控数据记录文件
        - userInfo 账号密码记录文件
    - getData.go 供外部调用的，查询数据的脚本
    - login.go 客户端登录脚本
    - storeServer.go 用于开启服务端监听的脚本