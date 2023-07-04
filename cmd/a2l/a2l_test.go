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
