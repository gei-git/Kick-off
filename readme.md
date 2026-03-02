# Kick-off

[![Go](https://img.shields.io/badge/Go-1.24-00ADD8.svg)](https://go.dev)
[![Gin](https://img.shields.io/badge/Gin-1.11-00ADD8.svg)](https://github.com/gin-gonic/gin)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-336791.svg)](https://www.postgresql.org)
[![Docker](https://img.shields.io/badge/Docker-Compose-2496ED.svg)](https://www.docker.com)

**一个干净、专业、开箱即用的 Go 后端 Todo / 任务管理 API 启动模板**

"Kick-off" 意为“启动”，专为快速搭建现代 Go Web 项目而设计，采用 **Clean Architecture** 分层结构，开箱即用。

---

## ✨ 核心特性

- **用户认证**：注册 + 登录（JWT HS256，24小时有效）
- **任务管理**：每个用户独立的任务列表（多用户隔离）
- **数据库**：PostgreSQL + GORM 自动迁移
- **架构**：Clean Architecture（handler / service / repository）
- **部署**：Docker Compose 一键启动（含 Postgres）
- **安全**：bcrypt 密码加密 + JWT 中间件
- **文档**：已预留 Swaggo/Swagger 支持（一键生成）

---

## 🛠️ 技术栈

- **语言**：Go 1.24
- **框架**：Gin v1.11
- **ORM**：GORM v2 + PostgreSQL 16
- **认证**：golang-jwt/jwt + bcrypt
- **配置**：godotenv
- **部署**：Docker + Docker Compose
- **其他**：Swaggo（API 文档）

---

## 🚀 快速开始（Docker 方式，推荐）

### 1. 克隆项目
```bash
git clone https://github.com/gei-git/Kick-off.git
cd Kick-off