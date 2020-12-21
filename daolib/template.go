package daolib

import (
	"fmt"
	"text/template"
)

func parseTemplateOrPanic(t string) *template.Template {
	tpl, err := template.New("output_template").Parse(t)
	if err != nil {
		panic(err)
	}
	return tpl
}

var commonTemplate = parseTemplateOrPanic(fmt.Sprintf(`
package {{.PkgName}}
type FieldData struct {
		Value interface{} %sjson:"value" form:"value"%s
		Symbol string %sjson:"symbol" form:"symbol"%s
	
}
`, "`", "`", "`", "`"))

var outputTemplate = parseTemplateOrPanic(fmt.Sprintf(`
package {{.PkgName}}
{{$TransformErr :=.TransformErr}}
import (
{{if $TransformErr}} "errors" {{end}}
	{{range .ImportPkgs}}
	"{{.Pkg}}"
	{{end}}
)


{{$LogName := .LogName}}
	
 {{if $TransformErr}} var(
		 ErrCreate{{.StructName}} = errors.New("create {{.StructName}} failed") 
		 ErrDelete{{.StructName}} = errors.New("delete {{.StructName}} failed") 
		 ErrGet{{.StructName}} = errors.New("get {{.StructName}} failed") 
		 ErrUpdate{{.StructName}} = errors.New("update {{.StructName}} failed") 
	)
{{end}}

	// {{.StructName}}Dao
	type {{.StructName}}Dao struct {
		Db *gorm.DB
		*{{.StructName}}
	}
	// Add add one record
	func (t *{{.StructName}}Dao) Add()(err error) {
		if err = t.Db.Create(t).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err){{end}}
			{{if $TransformErr}} err = ErrCreate{{.StructName}}{{end}}
			return
		}
		return 
	}

	// Delete delete record
	func (t *{{.StructName}}Dao) Delete()(err error) {
		if err = t.Db.Delete(t).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
			{{if $TransformErr}} err = ErrDelete{{.StructName}} {{end}}
			return
		}
		return
	}
	
	// Updates update record
	func (t *{{.StructName}}Dao) Updates(m map[string]interface{})(err error) {
		if err = t.Db.Model(&{{.StructName}}{}).Where("id = ?",t.ID).Updates(m).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
			{{if $TransformErr}} err = ErrUpdate{{.StructName}} {{end}}
			return
		}
		return 
	}

	// GetAll get all record
	func (t *{{.StructName}}Dao) GetAll()(ret []*{{.StructName}},err error){
		if err = t.Db.Find(&ret).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
			{{if $TransformErr}} err = ErrGet{{.StructName}} {{end}}
			return
		}
		return
	}
	
	// GetCount get count
	func (t *{{.StructName}}Dao) GetCount()(ret int64){
		t.Db.Model(&{{.StructName}}{}).Count(&ret)
		return
	}

	{{$StructName := .StructName}}
	type Query{{$StructName}}Form struct{
	{{range .OptionFields}} {{.FieldName}} *FieldData %sjson:"{{.HumpName}}" form:"{{.HumpName}}"%s; {{end}}
		Order []string %sjson:"order" form:"order"%s
		PageNum int %sjson:"pageNum" form:"pageNum"%s
		PageSize int %sjson:"pageSize" form:"pageSize"%s
		}

	//  GetList get list some field value or some condition
	func (t *{{.StructName}}Dao) GetList(q *Query{{$StructName}}Form)(ret []*{{$StructName}},err error){
		// order
		if len(q.Order)>0{
			for _,v:=range q.Order {
				t.Db = t.Db.Order(v)
			}
		}
		// pageSize
		if q.PageSize!=0{
			t.Db = t.Db.Limit(q.PageSize)
		}
		// pageNum
		if q.PageNum!=0{
			q.PageNum = (q.PageNum - 1) * q.PageSize
			t.Db = t.Db.Offset(q.PageNum)
		}
	{{range .OptionFields}} 
		// {{.FieldName}}
		if q.{{.FieldName}}!=nil{
			t.Db = t.Db.Where("{{.ColumnName}}" +q.{{.FieldName}}.Symbol +"?",q.{{.FieldName}}.Value)
		}  ; {{end}}
		if err = t.Db.Find(&ret).Error;err!=nil{
			return	
		}
		return 
	}
	{{range .OnlyFields}}
		// QueryBy{{.FieldName}} query cond by {{.FieldName}}
		func (t *{{$StructName}}Dao) SetQueryBy{{.FieldName}}({{.ColumnName}} {{.FieldType}})*{{$StructName}} {
			t.{{.FieldName}} = {{.ColumnName}}
			return  t.{{$StructName}}
		}
		// GetBy{{.FieldName}} get one record by {{.FieldName}}
		func (t *{{$StructName}}Dao)GetBy{{.FieldName}}()(err error){
			if err = t.Db.First(t,"{{.ColumnName}} = ?",t.{{.FieldName}}).Error;err!=nil{
				{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
				{{if $TransformErr}} err = ErrGet{{$StructName}} {{end}}
				return
			}
			return
		}
		// DeleteBy{{.FieldName}} delete record by {{.FieldName}}
		func (t *{{$StructName}}Dao) DeleteBy{{.FieldName}}()(err error) {
			if err= t.Db.Delete(t,"{{.ColumnName}} = ?",t.{{.FieldName}}).Error;err!=nil{
				{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
				{{if $TransformErr}} err = ErrDelete{{$StructName}} {{end}}
				return
				}
			return
		}
	{{end}}
`, "`", "`", "`", "`", "`", "`", "`", "`"))
