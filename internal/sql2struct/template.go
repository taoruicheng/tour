package sql2struct

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/taoruicheng/tour/internal/word"
)

const strcutTpl = `type {{.TableName | ToCamelCase}} struct {
	{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
		{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
	{{end}}}
	func (model {{.TableName | ToCamelCase}}) TableName() string {
		return "{{.TableName}}"
	}`

type StructTemplate struct {
	strcutTpl string
}

type StructColumn struct {
	Name    string //列名（ColumnName）：is_deleted
	Type    string //字段类型（DataType）：bigint，根据DBTypeToStructType查找对应关系
	Tag     string //转化json的tag，取的ColumnName
	Comment string //字段备注信息（ColumnComment）："是否已删除 0/正常，1/删除"
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{strcutTpl: strcutTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}

	return tplColumns
}

func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) (string, error) {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.strcutTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	var b strings.Builder

	err := tpl.Execute(&b, tplDB)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
