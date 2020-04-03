package nullz

import (
	"github.com/vmihailenco/msgpack"
)

var _ msgpack.CustomEncoder = (*StringAlias)(nil)
var _ msgpack.CustomEncoder = (*Uint64Alias)(nil)
var _ msgpack.CustomEncoder = (*BoolAlias)(nil)
var _ msgpack.CustomEncoder = (*TimeAlias)(nil)

func (s *StringAlias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if s.Valid {
		return enc.EncodeString(s.String)
	}

	return enc.EncodeNil()
}

func (u *Uint64Alias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if u.Valid {
		return enc.EncodeUint64(u.Uint64)
	}

	return enc.EncodeNil()
}

func (b *BoolAlias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if b.Valid {
		return enc.EncodeBool(b.Bool)
	}

	return enc.EncodeNil()
}

func (t *TimeAlias) EncodeMsgpack(enc *msgpack.Encoder) error {
	if t.Valid {
		return enc.EncodeTime(t.Time)
	}

	return enc.EncodeNil()
}
