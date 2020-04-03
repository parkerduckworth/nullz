package nullz

import (
	"github.com/vmihailenco/msgpack"
)

var _ msgpack.CustomDecoder = (*StringAlias)(nil)
var _ msgpack.CustomDecoder = (*Uint64Alias)(nil)
var _ msgpack.CustomDecoder = (*BoolAlias)(nil)
var _ msgpack.CustomDecoder = (*TimeAlias)(nil)

func (s *StringAlias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if s != nil {
		return dec.Decode(&s.String)
	}

	return nil
}

func (u *Uint64Alias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if u != nil {
		return dec.Decode(&u.Uint64)
	}

	return nil
}

func (b *BoolAlias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if b != nil {
		return dec.Decode(&b.Bool)
	}

	return nil
}

func (t *TimeAlias) DecodeMsgpack(dec *msgpack.Decoder) error {
	if t != nil {
		return dec.Decode(&t.Time)
	}

	return nil
}
