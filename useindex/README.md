# useindex

> 为 gorm 添加 mysql 的 `use index` 支持

## 用法

```go

// 注册插件
useindex.Register(db)

// 支持 select,delete,update
// select * from user use index(key1) where name = 'jinzhu'
db.Where("name = ?", "jinzhu").Set("use index","key1").First(&user)
```
