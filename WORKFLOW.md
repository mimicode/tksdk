# tksdk 开发工作流程

本文档定义 tksdk 项目的标准开发流程和规范，所有开发者必须遵循。

---

## 标准开发流程

### 1. 领取任务
- 从 GitHub Issues 中选择任务
- 确认任务编号（Issue ID）

### 2. 创建任务分支
```bash
# 从 main 分支创建新分支
git checkout main
git pull origin main
git checkout -b task_<issue_id>
```

### 3. 开发与提交
```bash
# 修改代码
# ...

# 提交代码
git add .
git commit -m "feat: 简洁描述修改内容"
```

### 4. 推送分支
```bash
git push origin task_<issue_id>
```

### 5. 创建 Pull Request
- 访问 GitHub 自动生成的 PR 链接
- 或手动创建 PR：从 `task_<issue_id>` 合并到 `main`

### 6. 代码审查与合并
- 等待代码审查
- 根据反馈修改代码
- 审查通过后合并到 main

### 7. 清理分支
```bash
# 合并后删除本地分支
git branch -d task_<issue_id>

# 删除远程分支
git push origin --delete task_<issue_id>
```

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
- 使用 `gofmt` 格式化代码
- 遵循项目现有代码风格

### 代码审查检查项
- [ ] 代码通过编译
- [ ] 代码符合项目规范
- [ ] 提交信息清晰准确
- [ ] 没有引入不必要的依赖
- [ ] 没有破坏现有功能

---

## 测试规范

### 单元测试
- 为新功能编写单元测试
- 确保测试覆盖率

### 接口测试
- 测试新接口功能
- 验证与 API 文档的一致性

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

---

## 参考资料

- [GitHub Flow](https://docs.github.com/en/get-started/quickstart/github-flow)
- [Conventional Commits](https://www.conventionalcommits.org/)
- [Effective Go](https://golang.org/doc/effective_go)
