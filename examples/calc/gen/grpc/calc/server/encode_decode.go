// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// calc GRPC server encoders and decoders
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design -o
// $(GOPATH)/src/goa.design/goa/examples/calc

package server

import (
	"context"

	calcsvc "goa.design/goa/examples/calc/gen/calc"
	"goa.design/goa/examples/calc/gen/grpc/calc/pb"
	goagrpc "goa.design/goa/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeAddResponse encodes responses from the "calc" service "add" endpoint.
func EncodeAddResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(int)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "add", "int", v)
	}
	resp := NewAddResponse(result)
	return resp, nil
}

// DecodeAddRequest decodes requests sent to "calc" service "add" endpoint.
func DecodeAddRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *pb.AddRequest
		ok      bool
		err     error
	)
	{
		if message, ok = v.(*pb.AddRequest); !ok {
			return nil, goagrpc.ErrInvalidType("calc", "add", "*pb.AddRequest", v)
		}
	}
	var (
		payload *calcsvc.AddPayload
	)
	{
		payload = NewAddPayload(message)
	}
	return payload, err
}