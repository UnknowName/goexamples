## gRPC use example

## Use Step

### Step 1

    definition protocol file

### step 2

generate `pb.go` and  `_gprc.pb.go`

```bash
 cd grpc/proto
 protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative data.proto
 protoc --go_out=. --go_opt=paths=source_relative data.proto
```

### step 3

    write go code, implement the method of service in proto file

    

