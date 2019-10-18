package useindex

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

const name = "use_index"

func Register(db *gorm.DB) {
	db.Callback().Query().After("gorm:query").Register(name, UseIndex)
	db.Callback().Update().After("gorm:update").Register(name, UseIndex)
	db.Callback().Delete().After("gorm:delete").Register(name, UseIndex)
}

func UseIndex(scope *gorm.Scope) {
	if v, ok := scope.Get(name); ok {
		scope.Search.Table(fmt.Sprintf("%s use index(%v)", scope.TableName(), v))
	}
}
