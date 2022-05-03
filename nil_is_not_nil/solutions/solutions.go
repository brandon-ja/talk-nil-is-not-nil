package main

import (
	"fmt"
	"io"
	"reflect"
)

func main() {
	//START OMIT
	tmp := 5
	i := interface{}(&tmp)

	if ii, ok := i.(*int); ok && ii != nil {
		fmt.Printf("i is %d\n", *ii)
	}
	//END OMIT

}

//START_FN OMIT
func isNilV(r io.Reader) bool {
	if r == nil {
		return true
	}

	// https://pkg.go.dev/reflect#Value.IsNil
	rv := reflect.ValueOf(r)
	if (rv.Kind() == reflect.Ptr && rv.IsNil()) ||
		(rv.Kind() == reflect.Chan && rv.IsNil()) ||
		(rv.Kind() == reflect.Func && rv.IsNil()) ||
		(rv.Kind() == reflect.Map && rv.IsNil()) ||
		(rv.Kind() == reflect.Interface && rv.IsNil()) ||
		(rv.Kind() == reflect.Slice && rv.IsNil()) {
		return true
	}

	return false
}

//END_FN OMIT
