package env

import (
	"encoding"
	"reflect"
	"strconv"
)

var defaultBuiltInParsers = map[reflect.Kind]ParserFunc{ //nolint:gochecknoglobals
	reflect.Bool: func(v string) (interface{}, error) {
		return strconv.ParseBool(v)
	},
	reflect.String: func(v string) (interface{}, error) {
		return v, nil
	},
	reflect.Int: func(v string) (interface{}, error) {
		i, err := strconv.ParseInt(v, 10, 32)
		return int(i), err
	},
	reflect.Int16: func(v string) (interface{}, error) {
		i, err := strconv.ParseInt(v, 10, 16)
		return int16(i), err
	},
	reflect.Int32: func(v string) (interface{}, error) {
		i, err := strconv.ParseInt(v, 10, 32)
		return int32(i), err
	},
	reflect.Int64: func(v string) (interface{}, error) {
		return strconv.ParseInt(v, 10, 64)
	},
	reflect.Int8: func(v string) (interface{}, error) {
		i, err := strconv.ParseInt(v, 10, 8)
		return int8(i), err
	},
	reflect.Uint: func(v string) (interface{}, error) {
		i, err := strconv.ParseUint(v, 10, 32)
		return uint(i), err
	},
	reflect.Uint16: func(v string) (interface{}, error) {
		i, err := strconv.ParseUint(v, 10, 16)
		return uint16(i), err
	},
	reflect.Uint32: func(v string) (interface{}, error) {
		i, err := strconv.ParseUint(v, 10, 32)
		return uint32(i), err
	},
	reflect.Uint64: func(v string) (interface{}, error) {
		i, err := strconv.ParseUint(v, 10, 64)
		return i, err
	},
	reflect.Uint8: func(v string) (interface{}, error) {
		i, err := strconv.ParseUint(v, 10, 8)
		return uint8(i), err
	},
	reflect.Float64: func(v string) (interface{}, error) {
		return strconv.ParseFloat(v, 64)
	},
	reflect.Float32: func(v string) (interface{}, error) {
		f, err := strconv.ParseFloat(v, 32)
		return float32(f), err
	},
}

func defaultTypeParsers() map[reflect.Type]ParserFunc { _ = "STUB: not implemented"; return nil }

func parseURL(v string) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func parseDuration(v string) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func parseLocation(v string) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

type ParserFunc func(v string) (interface{}, error)

type OnSetFn func(tag string, value interface{}, isDefault bool)

type processFieldFn func(
	refField reflect.Value,
	refTypeField reflect.StructField,
	opts Options,
	fieldParams FieldParams,
) error

type Options struct {
	Environment map[string]string

	TagName string

	PrefixTagName string

	DefaultValueTagName string

	RequiredIfNoDef bool

	OnSet OnSetFn

	Prefix string

	UseFieldNameByDefault bool

	SetDefaultsForZeroValuesOnly bool

	FuncMap map[reflect.Type]ParserFunc

	rawEnvVars map[string]string
}

func (opts *Options) getRawEnv(s string) string { _ = "STUB: not implemented"; return "" }

func defaultOptions() Options { _ = "STUB: not implemented"; return *new(Options) }

func mergeOptions[T any](target, source *T) { _ = "STUB: not implemented"; return }

func isZero(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

func customOptions(opts Options) Options { _ = "STUB: not implemented"; return *new(Options) }

func optionsWithSliceEnvPrefix(opts Options, index int) Options {
	_ = "STUB: not implemented"
	return *new(Options)
}

func optionsWithEnvPrefix(field reflect.StructField, opts Options) Options {
	_ = "STUB: not implemented"
	return *new(Options)
}

func Parse(v interface{}) error { _ = "STUB: not implemented"; return nil }

func ParseWithOptions(v interface{}, opts Options) error { _ = "STUB: not implemented"; return nil }

func ParseAs[T any]() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func ParseAsWithOptions[T any](opts Options) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func Must[T any](t T, err error) T { _ = "STUB: not implemented"; return *new(T) }

func GetFieldParams(v interface{}) ([]FieldParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetFieldParamsWithOptions(v interface{}, opts Options) ([]FieldParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseInternal(v interface{}, processField processFieldFn, opts Options) error {
	_ = "STUB: not implemented"
	return nil
}

func doParse(ref reflect.Value, processField processFieldFn, opts Options) error {
	_ = "STUB: not implemented"
	return nil
}

func doParseField(
	refField reflect.Value,
	refTypeField reflect.StructField,
	processField processFieldFn,
	opts Options,
) error {
	_ = "STUB: not implemented"
	return nil
}

func isSliceOfStructs(refTypeField reflect.StructField) bool {
	_ = "STUB: not implemented"
	return false
}

func doParseSlice(ref reflect.Value, processField processFieldFn, opts Options) error {
	_ = "STUB: not implemented"
	return nil
}

func setField(refField reflect.Value, refTypeField reflect.StructField, opts Options, fieldParams FieldParams) error {
	_ = "STUB: not implemented"
	return nil
}

const underscore rune = '_'

func toEnvName(input string) string { _ = "STUB: not implemented"; return "" }

type FieldParams struct {
	OwnKey          string
	Key             string
	DefaultValue    string
	HasDefaultValue bool
	Required        bool
	LoadFile        bool
	Unset           bool
	NotEmpty        bool
	Expand          bool
	Init            bool
	Ignored         bool
}

func parseFieldParams(field reflect.StructField, opts Options) (FieldParams, error) {
	_ = "STUB: not implemented"
	return *new(FieldParams), nil
}

func get(fieldParams FieldParams, opts Options) (val string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func parseKeyForOption(key string) (string, []string) { _ = "STUB: not implemented"; return "", nil }

func getFromFile(filename string) (value string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getOr(key, defaultValue string, defExists bool, envs map[string]string) (val string, exists, isDefault bool) {
	_ = "STUB: not implemented"
	return "", false, false
}

func set(field reflect.Value, sf reflect.StructField, value string, funcMap map[reflect.Type]ParserFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func handleSlice(field reflect.Value, value string, sf reflect.StructField, funcMap map[reflect.Type]ParserFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func handleMap(field reflect.Value, value string, sf reflect.StructField, funcMap map[reflect.Type]ParserFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func asTextUnmarshaler(field reflect.Value) encoding.TextUnmarshaler {
	_ = "STUB: not implemented"
	return *new(encoding.TextUnmarshaler)
}

func parseTextUnmarshalers(field reflect.Value, data []string, sf reflect.StructField) error {
	_ = "STUB: not implemented"
	return nil
}

func ToMap(env []string) map[string]string { _ = "STUB: not implemented"; return nil }

func isInvalidPtr(v reflect.Value) bool { _ = "STUB: not implemented"; return false }
