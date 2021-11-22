package table

import (
	"bytes"
	"encoding/binary"

	"github.com/regen-network/regen-ledger/orm/v2/internal/key"
	"github.com/regen-network/regen-ledger/orm/v2/types/kvlayout"
	"github.com/regen-network/regen-ledger/orm/v2/types/ormerrors"
	"google.golang.org/protobuf/proto"
)

func (s Store) Decode(k []byte, v []byte) (kvlayout.Entry, error) {
	if bytes.HasPrefix(k, s.Prefix) {
		if bytes.HasPrefix(k, s.PkPrefix) {
			pkValues, err := s.PkCodec.Decode(bytes.NewReader(k))
			if err != nil {
				return nil, err
			}

			msg := s.MsgType.New().Interface()
			err = proto.Unmarshal(v, msg)
			if err != nil {
				return nil, err
			}

			// NOTE: here we skip rehydrating the primary key
			return kvlayout.PrimaryEntry{
				Key:   pkValues,
				Value: msg,
			}, nil
		} else {
			r := bytes.NewReader(k)
			err := key.SkipPrefix(r, s.Prefix)
			if err != nil {
				return nil, err
			}

			id, err := binary.ReadUvarint(r)
			if err != nil {
				return nil, err
			}

			idx, ok := s.IndexersById[uint32(id)]
			if !ok {
				return nil, ormerrors.CantFindIndexer.Wrapf("trying to decode entry with id %d", id)
			}

			r.Reset(k)
			values, err := idx.Codec.Decode(r)
			if err != nil {
				return nil, err
			}

			idxKey, pk, err := idx.Codec.SplitIndexPK(values)
			if err != nil {
				return nil, err
			}

			return kvlayout.IndexEntry{
				IndexFields: idx.FieldNames,
				Key:         idxKey,
				PrimaryKey:  pk,
			}, nil
		}
	} else {
		return nil, ormerrors.UnexpectedDecodePrefix
	}
}