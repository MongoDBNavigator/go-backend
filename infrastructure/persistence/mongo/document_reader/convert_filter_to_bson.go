package document_reader

import (
	"github.com/mongodb/mongo-go-driver/bson"
)

func (rcv *documentReader) convertFilterToBson(data map[string]interface{}) *bson.Document {
	document := bson.NewDocument()

	//element := bson.EC.ArrayFromElements("$or",
	//	bson.VC.DocumentFromElements(bson.EC.Interface("name", "Roman")),
	//	bson.VC.DocumentFromElements(bson.EC.Interface("name", 123)),
	//)
	//
	//document.Append(element)
	//if regex, ok := data["$regex"]; ok {
	//	var options string
	//	if opts, ok := data["$options"]; ok {
	//		options = opts.(string)
	//	}
	//
	//	bson.EC.Regex("bla", regex.(string), options)
	//	bson.EC.String("$or", "aaa")
	//}
	//

	//for k, v := range data {
	//
	//	vType := reflect.TypeOf(v)
	//
	//	log.Println(vType.String())
	//
	//	if vType.String() == "map[string]interface {}" {
	//		if element := rcv.convertFilterToElement(k, v.(map[string]interface{})); element != nil {
	//			document.Append(element)
	//		}
	//	} else {
	//		document.Append(bson.EC.Interface(k, v))
	//	}
	//}

	return document
}

func (rcv *documentReader) convertFilterToElement(name string, data map[string]interface{}) *bson.Element {
	if element := rcv.convertFilterToRegex(name, data); element != nil {
		return element
	}

	return nil
}

func (rcv *documentReader) convertFilterToRegex(name string, data map[string]interface{}) *bson.Element {
	if regex, ok := data["$regex"]; ok {
		var options string
		if opts, ok := data["$options"]; ok {
			options = opts.(string)
		}

		return bson.EC.Regex(name, regex.(string), options)
	}

	return nil
}
