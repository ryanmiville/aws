// Code generated by counterfeiter. DO NOT EDIT.
package dynamodbfakes

import (
	"context"
	"sync"

	"github.com/ryanmiville/aws/dynamodb"
	"github.com/aws/aws-sdk-go/aws/request"
	dynamodba "github.com/aws/aws-sdk-go/service/dynamodb"
)

type FakeDynamoDBClient struct {
	PutItemWithContextStub        func(context.Context, *dynamodba.PutItemInput, ...request.Option) (*dynamodba.PutItemOutput, error)
	putItemWithContextMutex       sync.RWMutex
	putItemWithContextArgsForCall []struct {
		arg1 context.Context
		arg2 *dynamodba.PutItemInput
		arg3 []request.Option
	}
	putItemWithContextReturns struct {
		result1 *dynamodba.PutItemOutput
		result2 error
	}
	putItemWithContextReturnsOnCall map[int]struct {
		result1 *dynamodba.PutItemOutput
		result2 error
	}
	QueryWithContextStub        func(context.Context, *dynamodba.QueryInput, ...request.Option) (*dynamodba.QueryOutput, error)
	queryWithContextMutex       sync.RWMutex
	queryWithContextArgsForCall []struct {
		arg1 context.Context
		arg2 *dynamodba.QueryInput
		arg3 []request.Option
	}
	queryWithContextReturns struct {
		result1 *dynamodba.QueryOutput
		result2 error
	}
	queryWithContextReturnsOnCall map[int]struct {
		result1 *dynamodba.QueryOutput
		result2 error
	}
	ScanWithContextStub        func(context.Context, *dynamodba.ScanInput, ...request.Option) (*dynamodba.ScanOutput, error)
	scanWithContextMutex       sync.RWMutex
	scanWithContextArgsForCall []struct {
		arg1 context.Context
		arg2 *dynamodba.ScanInput
		arg3 []request.Option
	}
	scanWithContextReturns struct {
		result1 *dynamodba.ScanOutput
		result2 error
	}
	scanWithContextReturnsOnCall map[int]struct {
		result1 *dynamodba.ScanOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDynamoDBClient) PutItemWithContext(arg1 context.Context, arg2 *dynamodba.PutItemInput, arg3 ...request.Option) (*dynamodba.PutItemOutput, error) {
	fake.putItemWithContextMutex.Lock()
	ret, specificReturn := fake.putItemWithContextReturnsOnCall[len(fake.putItemWithContextArgsForCall)]
	fake.putItemWithContextArgsForCall = append(fake.putItemWithContextArgsForCall, struct {
		arg1 context.Context
		arg2 *dynamodba.PutItemInput
		arg3 []request.Option
	}{arg1, arg2, arg3})
	fake.recordInvocation("PutItemWithContext", []interface{}{arg1, arg2, arg3})
	fake.putItemWithContextMutex.Unlock()
	if fake.PutItemWithContextStub != nil {
		return fake.PutItemWithContextStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.putItemWithContextReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDynamoDBClient) PutItemWithContextCallCount() int {
	fake.putItemWithContextMutex.RLock()
	defer fake.putItemWithContextMutex.RUnlock()
	return len(fake.putItemWithContextArgsForCall)
}

func (fake *FakeDynamoDBClient) PutItemWithContextCalls(stub func(context.Context, *dynamodba.PutItemInput, ...request.Option) (*dynamodba.PutItemOutput, error)) {
	fake.putItemWithContextMutex.Lock()
	defer fake.putItemWithContextMutex.Unlock()
	fake.PutItemWithContextStub = stub
}

func (fake *FakeDynamoDBClient) PutItemWithContextArgsForCall(i int) (context.Context, *dynamodba.PutItemInput, []request.Option) {
	fake.putItemWithContextMutex.RLock()
	defer fake.putItemWithContextMutex.RUnlock()
	argsForCall := fake.putItemWithContextArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDynamoDBClient) PutItemWithContextReturns(result1 *dynamodba.PutItemOutput, result2 error) {
	fake.putItemWithContextMutex.Lock()
	defer fake.putItemWithContextMutex.Unlock()
	fake.PutItemWithContextStub = nil
	fake.putItemWithContextReturns = struct {
		result1 *dynamodba.PutItemOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeDynamoDBClient) PutItemWithContextReturnsOnCall(i int, result1 *dynamodba.PutItemOutput, result2 error) {
	fake.putItemWithContextMutex.Lock()
	defer fake.putItemWithContextMutex.Unlock()
	fake.PutItemWithContextStub = nil
	if fake.putItemWithContextReturnsOnCall == nil {
		fake.putItemWithContextReturnsOnCall = make(map[int]struct {
			result1 *dynamodba.PutItemOutput
			result2 error
		})
	}
	fake.putItemWithContextReturnsOnCall[i] = struct {
		result1 *dynamodba.PutItemOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeDynamoDBClient) QueryWithContext(arg1 context.Context, arg2 *dynamodba.QueryInput, arg3 ...request.Option) (*dynamodba.QueryOutput, error) {
	fake.queryWithContextMutex.Lock()
	ret, specificReturn := fake.queryWithContextReturnsOnCall[len(fake.queryWithContextArgsForCall)]
	fake.queryWithContextArgsForCall = append(fake.queryWithContextArgsForCall, struct {
		arg1 context.Context
		arg2 *dynamodba.QueryInput
		arg3 []request.Option
	}{arg1, arg2, arg3})
	fake.recordInvocation("QueryWithContext", []interface{}{arg1, arg2, arg3})
	fake.queryWithContextMutex.Unlock()
	if fake.QueryWithContextStub != nil {
		return fake.QueryWithContextStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryWithContextReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDynamoDBClient) QueryWithContextCallCount() int {
	fake.queryWithContextMutex.RLock()
	defer fake.queryWithContextMutex.RUnlock()
	return len(fake.queryWithContextArgsForCall)
}

func (fake *FakeDynamoDBClient) QueryWithContextCalls(stub func(context.Context, *dynamodba.QueryInput, ...request.Option) (*dynamodba.QueryOutput, error)) {
	fake.queryWithContextMutex.Lock()
	defer fake.queryWithContextMutex.Unlock()
	fake.QueryWithContextStub = stub
}

func (fake *FakeDynamoDBClient) QueryWithContextArgsForCall(i int) (context.Context, *dynamodba.QueryInput, []request.Option) {
	fake.queryWithContextMutex.RLock()
	defer fake.queryWithContextMutex.RUnlock()
	argsForCall := fake.queryWithContextArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDynamoDBClient) QueryWithContextReturns(result1 *dynamodba.QueryOutput, result2 error) {
	fake.queryWithContextMutex.Lock()
	defer fake.queryWithContextMutex.Unlock()
	fake.QueryWithContextStub = nil
	fake.queryWithContextReturns = struct {
		result1 *dynamodba.QueryOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeDynamoDBClient) QueryWithContextReturnsOnCall(i int, result1 *dynamodba.QueryOutput, result2 error) {
	fake.queryWithContextMutex.Lock()
	defer fake.queryWithContextMutex.Unlock()
	fake.QueryWithContextStub = nil
	if fake.queryWithContextReturnsOnCall == nil {
		fake.queryWithContextReturnsOnCall = make(map[int]struct {
			result1 *dynamodba.QueryOutput
			result2 error
		})
	}
	fake.queryWithContextReturnsOnCall[i] = struct {
		result1 *dynamodba.QueryOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeDynamoDBClient) ScanWithContext(arg1 context.Context, arg2 *dynamodba.ScanInput, arg3 ...request.Option) (*dynamodba.ScanOutput, error) {
	fake.scanWithContextMutex.Lock()
	ret, specificReturn := fake.scanWithContextReturnsOnCall[len(fake.scanWithContextArgsForCall)]
	fake.scanWithContextArgsForCall = append(fake.scanWithContextArgsForCall, struct {
		arg1 context.Context
		arg2 *dynamodba.ScanInput
		arg3 []request.Option
	}{arg1, arg2, arg3})
	fake.recordInvocation("ScanWithContext", []interface{}{arg1, arg2, arg3})
	fake.scanWithContextMutex.Unlock()
	if fake.ScanWithContextStub != nil {
		return fake.ScanWithContextStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.scanWithContextReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDynamoDBClient) ScanWithContextCallCount() int {
	fake.scanWithContextMutex.RLock()
	defer fake.scanWithContextMutex.RUnlock()
	return len(fake.scanWithContextArgsForCall)
}

func (fake *FakeDynamoDBClient) ScanWithContextCalls(stub func(context.Context, *dynamodba.ScanInput, ...request.Option) (*dynamodba.ScanOutput, error)) {
	fake.scanWithContextMutex.Lock()
	defer fake.scanWithContextMutex.Unlock()
	fake.ScanWithContextStub = stub
}

func (fake *FakeDynamoDBClient) ScanWithContextArgsForCall(i int) (context.Context, *dynamodba.ScanInput, []request.Option) {
	fake.scanWithContextMutex.RLock()
	defer fake.scanWithContextMutex.RUnlock()
	argsForCall := fake.scanWithContextArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDynamoDBClient) ScanWithContextReturns(result1 *dynamodba.ScanOutput, result2 error) {
	fake.scanWithContextMutex.Lock()
	defer fake.scanWithContextMutex.Unlock()
	fake.ScanWithContextStub = nil
	fake.scanWithContextReturns = struct {
		result1 *dynamodba.ScanOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeDynamoDBClient) ScanWithContextReturnsOnCall(i int, result1 *dynamodba.ScanOutput, result2 error) {
	fake.scanWithContextMutex.Lock()
	defer fake.scanWithContextMutex.Unlock()
	fake.ScanWithContextStub = nil
	if fake.scanWithContextReturnsOnCall == nil {
		fake.scanWithContextReturnsOnCall = make(map[int]struct {
			result1 *dynamodba.ScanOutput
			result2 error
		})
	}
	fake.scanWithContextReturnsOnCall[i] = struct {
		result1 *dynamodba.ScanOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeDynamoDBClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.putItemWithContextMutex.RLock()
	defer fake.putItemWithContextMutex.RUnlock()
	fake.queryWithContextMutex.RLock()
	defer fake.queryWithContextMutex.RUnlock()
	fake.scanWithContextMutex.RLock()
	defer fake.scanWithContextMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDynamoDBClient) recordInvocation(key string, args []interface{}) {
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

var _ dynamodb.DynamoDBClient = new(FakeDynamoDBClient)