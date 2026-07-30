package main

import (
	"C"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sauci/a2l-grpc/pkg/a2l"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"sync"
)

var protocolSizeMargin = 256

func chunkifyBySize(data []byte, chunkSize int) [][]byte {
	var chunks [][]byte
	for start := 0; start < len(data); start += chunkSize {
		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[start:end])
	}
	return chunks
}

func getTreeFromString(a2lString string) (result *a2l.RootNodeType, err error) {
	return a2l.GetTreeFromString(a2lString)
}

type grpcA2LImplType struct {
	a2l.UnimplementedA2LServer
	chunkSize int
}

func (s *grpcA2LImplType) GetTreeFromA2L(stream a2l.A2L_GetTreeFromA2LServer) error {
	var buffer bytes.Buffer
	var parseError error
	var err error
	var serializedTree []byte
	var chunk []byte
	var request *a2l.TreeFromA2LRequest
	tree := &a2l.RootNodeType{}
	response := &a2l.TreeResponse{}

	for {
		request, err = stream.Recv()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		buffer.Write(request.A2L)
	}

	if err == nil {
		if tree, parseError = getTreeFromString(buffer.String()); parseError == nil {
			if serializedTree, err = proto.Marshal(tree); err == nil {
				for _, chunk = range chunkifyBySize(serializedTree, s.chunkSize) {
					response.SerializedTreeChunk = chunk
					if err = stream.Send(response); err != nil {
						break
					}
				}
			} else {
				response.Error = proto.String(fmt.Sprintf("An error occured during serialization of Tree: %v", err))
				err = stream.Send(response)
			}
		} else {
			errString := parseError.Error()
			response.Error = &errString
			err = stream.Send(response)
		}
	}

	return err
}

func (s *grpcA2LImplType) GetJSONFromTree(stream a2l.A2L_GetJSONFromTreeServer) (err error) {
	var rawData []byte
	var buffer bytes.Buffer
	var chunk []byte
	var parseError error
	var request *a2l.JSONFromTreeRequest
	tree := &a2l.RootNodeType{}
	response := &a2l.JSONResponse{}
	indent := ""
	allowPartial := false
	emitUnpopulated := false
	// Note: optionsParsed := false avoid to parse option for each chunk
	optionsParsed := false

	for {
		request, err = stream.Recv()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if !optionsParsed {
			if request.Indent != nil {
				for i := uint32(0); i < *request.Indent; i++ {
					indent += " "
				}
			}
			if request.AllowPartial != nil {
				allowPartial = *request.AllowPartial
			}
			if request.EmitUnpopulated != nil {
				emitUnpopulated = *request.EmitUnpopulated
			}
			optionsParsed = true
		}
		// Use a buffer to receive all chunks
		buffer.Write(request.Tree)
	}
	if err == nil {
		if parseError = proto.Unmarshal(buffer.Bytes(), tree); parseError == nil {
			opt := protojson.MarshalOptions{
				AllowPartial:    allowPartial,
				EmitUnpopulated: emitUnpopulated}

			if rawData, parseError = opt.Marshal(tree); parseError == nil {
				// Note: see https://github.com/golang/protobuf/issues/1121
				var indentedBuffer bytes.Buffer
				if err = json.Indent(&indentedBuffer, rawData, "", indent); err == nil {
					rawData = indentedBuffer.Bytes()
					for _, chunk = range chunkifyBySize(rawData, s.chunkSize) {
						response.Json = chunk
						if err = stream.Send(response); err != nil {
							break
						}
					}
				} else {
					response.Error = proto.String(fmt.Sprintf("An error occured during json indent: %v", err))
					err = stream.Send(response)
				}
			} else {
				errString := parseError.Error()
				response.Error = &errString
				err = stream.Send(response)
			}
		} else {
			errString := parseError.Error()
			response.Error = &errString
			err = stream.Send(response)
		}
	}

	return err
}

func (s *grpcA2LImplType) GetTreeFromJSON(stream a2l.A2L_GetTreeFromJSONServer) (err error) {
	var parseError error
	var buffer bytes.Buffer
	var serializedTree []byte
	var chunk []byte
	var request *a2l.TreeFromJSONRequest
	tree := &a2l.RootNodeType{}
	response := &a2l.TreeResponse{}
	allowPartial := false
	// Note: optionsParsed := false avoid to parse option for each chunk
	optionsParsed := false

	for {
		request, err = stream.Recv()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		if !optionsParsed {
			if request.AllowPartial != nil {
				allowPartial = *request.AllowPartial
			}
			optionsParsed = true
		}
		// Use a buffer to receive all chunks
		buffer.Write(request.Json)
	}
	if err == nil {
		opt := protojson.UnmarshalOptions{
			AllowPartial: allowPartial,
		}

		if parseError = opt.Unmarshal(buffer.Bytes(), tree); parseError == nil {
			if serializedTree, err = proto.Marshal(tree); err == nil {
				for _, chunk = range chunkifyBySize(serializedTree, s.chunkSize) {
					response.SerializedTreeChunk = chunk
					if err = stream.Send(response); err != nil {
						break
					}
				}
			} else {
				response.Error = proto.String(fmt.Sprintf("An error occured during serialization of Tree: %v", err))
				err = stream.Send(response)
			}
		} else {
			errString := parseError.Error()
			response.Error = &errString
			err = stream.Send(response)
		}
	}

	return err
}

func (s *grpcA2LImplType) GetA2LFromTree(stream a2l.A2L_GetA2LFromTreeServer) (err error) {
	var buffer bytes.Buffer
	var request *a2l.A2LFromTreeRequest
	var chunk []byte
	var a2lDataBytes []byte
	tree := &a2l.RootNodeType{}
	response := &a2l.A2LResponse{}
	indent := ""
	sorted := false
	// Note: optionsParsed := false avoid to parse option for each chunk
	optionsParsed := false

	for {
		request, err = stream.Recv()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		if !optionsParsed {
			if request.Indent != nil {
				for i := uint32(0); i < *request.Indent; i++ {
					indent += " "
				}
			}
			if request.Sorted != nil {
				sorted = *request.Sorted
			}
			optionsParsed = true
		}
		// Use a buffer to receive all chunks
		buffer.Write(request.Tree)
	}
	if err == nil {
		if err = proto.Unmarshal(buffer.Bytes(), tree); err == nil {
			a2lDataBytes = []byte(tree.MarshalA2L(0, indent, sorted))
			for _, chunk = range chunkifyBySize(a2lDataBytes, s.chunkSize) {
				response.A2L = chunk
				if err = stream.Send(response); err != nil {
					break
				}
			}
		}
	}

	return err
}

//export GetJSONByteArrayFromA2LByteArray
func GetJSONByteArrayFromA2LByteArray(a2lByteArray []byte) {
	_ = a2lByteArray
}

func recoverUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = status.Errorf(codes.Internal, "%v: %v", info.FullMethod, r)
		}
	}()

	return handler(ctx, req)
}

func recoverStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = status.Errorf(codes.Internal, "%v: %v", info.FullMethod, r)
		}
	}()

	return handler(srv, ss)
}

var serverMutex sync.Mutex

var server *grpc.Server

//export Create
func Create(port C.int, maxMsgSize C.int) (result C.int) {
	var err error
	var listener net.Listener

	serverMutex.Lock()
	defer serverMutex.Unlock()

	if server != nil {
		result = 1
	} else {
		result = 0
	}

	if result == 0 {
		if listener, err = net.Listen("tcp", fmt.Sprintf(":%v", port)); err == nil {
			server = grpc.NewServer(
				grpc.MaxRecvMsgSize(int(maxMsgSize)),
				grpc.MaxSendMsgSize(int(maxMsgSize)),
				grpc.ChainUnaryInterceptor(recoverUnaryInterceptor),
				grpc.ChainStreamInterceptor(recoverStreamInterceptor))

			a2l.RegisterA2LServer(server, &grpcA2LImplType{chunkSize: int(maxMsgSize) - protocolSizeMargin})

			go func() {
				err = server.Serve(listener)
			}()
		}
	}

	return result
}

//export Close
func Close() (result C.int) {
	serverMutex.Lock()
	defer serverMutex.Unlock()

	if server == nil {
		result = 1
	} else {
		server.Stop()

		server = nil

		result = 0
	}

	return result
}

func main() {
	Create(3333, 4*1024*1024)

	for {
		select {}
	}
}
