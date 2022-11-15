package client

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
	"io"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . StorageClient
type StorageClient interface {
	Upload(
		context context.Context,
		body io.ReadSeekCloser,
		options *blockblob.UploadOptions) (blockblob.UploadResponse, error)
}

type DefaultStorageClient struct {
	azClient blockblob.Client
}

func NewStorageClient(azClient blockblob.Client) (StorageClient, error) {
	return DefaultStorageClient{azClient: azClient}, nil
}

func (dsc DefaultStorageClient) Upload(
	context context.Context,
	body io.ReadSeekCloser,
	options *blockblob.UploadOptions) (blockblob.UploadResponse, error) {

	response, err := dsc.azClient.Upload(context, body, options)
	return response, err
}
