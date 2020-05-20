package rgz

import (
	"errors"
	"reflect"
)

var ErrTagNotFound = errors.New("rgz: Given tag was not found in type registry")
var ErrTypeNotFound = errors.New("rgz: Given type was not found in type registry")

type Tag uint32 // 2^32 should be enough distinct types

type TypeToTag interface {
	TypeToTag(reflect.Type) (Tag, error)
}

type TagToType interface {
	TagToType(Tag) (reflect.Type, error)
}

type TagHandler interface {
	TagToType
	TagToType
}

func NewTypeRegistry() *TypeRegistry {
	return &TypeRegistry{
		typeToTag: make(map[reflect.Type]Tag),
		tagToType: make(map[Tag]reflect.Type),
	}
}

type TypeRegistry struct {
	typeToTag map[reflect.Type]Tag
	tagToType map[Tag]reflect.Type
}

func (tr *TypeRegistry) RegisterType(ty reflect.Type, tag Tag) {
	tr.typeToTag[ty] = tag
	tr.tagToType[tag] = ty
}

func (tr *TypeRegistry) TypeToTag(ty reflect.Type) (Tag, error) {
	tg, ok := tr.typeToTag[ty]
	var err error
	if !ok {
		err = ErrTypeNotFound
	}
	return tg, err
}

func (tr *TypeRegistry) TagToType(tg Tag) (reflect.Type, error) {
	ty, ok := tr.tagToType[tg]
	var err error
	if !ok {
		err = ErrTypeNotFound
	}
	return ty, err
}
