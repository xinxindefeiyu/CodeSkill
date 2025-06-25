# CodeSkill - 编程技巧收集项目

## 项目简介

CodeSkill 是一个用于记录和展示实用编程技巧的Go项目，使用Gin框架构建REST API。

## 项目结构

```
CodeSkill/
├── main.go                     # 程序入口
├── go.mod                      # Go模块文件
├── internal/                   # 内部包
│   ├── handlers/              # HTTP处理器
│   │   ├── handlers.go        # 基础处理器
│   │   └── user_handler.go    # 用户相关处理器
│   ├── models/                # 数据模型
│   │   └── user.go           # 用户模型
│   ├── services/             # 业务逻辑层
│   │   └── user_service.go   # 用户服务
│   └── router/               # 路由配置
│       └── router.go         # 路由设置
└── README.md                 # 项目说明
```

## 技巧展示

### 1. 接口参数类型升级技巧

**问题场景**: 接口最初接收类型A参数，后续升级需要接收类型B参数，同时保持向后兼容。

**解决方案**: 定义通用接口，让新旧类型都实现该接口的方法。

**实现位置**: 
- 接口定义: `internal/models/user.go` - `UserRequest`接口
- 旧版本实现: `BasicUserRequest`结构体
- 新版本实现: `DetailedUserRequest`结构体
- 业务逻辑: `internal/services/user_service.go` - `CreateUser`方法

**核心思想**:
1. 定义通用接口，包含获取参数的方法
2. 旧类型实现接口，新增字段返回默认值
3. 新类型实现接口，提供完整信息
4. 业务逻辑统一使用接口方法处理

## 快速开始

1. 安装依赖
```bash
go mod tidy
```

2. 启动服务
```bash
go run main.go
```

3. 测试接口

健康检查:
```bash
curl http://localhost:8080/health
```

创建用户(旧版本格式):
```bash
curl -X POST http://localhost:8080/api/v1/user/create \
  -H "Content-Type: application/json" \
  -d '{"name":"张三","email":"zhangsan@example.com","age":25}'
```

创建用户(新版本格式):
```bash
curl -X POST http://localhost:8080/api/v1/user/create \
  -H "Content-Type: application/json" \
  -d '{"name":"李四","email":"lisi@example.com","age":30,"phone":"13800138000","address":"北京市朝阳区","company":"某某公司"}'
```

获取技巧列表:
```bash
curl http://localhost:8080/api/v1/skills
```

## 技术栈

- **框架**: Gin (Go Web框架)
- **语言**: Go 1.21+
- **架构**: 分层架构 (Handler -> Service -> Model)

## 扩展计划

后续将继续添加更多实用的编程技巧示例，包括但不限于:
- 设计模式应用
- 性能优化技巧
- 错误处理最佳实践
- 并发编程模式
- 等等...