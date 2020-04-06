package nullz

import "github.com/vmihailenco/msgpack"

var _ msgpack.CustomDecoder = (*Bool)(nil)
var _ msgpack.CustomDecoder = (*Byte)(nil)
var _ msgpack.CustomDecoder = (*Bytes)(nil)
var _ msgpack.CustomDecoder = (*Float32)(nil)
var _ msgpack.CustomDecoder = (*Float64)(nil)
var _ msgpack.CustomDecoder = (*Int)(nil)
var _ msgpack.CustomDecoder = (*Int16)(nil)
var _ msgpack.CustomDecoder = (*Int32)(nil)
var _ msgpack.CustomDecoder = (*Int64)(nil)
var _ msgpack.CustomDecoder = (*Int8)(nil)
var _ msgpack.CustomDecoder = (*JSON)(nil)
var _ msgpack.CustomDecoder = (*String)(nil)
var _ msgpack.CustomDecoder = (*Time)(nil)
var _ msgpack.CustomDecoder = (*Uint)(nil)
var _ msgpack.CustomDecoder = (*Uint16)(nil)
var _ msgpack.CustomDecoder = (*Uint32)(nil)
var _ msgpack.CustomDecoder = (*Uint64)(nil)
var _ msgpack.CustomDecoder = (*Uint8)(nil)

func (b *Bool) DecodeMsgpack(dec *msgpack.Decoder) error {
	if b != nil {
		return dec.Decode(&b.Bool)
	}

	return nil
}

func (b *Byte) DecodeMsgpack(dec *msgpack.Decoder) error {
	if b != nil {
		return dec.Decode(&b.Byte)
	}

	return nil
}

func (b *Bytes) DecodeMsgpack(dec *msgpack.Decoder) error {
	if b != nil {
		return dec.Decode(&b.Bytes)
	}

	return nil
}

func (f *Float32) DecodeMsgpack(dec *msgpack.Decoder) error {
	if f != nil {
		return dec.Decode(&f.Float32)
	}

	return nil
}

func (f *Float64) DecodeMsgpack(dec *msgpack.Decoder) error {
	if f != nil {
		return dec.Decode(&f.Float64)
	}

	return nil
}

func (i *Int) DecodeMsgpack(dec *msgpack.Decoder) error {
	if i != nil {
		return dec.Decode(&i.Int)
	}

	return nil
}

func (i *Int16) DecodeMsgpack(dec *msgpack.Decoder) error {
	if i != nil {
		return dec.Decode(&i.Int16)
	}

	return nil
}

func (i *Int32) DecodeMsgpack(dec *msgpack.Decoder) error {
	if i != nil {
		return dec.Decode(&i.Int32)
	}

	return nil
}

func (i *Int64) DecodeMsgpack(dec *msgpack.Decoder) error {
	if i != nil {
		return dec.Decode(&i.Int64)
	}

	return nil
}

func (i *Int8) DecodeMsgpack(dec *msgpack.Decoder) error {
	if i != nil {
		return dec.Decode(&i.Int8)
	}

	return nil
}

func (j *JSON) DecodeMsgpack(dec *msgpack.Decoder) error {
	if j != nil {
		return dec.Decode(&j.JSON)
	}

	return nil
}

func (s *String) DecodeMsgpack(dec *msgpack.Decoder) error {
	if s != nil {
		return dec.Decode(&s.String)
	}

	return nil
}

func (t *Time) DecodeMsgpack(dec *msgpack.Decoder) error {
	if t != nil {
		return dec.Decode(&t.Time)
	}

	return nil
}

func (u *Uint) DecodeMsgpack(dec *msgpack.Decoder) error {
	if u != nil {
		return dec.Decode(&u.Uint)
	}

	return nil
}

func (u *Uint16) DecodeMsgpack(dec *msgpack.Decoder) error {
	if u != nil {
		return dec.Decode(&u.Uint16)
	}

	return nil
}

func (u *Uint32) DecodeMsgpack(dec *msgpack.Decoder) error {
	if u != nil {
		return dec.Decode(&u.Uint32)
	}

	return nil
}

func (u *Uint64) DecodeMsgpack(dec *msgpack.Decoder) error {
	if u != nil {
		return dec.Decode(&u.Uint64)
	}

	return nil
}

func (u *Uint8) DecodeMsgpack(dec *msgpack.Decoder) error {
	if u != nil {
		return dec.Decode(&u.Uint8)
	}

	return nil
}
