package duplicate

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

const name = "on_duplicate_key_update"

func Register(db *gorm.DB) {
	db.Callback().Create().Before("gorm:create").Register(name, OnDuplicateKey)
}

func OnDuplicateKey(scope *gorm.Scope) {
	if _, ok := scope.Get(name); ok {
		scope.Set("gorm:insert_option", &Scope{scope})
	}
}

type Scope struct {
	*gorm.Scope
}

type Update struct {
	s    string
	args []interface{}
}

func Exec(s string, args ...interface{}) *Update {
	return &Update{
		s:    s,
		args: args,
	}
}

func Cols(s ...string) []string {
	return s
}

func (o *Scope) String() string {
	s, ok := o.Get(name)
	if !ok {
		return ""
	}
	switch c := s.(type) {
	case []string:
		return o.fromCols(c)
	case *Update:
		return o.fromExec(c)
	case nil:
		return o.fromCols(nil)
	default:
		return ""
	}
}

func (o *Scope) fromCols(cols []string) string {
	if cols == nil || len(cols) == 0 {
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

func (o *Scope) fromExec(e *Update) string {
	if e.args != nil && len(e.args) > 0 {
		o.SQLVars = append(o.SQLVars, e.args...)
	}
	return "on duplicate key update " + e.s
}
