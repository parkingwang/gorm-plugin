# duplicate

> 为 gorm 添加 mysql的 `on_duplicate_key_update` 支持
## 用法

```go

// 注册插件
duplicate.Register(db)

// 更新指定字段
updates:=[]string{"name", "address"}
db.Set("irain:on_duplicate_key_update", updates).Create(&user)

// 更新所有字段
db.Set("irain:on_duplicate_key_update",nil).Create(&user)
```