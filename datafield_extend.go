package clncore

import (
	"reflect"
)

func (df *DataField) Tag(tagid string) string {
	return reflect.StructTag(df.FieldTag).Get(tagid)
}
