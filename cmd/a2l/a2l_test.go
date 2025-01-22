package main

import (
	"context"
	"github.com/sauci/a2l-grpc/pkg/a2l"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"io"
	"testing"
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
			grpc := grpcA2LImplType{}

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
			grpc := grpcA2LImplType{}

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
