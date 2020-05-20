package rgz

type MsgMapper func(msg interface{}) (newMsg interface{}, err error)

func (f MsgMapper) ProxyPolyMarshaler(pm PolyMarshaler) PolyMarshaler {
	return PolyMarshalerFunc(func(msg interface{}) (data []byte, err error) {
		msg, err = f(msg)
		if err != nil {
			return
		}
		return pm.PolyMarshal(msg)
	})
}

func (f MsgMapper) ProxyPolyUnmarshaler(pum PolyUnmarshaler) PolyUnmarshaler {
	return PolyUnmarshalerFunc(func(data []byte) (msg interface{}, err error) {
		msg, err = pum.PolyUnmarshal(data)
		if err != nil {
			return
		}
		return f(msg)
	})
}

func (f MsgMapper) ProxyMarshaler(m Marshaler) Marshaler {
	return MarshalerFunc(func(msg interface{}) (data []byte, err error) {
		msg, err = f(msg)
		if err != nil {
			return
		}
		return m.Marshal(msg)
	})
}
