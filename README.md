# Jsonix Kit

[English Document](https://github.com/JsonLee12138/jsonix-kit/blob/main/README.en.md)

通过 jsonix 命令行工具生成项目代码, 通过 fiber 开发 web 后端应用, 支持使用 dig 依赖注入。

## 概览

- 支持 `swagger` 生成, 自动上传 `apifox`, 支持创建 `openapi` 服务。
- 支持 `jsonix` 命令行工具, 生成项目代码, 自动迁移数据库, 自动生成 `gorm` 代码([jsonix 文档](https://github.com/JsonLee12138/jsonix/blob/main/README.md))。
- 支持 `fiber` 开发 web 后端应用。
- 支持 `gorm` 数据库 ORM(目前只做了 `mysql` 的配置, 不过仍可以自己通过 `gorm` 做其他数据库的适配)。
- 支持 `redis` 缓存。
- 支持 `i18n` 国际化。
- 支持 `logger` 日志。
- 支持 `cors` 跨域。
- 支持 `jsonix` 热重载。
- 支持使用 `dig` 依赖注入。

### jsonix 命令
```
jsonix migrate 项目根目录运行, 生成自动迁移文件
jsonix gen 项目apps目录运行, 生成模块代码
jsonix server 项目根目录运行, 启动服务
jsonix server -e env 项目根目录运行, 指定运行环境启动服务
jsonix server -w 项目根目录运行, 热重载服务
jsonix server -s prot 查看端口是否被占用
jsonix server -k prot 杀死端口进程
```

### 项目文件结构

```
├─ .vscode                     # VSCode 推荐配置
├─ .run                        # GoLand 运行命令
├─ apps                        # 应用模块
│  ├─ example                  # 示例模块
│  │  ├─ controller           # 控制器
│  │  ├─ repository          # 仓库
│  │  ├─ service             # 服务
│  │  └─ entry               # 入口文件
│  └─ ...                     # 其他模块
├─ auto_migrate               # 自动迁移目录(可通过 jsonix migrate 命令生成, 无需改动)
│  └─ ...                     # 自动迁移文件
├─ config                     # 配置文件目录
│  ├─ config.yaml            # 配置文件
│  ├─ config.dev.yaml        # 开发环境配置
│  ├─ config.test.yaml       # 测试环境配置
│  ├─ config.prod.yaml       # 生产环境配置
│  ├─ config.local.yaml      # 本地环境配置
│  ├─ config.dev.local.yaml  # 开发本地环境配置
│  ├─ config.test.local.yaml # 测试本地环境配置
│  ├─ config.prod.local.yaml # 生产本地环境配置
│  ├─ regexes.yaml           # uaparser 配置文件
│  └─ ...                     # 其他配置文件
├─ configs                    # 配置文件实例目录
│  ├─ configs.yaml           # 全部配置实例
│  └─ ...                     # 其他配置实例文件
├─ core                       # 核心模块目录
├─ docs                       # swagger 文档目录
├─ locales                    # 语言包目录
├─ logs                       # 日志目录
├─ middleware                 # 中间件目录
├─ tmp                        # air 启动临时文件(不要改动, 不要提交)
├─ utils                      # 工具函数目录
├─ .air.toml                  # air 配置文件
├─ main.go                    # 主函数
├─ go.mod                     # Go 模块文件
└─ go.sum                     # Go 模块文件
```
