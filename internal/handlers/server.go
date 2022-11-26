package handlers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/NZ-WEB/pdfcompose/pkg/composer"
	"github.com/NZ-WEB/pdfcompose/pkg/pdfservice"
	"io"
)

type PdfServiceServer struct {
	pdfservice.UnimplementedPdfServiceServer
}

func (s *PdfServiceServer) Send(ctx context.Context, params *pdfservice.PostSendParams) (*pdfservice.PostSendResponse, error) {
	fmt.Printf("Params: %v+\n", params)

	param1Reader := bytes.NewReader(params.Upfile1)
	param1ReadCloser := io.NopCloser(param1Reader)

	file, err := composer.ComposeFromFiles([]io.ReadCloser{param1ReadCloser})
	if err != nil {
		return nil, err
	}
	defer file.Close()
	res, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return &pdfservice.PostSendResponse{Payload: res}, nil
}
