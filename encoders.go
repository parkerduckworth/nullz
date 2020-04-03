package nullz

import "github.com/vmihailenco/msgpack"

var _ msgpack.CustomEncoder = (*BoolAlias)(nil)
var _ msgpack.CustomEncoder = (*ByteAlias)(nil)
var _ msgpack.CustomEncoder = (*BytesAlias)(nil)
var _ msgpack.CustomEncoder = (*Float32Alias)(nil)
var _ msgpack.CustomEncoder = (*Float64Alias)(nil)
var _ msgpack.CustomEncoder = (*IntAlias)(nil)
var _ msgpack.CustomEncoder = (*Int16Alias)(nil)
var _ msgpack.CustomEncoder = (*Int32Alias)(nil)
var _ msgpack.CustomEncoder = (*Int64Alias)(nil)
var _ msgpack.CustomEncoder = (*Int8Alias)(nil)
var _ msgpack.CustomEncoder = (*JsonAlias)(nil)
var _ msgpack.CustomEncoder = (*StringAlias)(nil)
var _ msgpack.CustomEncoder = (*TimeAlias)(nil)
var _ msgpack.CustomEncoder = (*UintAlias)(nil)
var _ msgpack.CustomEncoder = (*Uint16Alias)(nil)
var _ msgpack.CustomEncoder = (*Uint32Alias)(nil)
var _ msgpack.CustomEncoder = (*Uint64Alias)(nil)
var _ msgpack.CustomEncoder = (*Uint8Alias)(nil)

func (b *BoolAlias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if b.Valid {
		return enc.EncodeBool(b.Bool)
	}

	return enc.EncodeNil()
}

func (b *ByteAlias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if b.Valid {
		return enc.EncodeBytes(make([]byte, b.Byte))
	}

	return enc.EncodeNil()
}

func (b *BytesAlias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if b.Valid {
		return enc.EncodeBytes(b.Bytes)
	}

	return enc.EncodeNil()
}

func (f *Float32Alias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if f.Valid {
		return enc.EncodeFloat32(f.Float32)
	}

	return enc.EncodeNil()
}

func (f *Float64Alias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if f.Valid {
		return enc.EncodeFloat64(f.Float64)
	}

	return enc.EncodeNil()
}

func (i *IntAlias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if i.Valid {
		return enc.EncodeInt(int64(i.Int))
	}

	return enc.EncodeNil()
}

func (i *Int16Alias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if i.Valid {
		return enc.EncodeInt(int64(i.Int16))
	}

	return enc.EncodeNil()
}

func (i *Int32Alias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if i.Valid {
		return enc.EncodeInt(int64(i.Int32))
	}

	return enc.EncodeNil()
}

func (i *Int64Alias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if i.Valid {
		return enc.EncodeInt(i.Int64)
	}

	return enc.EncodeNil()
}

func (i *Int8Alias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if i.Valid {
		return enc.EncodeInt(int64(i.Int8))
	}

	return enc.EncodeNil()
}

func (j *JsonAlias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if j.Valid {
		return enc.EncodeBytes(j.JSON)
	}

	return enc.EncodeNil()
}

func (s *StringAlias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if s.Valid {
		return enc.EncodeString(s.String)
	}

	return enc.EncodeNil()
}

func (t *TimeAlias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if t.Valid {
		return enc.EncodeTime(t.Time)
	}

	return enc.EncodeNil()
}

func (u *UintAlias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if u.Valid {
		return enc.EncodeUint64(uint64(u.Uint))
	}

	return enc.EncodeNil()
}

func (u *Uint16Alias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if u.Valid {
		return enc.EncodeUint64(uint64(u.Uint16))
	}

	return enc.EncodeNil()
}

func (u *Uint32Alias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if u.Valid {
		return enc.EncodeUint64(uint64(u.Uint32))
	}

	return enc.EncodeNil()
}

func (u *Uint64Alias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if u.Valid {
		return enc.EncodeUint64(u.Uint64)
	}

	return enc.EncodeNil()
}

func (u *Uint8Alias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if u.Valid {
		return enc.EncodeUint64(uint64(u.Uint8))
	}

	return enc.EncodeNil()
}
