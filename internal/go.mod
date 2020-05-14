module github.com/bubble-hook/lightnet-assignment/internal

require (
	github.com/bubble-hook/lightnet-assignment/calculator_proto v0.0.0
	github.com/bubble-hook/lightnet-assignment/shared v0.0.0
	google.golang.org/grpc v1.29.1
)

replace github.com/bubble-hook/lightnet-assignment/shared => ../shared

replace github.com/bubble-hook/lightnet-assignment/calculator_proto => ../calculator_proto

go 1.12
