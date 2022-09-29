protoc -I ./proto \
    --go_out ./proto --go_opt paths=source_relative \
    --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
    ./proto/helloworld/hello_world.proto

protoc -I ./proto \
    --go_out ./proto --go_opt paths=source_relative \
    --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
    ./proto/ping/ping.proto

# With gateway
protoc -I ./proto \
    --go_out ./proto --go_opt paths=source_relative \
    --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
    ./proto/helloworld/hello_world.proto

protoc -I ./proto \
    --go_out ./proto --go_opt paths=source_relative \
    --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
    ./proto/ping/ping.proto

protoc -I ./proto \
    --go_out ./proto --go_opt paths=source_relative \
    --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
    ./proto/register/register.proto