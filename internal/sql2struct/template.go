package sql2struct

import (
	"fmt"
	"html/template"
	"os"
	"permission-cat/internal/word"
)

const structTemplate = `
type {{.TableName | ToCamelCase}} struct {
	{{range .Columns}} {{$length := len .Comment}} {{ if gt $length 0}}//{{.Comment}} {{else}}// {{.Name}} {{end}}
	{{$typeLen :=len .Type}} {{if gt $typeLen 0 }}{{.Name | ToCamelCase}} {{.Type}}  {{.Tag}}{{else}}{{.Name}}{{end}}
	{{end}}
}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}
`

type StructTemplate struct {
	structTemplate string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTemplate: structTemplate}
}

func (t *StructTemplate) AssemblyColumns(tableColumns []*TableColumn) []*StructColumn {
	templateColumns := make([]*StructColumn, 0, len(tableColumns))

	for _, column := range tableColumns {
		templateColumns = append(templateColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     fmt.Sprintf("`json:"+"%s"+"`", column.ColumnName),
			Comment: column.ColumnComment,
		})
	}

	return templateColumns
}

func (t *StructTemplate) Generate(tableName string, templateColumns []*StructColumn) error {
	templateSql := template.Must(
		template.New("sql2struct").Funcs(template.FuncMap{
			"ToCamelCase": word.UnderscoreToUpperCamelCase,
		}).Parse(t.structTemplate),
	)

	templateDB := StructTemplateDB{
		TableName: tableName,
		Columns:   templateColumns,
	}

	err := templateSql.Execute(os.Stdout, templateDB)
	if err != nil {
		return err
	}

	return nil
}
