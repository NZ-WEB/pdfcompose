package main

import (
	"bytes"
	"context"
	"github.com/NZ-WEB/pdfcompose/pkg/pdfservice"
	"google.golang.org/grpc"
	"io"
	"time"
)

func main() {
	cwt, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := grpc.DialContext(cwt, "localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	pdfS := pdfservice.NewPdfServiceClient(conn)

	body := &pdfservice.PostSendParams{
		Upfile1: []byte("1 param"),
		Upfile2: []byte("2 param"),
		Upfile3: []byte("3 param"),
	}

	res, err := pdfS.Send(cwt, body)
	if err != nil {
		panic(err)
	}

	resReader := bytes.NewReader(res.Payload)
	resFile := io.NopCloser(resReader)

	println(resFile)
	return
}
