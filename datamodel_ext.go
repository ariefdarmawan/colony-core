package clncore

func (dm *DataModel) FieldArray() []*DataField {
	fs := []*DataField{}
	for _, v := range dm.Fields {
		fs = append(fs, v)
	}
	return fs
}
