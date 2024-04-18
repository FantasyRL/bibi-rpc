package db

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *Comment) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "i"
	o = append(o, 0x86, 0xa1, 0x69)
	o = msgp.AppendInt64(o, z.ID)
	// string "v"
	o = append(o, 0xa1, 0x76)
	o = msgp.AppendInt64(o, z.VideoID)
	// string "p"
	o = append(o, 0xa1, 0x70)
	o = msgp.AppendInt64(o, z.ParentID)
	// string "u"
	o = append(o, 0xa1, 0x75)
	o = msgp.AppendInt64(o, z.Uid)
	// string "c"
	o = append(o, 0xa1, 0x63)
	o = msgp.AppendString(o, z.Content)
	// string "pu"
	o = append(o, 0xa2, 0x70, 0x75)
	o = msgp.AppendTime(o, z.CreatedAt)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Comment) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "i":
			z.ID, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ID")
				return
			}
		case "v":
			z.VideoID, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "VideoID")
				return
			}
		case "p":
			z.ParentID, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ParentID")
				return
			}
		case "u":
			z.Uid, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Uid")
				return
			}
		case "c":
			z.Content, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Content")
				return
			}
		case "pu":
			z.CreatedAt, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedAt")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Comment) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.StringPrefixSize + len(z.Content) + 3 + msgp.TimeSize
	return
}