package validator

import (
	"github.com/google/uuid"
	"github.com/sdkopen/sdkopen-go-webserver/types"
)

func registerCustomTypes() {
	registerUUIDCustomType()
	registerDateTimeCustomType()
	registerIsoDateCustomType()
	registerIsoTimeCustomType()
	registerNullBoolCustomType()
	registerNullDateTimeCustomType()
	registerNullFloat64CustomType()
	registerNullInt16CustomType()
	registerNullInt32CustomType()
	registerNullInt64CustomType()
	registerNullIsoDateCustomType()
	registerNullIsoTimeCustomType()
}

func registerUUIDCustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (any, error) {
		return uuid.Parse(vals[0])
	}, uuid.UUID{})
}

func registerDateTimeCustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return types.ParseDateTime(vals[0])
	}, types.DateTime{})
}

func registerIsoDateCustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return types.ParseIsoDate(vals[0])
	}, types.IsoDate{})
}

func registerIsoTimeCustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return types.ParseIsoTime(vals[0])
	}, types.IsoTime{})
}

func registerNullBoolCustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return types.ParseNullBool(vals[0])
	}, types.NullBool{})
}

func registerNullDateTimeCustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return types.ParseNullDateTime(vals[0])
	}, types.NullDateTime{})
}

func registerNullFloat64CustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return types.ParseNullFloat64(vals[0])
	}, types.NullFloat64{})
}

func registerNullInt16CustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return types.ParseNullInt16(vals[0])
	}, types.NullInt16{})
}

func registerNullInt32CustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return types.ParseNullInt32(vals[0])
	}, types.NullInt32{})
}

func registerNullInt64CustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return types.ParseNullInt64(vals[0])
	}, types.NullInt64{})
}

func registerNullIsoDateCustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return types.ParseNullIsoDate(vals[0])
	}, types.NullIsoDate{})
}

func registerNullIsoTimeCustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return types.ParseNullIsoTime(vals[0])
	}, types.NullIsoTime{})
}
