## 平台简介
* 基于GF(Go Frame)的后台管理系统
* 前端采用ruoyi-ui 、Vue、Element UI。
* 后端采用GO语言 框架 GF(Go Frame)。


## 内置功能

1.  用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2.  部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3.  岗位管理：配置系统用户所属担任职务。
4.  菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5.  角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6.  字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7.  参数管理：对系统动态配置常用参数。
8.  操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
9. 登录日志：系统登录日志记录查询包含登录异常。
10. 在线用户：当前系统中活跃用户状态监控。
11. 定时任务：在线（添加、修改、删除)任务调度包含执行结果日志。
12. 代码生成：前后端代码的生成。
13. 服务监控：监视当前系统CPU、内存、磁盘、堆栈等相关信息。
14. 在线构建器：拖动表单元素生成相应的HTML代码。
15. 文件上传,缓存标签等。
16. 正在开发中的功能：cms管理、模型管理、工作流引擎

## 演示地址
[https://dev.tp229.com](https://dev.tp229.com/)
账号：demo  密码：demo#1234

## 配置

项目数据库文件 ./data/db.sql 创建数据库导入后修改配置/config/config.toml

其中jwt配置

[gToken]

    CacheMode = 2 此处若使用了redis配置为2 若没使用redis配置1
    CacheKey = "SMSToken:"
    Timeout = 0
    MaxRefresh = 0
    TokenDelimiter="_"
    EncryptKey = "48ea90e65be432e7391a002a40c0fdcd"
    AuthFailMsg = "登录超时，请重新登录"
    MultiLogin = true

##运行
go run main.go 直接访问 `http://localhost:2975`

账号：demo  密码：demo#1234

项目为前后端分离
前端git地址：


## 演示图



## 感谢(排名不分先后)
> gf框架 [https://github.com/gogf/gf](https://github.com/gogf/gf) 
>
> RuoYi-Vue [https://gitee.com/y_project/RuoYi-Vue](https://gitee.com/y_project/RuoYi-Vue)
>
> swaggo [https://github.com/swaggo/swag](https://github.com/swaggo/swag)
>
> tpflow [https://gitee.com/ntdgg/tpflow](https://gitee.com/ntdgg/tpflow)  [这个是PHP的项目]
>
> gtoken [https://github.com/goflyfox/gtoken](https://github.com/goflyfox/gtoken)
>
> casbin [https://github.com/casbin/casbin](https://github.com/casbin/casbin)
>
> 云捷go [https://gitee.com/yunjieg/yjgo](https://gitee.com/yunjieg/yjgo)


