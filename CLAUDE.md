# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 构建与测试

```bash
# 构建
go build ./...

# 格式化
gofmt -l .
gofmt -w <文件路径>

# 静态检查
go vet ./...

# 运行全部测试
go test ./...

# 运行单个包的测试
go test github.com/haolie/goUtil/logUtil

# 运行单个测试函数
go test github.com/haolie/goUtil/logUtil -run TestLog
```

## 架构

该仓库是一个 Go 工具库，供其他项目（如 SparkEven）引用，模块路径为 `github.com/haolie/goUtil`。

### 包结构

| 包 | 说明 |
|---|---|
| `logUtil` | 基于 zap 的日志封装，支持文件轮转与控制台双输出 |
| `idUtil` | 基于原子操作的自增 ID 工厂 |

### logUtil

- `InitLog()` / `GetLogger()` 初始化全局 `*zap.Logger`，同时写入文件和控制台
- 日志文件默认输出至 `./log/`，按小时轮转，格式 `%Y-%m-%d_%H.log`，保留 30 天
- 对外 API：`InfoLog`、`ErrLog`、`DebugLog`、`FailLog`
- 文件日志级别：`InfoLevel`；控制台日志级别：`DebugLevel`

### idUtil

- `NumIdFactory` 使用 `sync/atomic` 实现线程安全的自增 ID
- 包 `init()` 自动创建默认工厂（seed=0, add=1）
- `InitDefaultNumFactory(seed, add int64)` 可重置默认工厂
- `GetDefault().CreateId()` 获取下一个 ID
