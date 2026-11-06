# be-2027-hw3-template：Gin 分层架构课后作业

本项目为西安交通大学挑战网社团技术部后端第三次培训课后作业模板。核心任务为基于 Controller 与 Service 分层架构规范实现用户注册业务，掌握结构体 DTO 声明式绑定、业务逻辑解耦与 RESTful 状态码映射。

---

## 一、作业内容（必做项）

在本项目骨架中遵循分层设计规范，实现用户注册业务接口：

1. **接口路径**：`POST /api/v1/users/register`
2. **请求格式**：`application/json`
3. **参数规范**：
   - `username`（string，必填）：长度 4 ~ 16 位
   - `password`（string，必填）：长度 6 ~ 32 位
   - `email`（string，必填）：合规邮箱格式
   - `age`（int，选填）：取值范围 1 ~ 120
4. **状态码映射规范**：
   - **注册成功**：返回 **`200 OK`**；
   - **参数缺失或规则不符**：返回 **`400 Bad Request`** 与具体校验错误信息；
   - **用户名已被占用（重复注册）**：返回 **`409 Conflict`** 与冲突提示。
5. **架构约束**：
   - 依赖流向单向流动：`Controller -> Service`；
   - **Service 层严禁引入 `*gin.Context`**（保持纯 Go 代码，面向单元测试与多协议复用）。

---

## 二、选做加分项（选做）

本项目提供两个选做中间件骨架，学员可按需实现其中一项或全部。未实现选做项不影响必做项的满分与 CI 机检通过。

### 1. 耗时统计中间件（洋葱模型）

- **文件位置**：`middleware/cost_time.go`
- **实现要求**：在 `c.Next()` 前后计算接口处理耗时，并通过 `log.Printf("[耗时统计] %s | %v", c.Request.URL.Path, time.Since(start))` 输出耗时审计日志；
- **挂载方式**：在 `main.go` 中通过 `r.Use(middleware.CostTimeMiddleware())` 全局挂载；
- **CI 验证**：GitHub Actions 将自动探测服务端日志中的耗时统计输出并记录。

### 2. Token 鉴权中间件（调用链拦截）

- **文件位置**：`middleware/auth.go`
- **实现要求**：提取请求头 `Authorization`。若凭证缺失或非法，返回 **`401 Unauthorized`** 并**显式调用 `c.Abort()`** 终止下游 Handler 执行；
- **挂载方式**：在 `main.go` 中挂载于受保护路由组：`v1.Group("/protected").Use(middleware.AuthMiddleware())`。

---

## 三、本地运行与调试

```bash
# 1. 整理与下载依赖
go mod tidy

# 2. 启动服务（默认监听 8080 端口）
go run main.go
```

---

## 四、工程实现指引

> **说明**：本次作业仅需修改 `controller/`、`service/` 与 `middleware/`（选做），其余配置请勿改动。

建议按以下顺序完善代码：

1. **`service/user_service.go`**：
   - 实现 `Register(...) error` 方法，使用 `sync.Map` 检查用户名是否已存在；
   - 用户名重复返回原生 `error`，不存在则写入存储并返回 `nil`。
2. **`controller/user_controller.go`**：
   - 完善 `RegisterRequest` 结构体字段及其 `json` 与 `binding` 标签；
   - 在 `Register(c *gin.Context)` 中调用 `c.ShouldBindJSON` 进行参数校验；
   - 调用下游 Service 业务方法，将返回的错误分别映射为 `400`、`409` 或 `200` 响应。
3. **`middleware/`（选做）**：
   - 完善选做中间件逻辑，并在 `main.go` 中解除对应挂载注释。

---

## 五、提交与验收

1. **仓库命名**：通过模板建仓（建仓方式与往常一致），建议统一命名为 `be-2027-hw3-姓名拼音`（例如 `be-2027-hw3-zhangsan`）；
2. **CI 验证**：代码 push 后 GitHub Actions 自动化测试通过（绿勾 ✅）；
3. **抓包截图**：将 Apifox 截图存放在 `docs/images/` 并更新至第六节。

---

## 六、抓包截图展示

截图存放规则：

1. 截图文件存放于本工程 `docs/images/` 目录；
2. 必须使用 Markdown 相对路径引用，不得使用操作系统本地绝对路径；
3. 截图内容应完整展示请求路径、入参、响应状态码与响应体。

### 1. 注册成功（HTTP 200）

*对应图片路径：`docs/images/register_200.png`*

![注册成功截图](docs/images/register_200.png)

### 2. 用户名重复冲突（HTTP 409）

*对应图片路径：`docs/images/register_409.png`*

![重复注册冲突截图](docs/images/register_409.png)

### 3. 参数校验失败（HTTP 400）

*对应图片路径：`docs/images/register_400.png`*

![参数校验失败截图](docs/images/register_400.png)

### 4. 选做中间件抓包（选做）

*若实现了选做中间件，对应图片路径：`docs/images/bonus_middleware.png`*

![选做中间件截图](docs/images/bonus_middleware.png)
