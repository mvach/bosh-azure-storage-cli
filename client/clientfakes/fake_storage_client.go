// Code generated by counterfeiter. DO NOT EDIT.
package clientfakes

import (
	"context"
	"io"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
	"github.com/mvach/bosh-azure-storage-cli/client"
)

type FakeStorageClient struct {
	UploadStub        func(context.Context, io.ReadSeekCloser, *blockblob.UploadOptions) (client.StorageResponse, error)
	uploadMutex       sync.RWMutex
	uploadArgsForCall []struct {
		arg1 context.Context
		arg2 io.ReadSeekCloser
		arg3 *blockblob.UploadOptions
	}
	uploadReturns struct {
		result1 client.StorageResponse
		result2 error
	}
	uploadReturnsOnCall map[int]struct {
		result1 client.StorageResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStorageClient) Upload(arg1 context.Context, arg2 io.ReadSeekCloser, arg3 *blockblob.UploadOptions) (client.StorageResponse, error) {
	fake.uploadMutex.Lock()
	ret, specificReturn := fake.uploadReturnsOnCall[len(fake.uploadArgsForCall)]
	fake.uploadArgsForCall = append(fake.uploadArgsForCall, struct {
		arg1 context.Context
		arg2 io.ReadSeekCloser
		arg3 *blockblob.UploadOptions
	}{arg1, arg2, arg3})
	stub := fake.UploadStub
	fakeReturns := fake.uploadReturns
	fake.recordInvocation("Upload", []interface{}{arg1, arg2, arg3})
	fake.uploadMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStorageClient) UploadCallCount() int {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return len(fake.uploadArgsForCall)
}

func (fake *FakeStorageClient) UploadCalls(stub func(context.Context, io.ReadSeekCloser, *blockblob.UploadOptions) (client.StorageResponse, error)) {
	fake.uploadMutex.Lock()
	defer fake.uploadMutex.Unlock()
	fake.UploadStub = stub
}

func (fake *FakeStorageClient) UploadArgsForCall(i int) (context.Context, io.ReadSeekCloser, *blockblob.UploadOptions) {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	argsForCall := fake.uploadArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStorageClient) UploadReturns(result1 client.StorageResponse, result2 error) {
	fake.uploadMutex.Lock()
	defer fake.uploadMutex.Unlock()
	fake.UploadStub = nil
	fake.uploadReturns = struct {
		result1 client.StorageResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeStorageClient) UploadReturnsOnCall(i int, result1 client.StorageResponse, result2 error) {
	fake.uploadMutex.Lock()
	defer fake.uploadMutex.Unlock()
	fake.UploadStub = nil
	if fake.uploadReturnsOnCall == nil {
		fake.uploadReturnsOnCall = make(map[int]struct {
			result1 client.StorageResponse
			result2 error
		})
	}
	fake.uploadReturnsOnCall[i] = struct {
		result1 client.StorageResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeStorageClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStorageClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ client.StorageClient = new(FakeStorageClient)
