# Travel-Assistant 智能旅游助手

> **基于 CloudWeGo (Hertz + Kitex) + 豆包大模型的企业级微服务项目**
> 参考 [bw-cli](https://github.com/BwCloudWeGo/bw-cli) DDD 脚手架规范构建

## 📌 项目特性

### ✅ 企业级工程化（已实现）

| 特性 | 实现方式 | 文件位置 |
|------|---------|---------|
| **Viper 配置管理** | YAML + 环境变量覆盖 + 多环境支持 | `pkg/config/config.go` |
| **Zap 日志系统** | 结构化日志 + Lumberjack 旋转 + **7天自动清理** ⭐ | `pkg/logger/logger.go` |
| **中间件统一封装** | Recovery/CORS/Logger/Auth/RequestID | `pkg/middleware/middleware.go` |
| **文件上传工具类** | Local/MinIO/OSS/COS 多后端支持 | `pkg/filex/uploader.go` |
| **参数校验器** | validator.v10 + 自定义规则（destination, travel_date等） | `pkg/validator/validator.go` |
| **统一错误码** | BizError 结构体 + HTTP状态映射 | `pkg/errors/errors.go` |
| **HTTP 响应封装** | 标准JSON格式 + 分页支持 | `pkg/httpx/response.go` |

## 🚀 快速开始

```bash
make tidy        # 整理依赖
make run-api     # 启动 API 网关
make run-travel  # 启动旅行服务
make run-agent   # 启动 Agent 服务
```

## 🎯 核心功能模块

- **配置管理**: Viper + YAML 多环境切换
- **日志系统**: Zap + Lumberjack 7天自动清理
- **中间件系统**: Recovery/CORS/Auth/RateLimiter 等
- **AI Agent 工作流**: 8步编排 (Parse → Generate → Validate → Respond)
- **文件上传**: Local/MinIO/OSS/COS 多后端

## 📄 License

MIT License