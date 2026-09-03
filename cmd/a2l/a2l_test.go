package main

import (
	"context"
	"fmt"
	"github.com/sauci/a2l-grpc/pkg/a2l"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"strings"
	"testing"
	"time"
)

type mockStreamBase[R any, S any] struct {
	requests  []R
	responses []S
	reqIndex  int
}

// newMockStreamBase creates a new mockStreamBase generic.
func newMockStreamBase[R any, S any]() *mockStreamBase[R, S] {
	return &mockStreamBase[R, S]{
		requests:  make([]R, 0),
		responses: make([]S, 0),
	}
}

// SendRequest simulates the sending of a request by the client
func (m *mockStreamBase[R, S]) SendRequest(req R) {
	m.requests = append(m.requests, req)
}

// RecvResponse simulates the reception of a client-side response
func (m *mockStreamBase[R, S]) RecvResponse() (S, error) {
	if len(m.responses) == 0 {
		var zero S
		return zero, io.EOF
	}
	resp := m.responses[0]
	m.responses = m.responses[1:]
	return resp, nil
}

// Close cleans resources
func (m *mockStreamBase[R, S]) Close() {
	m.requests = nil
	m.responses = nil
	m.reqIndex = 0
}

// Server-side functions

// Recv reads the following request (server-side)
func (m *mockStreamBase[R, S]) Recv() (R, error) {
	if m.reqIndex >= len(m.requests) {
		var zero R
		return zero, io.EOF
	}
	req := m.requests[m.reqIndex]
	m.reqIndex++
	return req, nil
}

// Send a response (server side)
func (m *mockStreamBase[R, S]) Send(resp S) error {
	m.responses = append(m.responses, resp)
	return nil
}

// Methods for satisfying the interface grpc.ServerStream
func (m *mockStreamBase[R, S]) SetHeader(metadata.MD) error  { return nil }
func (m *mockStreamBase[R, S]) SendHeader(metadata.MD) error { return nil }
func (m *mockStreamBase[R, S]) SetTrailer(metadata.MD)       {}
func (m *mockStreamBase[R, S]) Context() context.Context     { return context.Background() }
func (m *mockStreamBase[R, S]) SendMsg(interface{}) error    { return nil }
func (m *mockStreamBase[R, S]) RecvMsg(interface{}) error    { return nil }

func Test_GetTreeFromA2L_ConversionOverflowReturnsError(t *testing.T) {
	a2lString := `/begin PROJECT _ "_"
 /begin MODULE _ "_"
  /begin CHARACTERISTIC c "_" VALUE 0xFFFFFFFFFFFFFFFF _ 0 _ 0 0
  /end CHARACTERISTIC
 /end MODULE
/end PROJECT`

	impl := newGrpcA2LImpl(4 * 1024 * 1024)

	mockStream := newMockStreamBase[*a2l.TreeFromA2LRequest, *a2l.TreeResponse]()
	defer mockStream.Close()

	mockStream.SendRequest(&a2l.TreeFromA2LRequest{A2L: []byte(a2lString)})

	if err := impl.GetTreeFromA2L(mockStream); err == nil {
		if response, err := mockStream.RecvResponse(); err == nil {
			if assert.NotNil(t, response.Error) {
				assert.Contains(t, *response.Error, "value out of range")
			}
		} else {
			t.Fatal(err)
		}
	} else {
		t.Fatal(err)
	}
}

func Test_getTreeFromString_ConversionOverflowReturnsError(t *testing.T) {
	tree, err := getTreeFromString(`/begin PROJECT _ "_"
 /begin MODULE _ "_"
  /begin CHARACTERISTIC c "_" VALUE 0xFFFFFFFFFFFFFFFF _ 0 _ 0 0
  /end CHARACTERISTIC
 /end MODULE
/end PROJECT`)

	assert.Nil(t, tree)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "value out of range")
}

func Test_RecoverUnaryInterceptor(t *testing.T) {
	response, err := recoverUnaryInterceptor(
		context.Background(),
		nil,
		&grpc.UnaryServerInfo{FullMethod: "/A2L/Test"},
		func(context.Context, interface{}) (interface{}, error) {
			panic("test panic")
		})

	assert.Nil(t, response)
	assert.Equal(t, codes.Internal, status.Code(err))
	assert.Contains(t, err.Error(), "test panic")
}

func Test_RecoverStreamInterceptor(t *testing.T) {
	err := recoverStreamInterceptor(
		nil,
		nil,
		&grpc.StreamServerInfo{FullMethod: "/A2L/Test"},
		func(interface{}, grpc.ServerStream) error {
			panic("test panic")
		})

	assert.Equal(t, codes.Internal, status.Code(err))
	assert.Contains(t, err.Error(), "test panic")
}

func Test_GetJSONFromTree(t *testing.T) {
	type testCaseType struct {
		tree            *a2l.RootNodeType
		indent          uint32
		allowPartial    bool
		emitUnpopulated bool
		expectedJSON    string
		testChunkSize   int
	}

	for _, tc := range []testCaseType{
		{
			tree: &a2l.RootNodeType{
				PROJECT: &a2l.ProjectType{
					Name:           &a2l.IdentType{Value: "_"},
					LongIdentifier: &a2l.StringType{Value: "_"},
				},
			},
			indent:          1,
			allowPartial:    false,
			emitUnpopulated: false,
			testChunkSize:   3,
			expectedJSON: `{
 "PROJECT": {
  "Name": {
   "Value": "_"
  },
  "LongIdentifier": {
   "Value": "_"
  }
 }
}`,
		},
		{
			tree: &a2l.RootNodeType{
				PROJECT: &a2l.ProjectType{
					Name:           &a2l.IdentType{Value: "_"},
					LongIdentifier: &a2l.StringType{Value: "_"},
				},
			},
			indent:          2,
			allowPartial:    false,
			emitUnpopulated: false,
			testChunkSize:   3,
			expectedJSON: `{
  "PROJECT": {
    "Name": {
      "Value": "_"
    },
    "LongIdentifier": {
      "Value": "_"
    }
  }
}`,
		},
		{
			tree: &a2l.RootNodeType{
				PROJECT: &a2l.ProjectType{
					Name:           &a2l.IdentType{Value: "_"},
					LongIdentifier: &a2l.StringType{Value: "_"},
				},
			},
			indent:          3,
			allowPartial:    false,
			emitUnpopulated: false,
			testChunkSize:   3,
			expectedJSON: `{
   "PROJECT": {
      "Name": {
         "Value": "_"
      },
      "LongIdentifier": {
         "Value": "_"
      }
   }
}`,
		},
		{
			tree: &a2l.RootNodeType{
				PROJECT: &a2l.ProjectType{
					Name:           &a2l.IdentType{Value: "_"},
					LongIdentifier: &a2l.StringType{Value: "_"},
				},
			},
			indent:          1,
			allowPartial:    false,
			emitUnpopulated: false,
			testChunkSize:   4 * 1024 * 1024,
			expectedJSON: `{
 "PROJECT": {
  "Name": {
   "Value": "_"
  },
  "LongIdentifier": {
   "Value": "_"
  }
 }
}`,
		},
		{
			tree: &a2l.RootNodeType{
				PROJECT: &a2l.ProjectType{
					Name:           &a2l.IdentType{Value: "_"},
					LongIdentifier: &a2l.StringType{Value: "_"},
				},
			},
			indent:          2,
			allowPartial:    false,
			emitUnpopulated: false,
			testChunkSize:   4 * 1024 * 1024,
			expectedJSON: `{
  "PROJECT": {
    "Name": {
      "Value": "_"
    },
    "LongIdentifier": {
      "Value": "_"
    }
  }
}`,
		},
		{
			tree: &a2l.RootNodeType{
				PROJECT: &a2l.ProjectType{
					Name:           &a2l.IdentType{Value: "_"},
					LongIdentifier: &a2l.StringType{Value: "_"},
				},
			},
			indent:          3,
			allowPartial:    false,
			emitUnpopulated: false,
			testChunkSize:   4 * 1024 * 1024,
			expectedJSON: `{
   "PROJECT": {
      "Name": {
         "Value": "_"
      },
      "LongIdentifier": {
         "Value": "_"
      }
   }
}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var serializedTree []byte
			var err error
			var receivedJSON []byte
			var response *a2l.JSONResponse
			var firstRequest *a2l.JSONFromTreeRequest
			grpc := newGrpcA2LImpl(4 * 1024 * 1024)

			mockStream := newMockStreamBase[*a2l.JSONFromTreeRequest, *a2l.JSONResponse]()
			defer mockStream.Close()

			// We test if a chunk is larger than the tree size
			if serializedTree, err = proto.Marshal(tc.tree); err == nil {
				// Send first chunk with specific option
				if len(serializedTree) <= tc.testChunkSize {
					firstRequest = &a2l.JSONFromTreeRequest{
						Tree:            serializedTree,
						Indent:          &tc.indent,
						AllowPartial:    &tc.allowPartial,
						EmitUnpopulated: &tc.emitUnpopulated,
					}
				} else {
					firstRequest = &a2l.JSONFromTreeRequest{
						Tree:            serializedTree[:tc.testChunkSize],
						Indent:          &tc.indent,
						AllowPartial:    &tc.allowPartial,
						EmitUnpopulated: &tc.emitUnpopulated,
					}
				}
				mockStream.SendRequest(firstRequest)

				for start := tc.testChunkSize; start < len(serializedTree); start += tc.testChunkSize {
					end := start + tc.testChunkSize
					if end > len(serializedTree) {
						end = len(serializedTree)
					}
					req := &a2l.JSONFromTreeRequest{
						Tree: serializedTree[start:end],
					}
					mockStream.SendRequest(req)
				}

				if err = grpc.GetJSONFromTree(mockStream); err == nil {
					for {
						response, err = mockStream.RecvResponse()
						if err != nil {
							if err == io.EOF {
								err = nil
							}
							break
						}
						receivedJSON = append(receivedJSON, response.Json...)
					}
					if err == nil {
						assert.Equal(t, tc.expectedJSON, string(receivedJSON))
					} else {
						t.Fatal(err)
					}
				} else {
					t.Fatal(err)
				}
			} else {
				t.Fatal(err)
			}
		})
	}
}

func TestGrpcA2LImplType_GetA2LFromTree(t *testing.T) {
	type testCaseType struct {
		tree          *a2l.RootNodeType
		indent        uint32
		sorted        bool
		expectedA2L   string
		testChunkSize int
	}

	for _, tc := range []testCaseType{
		{
			tree: &a2l.RootNodeType{
				PROJECT: &a2l.ProjectType{
					Name:           &a2l.IdentType{Value: "_"},
					LongIdentifier: &a2l.StringType{Value: "_"},
					MODULE: []*a2l.ModuleType{
						{
							Name:           &a2l.IdentType{Value: "_"},
							LongIdentifier: &a2l.StringType{Value: "_"},
							CHARACTERISTIC: []*a2l.CharacteristicType{
								{
									Name:           &a2l.IdentType{Value: "c[1]"},
									LongIdentifier: &a2l.StringType{Value: "_"},
									Type:           "VALUE",
									Address:        &a2l.LongType{Value: 0x00000000},
									Deposit:        &a2l.IdentType{Value: "_"},
									MaxDiff:        &a2l.FloatType{Value: 0.0},
									Conversion:     &a2l.IdentType{Value: "_"},
									LowerLimit:     &a2l.FloatType{Value: 0.0},
									UpperLimit:     &a2l.FloatType{Value: 0.0},
								},
								{
									Name:           &a2l.IdentType{Value: "b[1]"},
									LongIdentifier: &a2l.StringType{Value: "_"},
									Type:           "VALUE",
									Address:        &a2l.LongType{Value: 0x00000000},
									Deposit:        &a2l.IdentType{Value: "_"},
									MaxDiff:        &a2l.FloatType{Value: 0.0},
									Conversion:     &a2l.IdentType{Value: "_"},
									LowerLimit:     &a2l.FloatType{Value: 0.0},
									UpperLimit:     &a2l.FloatType{Value: 0.0},
								},
								{
									Name:           &a2l.IdentType{Value: "a[0]"},
									LongIdentifier: &a2l.StringType{Value: "_"},
									Type:           "VALUE",
									Address:        &a2l.LongType{Value: 0x00000000},
									Deposit:        &a2l.IdentType{Value: "_"},
									MaxDiff:        &a2l.FloatType{Value: 0.0},
									Conversion:     &a2l.IdentType{Value: "_"},
									LowerLimit:     &a2l.FloatType{Value: 0.0},
									UpperLimit:     &a2l.FloatType{Value: 0.0},
								},
								{
									Name:           &a2l.IdentType{Value: "c[0]"},
									LongIdentifier: &a2l.StringType{Value: "_"},
									Type:           "VALUE",
									Address:        &a2l.LongType{Value: 0x00000000},
									Deposit:        &a2l.IdentType{Value: "_"},
									MaxDiff:        &a2l.FloatType{Value: 0.0},
									Conversion:     &a2l.IdentType{Value: "_"},
									LowerLimit:     &a2l.FloatType{Value: 0.0},
									UpperLimit:     &a2l.FloatType{Value: 0.0},
								},
								{
									Name:           &a2l.IdentType{Value: "b[10]"},
									LongIdentifier: &a2l.StringType{Value: "_"},
									Type:           "VALUE",
									Address:        &a2l.LongType{Value: 0x00000000},
									Deposit:        &a2l.IdentType{Value: "_"},
									MaxDiff:        &a2l.FloatType{Value: 0.0},
									Conversion:     &a2l.IdentType{Value: "_"},
									LowerLimit:     &a2l.FloatType{Value: 0.0},
									UpperLimit:     &a2l.FloatType{Value: 0.0},
								},
							},
						},
					},
				},
			},
			indent:        1,
			sorted:        true,
			testChunkSize: 3,
			expectedA2L: `/begin PROJECT _ "_"
 /begin MODULE _ "_"
  /begin CHARACTERISTIC a[0] "_" VALUE 0x0 _ 0 _ 0 0
  /end CHARACTERISTIC
  /begin CHARACTERISTIC b[1] "_" VALUE 0x0 _ 0 _ 0 0
  /end CHARACTERISTIC
  /begin CHARACTERISTIC b[10] "_" VALUE 0x0 _ 0 _ 0 0
  /end CHARACTERISTIC
  /begin CHARACTERISTIC c[0] "_" VALUE 0x0 _ 0 _ 0 0
  /end CHARACTERISTIC
  /begin CHARACTERISTIC c[1] "_" VALUE 0x0 _ 0 _ 0 0
  /end CHARACTERISTIC
 /end MODULE
/end PROJECT`,
		},
		{
			tree: &a2l.RootNodeType{
				PROJECT: &a2l.ProjectType{
					Name:           &a2l.IdentType{Value: "_"},
					LongIdentifier: &a2l.StringType{Value: "_"},
					MODULE: []*a2l.ModuleType{
						{
							Name:           &a2l.IdentType{Value: "_"},
							LongIdentifier: &a2l.StringType{Value: "_"},
							CHARACTERISTIC: []*a2l.CharacteristicType{
								{
									Name:           &a2l.IdentType{Value: "c[1]"},
									LongIdentifier: &a2l.StringType{Value: "_"},
									Type:           "VALUE",
									Address:        &a2l.LongType{Value: 0x00000000},
									Deposit:        &a2l.IdentType{Value: "_"},
									MaxDiff:        &a2l.FloatType{Value: 0.0},
									Conversion:     &a2l.IdentType{Value: "_"},
									LowerLimit:     &a2l.FloatType{Value: 0.0},
									UpperLimit:     &a2l.FloatType{Value: 0.0},
								},
								{
									Name:           &a2l.IdentType{Value: "b[1]"},
									LongIdentifier: &a2l.StringType{Value: "_"},
									Type:           "VALUE",
									Address:        &a2l.LongType{Value: 0x00000000},
									Deposit:        &a2l.IdentType{Value: "_"},
									MaxDiff:        &a2l.FloatType{Value: 0.0},
									Conversion:     &a2l.IdentType{Value: "_"},
									LowerLimit:     &a2l.FloatType{Value: 0.0},
									UpperLimit:     &a2l.FloatType{Value: 0.0},
								},
								{
									Name:           &a2l.IdentType{Value: "a[0]"},
									LongIdentifier: &a2l.StringType{Value: "_"},
									Type:           "VALUE",
									Address:        &a2l.LongType{Value: 0x00000000},
									Deposit:        &a2l.IdentType{Value: "_"},
									MaxDiff:        &a2l.FloatType{Value: 0.0},
									Conversion:     &a2l.IdentType{Value: "_"},
									LowerLimit:     &a2l.FloatType{Value: 0.0},
									UpperLimit:     &a2l.FloatType{Value: 0.0},
								},
								{
									Name:           &a2l.IdentType{Value: "c[0]"},
									LongIdentifier: &a2l.StringType{Value: "_"},
									Type:           "VALUE",
									Address:        &a2l.LongType{Value: 0x00000000},
									Deposit:        &a2l.IdentType{Value: "_"},
									MaxDiff:        &a2l.FloatType{Value: 0.0},
									Conversion:     &a2l.IdentType{Value: "_"},
									LowerLimit:     &a2l.FloatType{Value: 0.0},
									UpperLimit:     &a2l.FloatType{Value: 0.0},
								},
								{
									Name:           &a2l.IdentType{Value: "b[10]"},
									LongIdentifier: &a2l.StringType{Value: "_"},
									Type:           "VALUE",
									Address:        &a2l.LongType{Value: 0x00000000},
									Deposit:        &a2l.IdentType{Value: "_"},
									MaxDiff:        &a2l.FloatType{Value: 0.0},
									Conversion:     &a2l.IdentType{Value: "_"},
									LowerLimit:     &a2l.FloatType{Value: 0.0},
									UpperLimit:     &a2l.FloatType{Value: 0.0},
								},
							},
						},
					},
				},
			},
			indent:        1,
			sorted:        true,
			testChunkSize: 4 * 1024 * 1024,
			expectedA2L: `/begin PROJECT _ "_"
 /begin MODULE _ "_"
  /begin CHARACTERISTIC a[0] "_" VALUE 0x0 _ 0 _ 0 0
  /end CHARACTERISTIC
  /begin CHARACTERISTIC b[1] "_" VALUE 0x0 _ 0 _ 0 0
  /end CHARACTERISTIC
  /begin CHARACTERISTIC b[10] "_" VALUE 0x0 _ 0 _ 0 0
  /end CHARACTERISTIC
  /begin CHARACTERISTIC c[0] "_" VALUE 0x0 _ 0 _ 0 0
  /end CHARACTERISTIC
  /begin CHARACTERISTIC c[1] "_" VALUE 0x0 _ 0 _ 0 0
  /end CHARACTERISTIC
 /end MODULE
/end PROJECT`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var serializedTree []byte
			var err error
			var receivedA2L []byte
			var response *a2l.A2LResponse
			var firstRequest *a2l.A2LFromTreeRequest
			grpc := newGrpcA2LImpl(4 * 1024 * 1024)

			mockStream := newMockStreamBase[*a2l.A2LFromTreeRequest, *a2l.A2LResponse]()
			defer mockStream.Close()

			// We test if a chunk is larger than the tree size
			if serializedTree, err = proto.Marshal(tc.tree); err == nil {
				// Send first chunk with specific option
				if len(serializedTree) <= tc.testChunkSize {
					firstRequest = &a2l.A2LFromTreeRequest{
						Tree:   serializedTree,
						Indent: &tc.indent,
						Sorted: &tc.sorted,
					}
				} else {
					firstRequest = &a2l.A2LFromTreeRequest{
						Tree:   serializedTree[:tc.testChunkSize],
						Indent: &tc.indent,
						Sorted: &tc.sorted,
					}
				}
				mockStream.SendRequest(firstRequest)

				for start := tc.testChunkSize; start < len(serializedTree); start += tc.testChunkSize {
					end := start + tc.testChunkSize
					if end > len(serializedTree) {
						end = len(serializedTree)
					}
					req := &a2l.A2LFromTreeRequest{
						Tree: serializedTree[start:end],
					}
					mockStream.SendRequest(req)
				}

				if err = grpc.GetA2LFromTree(mockStream); err == nil {
					for {
						response, err = mockStream.RecvResponse()
						if err != nil {
							if err == io.EOF {
								err = nil
							}
							break
						}
						receivedA2L = append(receivedA2L, response.A2L...)
					}
					if err == nil {
						assert.Equal(t, tc.expectedA2L, string(receivedA2L))
					} else {
						t.Fatal(err)
					}
				} else {
					t.Fatal(err)
				}
			} else {
				t.Fatal(err)
			}
		})
	}
}

// The responses carry the serialized tree in chunks of maxMsgSize - protocolSizeMargin bytes. A
// chunk size which is not positive would make the loop of chunkifyBySize spin forever or slice
// backwards, so Create rejects such a configuration and the helper itself does not split.
func Test_ChunkifyBySize(t *testing.T) {
	data := []byte("0123456789")

	t.Run("splits into chunks of the requested size", func(t *testing.T) {
		assert.Equal(t, [][]byte{[]byte("0123"), []byte("4567"), []byte("89")},
			chunkifyBySize(data, 4))
	})

	t.Run("a chunk size which is not positive does not split", func(t *testing.T) {
		for _, chunkSize := range []int{0, -1, -256} {
			assert.Equal(t, [][]byte{data}, chunkifyBySize(data, chunkSize),
				"chunk size %d", chunkSize)
		}
	})

	t.Run("no data yields no chunk", func(t *testing.T) {
		for _, chunkSize := range []int{4, 0, -1} {
			assert.Empty(t, chunkifyBySize(nil, chunkSize), "chunk size %d", chunkSize)
		}
	})
}

// createServer returns 0 on success and 1 on failure. It used to report success whenever no
// server was running yet, even when the port could not be bound, so that a caller checking the
// result went on to connect to a port nothing was listening on.
func Test_CreateServer(t *testing.T) {
	const port = 33331
	const maxMsgSize = 4 * 1024 * 1024

	t.Run("fails when the message size leaves no room for a chunk", func(t *testing.T) {
		for _, size := range []int{0, 1, protocolSizeMargin, -1} {
			assert.Equal(t, 1, createServer(port, size), "maximum message size %d", size)
			assert.Nil(t, server, "no server should have been started")
		}
	})

	t.Run("fails when the port cannot be bound", func(t *testing.T) {
		listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", port))
		if !assert.NoError(t, err) {
			return
		}
		defer func() { _ = listener.Close() }()

		assert.Equal(t, 1, createServer(port, maxMsgSize))
		assert.Nil(t, server, "no server should have been started")
	})

	t.Run("serves the loopback interface only", func(t *testing.T) {
		if !assert.Equal(t, 0, createServer(port, maxMsgSize)) {
			return
		}
		defer func() { assert.Equal(t, 0, closeServer()) }()

		// a second server cannot be started while the first one is running
		assert.Equal(t, 1, createServer(port+1, maxMsgSize))

		connection, err := net.DialTimeout("tcp", fmt.Sprintf("127.0.0.1:%v", port), time.Second)
		if assert.NoError(t, err, "the loopback interface should be served") {
			assert.NoError(t, connection.Close())
		}
	})

	t.Run("closeServer reports that no server is running", func(t *testing.T) {
		assert.Equal(t, 1, closeServer())
	})
}

// A tree which lacks an element the parser always fills, but which a client building a tree by
// hand may leave out, used to make the marshaller dereference a nil node and lose the stream to an
// internal error. The failure is reported through the error field of the response instead.
func Test_GetA2LFromTree_MalformedTreeReturnsError(t *testing.T) {
	tree := &a2l.RootNodeType{PROJECT: &a2l.ProjectType{}}

	serializedTree, err := proto.Marshal(tree)
	if !assert.NoError(t, err) {
		return
	}

	impl := newGrpcA2LImpl(4 * 1024 * 1024)

	mockStream := newMockStreamBase[*a2l.A2LFromTreeRequest, *a2l.A2LResponse]()
	defer mockStream.Close()

	mockStream.SendRequest(&a2l.A2LFromTreeRequest{Tree: serializedTree})

	if !assert.NoError(t, impl.GetA2LFromTree(mockStream)) {
		return
	}

	response, err := mockStream.RecvResponse()
	if !assert.NoError(t, err, "a response should have been sent") {
		return
	}

	if assert.NotNil(t, response.Error, "the failure should be reported in the error field") {
		assert.Contains(t, *response.Error, "serialization of the tree to A2L")
	}
}

// treeResponses runs GetTreeFromA2L on the passed content and returns every response of the stream.
func treeResponses(t *testing.T, impl *grpcA2LImplType, a2lString string) (result []*a2l.TreeResponse) {
	t.Helper()

	mockStream := newMockStreamBase[*a2l.TreeFromA2LRequest, *a2l.TreeResponse]()
	defer mockStream.Close()

	mockStream.SendRequest(&a2l.TreeFromA2LRequest{A2L: []byte(a2lString)})

	if err := impl.GetTreeFromA2L(mockStream); !assert.NoError(t, err) {
		return nil
	}

	for {
		response, err := mockStream.RecvResponse()
		if err != nil {
			return result
		}

		result = append(result, response)
	}
}

// Every response must fit the maximum message size the server was created with, whatever the file
// produces. The warnings used to ride on the first response, on top of a chunk of tree sized to
// fill it: a large file with many of them made gRPC reject the response with RESOURCE_EXHAUSTED.
func Test_GetTreeFromA2L_WarningsAreSpreadOverResponsesWhichFit(t *testing.T) {
	const maxMsgSize = 2048
	const measurements = 200

	// a repeated single occurrence keyword yields one warning per measurement
	var a2lString strings.Builder
	a2lString.WriteString("/begin PROJECT p \"\"\n/begin MODULE m \"\"\n")
	for i := 0; i < measurements; i++ {
		fmt.Fprintf(&a2lString, "/begin MEASUREMENT m%d \"\" UWORD cm 1 0 0 100\n"+
			"BYTE_ORDER MSB_LAST\nBYTE_ORDER MSB_FIRST\n/end MEASUREMENT\n", i)
	}
	a2lString.WriteString("/end MODULE\n/end PROJECT")

	responses := treeResponses(t, newGrpcA2LImpl(maxMsgSize), a2lString.String())
	if !assert.NotEmpty(t, responses) {
		return
	}

	warnings, chunks, firstChunk := 0, 0, -1
	var serializedTree []byte

	for i, response := range responses {
		assert.LessOrEqual(t, proto.Size(response), maxMsgSize, "response %d", i)
		assert.Nil(t, response.Error, "response %d", i)

		if len(response.SerializedTreeChunk) > 0 {
			if firstChunk < 0 {
				firstChunk = i
			}
			chunks++
			serializedTree = append(serializedTree, response.SerializedTreeChunk...)
			assert.Empty(t, response.Warnings, "a chunk fills its response, response %d", i)
		} else {
			assert.Less(t, firstChunk, 0, "the warnings precede the chunks, response %d", i)
			assert.NotEmpty(t, response.Warnings, "response %d carries nothing", i)
		}

		warnings += len(response.Warnings)
	}

	assert.Equal(t, measurements, warnings, "every warning should be delivered")
	assert.Greater(t, firstChunk, 1, "the warnings should need several responses at this size")
	assert.Greater(t, chunks, 0)

	tree := &a2l.RootNodeType{}
	if assert.NoError(t, proto.Unmarshal(serializedTree, tree)) {
		assert.Len(t, tree.PROJECT.MODULE[0].MEASUREMENT, measurements)
	}
}

// An error used to be sent in one piece whatever its length; a badly broken large file made gRPC
// reject it. It is shortened to its first lines, which name the cause, and counts the rest.
func Test_GetTreeFromA2L_ErrorIsBoundedToTheMaximumMessageSize(t *testing.T) {
	const maxMsgSize = 1024

	// every per cent sign is a token the lexer cannot recognise, and yields one error
	a2lString := "/begin PROJECT p \"\"\n" + strings.Repeat("%\n", 500) + "/end PROJECT"

	responses := treeResponses(t, newGrpcA2LImpl(maxMsgSize), a2lString)
	if !assert.Len(t, responses, 1) {
		return
	}

	response := responses[0]
	assert.LessOrEqual(t, proto.Size(response), maxMsgSize)

	if assert.NotNil(t, response.Error) {
		assert.Contains(t, *response.Error, "token recognition error", "the first errors are kept")
		assert.Regexp(t, `\.\.\. [0-9]+ further errors omitted$`, *response.Error)
	}
}

func Test_boundedError(t *testing.T) {
	// twenty errors of forty characters each, as the parser reports them: one per line
	lines := make([]string, 20)
	for i := range lines {
		lines[i] = fmt.Sprintf("%3d:1 %s", i, strings.Repeat("x", 34))
	}
	message := strings.Join(lines, "\n")

	t.Run("a message which fits is untouched", func(t *testing.T) {
		assert.Equal(t, message, boundedError(message, 1024))
	})

	t.Run("the first lines are kept and the rest is counted", func(t *testing.T) {
		const budget = 300

		bounded := boundedError(message, budget)
		assert.LessOrEqual(t, stringFieldSize(bounded), budget)

		parts := strings.Split(bounded, "\n")
		kept, tail := parts[:len(parts)-1], parts[len(parts)-1]

		if !assert.Greater(t, len(kept), 0) || !assert.Less(t, len(kept), len(lines)) {
			return
		}

		assert.Equal(t, lines[:len(kept)], kept, "the first lines are kept as they are")
		assert.Equal(t, fmt.Sprintf("... %d further errors omitted", len(lines)-len(kept)), tail)
	})

	t.Run("a single line longer than the budget is cut", func(t *testing.T) {
		bounded := boundedError(strings.Repeat("x", 100), 20)

		assert.True(t, strings.HasSuffix(bounded, "..."))
		assert.LessOrEqual(t, stringFieldSize(bounded), 20)
	})
}

// The responses are built to fit, so an oversize one is a bug of the server. It is reported through
// the error field instead of surfacing as a transport error on the client.
func Test_sendWithin_ReportsAnOversizeResponse(t *testing.T) {
	mockStream := newMockStreamBase[*a2l.TreeFromA2LRequest, *a2l.TreeResponse]()
	defer mockStream.Close()

	oversize := &a2l.TreeResponse{SerializedTreeChunk: make([]byte, 100)}
	if !assert.NoError(t, sendWithin(mockStream, oversize, 50, treeErrorResponse)) {
		return
	}

	response, err := mockStream.RecvResponse()
	if assert.NoError(t, err) && assert.NotNil(t, response.Error) {
		assert.Contains(t, *response.Error, "exceeds the maximum message size of 50 bytes")
		assert.Empty(t, response.SerializedTreeChunk)
	}
}
