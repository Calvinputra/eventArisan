package helper

import (
	"fmt"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
)

func IsDateFormat(fl validator.FieldLevel) bool {
	_, err := time.Parse("20060102", fl.Field().String())
	return err == nil
}

func IsDateTimeFormat(fl validator.FieldLevel) bool {
	_, err := time.Parse("20060102 15:04:05", fl.Field().String())
	return err == nil
}

func CopyStruct(src interface{}, dst interface{}) error {
	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()

	if srcVal.Kind() != reflect.Struct || dstVal.Kind() != reflect.Struct {
		return fmt.Errorf("expected both src and dst to be structs, got %s and %s", srcVal.Kind(), dstVal.Kind())
	}
	for i := 0; i < srcVal.NumField(); i++ {
		field := srcVal.Field(i)
		dstField := dstVal.FieldByName(srcVal.Type().Field(i).Name)
		if dstField.IsValid() && dstField.CanSet() {
			if field.Type() == dstField.Type() {
				dstField.Set(field)
			} else if field.Type().Kind() == reflect.Interface && field.Type().Implements(dstField.Type()) {
				dstField.Set(field)
			}
		}
	}
	return nil
}

///*
//RequiredIfEmpty
//Validation required if another field is empty
//*/
//func RequiredIfEmpty(fl validator.FieldLevel) bool {
//	fmt.Printf("[RequiredIfEmpty] parent field: %s\n", fl.Parent().FieldByName(fl.Param()).String())
//	//fmt.Printf("[RequiredIfEmpty] field: %s\n", reflect.TypeOf(fl.Field()).Kind())
//	if fl.Parent().FieldByName(fl.Param()).String() == "" {
//		fmt.Printf("[RequiredIfEmpty] masuk if\n")
//		return fl.Field().String() == ""
//	}
//	return true
//}

///*
//RequiredIfEmptyStruct
//Validation required if another field is empty
//*/
//func RequiredIfEmptyStruct(fl validator.FieldLevel) bool {
//	fmt.Printf("[RequiredIfEmptyStruct] param field is zero: %v\n", fl.Parent().FieldByName(fl.Param()).IsZero())
//	//fmt.Printf("[RequiredIfEmptyStruct] field: %s\n", reflect.TypeOf(fl.Field()).Kind())
//	if fl.Parent().FieldByName(fl.Param()).IsZero() {
//		fmt.Printf("[RequiredIfEmptyStruct] masuk if\n")
//		return fl.Field().String() == ""
//	}
//	return true
//}
