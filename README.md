# Package description
This package aims to provide a fast, language-agnostic [A2L](https://www.asam.net/standards/detail/mcd-2-mc/) file 
parser.

---

# Implementation
The parser/lexer code is generated using [*ANTLR4*](https://www.antlr.org/). Once the AST has been generated, the 
package generates a nested structure tree using [*gRPC*](https://grpc.io/) messages. The idea behind this is to provide 
the AST to any client written in [*any language*](https://grpc.io/docs/languages/), through a well-documented, 
supported and structured message format.

---

# Example
Let's say you must parse the following A2L content using *Python*:
```
ASAP2_VERSION 1 51
/begin PROJECT my_project_name "my_project_long_id"
/end PROJECT
```
## Getting the pre-built artefacts
The first thing to do is to retrieve the pre-built artefacts from the 
[releases page](https://github.com/Sauci/a2l-grpc/releases). To do so, follow 
[this link](https://github.com/Sauci/a2l-grpc/releases/latest) and download the file named `a2l_grpc.tar.gz`. This 
archive contains the following files/directories:
```
a2l_grpc.tar.gz
├── protobuf
│   ├── A2L.proto
│   ├── A2ML.proto
│   ├── API.proto
│   ├── IF_DATA.proto
│   └── shared.proto
├── a2l_grpc_(os)_(arch).h
├── a2l_grpc_(os)_(arch).dll
├── a2l_grpc_(os)_(arch).dylib
└── a2l_grpc_(os)_(arch).so
```
The `a2l_grpc.tar.gz/protobuf` folder contains all the *gRPC* data structure definitions. It will be used to generate
the structure definitions, as well as their serialization/deserialization methods for the desired language (*Python* in
this case).

The `a2l_grpc_(os)_(arch).h` file contains the API definition of the corresponding *Dynamic Link Library*/*Shared 
Object*.

Finally, the `a2l_grpc_(os)_(arch).dll`/`.so`/`.dylib` file serves this API.

## Integrating the pre-built artefacts
Now that we have the artefacts, we need to integrate them into the *Python* project. Simply extract the archive and 
place it into the project to have the following file/directory structure:
```
src
├── a2l_grpc
│   ├── protobuf
│   │   ├── A2L.proto
│   │   ├── A2ML.proto
│   │   ├── API.proto
│   │   ├── IF_DATA.proto
│   │   └── shared.proto
│   ├── a2l_grpc_(os)_(arch).h
│   ├── a2l_grpc_(os)_(arch).dll
│   ├── a2l_grpc_(os)_(arch).dylib
│   └── a2l_grpc_(os)_(arch).so
└── main.py
```
## Generating the *gRPC* sources
The next step is to generate the *gRPC* sources for our language, in this case *Python*. To do that, we first need to 
install the code generator for *Python* with the command `python -m pip install grpcio-tools`

Then, we generate the sources with the following command (running from the `src` directory):

`python -m grpc_tools.protoc -I./a2l_grpc --python_out=. --pyi_out=. --grpc_python_out=. ./a2l_grpc/protobuf/*.proto`

This command will create a `protobuf` directory in the `src` folder, containing all the *Python* sources.

## Writing the *gRPC* client
Now that all sources are available, we can start writing the *gRPC* client in *Python*. Note that all the *gRPC* 
methods are bidirectional streams: the A2L content is sent as a sequence of requests, and the resulting tree is 
received as a sequence of chunks which must be concatenated before being deserialized. Here is the code:
```python
import ctypes
import os
import sys

import grpc

from protobuf.A2L_pb2 import RootNodeType
from protobuf.API_pb2 import TreeFromA2LRequest
from protobuf.API_pb2_grpc import A2LStub

PORT = 3333
MAX_MESSAGE_SIZE = 4 * 1024 * 1024
CHUNK_SIZE = MAX_MESSAGE_SIZE - 1024


def get_shared_object_name() -> str:
    if os.name == 'nt':
        shared_object = 'a2l_grpc_windows_amd64.dll'
    elif os.name == 'posix':
        if sys.platform == 'darwin':
            shared_object = 'a2l_grpc_darwin_arm64.dylib'
        else:
            shared_object = 'a2l_grpc_linux_amd64.so'
    else:
        raise Exception(f'unsupported operating system {os.name}')
    return shared_object


class A2lParser(object):
    def __init__(self, string):
        self._dll = ctypes.cdll.LoadLibrary(
            os.path.join(os.path.dirname(__file__), 'a2l_grpc', get_shared_object_name()))
        if self._dll.Create(PORT, MAX_MESSAGE_SIZE):
            raise Exception('unable to start the gRPC server')

        self._channel = grpc.insecure_channel(f'localhost:{PORT}', options=[
            ('grpc.max_receive_message_length', MAX_MESSAGE_SIZE),
            ('grpc.max_send_message_length', MAX_MESSAGE_SIZE)])
        client = A2LStub(self._channel)

        payload = string.encode()
        requests = (TreeFromA2LRequest(a2l=payload[i:i + CHUNK_SIZE])
                    for i in range(0, len(payload), CHUNK_SIZE))

        serialized_tree = b''
        self.warnings = []
        for response in client.GetTreeFromA2L(requests):
            if response.HasField('error'):
                raise Exception(response.error)
            self.warnings.extend(response.warnings)
            serialized_tree += response.serializedTreeChunk

        self.ast = RootNodeType()
        self.ast.ParseFromString(serialized_tree)

    def close(self):
        self._channel.close()
        self._dll.Close()


if __name__ == '__main__':
    a2l_string = """ASAP2_VERSION 1 51
/begin PROJECT my_project_name "my_project_long_id"
/end PROJECT"""

    p = A2lParser(a2l_string)
    print(f'ASAP2 version number = {p.ast.ASAP2_VERSION.VersionNo.Value}')
    print(f'ASAP2 upgrade number = {p.ast.ASAP2_VERSION.UpgradeNo.Value}')
    print(p.ast)
    p.close()
```
Running this code will produce the following output:
```
ASAP2 version number = 1
ASAP2 upgrade number = 51
ASAP2_VERSION {
  VersionNo {
    Value: 1
    Base: 10
    Size: 1
  }
  UpgradeNo {
    Value: 51
    Base: 10
    Size: 2
  }
}
PROJECT {
  Name {
    Value: "my_project_name"
  }
  LongIdentifier {
    Value: "my_project_long_id"
  }
}
```

---

# Hints
- As A2L files are often quite large, it might happen that *gRPC* raises an error because the message size is too 
  large. The maximum message size accepted by the server is the second argument of `Create`, and it also defines the 
  size of the chunks the server sends back. On the client side, the same limit must be set with the 
  `grpc.max_receive_message_length` and `grpc.max_send_message_length` options while instantiating the *gRPC* channel.
- As shown in the above example, the numerical values are held in the *Value* field. The other fields hold metadata in 
  case the value must be dumped. In the case of the *UpgradeNo* for instance, the *Base* field says that the original 
  value was defined in numerical base 10, and the *Size* field says that it has 2 digits.
- A string is held in the *Value* field as it is written in the file, without its delimiters: the escape sequences 
  (`\"`, `\'`, `\\`, `\n`, `\r`, `\t` and a doubled `""`) are left in place, so that the value is reproduced verbatim 
  when the tree is serialized back to A2L. A client which needs the characters the string stands for, rather than its 
  source form, has to resolve the sequences itself: `"a \"b\""` is delivered as `a \"b\"`.
- The parser is given the content of a single file and has no access to a file system, so it cannot resolve an 
  `/include` statement. Each of them is reported as an error naming the file it refers to; substituting the content of 
  the included files is up to the caller.
- `TreeFromA2LRequest` accepts an optional `enforce_version_check` flag. When set, keywords requiring a newer ASAP2 
  version than the one declared by the file are rejected as errors; otherwise they are reported in the `warnings` field 
  of the first response of the stream.
