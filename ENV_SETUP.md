# 环境变量配置说明

## 文件说明

### `.env` 文件
包含实际的敏感信息（密码、密钥等），**不会被提交到Git仓库**。

### `.env.example` 文件
环境变量的模板文件，**会被提交到Git仓库**，供其他开发者参考。

### `config.yaml` 文件
使用环境变量占位符（`${VARIABLE_NAME}`），实际值从`.env`文件或系统环境变量中读取。

---

## 配置步骤

### 1. 复制环境变量模板

```bash
cp .env.example .env
```

### 2. 编辑 `.env` 文件

根据你的实际情况修改以下配置：

```bash
# 服务器配置
SERVER_PORT=8080
SERVER_MODE=debug

# 数据库配置
DB_HOST=localhost
DB_PORT=7809
DB_USER=root
DB_PASSWORD=your_mysql_password_here
DB_NAME=taskdb
DB_CHARSET=utf8mb4

# Redis配置
REDIS_HOST=localhost
REDIS_PORT=7810
REDIS_PASSWORD=
REDIS_DB=0

# JWT配置
# 使用命令生成: openssl rand -base64 32
JWT_SECRET=your_jwt_secret_here
JWT_EXPIRE=24h

# 限流配置
RATE_LIMIT_REQUESTS_PER_MINUTE=60
```

### 3. 生成JWT Secret

```bash
openssl rand -base64 32
```

将生成的值填入 `.env` 文件的 `JWT_SECRET` 字段。

---

## 安全注意事项

### ✅ 应该做的

- 将 `.env` 添加到 `.gitignore` 文件
- 使用强密码和复杂的JWT secret
- 定期更换密钥
- 在生产环境使用环境变量而非文件
- 为不同环境（开发、测试、生产）使用不同的配置

### ❌ 不应该做的

- 将 `.env` 文件提交到Git仓库
- 在代码中硬编码密码或密钥
- 使用简单的密码（如 "123456"）
- 在多个环境使用相同的密钥

---

## 生产环境部署

在生产环境中，推荐使用系统环境变量而非 `.env` 文件：

```bash
# Docker Compose
environment:
  - DB_PASSWORD=${DB_PASSWORD}
  - JWT_SECRET=${JWT_SECRET}

# Kubernetes
env:
  - name: DB_PASSWORD
    valueFrom:
      secretKeyRef:
        name: db-secret
        key: password

# 直接设置环境变量
export DB_PASSWORD=your_password
export JWT_SECRET=your_secret
./server
```

---

## 配置优先级

配置加载的优先级（从高到低）：

1. 系统环境变量
2. `.env` 文件
3. `config.yaml` 文件中的默认值

---

## 常见问题

### Q: 为什么配置文件中的密码是 `${DB_PASSWORD}` 而不是实际密码？

A: 这是环境变量占位符。程序启动时会从 `.env` 文件或系统环境变量中读取实际值。这样可以避免将敏感信息提交到代码仓库。

### Q: `.env.example` 文件有什么用？

A: 它是环境变量的模板，告诉其他开发者需要配置哪些环境变量。每个开发者应该复制它并创建自己的 `.env` 文件。

### Q: 如何在Docker中使用环境变量？

A: 在 `docker-compose.yml` 中使用 `env_file` 或 `environment` 字段：

```yaml
services:
  app:
    env_file:
      - .env
    # 或者
    environment:
      - DB_PASSWORD=${DB_PASSWORD}
```

### Q: 忘记了数据库密码怎么办？

A: 检查 `.env` 文件，或者查看Docker容器的环境变量：

```bash
docker inspect mysql-taskdb | grep -A 10 Env
```
