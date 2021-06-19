package env

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/cheesycoffee/go-utilities/errorutil"
)

// MustParseEnv must parse env to struct, panic if env from target struct tag is not found
func MustParseEnv(target interface{}) {
	pValue := reflect.ValueOf(target)
	pValue = pValue.Elem()
	pType := reflect.TypeOf(target).Elem()
	mErrs := errorutil.NewMultiError()
	for i := 0; i < pValue.NumField(); i++ {
		field := pValue.Field(i)
		if !field.CanSet() { // skip if field cannot set a value (usually an unexported field in struct), to avoid a panic
			continue
		}

		typ := pType.Field(i)
		if typ.Anonymous ||
			(typ.Type.Kind() == reflect.Struct && !reflect.DeepEqual(field.Interface(), time.Time{})) { // embedded struct or struct field
			MustParseEnv(field.Addr().Interface())
			continue
		}

		key := typ.Tag.Get("env")
		if key == "" || key == "-" {
			continue
		}

		// split key by semicolon to check if it has required option or not
		// if required is given then it will panic when the value is empty
		keyTags := strings.Split(key, ";")
		var isRequired bool
		for _, v := range keyTags {
			if strings.TrimSpace(strings.ToLower(v)) == "required" {
				isRequired = true
				break
			}
		}

		val, ok := os.LookupEnv(key)
		if !ok && isRequired {
			mErrs.Append(key, fmt.Errorf("missing %s environment", keyTags[0]))
			continue
		}

		switch field.Interface().(type) {
		case time.Duration:
			dur, err := time.ParseDuration(val)
			if err != nil {
				mErrs.Append(key, fmt.Errorf("env '%s': %v", key, err))
				continue
			}
			field.Set(reflect.ValueOf(dur))
		case time.Time:
			t, err := time.Parse(time.RFC3339, val)
			if err != nil {
				mErrs.Append(key, fmt.Errorf("env '%s': %v", key, err))
				continue
			}
			field.Set(reflect.ValueOf(t))
		case int32, int, int64:
			vInt, err := strconv.Atoi(val)
			if err != nil {
				mErrs.Append(key, fmt.Errorf("env '%s': %v", key, err))
				continue
			}
			field.SetInt(int64(vInt))
		case float32, float64:
			vFloat, err := strconv.ParseFloat(val, 64)
			if err != nil {
				mErrs.Append(key, fmt.Errorf("env '%s': %v", key, err))
				continue
			}
			field.SetFloat(vFloat)
		case bool:
			vBool, err := strconv.ParseBool(val)
			if err != nil {
				mErrs.Append(key, fmt.Errorf("env '%s': %v", key, err))
				continue
			}
			field.SetBool(vBool)
		case string:
			field.SetString(val)
		}
	}

	if mErrs.HasError() {
		panic("Environment error: \n" + mErrs.Error())
	}
}
