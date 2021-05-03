# Integration Server

## 文件结构
> main.go 主程序
> 
> Dockerfile Docker构建
> 
> config.json 配置文件
> 
> default.jpg 用户默认头像
> 
> > utils 通用接口
> 
> > router 路由组
> 
> > models 数据库模型
> 
> > integration_handlers 中间件及数据库处理程序
> 
> > database 数据库连接和初始化
> 
> > controllers 上下文控制器
> 
> > config 配置加载器


## Config

在根目录下`config.json`配置信息, 包含以下信息:

- 服务端口(从容器出方向)
- DB用户名
- DB密码
- DB IP
- DB 端口
- DB名

