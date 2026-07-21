package env

import (
	"reflect"
)

type AggregateError struct {
	Errors []error
}

func newAggregateError(initErr error) error { _ = "STUB: not implemented"; return nil }

func (e AggregateError) Error() string { _ = "STUB: not implemented"; return "" }

func (e AggregateError) Unwrap() []error { _ = "STUB: not implemented"; return nil }

func (e AggregateError) Is(err error) bool { _ = "STUB: not implemented"; return false }

type ParseError struct {
	Name string
	Type reflect.Type
	Err  error
}

func newParseError(sf reflect.StructField, err error) error { _ = "STUB: not implemented"; return nil }

func (e ParseError) Error() string { _ = "STUB: not implemented"; return "" }

type NotStructPtrError struct{}

func (e NotStructPtrError) Error() string { _ = "STUB: not implemented"; return "" }

type NoParserError struct {
	Name string
	Type reflect.Type
}

func newNoParserError(sf reflect.StructField) error { _ = "STUB: not implemented"; return nil }

func (e NoParserError) Error() string { _ = "STUB: not implemented"; return "" }

type NoSupportedTagOptionError struct {
	Tag string
}

func newNoSupportedTagOptionError(tag string) error { _ = "STUB: not implemented"; return nil }

func (e NoSupportedTagOptionError) Error() string { _ = "STUB: not implemented"; return "" }

type EnvVarIsNotSetError = VarIsNotSetError

type VarIsNotSetError struct {
	Key string
}

func newVarIsNotSetError(key string) error { _ = "STUB: not implemented"; return nil }

func (e VarIsNotSetError) Error() string { _ = "STUB: not implemented"; return "" }

type EmptyEnvVarError = EmptyVarError

type EmptyVarError struct {
	Key string
}

func newEmptyVarError(key string) error { _ = "STUB: not implemented"; return nil }

func (e EmptyVarError) Error() string { _ = "STUB: not implemented"; return "" }

type LoadFileContentError struct {
	Filename string
	Key      string
	Err      error
}

func newLoadFileContentError(filename, key string, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func (e LoadFileContentError) Error() string { _ = "STUB: not implemented"; return "" }

type ParseValueError struct {
	Msg string
	Err error
}

func newParseValueError(message string, err error) error { _ = "STUB: not implemented"; return nil }

func (e ParseValueError) Error() string { _ = "STUB: not implemented"; return "" }
