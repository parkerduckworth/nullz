package nullz

import "github.com/vmihailenco/msgpack"

var _ msgpack.CustomDecoder = (*BoolAlias)(nil)
var _ msgpack.CustomDecoder = (*ByteAlias)(nil)
var _ msgpack.CustomDecoder = (*BytesAlias)(nil)
var _ msgpack.CustomDecoder = (*Float32Alias)(nil)
var _ msgpack.CustomDecoder = (*Float64Alias)(nil)
var _ msgpack.CustomDecoder = (*IntAlias)(nil)
var _ msgpack.CustomDecoder = (*Int16Alias)(nil)
var _ msgpack.CustomDecoder = (*Int32Alias)(nil)
var _ msgpack.CustomDecoder = (*Int64Alias)(nil)
var _ msgpack.CustomDecoder = (*Int8Alias)(nil)
var _ msgpack.CustomDecoder = (*JsonAlias)(nil)
var _ msgpack.CustomDecoder = (*StringAlias)(nil)
var _ msgpack.CustomDecoder = (*TimeAlias)(nil)
var _ msgpack.CustomDecoder = (*UintAlias)(nil)
var _ msgpack.CustomDecoder = (*Uint16Alias)(nil)
var _ msgpack.CustomDecoder = (*Uint32Alias)(nil)
var _ msgpack.CustomDecoder = (*Uint64Alias)(nil)
var _ msgpack.CustomDecoder = (*Uint8Alias)(nil)

func (b *BoolAlias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if b != nil {
		return dec.Decode(&b.Bool)
	}

	return nil
}

func (b *ByteAlias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if b != nil {
		return dec.Decode(&b.Byte)
	}

	return nil
}

func (b *BytesAlias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if b != nil {
		return dec.Decode(&b.Bytes)
	}

	return nil
}

func (f *Float32Alias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if f != nil {
		return dec.Decode(&f.Float32)
	}

	return nil
}

func (f *Float64Alias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if f != nil {
		return dec.Decode(&f.Float64)
	}

	return nil
}

func (i *IntAlias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if i != nil {
		return dec.Decode(&i.Int)
	}

	return nil
}

func (i *Int16Alias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if i != nil {
		return dec.Decode(&i.Int16)
	}

	return nil
}

func (i *Int32Alias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if i != nil {
		return dec.Decode(&i.Int32)
	}

	return nil
}

func (i *Int64Alias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if i != nil {
		return dec.Decode(&i.Int64)
	}

	return nil
}

func (i *Int8Alias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if i != nil {
		return dec.Decode(&i.Int8)
	}

	return nil
}

func (j *JsonAlias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if j != nil {
		return dec.Decode(&j.JSON)
	}

	return nil
}

func (s *StringAlias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if s != nil {
		return dec.Decode(&s.String)
	}

	return nil
}

func (t *TimeAlias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if t != nil {
		return dec.Decode(&t.Time)
	}

	return nil
}

func (u *UintAlias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if u != nil {
		return dec.Decode(&u.Uint)
	}

	return nil
}

func (u *Uint16Alias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if u != nil {
		return dec.Decode(&u.Uint16)
	}

	return nil
}

func (u *Uint32Alias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if u != nil {
		return dec.Decode(&u.Uint32)
	}

	return nil
}

func (u *Uint64Alias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if u != nil {
		return dec.Decode(&u.Uint64)
	}

	return nil
}

func (u *Uint8Alias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if u != nil {
		return dec.Decode(&u.Uint8)
	}

	return nil
}
