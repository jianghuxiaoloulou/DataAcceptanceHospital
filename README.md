# 项目
# ****PACS集成平台医院数据上传****

# 项目描述
* 模拟第三方PACS上传申请单数据
# 对象管理

# 公共组件
配置管理
数据库连接

Oracle数据库生产环境需要配置客户端程序Tool/instantcloent_21_9

日志写入

# 文件配置文件读取：go get -u github.com/spf13/viper
Viper 是适用于GO 应用程序的完整配置解决方案

# 日志：go get -u gopkg.in/natefinch/lumberjack.v2
它的核心功能是将日志写入滚动文件中，该库支持设置所允许单日志文件的最大占用空间、最大生存周期、允许保留的最多旧文件数，
如果出现超出设置项的情况，就会对日志文件进行滚动处理。

# 生成接口文档
Swagger 相关的工具集会根据 OpenAPI 规范去生成各式各类的与接口相关联的内容，
常见的流程是编写注解 =》调用生成库-》生成标准描述文件 =》生成/导入到对应的 Swagger 工具
$ go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
$ go get -u github.com/swaggo/gin-swagger@v1.2.0 
$ go get -u github.com/swaggo/files
$ go get -u github.com/alecthomas/template

@Summary	摘要
@Produce	API 可以产生的 MIME 类型的列表，MIME 类型你可以简单的理解为响应类型，例如：json、xml、html 等等
@Param	参数格式，从左到右分别为：参数名、入参类型、数据类型、是否必填、注释
@Success	响应成功，从左到右分别为：状态码、参数类型、数据类型、注释
@Failure	响应失败，从左到右分别为：状态码、参数类型、数据类型、注释
@Router	路由，从左到右分别为：路由地址，HTTP 方法

swag init

http://127.0.0.1:8000/swagger/index.html



## 第二次相同项目提交文件到github
# git add README.md
# git commit -m "first commit"
# git push -u origin master

# 代码编译命令
go build -ldflags="-H windowsgui" -o .\DataAcceptanceHospital\DataAcceptanceHospital.exe .\main.go


# 修改记录
* 2023/07/12 创建项目，该项目用于模拟第三方PACS上传申请单数据接口