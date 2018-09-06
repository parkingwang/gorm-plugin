# duplicate

> 为 gorm 添加 mysql的 `on_duplicate_key_update` 支持
## 用法

```go

// 注册插件
duplicate.Register(db)

// 更新指定字段
up := []string{"name", "address"}
// or
up := duplicate.Cols("name", "address")
db.Set("on_duplicate_key_update", up).Create(&user)

// 更新所有字段
db.Set("on_duplicate_key_update",nil).Create(&user)

// 自定义更新属性
up := duplicate.Exec("name = ?, age = ?, num = num + 1", "afocus", 30)
db.Set("on_duplicate_key_update", up).Create(&user)
```