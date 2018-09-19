// Code generated by counterfeiter. DO NOT EDIT.
package containerserverfakes

import (
	sync "sync"

	containerserver "github.com/concourse/concourse/atc/api/containerserver"
)

type FakeInterceptTimeoutFactory struct {
	NewInterceptTimeoutStub        func() containerserver.InterceptTimeout
	newInterceptTimeoutMutex       sync.RWMutex
	newInterceptTimeoutArgsForCall []struct {
	}
	newInterceptTimeoutReturns struct {
		result1 containerserver.InterceptTimeout
	}
	newInterceptTimeoutReturnsOnCall map[int]struct {
		result1 containerserver.InterceptTimeout
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInterceptTimeoutFactory) NewInterceptTimeout() containerserver.InterceptTimeout {
	fake.newInterceptTimeoutMutex.Lock()
	ret, specificReturn := fake.newInterceptTimeoutReturnsOnCall[len(fake.newInterceptTimeoutArgsForCall)]
	fake.newInterceptTimeoutArgsForCall = append(fake.newInterceptTimeoutArgsForCall, struct {
	}{})
	fake.recordInvocation("NewInterceptTimeout", []interface{}{})
	fake.newInterceptTimeoutMutex.Unlock()
	if fake.NewInterceptTimeoutStub != nil {
		return fake.NewInterceptTimeoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newInterceptTimeoutReturns
	return fakeReturns.result1
}

func (fake *FakeInterceptTimeoutFactory) NewInterceptTimeoutCallCount() int {
	fake.newInterceptTimeoutMutex.RLock()
	defer fake.newInterceptTimeoutMutex.RUnlock()
	return len(fake.newInterceptTimeoutArgsForCall)
}

func (fake *FakeInterceptTimeoutFactory) NewInterceptTimeoutReturns(result1 containerserver.InterceptTimeout) {
	fake.NewInterceptTimeoutStub = nil
	fake.newInterceptTimeoutReturns = struct {
		result1 containerserver.InterceptTimeout
	}{result1}
}

func (fake *FakeInterceptTimeoutFactory) NewInterceptTimeoutReturnsOnCall(i int, result1 containerserver.InterceptTimeout) {
	fake.NewInterceptTimeoutStub = nil
	if fake.newInterceptTimeoutReturnsOnCall == nil {
		fake.newInterceptTimeoutReturnsOnCall = make(map[int]struct {
			result1 containerserver.InterceptTimeout
		})
	}
	fake.newInterceptTimeoutReturnsOnCall[i] = struct {
		result1 containerserver.InterceptTimeout
	}{result1}
}

func (fake *FakeInterceptTimeoutFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newInterceptTimeoutMutex.RLock()
	defer fake.newInterceptTimeoutMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInterceptTimeoutFactory) recordInvocation(key string, args []interface{}) {
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

var _ containerserver.InterceptTimeoutFactory = new(FakeInterceptTimeoutFactory)
