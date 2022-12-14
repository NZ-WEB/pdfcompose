package main

import (
	"fmt"
	"github.com/NZ-WEB/pdfcompose/internal/handlers"
	"github.com/NZ-WEB/pdfcompose/pkg/pdfservice"
	"google.golang.org/grpc"
	"net"
)

func main() {
	fmt.Println("Starting server ...")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	server := &handlers.PdfServiceServer{}
	pdfservice.RegisterPdfServiceServer(s, server)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
