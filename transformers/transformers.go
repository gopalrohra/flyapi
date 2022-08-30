package transformers

import (
	"reflect"

	"github.com/gopalrohra/flyapi/log"
	"github.com/gopalrohra/flyapi/util"
)

type TransformerFunc = func(reflect.Value, string)

var Transformers = map[string]TransformerFunc{
	"int":       IntTransformer,
	"string":    StringTransformer,
	"time.Time": TimeTransformer,
}

func IntTransformer(f reflect.Value, v string) {
	log.Debugf("Value of request param:%v\n", v)
	intV := util.ToInt(v)
	if !f.OverflowInt(int64(intV)) {
		f.SetInt(int64(intV))
	}
}
func StringTransformer(f reflect.Value, v string) {
	log.Info(v)
	f.SetString(v)
}
func TimeTransformer(f reflect.Value, v string) {
	f.Set(reflect.ValueOf(util.ToDate(v)))
}
