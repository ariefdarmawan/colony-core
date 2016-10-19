package clncore

import (
	"reflect"

	"github.com/eaciit/toolkit"
)

func (dm *DataModel) FieldArray() []*DataField {
	fs := []*DataField{}
	for _, v := range dm.Fields {
		fs = append(fs, v)
	}
	return fs
}

func DataModelFrom(obj interface{}) *DataModel {
	dm := NewDataModel()
	rval := reflect.Indirect(reflect.ValueOf(obj))
	rtype := rval.Type()

	fieldCount := rtype.NumField()
	for fieldIdx := 0; fieldIdx < fieldCount; fieldIdx++ {
		filter := rtype.FieldByIndex([]int{fieldIdx})
		fieldName := filter.Name

		field := new(DataField)
		field.ID = fieldName
		field.FieldType = rtype.String()
		field.FieldTag = string(filter.Tag)

		field.GridColumn = toolkit.ToInt(field.Tag("gridcolumn"), toolkit.RoundingAuto)
		field.GridShow = field.GridColumn >= 0
		field.GridWidth = toolkit.ToInt(field.Tag("gridwidth"), toolkit.RoundingAuto)
		field.GridFormat = field.Tag("gridformat")
		field.GridTemplate = field.Tag("gridtemplate")
		if field.GridFormat == "" && field.GridTemplate != "" {
			field.GridUseTemplate = true
		}

		field.Label = field.TagDefault("label", field.ID)

		dm.Fields[fieldName] = field
	}

	return dm
}
