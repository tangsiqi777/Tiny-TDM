#### 开发原因， 官方的推荐的TdengineGUI存在一些问题以及Tiny TDM的解决方法

- 版本兼容性差，很久没更新了，最新的Tdengine已经连不上

- 当数据库表名有大写会无法访问

- 卡顿

- 安装包过大

- 界面风格过于老旧

- [x]  如何解决版本兼容性差的问题

不能采用原生的golang连接，跨版本会导致直接连接不上。采用的rest接口，但其实TdengineGUI也是采用的Rest API的方式为啥还是有版本问题？以及如何解决？

目前看来v2和v3版本rest接口返回的响应码和数据类型有些许差异。

解决方法采用动态加载脚本方式实现（目前做了部分，后期可以通过加载文件js脚本扩展），目前已实现v2和v3两个版本返回数据处理的js脚本。代码在`frontend/version`目录下，在输入连接时会让用户选择版本，后面请求数据会根据选择的版本调用不同版本的数据处理函数。

v2版本返回

> {
>  "status": "succ",
>  "head": [
>  "name",
>  "status"
>  ],
>  "data": [
>  [
>  "log",
>  "ready"
>  ]
>  ],
>  "rows": 1
> }

v3版本返回

> {
>  "status": "succ",
>  "head": [
>  "name",
>  "status"
>  ],
>  "data": [
>  [
>  "log",
>  "ready"
>  ]
>  ],
>  "rows": 1
> }



- [x]  如何当数据库表名有大写会无法访问的问题

当数据库或表名有大写时能显示但是查询时因为没使用````转义会导致查询不到导致报错

解决方法：执行sql时给数据库表名称都进行转义



- [x]  卡顿

Tdengine数据库不同于mysql等关系型数据库存在有：表超多，数据量超大的问题，所以当表过多，查询范围过大会导致渲染卡顿

解决方法：对表显示进行分页，数据显示进行分页

#### 技术栈以及解决上述问题方案

| 技术    | 说明              | Github地址                                                           | 优点                        |
| ----- | --------------- | ------------------------------------------------------------------ | ------------------------- |
| wails | 基于Golang桌面端开发框架 | https://github.com/wailsapp/wailshttps://github.com/wailsapp/wails | wails是相比于electron打包更小，好上手 |
| Vue3  | 核心前端框架          | https://github.com/vuejs/vue                                       | 大家都在用                     |
| arco  | 字节前端UI框架        | https://github.com/arco-design/arco-design-vue                     | 相比element plus易用，风格更好看    |

#### 实现功能

**建立连接**

可以新建连接，查看连接详情和删除（没有更新，暂时没时间做）

![img](https://github.com/tangsiqi777/Tiny-TDM/blob/main/doc/connection.png?raw=true)

**数据库列表**
可以查看数据库信息，暂时没提供删除（可以自行通过sql执行）

![img](https://github.com/tangsiqi777/Tiny-TDM/blob/main/doc/database.png?raw=true)

**超级表列表**
可以查看超级表列表，以及超级表信息，同时考虑到超级表会比较多提供了搜索功能

![img](https://github.com/tangsiqi777/Tiny-TDM/blob/main/doc/superTable.png?raw=true)

**子表列表**

可以查看子表列表，以及子表信息，同时考虑到子表会比较多提供了搜索功能，并进行了分页显示，避免字表过多卡顿
![img](https://github.com/tangsiqi777/Tiny-TDM/blob/main/doc/sqlQuery.png?raw=true)

**自定义Sql查询功能**![img](https://github.com/tangsiqi777/Tiny-TDM/blob/main/doc/sqlQuery.png?raw=true)

点击搜索框右边按钮进入自定义Sql查询功能