// Code generated by counterfeiter. DO NOT EDIT.
package mfpfakes

import (
	sync "sync"

	mfp "github.com/ryanmiville/amigobot/mfp"
)

type FakeFetcher struct {
	FetchStub        func(string) (*mfp.Diary, error)
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct {
		arg1 string
	}
	fetchReturns struct {
		result1 *mfp.Diary
		result2 error
	}
	fetchReturnsOnCall map[int]struct {
		result1 *mfp.Diary
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFetcher) Fetch(arg1 string) (*mfp.Diary, error) {
	fake.fetchMutex.Lock()
	ret, specificReturn := fake.fetchReturnsOnCall[len(fake.fetchArgsForCall)]
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Fetch", []interface{}{arg1})
	fake.fetchMutex.Unlock()
	if fake.FetchStub != nil {
		return fake.FetchStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.fetchReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFetcher) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *FakeFetcher) FetchCalls(stub func(string) (*mfp.Diary, error)) {
	fake.fetchMutex.Lock()
	defer fake.fetchMutex.Unlock()
	fake.FetchStub = stub
}

func (fake *FakeFetcher) FetchArgsForCall(i int) string {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	argsForCall := fake.fetchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFetcher) FetchReturns(result1 *mfp.Diary, result2 error) {
	fake.fetchMutex.Lock()
	defer fake.fetchMutex.Unlock()
	fake.FetchStub = nil
	fake.fetchReturns = struct {
		result1 *mfp.Diary
		result2 error
	}{result1, result2}
}

func (fake *FakeFetcher) FetchReturnsOnCall(i int, result1 *mfp.Diary, result2 error) {
	fake.fetchMutex.Lock()
	defer fake.fetchMutex.Unlock()
	fake.FetchStub = nil
	if fake.fetchReturnsOnCall == nil {
		fake.fetchReturnsOnCall = make(map[int]struct {
			result1 *mfp.Diary
			result2 error
		})
	}
	fake.fetchReturnsOnCall[i] = struct {
		result1 *mfp.Diary
		result2 error
	}{result1, result2}
}

func (fake *FakeFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFetcher) recordInvocation(key string, args []interface{}) {
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

var _ mfp.Fetcher = new(FakeFetcher)
