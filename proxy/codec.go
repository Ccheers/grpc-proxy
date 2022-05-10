package proxy

import (
	"google.golang.org/grpc/encoding"
	eproto "google.golang.org/grpc/encoding/proto"
	"google.golang.org/protobuf/proto"
)

// Codec returns a proxying grpc.Codec with the default protobuf codec as parent.
// See CodecWithParent.
func Codec() encoding.Codec {
	return CodecWithParent(&protoCodec{})
}

// CodecWithParent returns a proxying grpc.Codec with a user provided codec as parent.
func CodecWithParent(fallback encoding.Codec) encoding.Codec {
	return &rawCodec{parentCodec: fallback}
}

type rawCodec struct {
	parentCodec encoding.Codec
}

func (c *rawCodec) Name() string {
	return eproto.Name
}

type frame struct {
	payload []byte
}

func (c *rawCodec) Marshal(v interface{}) ([]byte, error) {
	out, ok := v.(*frame)
	if !ok {
		return c.parentCodec.Marshal(v)
	}
	return out.payload, nil

}

func (c *rawCodec) Unmarshal(data []byte, v interface{}) error {
	dst, ok := v.(*frame)
	if !ok {
		return c.parentCodec.Unmarshal(data, v)
	}
	dst.payload = data
	return nil
}

// protoCodec is a Codec implementation with protobuf. It is the default rawCodec for gRPC.
type protoCodec struct{}

func (c protoCodec) Name() string {
	return ""
}

func (protoCodec) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (protoCodec) Unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}
