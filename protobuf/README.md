# Battery of questions supported with code about the use of Protobuf and Go

# Environment



## Q: How does protobuf handle versions?

The following scenario assumes the following files:

`proto/message.proto`:

```protobuf
syntax = "proto3";

package messaging;

import "google/protobuf/timestamp.proto";

message Transcript {
  google.protobuf.Timestamp received_at = 1;
  string body = 2;
}
```

[run protobuf compiler](https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers)

```console
protoc -I=proto --go_out=lib/ proto/*.proto
```

Now with this as baseline let's alter the proto file and use git diff to analyze the changes to the generated code.

### Q: What happens when we add a field?
