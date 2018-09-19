// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	sync "sync"

	atc "github.com/concourse/concourse/atc"
	db "github.com/concourse/concourse/atc/db"
)

type FakeUsedResourceCache struct {
	BaseResourceTypeStub        func() *db.UsedBaseResourceType
	baseResourceTypeMutex       sync.RWMutex
	baseResourceTypeArgsForCall []struct {
	}
	baseResourceTypeReturns struct {
		result1 *db.UsedBaseResourceType
	}
	baseResourceTypeReturnsOnCall map[int]struct {
		result1 *db.UsedBaseResourceType
	}
	DestroyStub        func(db.Tx) (bool, error)
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		arg1 db.Tx
	}
	destroyReturns struct {
		result1 bool
		result2 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 int
	}
	iDReturnsOnCall map[int]struct {
		result1 int
	}
	ResourceConfigStub        func() db.ResourceConfig
	resourceConfigMutex       sync.RWMutex
	resourceConfigArgsForCall []struct {
	}
	resourceConfigReturns struct {
		result1 db.ResourceConfig
	}
	resourceConfigReturnsOnCall map[int]struct {
		result1 db.ResourceConfig
	}
	VersionStub        func() atc.Version
	versionMutex       sync.RWMutex
	versionArgsForCall []struct {
	}
	versionReturns struct {
		result1 atc.Version
	}
	versionReturnsOnCall map[int]struct {
		result1 atc.Version
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUsedResourceCache) BaseResourceType() *db.UsedBaseResourceType {
	fake.baseResourceTypeMutex.Lock()
	ret, specificReturn := fake.baseResourceTypeReturnsOnCall[len(fake.baseResourceTypeArgsForCall)]
	fake.baseResourceTypeArgsForCall = append(fake.baseResourceTypeArgsForCall, struct {
	}{})
	fake.recordInvocation("BaseResourceType", []interface{}{})
	fake.baseResourceTypeMutex.Unlock()
	if fake.BaseResourceTypeStub != nil {
		return fake.BaseResourceTypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.baseResourceTypeReturns
	return fakeReturns.result1
}

func (fake *FakeUsedResourceCache) BaseResourceTypeCallCount() int {
	fake.baseResourceTypeMutex.RLock()
	defer fake.baseResourceTypeMutex.RUnlock()
	return len(fake.baseResourceTypeArgsForCall)
}

func (fake *FakeUsedResourceCache) BaseResourceTypeReturns(result1 *db.UsedBaseResourceType) {
	fake.BaseResourceTypeStub = nil
	fake.baseResourceTypeReturns = struct {
		result1 *db.UsedBaseResourceType
	}{result1}
}

func (fake *FakeUsedResourceCache) BaseResourceTypeReturnsOnCall(i int, result1 *db.UsedBaseResourceType) {
	fake.BaseResourceTypeStub = nil
	if fake.baseResourceTypeReturnsOnCall == nil {
		fake.baseResourceTypeReturnsOnCall = make(map[int]struct {
			result1 *db.UsedBaseResourceType
		})
	}
	fake.baseResourceTypeReturnsOnCall[i] = struct {
		result1 *db.UsedBaseResourceType
	}{result1}
}

func (fake *FakeUsedResourceCache) Destroy(arg1 db.Tx) (bool, error) {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		arg1 db.Tx
	}{arg1})
	fake.recordInvocation("Destroy", []interface{}{arg1})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.destroyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUsedResourceCache) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeUsedResourceCache) DestroyArgsForCall(i int) db.Tx {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	argsForCall := fake.destroyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUsedResourceCache) DestroyReturns(result1 bool, result2 error) {
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeUsedResourceCache) DestroyReturnsOnCall(i int, result1 bool, result2 error) {
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeUsedResourceCache) ID() int {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.iDReturns
	return fakeReturns.result1
}

func (fake *FakeUsedResourceCache) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeUsedResourceCache) IDReturns(result1 int) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeUsedResourceCache) IDReturnsOnCall(i int, result1 int) {
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeUsedResourceCache) ResourceConfig() db.ResourceConfig {
	fake.resourceConfigMutex.Lock()
	ret, specificReturn := fake.resourceConfigReturnsOnCall[len(fake.resourceConfigArgsForCall)]
	fake.resourceConfigArgsForCall = append(fake.resourceConfigArgsForCall, struct {
	}{})
	fake.recordInvocation("ResourceConfig", []interface{}{})
	fake.resourceConfigMutex.Unlock()
	if fake.ResourceConfigStub != nil {
		return fake.ResourceConfigStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.resourceConfigReturns
	return fakeReturns.result1
}

func (fake *FakeUsedResourceCache) ResourceConfigCallCount() int {
	fake.resourceConfigMutex.RLock()
	defer fake.resourceConfigMutex.RUnlock()
	return len(fake.resourceConfigArgsForCall)
}

func (fake *FakeUsedResourceCache) ResourceConfigReturns(result1 db.ResourceConfig) {
	fake.ResourceConfigStub = nil
	fake.resourceConfigReturns = struct {
		result1 db.ResourceConfig
	}{result1}
}

func (fake *FakeUsedResourceCache) ResourceConfigReturnsOnCall(i int, result1 db.ResourceConfig) {
	fake.ResourceConfigStub = nil
	if fake.resourceConfigReturnsOnCall == nil {
		fake.resourceConfigReturnsOnCall = make(map[int]struct {
			result1 db.ResourceConfig
		})
	}
	fake.resourceConfigReturnsOnCall[i] = struct {
		result1 db.ResourceConfig
	}{result1}
}

func (fake *FakeUsedResourceCache) Version() atc.Version {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct {
	}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.versionReturns
	return fakeReturns.result1
}

func (fake *FakeUsedResourceCache) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeUsedResourceCache) VersionReturns(result1 atc.Version) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeUsedResourceCache) VersionReturnsOnCall(i int, result1 atc.Version) {
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 atc.Version
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeUsedResourceCache) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.baseResourceTypeMutex.RLock()
	defer fake.baseResourceTypeMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.resourceConfigMutex.RLock()
	defer fake.resourceConfigMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUsedResourceCache) recordInvocation(key string, args []interface{}) {
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

var _ db.UsedResourceCache = new(FakeUsedResourceCache)
