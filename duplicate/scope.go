package duplicate

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

const name = "irain:on_duplicate_key_update"

func Register(db *gorm.DB) {
	db.Callback().Create().Before("gorm:create").Register(name, OnDuplicateKey)
}

func OnDuplicateKey(scope *gorm.Scope) {
	scope.Set("gorm:insert_option", &Scope{scope})
}

type Scope struct {
	*gorm.Scope
}

func (o *Scope) String() string {
	s, ok := o.Get(name)
	if !ok {
		return ""
	}
	cols, ok := s.([]string)
	if !ok || len(cols) == 0 {
		for _, v := range o.Fields() {
			if v.IsBlank || v.IsPrimaryKey || v.IsIgnored {
				continue
			}
			cols = append(cols, v.DBName)
		}
	}
	var kv []string
	for _, col := range cols {
		field, ok := o.FieldByName(col)
		if !ok {
			continue
		}
		kv = append(kv, fmt.Sprintf("%v = %v", o.Quote(field.DBName), o.AddToVars(field.Field.Interface())))
	}
	return "on duplicate key update " + strings.Join(kv, ",")

}
