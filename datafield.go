package clncore

import (
	"github.com/eaciit/dbox"
	"github.com/eaciit/orm"
	"github.com/eaciit/toolkit"
)

type DataField struct {
	orm.ModelBase   `bson:"-" json:"-"`
	ID              string `bson:"_id"`
	Label           string
	FieldType       string
	GridShow        bool
	GridColumn      int
	GridWidth       int
	GridFormat      string
	GridUseTemplate bool
	GridTemplate    string
	GridAlign       string
	DatabaseFieldID string
	MappingFields   []string
	FieldTag        string
	Config          toolkit.M
}

func (o *DataField) TableName() string {
	return "datafields"
}
func NewDataField() *DataField {
	o := new(DataField)
	return o
}
func DataFieldFind(filter *dbox.Filter, fields, orders string, limit, skip int) dbox.ICursor {
	config := makeFindConfig(fields, orders, skip, limit)
	if filter != nil {
		config.Set("where", filter)
	}
	c, _ := DB().Find(new(DataField), config)
	return c
}
func DataFieldGet(filter *dbox.Filter, orders string, skip int) (emp *DataField, err error) {
	config := makeFindConfig("", orders, skip, 1)
	if filter != nil {
		config.Set("where", filter)
	}
	c, ecursor := DB().Find(new(DataField), config)
	if ecursor != nil {
		return nil, ecursor
	}
	defer c.Close()

	emp = new(DataField)
	err = c.Fetch(emp, 1, false)
	return emp, err
}
