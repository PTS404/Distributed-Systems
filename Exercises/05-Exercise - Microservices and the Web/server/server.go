package main

import (
	proto "gRPCIntro/grpc"
)

type Server struct {
	proto.AskForCourseMessage
	name string
	port int
}
