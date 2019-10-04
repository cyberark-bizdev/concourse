// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"
	"time"

	"github.com/concourse/concourse/atc/db"
)

type FakeComponentFactory struct {
	FindStub        func(string) (db.Component, bool, error)
	findMutex       sync.RWMutex
	findArgsForCall []struct {
		arg1 string
	}
	findReturns struct {
		result1 db.Component
		result2 bool
		result3 error
	}
	findReturnsOnCall map[int]struct {
		result1 db.Component
		result2 bool
		result3 error
	}
	UpdateIntervalsStub        func(map[string]time.Duration) error
	updateIntervalsMutex       sync.RWMutex
	updateIntervalsArgsForCall []struct {
		arg1 map[string]time.Duration
	}
	updateIntervalsReturns struct {
		result1 error
	}
	updateIntervalsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeComponentFactory) Find(arg1 string) (db.Component, bool, error) {
	fake.findMutex.Lock()
	ret, specificReturn := fake.findReturnsOnCall[len(fake.findArgsForCall)]
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Find", []interface{}{arg1})
	fake.findMutex.Unlock()
	if fake.FindStub != nil {
		return fake.FindStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.findReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeComponentFactory) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *FakeComponentFactory) FindCalls(stub func(string) (db.Component, bool, error)) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = stub
}

func (fake *FakeComponentFactory) FindArgsForCall(i int) string {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	argsForCall := fake.findArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeComponentFactory) FindReturns(result1 db.Component, result2 bool, result3 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 db.Component
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeComponentFactory) FindReturnsOnCall(i int, result1 db.Component, result2 bool, result3 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	if fake.findReturnsOnCall == nil {
		fake.findReturnsOnCall = make(map[int]struct {
			result1 db.Component
			result2 bool
			result3 error
		})
	}
	fake.findReturnsOnCall[i] = struct {
		result1 db.Component
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeComponentFactory) UpdateIntervals(arg1 map[string]time.Duration) error {
	fake.updateIntervalsMutex.Lock()
	ret, specificReturn := fake.updateIntervalsReturnsOnCall[len(fake.updateIntervalsArgsForCall)]
	fake.updateIntervalsArgsForCall = append(fake.updateIntervalsArgsForCall, struct {
		arg1 map[string]time.Duration
	}{arg1})
	fake.recordInvocation("UpdateIntervals", []interface{}{arg1})
	fake.updateIntervalsMutex.Unlock()
	if fake.UpdateIntervalsStub != nil {
		return fake.UpdateIntervalsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateIntervalsReturns
	return fakeReturns.result1
}

func (fake *FakeComponentFactory) UpdateIntervalsCallCount() int {
	fake.updateIntervalsMutex.RLock()
	defer fake.updateIntervalsMutex.RUnlock()
	return len(fake.updateIntervalsArgsForCall)
}

func (fake *FakeComponentFactory) UpdateIntervalsCalls(stub func(map[string]time.Duration) error) {
	fake.updateIntervalsMutex.Lock()
	defer fake.updateIntervalsMutex.Unlock()
	fake.UpdateIntervalsStub = stub
}

func (fake *FakeComponentFactory) UpdateIntervalsArgsForCall(i int) map[string]time.Duration {
	fake.updateIntervalsMutex.RLock()
	defer fake.updateIntervalsMutex.RUnlock()
	argsForCall := fake.updateIntervalsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeComponentFactory) UpdateIntervalsReturns(result1 error) {
	fake.updateIntervalsMutex.Lock()
	defer fake.updateIntervalsMutex.Unlock()
	fake.UpdateIntervalsStub = nil
	fake.updateIntervalsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeComponentFactory) UpdateIntervalsReturnsOnCall(i int, result1 error) {
	fake.updateIntervalsMutex.Lock()
	defer fake.updateIntervalsMutex.Unlock()
	fake.UpdateIntervalsStub = nil
	if fake.updateIntervalsReturnsOnCall == nil {
		fake.updateIntervalsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateIntervalsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeComponentFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	fake.updateIntervalsMutex.RLock()
	defer fake.updateIntervalsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeComponentFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.ComponentFactory = new(FakeComponentFactory)
