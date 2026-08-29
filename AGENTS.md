# AGENTS.md — tksdk 开发指南

本文件是 tksdk 项目的开发工作流程与接口开发规范，供所有开发者与 AI 编码代理遵循。
（本文档取代原 WORKFLOW.md）

---

## 项目简介

tksdk 是多平台电商 CPS 聚合 SDK（Go 1.16），目前覆盖京东（jdopensdk）、淘宝（tbopensdk）、
拼多多（pddopensdk）、苏宁（snopensdk）、唯品会（vipopensdk）、阿里本地生活（alscopensdk）等平台。

每个平台的 SDK 由三部分组成：
- `request/`：请求封装（GetApiName / CheckParameters / AddParameter）
- `response/<接口名>/`：响应解析（WrapResult 二次解析 queryResult）
- `generate/`：代码生成器配置（接口条目 + 模板）

---

## 标准开发流程

### 1. 领取任务
- 从 GitHub Issues 中选择任务：`gh issue view <issue_id> --repo mimicode/tksdk`
- 确认任务编号（Issue ID）与任务范围

### 2. 创建任务分支
```bash
# 从 main 分支创建新分支
git checkout main
git pull origin main
git checkout -b task_<issue_id>
```

### 3. 查阅官方接口文档（接口任务必做）
- **必须通过浏览器打开对应平台的官方接口文档页面**（如京东联盟：
  `https://union.jd.com/openplatform/api/v2?apiName=<接口名>`）
- **禁止通过网络搜索获取接口定义**，一切以官方文档页面为准
- 需要完整读取：接口描述、业务参数表、返回结果字段表、请求示例、响应示例

### 4. 开发接口
按下方「接口开发规范」实现 request / response / 生成器配置 / README / 测试。

### 5. 本地验证
```bash
go build ./...
go vet ./...
go test ./...
```

### 6. 提交与推送
```bash
git add .
git commit -m "feat: 简洁描述修改内容"
git push origin task_<issue_id>
```

### 7. 创建 Pull Request
```bash
gh pr create --repo mimicode/tksdk --base main --head task_<issue_id> \
  --title "<type>: <描述>" --body "..."
```
- PR body 中使用 `Closes #<issue_id>` / `Fixes #<issue_id>` 关联任务
- 说明关键建模决策与文档不一致之处的取舍

### 8. 代码审查与合并
- 等待代码审查，根据反馈修改代码
- 审查通过后合并到 main

### 9. 清理分支
```bash
git branch -d task_<issue_id>
git push origin --delete task_<issue_id>
```

---

## 接口开发规范 ⚠️ 核心规范，必须遵循

### 参数键名
- `360buy_param_json` 内层业务参数的**键名必须与官方文档「业务参数」表格「名称」列完全一致，
  区分大小写**（如 `goodsReq`、`RankGoodsReq`、`couponReq`）
- 不得按其他接口的惯例推测键名；当文档表格名称列与请求示例 setter 命名不一致时，
  **以业务参数表格为准**

### 目录与命名
- 请求文件：`jdopensdk/request/<接口名去点小写>.go`，类型名 = 接口名按 `.` 分段后首字母大写拼接
  + `Request`（如 `jd.union.open.goods.rank.query` → `JdUnionOpenGoodsRankQueryRequest`，
  与 `utils.StrFirstToUpper(apiName, ".")` 一致）
- 响应包：`jdopensdk/response/<接口名去点小写>/`，包内主文件与目录同名
- 新增接口必须同步更新 `jdopensdk/generate/main.go` 条目与 `jdopensdk/README.MD` api 列表
- 客户端测试按平台追加在 `jdopensdk/client_test.go`（其他平台同理）

### 响应结构模式
- 顶层 `Response` 内嵌 `response.TopResponse`，业务字段 JSON tag 为 `<接口名点转下划线>_responce`
- 外层 `queryResult` 是字符串，需在 `WrapResult` 中二次 `json.Unmarshal` 到 `QueryResult`
- `WrapResult` 统一填充 `Body`（原始报文）与 `ErrorResponse`（code/message/requestId）

### 字段类型映射
- 文档 `Number` → `int64`（整数语义）或 `float64`（金额/比率），`String` → `string`，
  `Xxx[]` → 切片，`String[]` → `[]string`
- 文档「响应示例」中数字被引号包裹、`data` 展示为单对象包裹（`{"xxxResp":{...}}`）属于文档
  渲染行为，**数据形状以「返回结果」字段表的类型声明为准**（数组即建模为切片），
  并参照仓库同类已上线接口的实现
- 字段注释保留文档中的中文含义与枚举说明

### 测试规范
- **联调测试**：在 `client_test.go` 按既有模式编写真实调用测试（凭证读取 `dev_env.json`），
  用于验证参数键名与响应解析；注意平台接口有每日调用配额，配额耗尽时测试只会打印
  `error_response`，需择日重跑确认
- **离线解析测试**：按官方文档响应字段形状构造 JSON，验证 `WrapResult` 字段映射
  （无需网络与凭证，可即时运行）
- 修复响应字段时，同步补充/更新离线解析测试

### 联调验证要求
- 新增/修改接口，必须真实调用一次并核对响应解析无字段丢失（配额受限时在 PR 中注明待验证项）

---

## 分支规范 ⚠️ 重要！必须遵循

- **任务分支**：`task_<issue_id>` - 必须使用此格式，不要使用 feat/、feature/ 等其他前缀
- **Bug 分支**：`bug_<issue_id>` - 必须使用此格式
- 所有分支从 `main` 分叉，完成后合并回 `main`

**示例：**
- 对应 Issue #6 的任务分支：task_6（不是 feat/tksdk-api-update）
- 对应 Issue #123 的任务分支：task_123
- 对应 Issue #45 的 Bug 修复分支：bug_45

---

## 提交信息规范

### 提交类型
- `feat:` 新功能
- `fix:` Bug 修复
- `docs:` 文档更新
- `style:` 代码格式调整（不影响功能）
- `refactor:` 重构
- `test:` 测试相关
- `chore:` 构建工具或辅助工具的变动

### 提交格式
```
<type>: <subject>

<body>

<footer>
```

**示例：**
```
feat: 更新物料接口&万能转链字段

- 物料搜索升级版接口 IncomeInfo 新增 SubsidyUpperLimit、SubsidyType 字段
- 商品详情查询升级版接口 IncomeInfo 新增 SubsidyUpperLimit、SubsidyType 字段
- 万能转链接口新增 SetRequiredLinkType 方法支持 required_link_type 参数

参考：千牛头条公告(2026.02.09)、GitHub Issue #6
```

---

## 代码规范

### Go 代码风格
- 遵循 [Effective Go](https://golang.org/doc/effective_go) 指南
- 使用 `gofmt` 格式化代码（注：`jdopensdk/request/` 下由生成器模板产出的文件保持模板
  原样即可，手写新文件应通过 gofmt）
- 遵循项目现有代码风格

### 代码审查检查项
- [ ] 代码通过编译
- [ ] 参数键名与官方文档业务参数表一致
- [ ] 响应字段与官方文档返回结果表一致
- [ ] 已通过离线解析测试；有条件时已通过真实联调验证
- [ ] 提交信息清晰准确
- [ ] 没有引入不必要的依赖
- [ ] 没有破坏现有功能

---

## 常见问题

### Q: 如果错误提交到 main 分支怎么办？
A:
```bash
# 1. 回滚 main 分支
git reset --hard <previous_commit>
git push origin main --force

# 2. 创建正确的任务分支
git checkout -b task_<issue_id>

# 3. 重新提交
# ...
```

### Q: 如何重命名分支？
A:
```bash
# 重命名本地分支
git branch -m <old_name> <new_name>

# 推送新分支
git push origin <new_name>

# 删除旧分支
git push origin --delete <old_name>
```

### Q: 接口调用配额耗尽无法联调怎么办？
A: 先完成离线解析测试并在 PR 中注明「配额恢复后待真实联调确认」；
   配额恢复（京东联盟为每日 0 点重置）后运行 `client_test.go` 中对应测试补齐验证。

---

## 参考资料

- [GitHub Flow](https://docs.github.com/en/get-started/quickstart/github-flow)
- [Conventional Commits](https://www.conventionalcommits.org/)
- [Effective Go](https://golang.org/doc/effective_go)
- [京东联盟开放平台](https://union.jd.com/openplatform/api)
