package main

import (
	"github.com/sauci/a2l-grpc/pkg/a2l"
	"github.com/stretchr/testify/assert"
	"testing"
)

func baseTypeToPointer[T bool | uint32 | string](baseType T) *T {
	return &baseType
}

func Test_GetJSONFromTree(t *testing.T) {
	type testCaseType struct {
		request  *a2l.JSONFromTreeRequest
		response *a2l.JSONResponse
	}

	for _, tc := range []testCaseType{
		{
			request: &a2l.JSONFromTreeRequest{
				Tree: &a2l.RootNodeType{
					PROJECT: &a2l.ProjectType{
						Name:           &a2l.IdentType{Value: "_"},
						LongIdentifier: &a2l.StringType{Value: "_"},
					},
				},
				Indent:          baseTypeToPointer[uint32](1),
				AllowPartial:    baseTypeToPointer[bool](false),
				EmitUnpopulated: baseTypeToPointer[bool](false),
			},
			response: &a2l.JSONResponse{
				Json: []byte(`{
 "PROJECT": {
  "Name": {
   "Value": "_"
  },
  "LongIdentifier": {
   "Value": "_"
  }
 }
}`),
				Error: nil,
			},
		},
		{
			request: &a2l.JSONFromTreeRequest{
				Tree: &a2l.RootNodeType{
					PROJECT: &a2l.ProjectType{
						Name:           &a2l.IdentType{Value: "_"},
						LongIdentifier: &a2l.StringType{Value: "_"},
					},
				},
				Indent:          baseTypeToPointer[uint32](2),
				AllowPartial:    baseTypeToPointer[bool](false),
				EmitUnpopulated: baseTypeToPointer[bool](false),
			},
			response: &a2l.JSONResponse{
				Json: []byte(`{
  "PROJECT": {
    "Name": {
      "Value": "_"
    },
    "LongIdentifier": {
      "Value": "_"
    }
  }
}`),
				Error: nil,
			},
		},
		{
			request: &a2l.JSONFromTreeRequest{
				Tree: &a2l.RootNodeType{
					PROJECT: &a2l.ProjectType{
						Name:           &a2l.IdentType{Value: "_"},
						LongIdentifier: &a2l.StringType{Value: "_"},
					},
				},
				Indent:          baseTypeToPointer[uint32](3),
				AllowPartial:    baseTypeToPointer[bool](false),
				EmitUnpopulated: baseTypeToPointer[bool](false),
			},
			response: &a2l.JSONResponse{
				Json: []byte(`{
   "PROJECT": {
      "Name": {
         "Value": "_"
      },
      "LongIdentifier": {
         "Value": "_"
      }
   }
}`),
				Error: nil,
			},
		},
	} {
		t.Run("", func(t *testing.T) {
			grpc := grpcA2LImplType{}

			if response, err := grpc.GetJSONFromTree(nil, tc.request); err == nil {
				assert.Equal(t, tc.response, response)
			} else {
				t.Fatal(err)
			}
		})
	}
}

func TestGrpcA2LImplType_GetA2LFromTree(t *testing.T) {
	type testCaseType struct {
		request  *a2l.A2LFromTreeRequest
		response *a2l.A2LResponse
	}

	for _, tc := range []testCaseType{
		{
			request: &a2l.A2LFromTreeRequest{
				Tree: &a2l.RootNodeType{
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
				Indent: baseTypeToPointer[uint32](1),
				Sorted: baseTypeToPointer[bool](true),
			},
			response: &a2l.A2LResponse{
				A2L: []byte(`/begin PROJECT _ "_"
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
/end PROJECT`),
				Error: nil,
			},
		},
	} {
		t.Run("", func(t *testing.T) {
			grpc := grpcA2LImplType{}

			if response, err := grpc.GetA2LFromTree(nil, tc.request); err == nil {
				assert.Equal(t, tc.response, response)
			} else {
				t.Fatal(err)
			}
		})
	}
}
