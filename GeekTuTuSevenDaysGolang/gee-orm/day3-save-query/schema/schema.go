package schema

import (
	"day1-database-sql/dialect"
	"go/ast"
	"reflect"
)

/*
定义对象object与表tabled的转换
*/

type Field struct {
	Name string // 字段命
	Type string // 类型
	Tag  string // 约束条件
}

type Schema struct {
	Model      any               // 被映射的对象 model
	Name       string            // 表名
	Fields     []*Field          // 字段 fields
	FieldNames []string          // 记录所有字段名
	fieldMap   map[string]*Field // 记录字段名和 Field 的映射关系
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// Parse 将任意对象解析为 Schema 实例
func Parse(dest any, d dialect.Dialect) *Schema {
	// 获取传入结构体的真实类型，比如 dest 为 &User{}， modelTpe 就是 User
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	// 初始化 Schema 结构
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}
	// 遍历 struct 所有字段
	for i := 0; i < modelType.NumField(); i++ {
		// p 是 reflect.StructField，包含： Name Type Tag Anonymous（是否匿名字段） Exported（是否大写开头字段）
		p := modelType.Field(i)
		// 跳过匿名字段和非导出字段
		if !p.Anonymous && ast.IsExported(p.Name) {
			// 构造 Field
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			// 解析 tag 获取约束条件
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			// 填入 Schema
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}
