# 项目主要功能介绍
1. 用户注册登录功能
2. 对备忘录进行增删改查
3. 存储每条备忘录的浏览次数
4. 分页功能

## 项目结构介绍
1. models：存放数据库模型
2. api：存放api接口
3. cache: 存放redis一些缓存
4. config：放入一些配置文件。里面包含了MySQL数据配的配置（用户名、用户密码、端口号、数据库名称、IP地址等等），Redis数据库配置和运行时的服务地址
5. middleware：存放中间件
6. utils：存放一些工具
7. routers：存放路由
8. serializer：数据转换成json格式
9. service：存放业务逻辑

## 开发步骤
1. 思考服务器、MySQL、Redis等的配置信息，用config.ini文件存储，然后在config.go文件中将config.ini中的配置信息读取出来并构建数据库连接的path
2. 在models中进行数据库连接
3. 在models建立表格(每一个表的创建都单独用一个文件)；然后再在migrate.go文件中进行创建表操作（所有表的创建），最后在连接数据库的文件最后调用创建表函数即可进行表创建
4. 建完表之后写路由。写路由时建议一起把对应的接口函数和业务处理函数一起写了。例如，对于用户这种资源，有Login和Register两种接口，Register的业务逻辑在于，首先判断用户名是否存在，不存在则存入密码然后创建用户成功；
5. 写完其中一个业务逻辑的路由和处理逻辑后，可以在apifox等接口软件中测试
6. 登录功能的业务逻辑在于首先要判断用户名是否存在，判断存在之后要给改成登录颁发一次token以验证后续操作是谁，这里就包括token生成和验证以及对应的序列化器、结构体的定义