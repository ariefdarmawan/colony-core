package clncore

import (
	"time"

	"github.com/eaciit/dbox"
	"github.com/eaciit/orm"
)

type DataModel struct {
	orm.ModelBase `bson:"-" json:"-"`
	ID            string `bson:"_id"`
	Title         string
	Fields        map[string]*DataField
	CreatedBy     string
	LastModified  time.Time
	UpdatedBy     string
}

func (o *DataModel) TableName() string {
	return "datamodels"
}
func NewDataModel() *DataModel {
	o := new(DataModel)
	return o
}
func DataModelFind(filter *dbox.Filter, fields, orders string, limit, skip int) dbox.ICursor {
	config := makeFindConfig(fields, orders, skip, limit)
	if filter != nil {
		config.Set("where", filter)
	}
	c, _ := DB().Find(new(DataModel), config)
	return c
}
func DataModelGet(filter *dbox.Filter, orders string, skip int) (emp *DataModel, err error) {
	config := makeFindConfig("", orders, skip, 1)
	if filter != nil {
		config.Set("where", filter)
	}
	c, ecursor := DB().Find(new(DataModel), config)
	if ecursor != nil {
		return nil, ecursor
	}
	defer c.Close()

	emp = new(DataModel)
	err = c.Fetch(emp, 1, false)
	return emp, err
}

func DataModelGetByID(pID string, orders string) (*DataModel, error) {
	return DataModelGet(dbox.Eq("_id", pID), orders, 0)
}

func DataModelGetByTitle(pTitle string, orders string) (*DataModel, error) {
	return DataModelGet(dbox.Eq("title", pTitle), orders, 0)
}

func DataModelFindByTitle(pTitle string, fields string, limit, skip int) dbox.ICursor {
	return DataModelFind(dbox.Eq("title", pTitle), orders, "", limit, skip)
}
