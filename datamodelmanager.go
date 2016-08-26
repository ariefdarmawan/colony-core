package clncore

import "reflect"

type DataModelManager struct {
	container map[string]*DataModel
}

func NewDataModelManager() *DataModelManager {
	dmm := new(DataModelManager)
	dmm.container = map[string]*DataModel{}
	return dmm
}

func (dmm *DataModelManager) Get(id string) *DataModel {
	return dmm.container[id]
}

func (dmm *DataModelManager) Set(id string, obj *DataModel) *DataModelManager {
	if obj == nil {
		delete(dmm.container, id)
	} else {
		dmm.container[id] = obj
	}
	return dmm
}

func (dmm *DataModelManager) SetObj(obj interface{}) *DataModelManager {
	return dmm.SetObjAndID("", obj)
}

func (dmm *DataModelManager) SetObjAndID(id string, obj interface{}) *DataModelManager {
	rval := reflect.Indirect(reflect.ValueOf(obj))
	rtype := rval.Type()
	if id == "" {
		id = rtype.String()
	}
	//toolkit.Println("Registering objects", id)

	dm := dmm.Get(id)
	if dm == nil {
		dm = NewDataModel()
	}

	fieldCount := rtype.NumField()
	for fieldIdx := 0; fieldIdx < fieldCount; fieldIdx++ {
		fiter := rtype.FieldByIndex([]int{fieldIdx})
		fieldName := fiter.Name
		//fieldValue := rval.FieldByIndex([]int{fieldIdx})

		field := new(DataField)
		field.ID = fieldName
		field.FieldType = fiter.Type.String()
		field.FieldTag = string(fiter.Tag)
		dm.Fields[fieldName] = field
	}
	dmm.Set(id, dm)

	return dmm
}
