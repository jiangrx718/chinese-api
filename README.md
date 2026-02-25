crm 项目，使用go语言编写，基于gin的web框架封装基础的相关信息，开箱即用

## 说明
```
# 后端项目地址
$ git clone https://github.com/jiangrx718/chinese-api.git

# 前端项目地址
$ git clone https://github.com/jiangrx718/chinese-page.git
```

## 1.正式项目目录运行时结构

```
crm/
├── commands/
│   └── agenerate # 生成SQL ORM目录
│   └── migrate   # 数据库建表 目录
│   └── server    # web服务运行 目录
├── config/       # 配置文件 目录
├── gopkg/        # 核心基础依赖 目录
├── handler/      # 路由API 目录
├── internal/     # 业务逻辑处理以及数据表 目录
├── README.md     # README 文件
└── main.go       # 入口文件
```

## 2.快速使用
在使用前需要修改module的名称，先查看对应的名称
```
# 查看当前模块名称：
go list -m

# 修改模块名称
go mod edit -module web1
```
## 3.Commands
```shell
生成SQL ORM
go run main.go generate

数据库建表，并初始化数据
go run main.go migrate up

生成API文档
go run main.go swag init
```

## 4.后端启动命令
```shell
# 进入到crm目录执行以下命令

fresh
或者
go run main.go 

```

## 5.项目部署
```shell
# 打镜像

docker buildx build \
  --platform linux/amd64 \
  -t chinese-api:1.0.0-amd64 \
  --output type=docker \
  .

# 保存镜像

docker save -o /Users/jiang/jiangrx816/docker-images/chinese-api-1.0.0-amd64.tar chinese-api:1.0.0-amd64

# 上传镜像：

scp /Users/jiang/jiangrx816/docker-images/chinese-api-1.0.0-amd64.tar root@xxx.xxx.xxx.xxx:/data/project/chinese

```