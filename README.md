## 2026-1-31

初始化本地仓库，并创建 `.gitignore` 文件；

搭建好了项目架构，并初始化 `Go Modules` ；

在 `./task-management-system/configs/config.yaml` 中设置 `server` 、 `database` 、 `redis` 、 `jwt`  、 `rate_limit` 等配置；

在根目录下创建 `.env` 和 `.env.example` ，分别用于存储私密信息（不上传到 github）和供其他开发者参考配置信息；

创建 Go 语言的配置管理模块， `./task-management-system/internal/config/config.go` ，用于从 `.env` 和 `config.yaml` 加载配置信息并替换环境变量。并安装好需要的包， `go get github.com/joho/godotenv github.com/spf13/viper` ；

进行了一次 `commit` ， ""

## 2026-1-30

输入提示词让 AI 生成的基于 Go 的后端项目，

搞懂不清晰的概念：✅

JWT（JSON Web Token），常用于微服务、RESTful API，像是互联网世界的数字通行证。组成部分，Header (头部)，声明类型和所使用的签名算法；Payload (负载)：存放实际的数据；Signature (签名)：防止数据被篡改的核心。库 `golang-jwt/jwt` 。

Redis（Remote Dictionary Server），常用于处理高并发请求，后端加速器，基于内存的键值型（Key-Value）数据库。用途，缓存层，减轻关系型数据库（如 MySQL）的压力；分布式锁，`SETNX` 指令，在多实例部署的 Go 服务中保证逻辑的原子性；消息队列，`List` 或 `Pub` / `Sub` 甚至 `Stream` 机制实现轻量级消息解耦；计数器/排行榜，`INCR` 和 `ZSet` 实现点赞数或实时排名。库，`redis/go-redis` 。

Worker Pool（工作池），Goroutine 是 Go 并发的“原子”，那么工作池就是管理这些原子的“调度室”，是一种并发设计模式，维护固定数量的 Goroutine 集合来处理一系列任务。为什么需要工作池，虽然 Go 的 Goroutine 很轻量，但在极端高并发下，会造成内存耗尽、调度开销（过多的上下文切换）、资源过载。则成，Jobs Channel（一般带缓冲），存放待处理任务的队列；Workers，一组运行中的 Goroutine，监听 Jobs 管道；Results Channel 用于接收 Workers 处理完后的返回结果。避免重复造轮子，`ants` 开源库。

| 并发设计模式对比      | 核心区别                                  | 适用场景                                        |
| --------------------- | ----------------------------------------- | ----------------------------------------------- |
| Worker Pool           | 固定数量的协程，任务排队                  | 保护下游资源（如数据库连接）、控制 CPU 占用     |
| Goroutine per Request | 按需创建，一个请求/任务一个协程           | 绝大多数普通的 I/O 密集型 Web 服务              |
| Semaphore (信号量)    | 不限制协程总数，但限制同时运行的任务数    | 简单的限流，不需要复用协程，只需要限流          |
| Thread Pool (线程池)  | Java/C++ 的概念，管理昂贵的操作系统线程。 | 解决线程创建成本高的问题（Go 原生已解决此问题） |

Middleware（中间件），是一种代码逻辑，在 Request(请求) 到达业务逻辑之前执行，或者Response（响应）返回给客户端之后执行。用途，身份认证；日志记录；错误处理；跨域处理；限流。在 Go 中，本质是一个接收 `http.Handler` 并返回 `http.Handler` 的函数。库，`net/http` 。

Viper，用途，帮你“读取和管理程序配置”的一站式工具。在结构体加上 `mapstructure` 标签，而不是普通 `json` 标签或 `gorm` 标签。在 Go 中可以对结构体字段写多个标签。在项目中，将从文件中读配置和与数据库对接这两个功能分开，分为两个结构体存储。

Zap，专门为高性能设计的日志组件，替代 Go 标准库 `log` 。

Validator，在处理用户提交的表单、API 请求的 JSON 数据时，先进行校验。库，`go-playground/validator`。通过“扫描”结构体 `validate` 标签来运行，将结构体实例作为参数传入 `Validate.Struct()`  。

Testify，代替Go 标准库 `test` 。