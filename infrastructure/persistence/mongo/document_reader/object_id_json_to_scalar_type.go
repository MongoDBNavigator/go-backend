package document_reader

import "reflect"

func (rcv *documentReader) objectIDToScalarType(objID interface{}) interface{} {
	if reflect.ValueOf(objID).Kind() == reflect.Map {
		if id, ok := objID.(map[string]interface{})["$oid"]; ok {
			return id.(string)
		}
	}

	return objID
}
