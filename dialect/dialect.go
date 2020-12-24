package dialect
// 数据库与Golang的类型映射

import "reflect"

type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	// 用于将 Go 语言的类型转换为该数据库的数据类型。
	TableExistSQL(tableName string) (string, []interface{})
}

var dialectsMap = map[string]Dialect{}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect,ok bool) {
	dialect, ok = dialectsMap[name]
	return dialect, ok
}