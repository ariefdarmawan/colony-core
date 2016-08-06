package clncore

import (
	"github.com/eaciit/dbox"
	"github.com/eaciit/orm"
	"github.com/eaciit/toolkit"
	"time"
)

type DataConnection struct {
	orm.ModelBase `bson:"-" json:"-"`
	ID            string `bson:"_id"`
	Title         string
	Host          string
	Database      string
	User          string
	Password      string
	Config        toolkit.M
	Created       time.Time
	CreatedBy     string
	LastModified  time.Time
	UpdatedBy     string
	Enable        bool
}

func (o *DataConnection) TableName() string {
	return "dataconnections"
}
func NewDataConnection() *DataConnection {
	o := new(DataConnection)
	o.Enable = true
	return o
}
func DataConnectionFind(filter *dbox.Filter, fields, orders string, limit, skip int) dbox.ICursor {
	config := makeFindConfig(fields, orders, skip, limit)
	if filter != nil {
		config.Set("where", filter)
	}
	c, _ := DB().Find(new(DataConnection), config)
	return c
}
func DataConnectionGet(filter *dbox.Filter, orders string, skip int) (emp *DataConnection, err error) {
	config := makeFindConfig("", orders, skip, 1)
	if filter != nil {
		config.Set("where", filter)
	}
	c, ecursor := DB().Find(new(DataConnection), config)
	if ecursor != nil {
		return nil, ecursor
	}
	defer c.Close()

	emp = new(DataConnection)
	err = c.Fetch(emp, 1, false)
	return emp, err
}

func DataConnectionGetByID(pID string, orders string) (*DataConnection, error) {
	return DataConnectionGet(dbox.Eq("_id", pID), orders, 0)
}

func DataConnectionGetByTitle(pTitle string, orders string) (*DataConnection, error) {
	return DataConnectionGet(dbox.Eq("title", pTitle), orders, 0)
}

func DataConnectionFindByTitle(pTitle string, fields string, limit, skip int) dbox.ICursor {
	return DataConnectionFind(dbox.Eq("title", pTitle), orders, "", limit, skip)
}
