package nullz

import "github.com/vmihailenco/msgpack"

var _ msgpack.CustomEncoder = (*Bool)(nil)
var _ msgpack.CustomEncoder = (*Byte)(nil)
var _ msgpack.CustomEncoder = (*Bytes)(nil)
var _ msgpack.CustomEncoder = (*Float32)(nil)
var _ msgpack.CustomEncoder = (*Float64)(nil)
var _ msgpack.CustomEncoder = (*Int)(nil)
var _ msgpack.CustomEncoder = (*Int16)(nil)
var _ msgpack.CustomEncoder = (*Int32)(nil)
var _ msgpack.CustomEncoder = (*Int64)(nil)
var _ msgpack.CustomEncoder = (*Int8)(nil)
var _ msgpack.CustomEncoder = (*JSON)(nil)
var _ msgpack.CustomEncoder = (*String)(nil)
var _ msgpack.CustomEncoder = (*Time)(nil)
var _ msgpack.CustomEncoder = (*Uint)(nil)
var _ msgpack.CustomEncoder = (*Uint16)(nil)
var _ msgpack.CustomEncoder = (*Uint32)(nil)
var _ msgpack.CustomEncoder = (*Uint64)(nil)
var _ msgpack.CustomEncoder = (*Uint8)(nil)

func (b *Bool) EncodeMsgpack(enc *msgpack.Encoder) error {
	if b.Valid {
		return enc.EncodeBool(b.Bool.Bool)
	}

	return enc.EncodeNil()
}

func (b *Byte) EncodeMsgpack(enc *msgpack.Encoder) error {
	if b.Valid {
		return enc.EncodeBytes(make([]byte, b.Byte.Byte))
	}

	return enc.EncodeNil()
}

func (b *Bytes) EncodeMsgpack(enc *msgpack.Encoder) error {
	if b.Valid {
		return enc.EncodeBytes(b.Bytes.Bytes)
	}

	return enc.EncodeNil()
}

func (f *Float32) EncodeMsgpack(enc *msgpack.Encoder) error {
	if f.Valid {
		return enc.EncodeFloat32(f.Float32.Float32)
	}

	return enc.EncodeNil()
}

func (f *Float64) EncodeMsgpack(enc *msgpack.Encoder) error {
	if f.Valid {
		return enc.EncodeFloat64(f.Float64.Float64)
	}

	return enc.EncodeNil()
}

func (i *Int) EncodeMsgpack(enc *msgpack.Encoder) error {
	if i.Valid {
		return enc.EncodeInt(int64(i.Int.Int))
	}

	return enc.EncodeNil()
}

func (i *Int16) EncodeMsgpack(enc *msgpack.Encoder) error {
	if i.Valid {
		return enc.EncodeInt(int64(i.Int16.Int16))
	}

	return enc.EncodeNil()
}

func (i *Int32) EncodeMsgpack(enc *msgpack.Encoder) error {
	if i.Valid {
		return enc.EncodeInt(int64(i.Int32.Int32))
	}

	return enc.EncodeNil()
}

func (i *Int64) EncodeMsgpack(enc *msgpack.Encoder) error {
	if i.Valid {
		return enc.EncodeInt(i.Int64.Int64)
	}

	return enc.EncodeNil()
}

func (i *Int8) EncodeMsgpack(enc *msgpack.Encoder) error {
	if i.Valid {
		return enc.EncodeInt(int64(i.Int8.Int8))
	}

	return enc.EncodeNil()
}

func (j *JSON) EncodeMsgpack(enc *msgpack.Encoder) error {
	if j.Valid {
		return enc.EncodeBytes(j.JSON.JSON)
	}

	return enc.EncodeNil()
}

func (s *String) EncodeMsgpack(enc *msgpack.Encoder) error {
	if s.Valid {
		return enc.EncodeString(s.String.String)
	}

	return enc.EncodeNil()
}

func (t *Time) EncodeMsgpack(enc *msgpack.Encoder) error {
	if t.Valid {
		return enc.EncodeTime(t.Time.Time)
	}

	return enc.EncodeNil()
}

func (u *Uint) EncodeMsgpack(enc *msgpack.Encoder) error {
	if u.Valid {
		return enc.EncodeUint64(uint64(u.Uint.Uint))
	}

	return enc.EncodeNil()
}

func (u *Uint16) EncodeMsgpack(enc *msgpack.Encoder) error {
	if u.Valid {
		return enc.EncodeUint64(uint64(u.Uint16.Uint16))
	}

	return enc.EncodeNil()
}

func (u *Uint32) EncodeMsgpack(enc *msgpack.Encoder) error {
	if u.Valid {
		return enc.EncodeUint64(uint64(u.Uint32.Uint32))
	}

	return enc.EncodeNil()
}

func (u *Uint64) EncodeMsgpack(enc *msgpack.Encoder) error {
	if u.Valid {
		return enc.EncodeUint64(u.Uint64.Uint64)
	}

	return enc.EncodeNil()
}

func (u *Uint8) EncodeMsgpack(enc *msgpack.Encoder) error {
	if u.Valid {
		return enc.EncodeUint64(uint64(u.Uint8.Uint8))
	}

	return enc.EncodeNil()
}
