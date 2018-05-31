# Build the net messages with protocol buffers

protoc --gofast_out=. connect_response.proto
protoc --gofast_out=. update_player.proto
protoc --gofast_out=. disconnect.proto
