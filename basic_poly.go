package rgz

import (
	"encoding/binary"
	"reflect"
)

func flattenType(ty reflect.Type) reflect.Type {
	for ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	return ty
}

const tagLengthBytes = 4

// BasicPolyMarshaler prepends tag as binary big endian integer
// to the message and marshals it with given marshaler.
func BasicPolyMarshaler(
	mar Marshaler,
	tt TypeToTag,
) PolyMarshaler {
	return PolyMarshalerFunc(func(msg interface{}) (data []byte, err error) {
		ty := flattenType(reflect.TypeOf(msg))
		tag, err := tt.TypeToTag(ty)
		if err != nil {
			return
		}
		msgData, err := mar.Marshal(msg)
		if err != nil {
			return
		}
		data = make([]byte, len(msgData)+tagLengthBytes)
		// TOOD(teawithsand): Implement memmove approach
		binary.BigEndian.PutUint32(data[:tagLengthBytes], uint32(tag))
		copy(data[tagLengthBytes:], msgData)

		return
	})
}

func BasicPolyUnmarshaler(
	umar Unmarshaler,
	tt TagToType,
) PolyUnmarshaler {
	return PolyUnmarshalerFunc(func(data []byte) (msg interface{}, err error) {
		if len(data) < tagLengthBytes {
			err = ErrTagNotFound
			return
		}
		tagRaw := data[:tagLengthBytes]
		data = data[tagLengthBytes:]
		tag := binary.BigEndian.Uint32(tagRaw)

		ty, err := tt.TagToType(Tag(tag))
		if err != nil {
			return
		}
		msg = reflect.New(ty).Interface()
		err = umar.Unmarshal(data, msg)

		return
	})
}

// TODO9(teawithsand): test these functions
