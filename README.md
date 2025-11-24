# Robot Manage 小娇智能运维平台

企业级智能运维后台管理系统，聚合 Dashboard、信息管理、智能交付、数字员工与系统管理五大模块，并扩展数字人交互能力。

## 技术栈

- **后端**：Go 1.22、Gin、GORM、Redis、Asynq、JWT、Zap、Viper、Swagger（可扩展）
- **前端**：Vue 3 + TypeScript + Vite、Pinia、Vue Router 4、Element Plus、ECharts、Three.js
- **数据层**：MySQL 8.0（业务库）、Redis、Chroma/Weaviate（向量库）

## 项目结构

```
robot-manage
├─ backend
│  ├─ cmd/server            # 入口
│  ├─ configs               # 多环境配置
│  ├─ internal
│  │  ├─ bootstrap          # 依赖注入
│  │  ├─ config             # 配置模型
│  │  ├─ handler/service/repository  # 分层实现
│  │  ├─ infra              # DB/Redis/Queue/Vector
│  │  ├─ logger             # Zap封装
│  │  └─ server             # Gin路由、中间件
│  └─ go.mod
└─ frontend
   ├─ src
   │  ├─ api / stores / types
   │  ├─ components/layouts/views
   │  └─ assets & composables
   └─ vite.config.ts
```

## 快速开始

### 方式一：Docker 容器化部署（推荐）

```bash
# 一键部署
./deploy.sh

# 访问应用
# 前端: http://localhost
# 后端: http://localhost:8088
```

详细说明请查看 [Docker 部署文档](./DOCKER_DEPLOY.md)

### 方式二：本地开发

```bash
# 后端
cd backend
cp configs/config.yaml configs/config.local.yaml
go mod tidy
ROBOT_MANAGE_CONFIG=configs/config.yaml go run ./cmd/server

# 前端
cd frontend
npm install
npm run dev
```

Vite dev server 已配置 `http://localhost:8088` 代理，生产环境通过 Nginx 转发 `/api`.

## API 约定

- 所有接口前缀 `/api/v1`
- JWT 认证（local 环境可跳过），请求头：`Authorization: Bearer <token>`
- 统一响应：

```json
{ "code": 0, "message": "success", "data": {} }
```

### 已实现

| 模块 | 方法 | 路径 | 说明 |
| --- | --- | --- | --- |
| Dashboard | GET | /dashboard/summary | 数据看板指标 |
| Dashboard | GET | /dashboard/knowledge | 知识库动态 |
| Dashboard | GET | /dashboard/timeline | 值班/投产 |
| Dashboard | GET | /dashboard/assistant-report?range=daily | 助手总结 |

其余模块（信息管理、智能交付、数字员工、系统管理）已经在路由层完成分组设计，并返回占位结构，方便后续对接。

## Dashboard 前端

- `MetricsPanel`：ECharts 漏斗/柱状图、巡检/告警趋势
- `KnowledgePanel`：知识库动态 & 热门文章
- `TimelinePanel`：值班、投产、事件流
- `AssistantReport`：日报/周报/月报切换
- `DigitalHuman`：Three.js 3D 模型 + Web Speech API

## 后续规划

- 对接真实 MySQL 表结构（CMDB、ITSM、巡检、知识库等）
- 落地 Redis 缓存与 Asynq 异步任务、工作流编排
- 集成 Swagger/OpenAPI 文档与单元测试
- 数字员工联动向量数据库，实现知识召回与会话增强

欢迎根据业务需要扩展模块、对接外部系统。Pull Request & Issue 均可。