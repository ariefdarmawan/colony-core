package clncore

import (
	"reflect"
)

func (df *DataField) Tag(tagid string) string {
	return reflect.StructTag(df.FieldTag).Get(tagid)
}

func (df *DataField) TagDefault(tagid string, def string) string {
	r := df.Tag(tagid)
	if r != "" {
		return r
	} else {
		return def
	}
}
