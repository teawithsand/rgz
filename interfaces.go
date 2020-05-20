package rgz

type Marshaler interface {
	Marshal(msg interface{}) ([]byte, error)
}

type MarshalerFunc func(msg interface{}) ([]byte, error)

func (f MarshalerFunc) Marshal(msg interface{}) ([]byte, error) {
	return f(msg)
}

type Unmarshaler interface {
	Unmarshal(data []byte, msg interface{}) error
}

type UnmarshalerFunc func(data []byte, msg interface{}) error

func (f UnmarshalerFunc) Unmarshal(data []byte, msg interface{}) error {
	return f(data, msg)
}

// poly

type PolyMarshaler interface {
	PolyMarshal(msg interface{}) ([]byte, error)
}

type PolyMarshalerFunc func(msg interface{}) ([]byte, error)

func (f PolyMarshalerFunc) PolyMarshal(msg interface{}) ([]byte, error) {
	return f(msg)
}

type PolyUnmarshaler interface {
	PolyUnmarshal(data []byte) (interface{}, error)
}

type PolyUnmarshalerFunc func(data []byte) (interface{}, error)

func (f PolyUnmarshalerFunc) PolyUnmarshal(data []byte) (interface{}, error) {
	return f(data)
}
