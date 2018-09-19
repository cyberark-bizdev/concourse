// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	sync "sync"
	time "time"

	atc "github.com/concourse/concourse/atc"
	db "github.com/concourse/concourse/atc/db"
)

type FakeWorker struct {
	ActiveContainersStub        func() int
	activeContainersMutex       sync.RWMutex
	activeContainersArgsForCall []struct {
	}
	activeContainersReturns struct {
		result1 int
	}
	activeContainersReturnsOnCall map[int]struct {
		result1 int
	}
	BaggageclaimURLStub        func() *string
	baggageclaimURLMutex       sync.RWMutex
	baggageclaimURLArgsForCall []struct {
	}
	baggageclaimURLReturns struct {
		result1 *string
	}
	baggageclaimURLReturnsOnCall map[int]struct {
		result1 *string
	}
	CertsPathStub        func() *string
	certsPathMutex       sync.RWMutex
	certsPathArgsForCall []struct {
	}
	certsPathReturns struct {
		result1 *string
	}
	certsPathReturnsOnCall map[int]struct {
		result1 *string
	}
	DeleteStub        func() error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	EphemeralStub        func() bool
	ephemeralMutex       sync.RWMutex
	ephemeralArgsForCall []struct {
	}
	ephemeralReturns struct {
		result1 bool
	}
	ephemeralReturnsOnCall map[int]struct {
		result1 bool
	}
	ExpiresAtStub        func() time.Time
	expiresAtMutex       sync.RWMutex
	expiresAtArgsForCall []struct {
	}
	expiresAtReturns struct {
		result1 time.Time
	}
	expiresAtReturnsOnCall map[int]struct {
		result1 time.Time
	}
	GardenAddrStub        func() *string
	gardenAddrMutex       sync.RWMutex
	gardenAddrArgsForCall []struct {
	}
	gardenAddrReturns struct {
		result1 *string
	}
	gardenAddrReturnsOnCall map[int]struct {
		result1 *string
	}
	HTTPProxyURLStub        func() string
	hTTPProxyURLMutex       sync.RWMutex
	hTTPProxyURLArgsForCall []struct {
	}
	hTTPProxyURLReturns struct {
		result1 string
	}
	hTTPProxyURLReturnsOnCall map[int]struct {
		result1 string
	}
	HTTPSProxyURLStub        func() string
	hTTPSProxyURLMutex       sync.RWMutex
	hTTPSProxyURLArgsForCall []struct {
	}
	hTTPSProxyURLReturns struct {
		result1 string
	}
	hTTPSProxyURLReturnsOnCall map[int]struct {
		result1 string
	}
	LandStub        func() error
	landMutex       sync.RWMutex
	landArgsForCall []struct {
	}
	landReturns struct {
		result1 error
	}
	landReturnsOnCall map[int]struct {
		result1 error
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	NoProxyStub        func() string
	noProxyMutex       sync.RWMutex
	noProxyArgsForCall []struct {
	}
	noProxyReturns struct {
		result1 string
	}
	noProxyReturnsOnCall map[int]struct {
		result1 string
	}
	PlatformStub        func() string
	platformMutex       sync.RWMutex
	platformArgsForCall []struct {
	}
	platformReturns struct {
		result1 string
	}
	platformReturnsOnCall map[int]struct {
		result1 string
	}
	PruneStub        func() error
	pruneMutex       sync.RWMutex
	pruneArgsForCall []struct {
	}
	pruneReturns struct {
		result1 error
	}
	pruneReturnsOnCall map[int]struct {
		result1 error
	}
	ReloadStub        func() (bool, error)
	reloadMutex       sync.RWMutex
	reloadArgsForCall []struct {
	}
	reloadReturns struct {
		result1 bool
		result2 error
	}
	reloadReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ResourceCertsStub        func() (*db.UsedWorkerResourceCerts, bool, error)
	resourceCertsMutex       sync.RWMutex
	resourceCertsArgsForCall []struct {
	}
	resourceCertsReturns struct {
		result1 *db.UsedWorkerResourceCerts
		result2 bool
		result3 error
	}
	resourceCertsReturnsOnCall map[int]struct {
		result1 *db.UsedWorkerResourceCerts
		result2 bool
		result3 error
	}
	ResourceTypesStub        func() []atc.WorkerResourceType
	resourceTypesMutex       sync.RWMutex
	resourceTypesArgsForCall []struct {
	}
	resourceTypesReturns struct {
		result1 []atc.WorkerResourceType
	}
	resourceTypesReturnsOnCall map[int]struct {
		result1 []atc.WorkerResourceType
	}
	RetireStub        func() error
	retireMutex       sync.RWMutex
	retireArgsForCall []struct {
	}
	retireReturns struct {
		result1 error
	}
	retireReturnsOnCall map[int]struct {
		result1 error
	}
	StartTimeStub        func() int64
	startTimeMutex       sync.RWMutex
	startTimeArgsForCall []struct {
	}
	startTimeReturns struct {
		result1 int64
	}
	startTimeReturnsOnCall map[int]struct {
		result1 int64
	}
	StateStub        func() db.WorkerState
	stateMutex       sync.RWMutex
	stateArgsForCall []struct {
	}
	stateReturns struct {
		result1 db.WorkerState
	}
	stateReturnsOnCall map[int]struct {
		result1 db.WorkerState
	}
	TagsStub        func() []string
	tagsMutex       sync.RWMutex
	tagsArgsForCall []struct {
	}
	tagsReturns struct {
		result1 []string
	}
	tagsReturnsOnCall map[int]struct {
		result1 []string
	}
	TeamIDStub        func() int
	teamIDMutex       sync.RWMutex
	teamIDArgsForCall []struct {
	}
	teamIDReturns struct {
		result1 int
	}
	teamIDReturnsOnCall map[int]struct {
		result1 int
	}
	TeamNameStub        func() string
	teamNameMutex       sync.RWMutex
	teamNameArgsForCall []struct {
	}
	teamNameReturns struct {
		result1 string
	}
	teamNameReturnsOnCall map[int]struct {
		result1 string
	}
	VersionStub        func() *string
	versionMutex       sync.RWMutex
	versionArgsForCall []struct {
	}
	versionReturns struct {
		result1 *string
	}
	versionReturnsOnCall map[int]struct {
		result1 *string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorker) ActiveContainers() int {
	fake.activeContainersMutex.Lock()
	ret, specificReturn := fake.activeContainersReturnsOnCall[len(fake.activeContainersArgsForCall)]
	fake.activeContainersArgsForCall = append(fake.activeContainersArgsForCall, struct {
	}{})
	fake.recordInvocation("ActiveContainers", []interface{}{})
	fake.activeContainersMutex.Unlock()
	if fake.ActiveContainersStub != nil {
		return fake.ActiveContainersStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.activeContainersReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) ActiveContainersCallCount() int {
	fake.activeContainersMutex.RLock()
	defer fake.activeContainersMutex.RUnlock()
	return len(fake.activeContainersArgsForCall)
}

func (fake *FakeWorker) ActiveContainersReturns(result1 int) {
	fake.ActiveContainersStub = nil
	fake.activeContainersReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorker) ActiveContainersReturnsOnCall(i int, result1 int) {
	fake.ActiveContainersStub = nil
	if fake.activeContainersReturnsOnCall == nil {
		fake.activeContainersReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.activeContainersReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorker) BaggageclaimURL() *string {
	fake.baggageclaimURLMutex.Lock()
	ret, specificReturn := fake.baggageclaimURLReturnsOnCall[len(fake.baggageclaimURLArgsForCall)]
	fake.baggageclaimURLArgsForCall = append(fake.baggageclaimURLArgsForCall, struct {
	}{})
	fake.recordInvocation("BaggageclaimURL", []interface{}{})
	fake.baggageclaimURLMutex.Unlock()
	if fake.BaggageclaimURLStub != nil {
		return fake.BaggageclaimURLStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.baggageclaimURLReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) BaggageclaimURLCallCount() int {
	fake.baggageclaimURLMutex.RLock()
	defer fake.baggageclaimURLMutex.RUnlock()
	return len(fake.baggageclaimURLArgsForCall)
}

func (fake *FakeWorker) BaggageclaimURLReturns(result1 *string) {
	fake.BaggageclaimURLStub = nil
	fake.baggageclaimURLReturns = struct {
		result1 *string
	}{result1}
}

func (fake *FakeWorker) BaggageclaimURLReturnsOnCall(i int, result1 *string) {
	fake.BaggageclaimURLStub = nil
	if fake.baggageclaimURLReturnsOnCall == nil {
		fake.baggageclaimURLReturnsOnCall = make(map[int]struct {
			result1 *string
		})
	}
	fake.baggageclaimURLReturnsOnCall[i] = struct {
		result1 *string
	}{result1}
}

func (fake *FakeWorker) CertsPath() *string {
	fake.certsPathMutex.Lock()
	ret, specificReturn := fake.certsPathReturnsOnCall[len(fake.certsPathArgsForCall)]
	fake.certsPathArgsForCall = append(fake.certsPathArgsForCall, struct {
	}{})
	fake.recordInvocation("CertsPath", []interface{}{})
	fake.certsPathMutex.Unlock()
	if fake.CertsPathStub != nil {
		return fake.CertsPathStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.certsPathReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) CertsPathCallCount() int {
	fake.certsPathMutex.RLock()
	defer fake.certsPathMutex.RUnlock()
	return len(fake.certsPathArgsForCall)
}

func (fake *FakeWorker) CertsPathReturns(result1 *string) {
	fake.CertsPathStub = nil
	fake.certsPathReturns = struct {
		result1 *string
	}{result1}
}

func (fake *FakeWorker) CertsPathReturnsOnCall(i int, result1 *string) {
	fake.CertsPathStub = nil
	if fake.certsPathReturnsOnCall == nil {
		fake.certsPathReturnsOnCall = make(map[int]struct {
			result1 *string
		})
	}
	fake.certsPathReturnsOnCall[i] = struct {
		result1 *string
	}{result1}
}

func (fake *FakeWorker) Delete() error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
	}{})
	fake.recordInvocation("Delete", []interface{}{})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeWorker) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorker) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorker) Ephemeral() bool {
	fake.ephemeralMutex.Lock()
	ret, specificReturn := fake.ephemeralReturnsOnCall[len(fake.ephemeralArgsForCall)]
	fake.ephemeralArgsForCall = append(fake.ephemeralArgsForCall, struct {
	}{})
	fake.recordInvocation("Ephemeral", []interface{}{})
	fake.ephemeralMutex.Unlock()
	if fake.EphemeralStub != nil {
		return fake.EphemeralStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.ephemeralReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) EphemeralCallCount() int {
	fake.ephemeralMutex.RLock()
	defer fake.ephemeralMutex.RUnlock()
	return len(fake.ephemeralArgsForCall)
}

func (fake *FakeWorker) EphemeralReturns(result1 bool) {
	fake.EphemeralStub = nil
	fake.ephemeralReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeWorker) EphemeralReturnsOnCall(i int, result1 bool) {
	fake.EphemeralStub = nil
	if fake.ephemeralReturnsOnCall == nil {
		fake.ephemeralReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.ephemeralReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeWorker) ExpiresAt() time.Time {
	fake.expiresAtMutex.Lock()
	ret, specificReturn := fake.expiresAtReturnsOnCall[len(fake.expiresAtArgsForCall)]
	fake.expiresAtArgsForCall = append(fake.expiresAtArgsForCall, struct {
	}{})
	fake.recordInvocation("ExpiresAt", []interface{}{})
	fake.expiresAtMutex.Unlock()
	if fake.ExpiresAtStub != nil {
		return fake.ExpiresAtStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.expiresAtReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) ExpiresAtCallCount() int {
	fake.expiresAtMutex.RLock()
	defer fake.expiresAtMutex.RUnlock()
	return len(fake.expiresAtArgsForCall)
}

func (fake *FakeWorker) ExpiresAtReturns(result1 time.Time) {
	fake.ExpiresAtStub = nil
	fake.expiresAtReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeWorker) ExpiresAtReturnsOnCall(i int, result1 time.Time) {
	fake.ExpiresAtStub = nil
	if fake.expiresAtReturnsOnCall == nil {
		fake.expiresAtReturnsOnCall = make(map[int]struct {
			result1 time.Time
		})
	}
	fake.expiresAtReturnsOnCall[i] = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeWorker) GardenAddr() *string {
	fake.gardenAddrMutex.Lock()
	ret, specificReturn := fake.gardenAddrReturnsOnCall[len(fake.gardenAddrArgsForCall)]
	fake.gardenAddrArgsForCall = append(fake.gardenAddrArgsForCall, struct {
	}{})
	fake.recordInvocation("GardenAddr", []interface{}{})
	fake.gardenAddrMutex.Unlock()
	if fake.GardenAddrStub != nil {
		return fake.GardenAddrStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.gardenAddrReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) GardenAddrCallCount() int {
	fake.gardenAddrMutex.RLock()
	defer fake.gardenAddrMutex.RUnlock()
	return len(fake.gardenAddrArgsForCall)
}

func (fake *FakeWorker) GardenAddrReturns(result1 *string) {
	fake.GardenAddrStub = nil
	fake.gardenAddrReturns = struct {
		result1 *string
	}{result1}
}

func (fake *FakeWorker) GardenAddrReturnsOnCall(i int, result1 *string) {
	fake.GardenAddrStub = nil
	if fake.gardenAddrReturnsOnCall == nil {
		fake.gardenAddrReturnsOnCall = make(map[int]struct {
			result1 *string
		})
	}
	fake.gardenAddrReturnsOnCall[i] = struct {
		result1 *string
	}{result1}
}

func (fake *FakeWorker) HTTPProxyURL() string {
	fake.hTTPProxyURLMutex.Lock()
	ret, specificReturn := fake.hTTPProxyURLReturnsOnCall[len(fake.hTTPProxyURLArgsForCall)]
	fake.hTTPProxyURLArgsForCall = append(fake.hTTPProxyURLArgsForCall, struct {
	}{})
	fake.recordInvocation("HTTPProxyURL", []interface{}{})
	fake.hTTPProxyURLMutex.Unlock()
	if fake.HTTPProxyURLStub != nil {
		return fake.HTTPProxyURLStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.hTTPProxyURLReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) HTTPProxyURLCallCount() int {
	fake.hTTPProxyURLMutex.RLock()
	defer fake.hTTPProxyURLMutex.RUnlock()
	return len(fake.hTTPProxyURLArgsForCall)
}

func (fake *FakeWorker) HTTPProxyURLReturns(result1 string) {
	fake.HTTPProxyURLStub = nil
	fake.hTTPProxyURLReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) HTTPProxyURLReturnsOnCall(i int, result1 string) {
	fake.HTTPProxyURLStub = nil
	if fake.hTTPProxyURLReturnsOnCall == nil {
		fake.hTTPProxyURLReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.hTTPProxyURLReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) HTTPSProxyURL() string {
	fake.hTTPSProxyURLMutex.Lock()
	ret, specificReturn := fake.hTTPSProxyURLReturnsOnCall[len(fake.hTTPSProxyURLArgsForCall)]
	fake.hTTPSProxyURLArgsForCall = append(fake.hTTPSProxyURLArgsForCall, struct {
	}{})
	fake.recordInvocation("HTTPSProxyURL", []interface{}{})
	fake.hTTPSProxyURLMutex.Unlock()
	if fake.HTTPSProxyURLStub != nil {
		return fake.HTTPSProxyURLStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.hTTPSProxyURLReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) HTTPSProxyURLCallCount() int {
	fake.hTTPSProxyURLMutex.RLock()
	defer fake.hTTPSProxyURLMutex.RUnlock()
	return len(fake.hTTPSProxyURLArgsForCall)
}

func (fake *FakeWorker) HTTPSProxyURLReturns(result1 string) {
	fake.HTTPSProxyURLStub = nil
	fake.hTTPSProxyURLReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) HTTPSProxyURLReturnsOnCall(i int, result1 string) {
	fake.HTTPSProxyURLStub = nil
	if fake.hTTPSProxyURLReturnsOnCall == nil {
		fake.hTTPSProxyURLReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.hTTPSProxyURLReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) Land() error {
	fake.landMutex.Lock()
	ret, specificReturn := fake.landReturnsOnCall[len(fake.landArgsForCall)]
	fake.landArgsForCall = append(fake.landArgsForCall, struct {
	}{})
	fake.recordInvocation("Land", []interface{}{})
	fake.landMutex.Unlock()
	if fake.LandStub != nil {
		return fake.LandStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.landReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) LandCallCount() int {
	fake.landMutex.RLock()
	defer fake.landMutex.RUnlock()
	return len(fake.landArgsForCall)
}

func (fake *FakeWorker) LandReturns(result1 error) {
	fake.LandStub = nil
	fake.landReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorker) LandReturnsOnCall(i int, result1 error) {
	fake.LandStub = nil
	if fake.landReturnsOnCall == nil {
		fake.landReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.landReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorker) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nameReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeWorker) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) NameReturnsOnCall(i int, result1 string) {
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) NoProxy() string {
	fake.noProxyMutex.Lock()
	ret, specificReturn := fake.noProxyReturnsOnCall[len(fake.noProxyArgsForCall)]
	fake.noProxyArgsForCall = append(fake.noProxyArgsForCall, struct {
	}{})
	fake.recordInvocation("NoProxy", []interface{}{})
	fake.noProxyMutex.Unlock()
	if fake.NoProxyStub != nil {
		return fake.NoProxyStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.noProxyReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) NoProxyCallCount() int {
	fake.noProxyMutex.RLock()
	defer fake.noProxyMutex.RUnlock()
	return len(fake.noProxyArgsForCall)
}

func (fake *FakeWorker) NoProxyReturns(result1 string) {
	fake.NoProxyStub = nil
	fake.noProxyReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) NoProxyReturnsOnCall(i int, result1 string) {
	fake.NoProxyStub = nil
	if fake.noProxyReturnsOnCall == nil {
		fake.noProxyReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.noProxyReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) Platform() string {
	fake.platformMutex.Lock()
	ret, specificReturn := fake.platformReturnsOnCall[len(fake.platformArgsForCall)]
	fake.platformArgsForCall = append(fake.platformArgsForCall, struct {
	}{})
	fake.recordInvocation("Platform", []interface{}{})
	fake.platformMutex.Unlock()
	if fake.PlatformStub != nil {
		return fake.PlatformStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.platformReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) PlatformCallCount() int {
	fake.platformMutex.RLock()
	defer fake.platformMutex.RUnlock()
	return len(fake.platformArgsForCall)
}

func (fake *FakeWorker) PlatformReturns(result1 string) {
	fake.PlatformStub = nil
	fake.platformReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) PlatformReturnsOnCall(i int, result1 string) {
	fake.PlatformStub = nil
	if fake.platformReturnsOnCall == nil {
		fake.platformReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.platformReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) Prune() error {
	fake.pruneMutex.Lock()
	ret, specificReturn := fake.pruneReturnsOnCall[len(fake.pruneArgsForCall)]
	fake.pruneArgsForCall = append(fake.pruneArgsForCall, struct {
	}{})
	fake.recordInvocation("Prune", []interface{}{})
	fake.pruneMutex.Unlock()
	if fake.PruneStub != nil {
		return fake.PruneStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pruneReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) PruneCallCount() int {
	fake.pruneMutex.RLock()
	defer fake.pruneMutex.RUnlock()
	return len(fake.pruneArgsForCall)
}

func (fake *FakeWorker) PruneReturns(result1 error) {
	fake.PruneStub = nil
	fake.pruneReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorker) PruneReturnsOnCall(i int, result1 error) {
	fake.PruneStub = nil
	if fake.pruneReturnsOnCall == nil {
		fake.pruneReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pruneReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorker) Reload() (bool, error) {
	fake.reloadMutex.Lock()
	ret, specificReturn := fake.reloadReturnsOnCall[len(fake.reloadArgsForCall)]
	fake.reloadArgsForCall = append(fake.reloadArgsForCall, struct {
	}{})
	fake.recordInvocation("Reload", []interface{}{})
	fake.reloadMutex.Unlock()
	if fake.ReloadStub != nil {
		return fake.ReloadStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.reloadReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorker) ReloadCallCount() int {
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	return len(fake.reloadArgsForCall)
}

func (fake *FakeWorker) ReloadReturns(result1 bool, result2 error) {
	fake.ReloadStub = nil
	fake.reloadReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) ReloadReturnsOnCall(i int, result1 bool, result2 error) {
	fake.ReloadStub = nil
	if fake.reloadReturnsOnCall == nil {
		fake.reloadReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.reloadReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) ResourceCerts() (*db.UsedWorkerResourceCerts, bool, error) {
	fake.resourceCertsMutex.Lock()
	ret, specificReturn := fake.resourceCertsReturnsOnCall[len(fake.resourceCertsArgsForCall)]
	fake.resourceCertsArgsForCall = append(fake.resourceCertsArgsForCall, struct {
	}{})
	fake.recordInvocation("ResourceCerts", []interface{}{})
	fake.resourceCertsMutex.Unlock()
	if fake.ResourceCertsStub != nil {
		return fake.ResourceCertsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.resourceCertsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeWorker) ResourceCertsCallCount() int {
	fake.resourceCertsMutex.RLock()
	defer fake.resourceCertsMutex.RUnlock()
	return len(fake.resourceCertsArgsForCall)
}

func (fake *FakeWorker) ResourceCertsReturns(result1 *db.UsedWorkerResourceCerts, result2 bool, result3 error) {
	fake.ResourceCertsStub = nil
	fake.resourceCertsReturns = struct {
		result1 *db.UsedWorkerResourceCerts
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorker) ResourceCertsReturnsOnCall(i int, result1 *db.UsedWorkerResourceCerts, result2 bool, result3 error) {
	fake.ResourceCertsStub = nil
	if fake.resourceCertsReturnsOnCall == nil {
		fake.resourceCertsReturnsOnCall = make(map[int]struct {
			result1 *db.UsedWorkerResourceCerts
			result2 bool
			result3 error
		})
	}
	fake.resourceCertsReturnsOnCall[i] = struct {
		result1 *db.UsedWorkerResourceCerts
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorker) ResourceTypes() []atc.WorkerResourceType {
	fake.resourceTypesMutex.Lock()
	ret, specificReturn := fake.resourceTypesReturnsOnCall[len(fake.resourceTypesArgsForCall)]
	fake.resourceTypesArgsForCall = append(fake.resourceTypesArgsForCall, struct {
	}{})
	fake.recordInvocation("ResourceTypes", []interface{}{})
	fake.resourceTypesMutex.Unlock()
	if fake.ResourceTypesStub != nil {
		return fake.ResourceTypesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.resourceTypesReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) ResourceTypesCallCount() int {
	fake.resourceTypesMutex.RLock()
	defer fake.resourceTypesMutex.RUnlock()
	return len(fake.resourceTypesArgsForCall)
}

func (fake *FakeWorker) ResourceTypesReturns(result1 []atc.WorkerResourceType) {
	fake.ResourceTypesStub = nil
	fake.resourceTypesReturns = struct {
		result1 []atc.WorkerResourceType
	}{result1}
}

func (fake *FakeWorker) ResourceTypesReturnsOnCall(i int, result1 []atc.WorkerResourceType) {
	fake.ResourceTypesStub = nil
	if fake.resourceTypesReturnsOnCall == nil {
		fake.resourceTypesReturnsOnCall = make(map[int]struct {
			result1 []atc.WorkerResourceType
		})
	}
	fake.resourceTypesReturnsOnCall[i] = struct {
		result1 []atc.WorkerResourceType
	}{result1}
}

func (fake *FakeWorker) Retire() error {
	fake.retireMutex.Lock()
	ret, specificReturn := fake.retireReturnsOnCall[len(fake.retireArgsForCall)]
	fake.retireArgsForCall = append(fake.retireArgsForCall, struct {
	}{})
	fake.recordInvocation("Retire", []interface{}{})
	fake.retireMutex.Unlock()
	if fake.RetireStub != nil {
		return fake.RetireStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.retireReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) RetireCallCount() int {
	fake.retireMutex.RLock()
	defer fake.retireMutex.RUnlock()
	return len(fake.retireArgsForCall)
}

func (fake *FakeWorker) RetireReturns(result1 error) {
	fake.RetireStub = nil
	fake.retireReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorker) RetireReturnsOnCall(i int, result1 error) {
	fake.RetireStub = nil
	if fake.retireReturnsOnCall == nil {
		fake.retireReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.retireReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorker) StartTime() int64 {
	fake.startTimeMutex.Lock()
	ret, specificReturn := fake.startTimeReturnsOnCall[len(fake.startTimeArgsForCall)]
	fake.startTimeArgsForCall = append(fake.startTimeArgsForCall, struct {
	}{})
	fake.recordInvocation("StartTime", []interface{}{})
	fake.startTimeMutex.Unlock()
	if fake.StartTimeStub != nil {
		return fake.StartTimeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.startTimeReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) StartTimeCallCount() int {
	fake.startTimeMutex.RLock()
	defer fake.startTimeMutex.RUnlock()
	return len(fake.startTimeArgsForCall)
}

func (fake *FakeWorker) StartTimeReturns(result1 int64) {
	fake.StartTimeStub = nil
	fake.startTimeReturns = struct {
		result1 int64
	}{result1}
}

func (fake *FakeWorker) StartTimeReturnsOnCall(i int, result1 int64) {
	fake.StartTimeStub = nil
	if fake.startTimeReturnsOnCall == nil {
		fake.startTimeReturnsOnCall = make(map[int]struct {
			result1 int64
		})
	}
	fake.startTimeReturnsOnCall[i] = struct {
		result1 int64
	}{result1}
}

func (fake *FakeWorker) State() db.WorkerState {
	fake.stateMutex.Lock()
	ret, specificReturn := fake.stateReturnsOnCall[len(fake.stateArgsForCall)]
	fake.stateArgsForCall = append(fake.stateArgsForCall, struct {
	}{})
	fake.recordInvocation("State", []interface{}{})
	fake.stateMutex.Unlock()
	if fake.StateStub != nil {
		return fake.StateStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stateReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) StateCallCount() int {
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	return len(fake.stateArgsForCall)
}

func (fake *FakeWorker) StateReturns(result1 db.WorkerState) {
	fake.StateStub = nil
	fake.stateReturns = struct {
		result1 db.WorkerState
	}{result1}
}

func (fake *FakeWorker) StateReturnsOnCall(i int, result1 db.WorkerState) {
	fake.StateStub = nil
	if fake.stateReturnsOnCall == nil {
		fake.stateReturnsOnCall = make(map[int]struct {
			result1 db.WorkerState
		})
	}
	fake.stateReturnsOnCall[i] = struct {
		result1 db.WorkerState
	}{result1}
}

func (fake *FakeWorker) Tags() []string {
	fake.tagsMutex.Lock()
	ret, specificReturn := fake.tagsReturnsOnCall[len(fake.tagsArgsForCall)]
	fake.tagsArgsForCall = append(fake.tagsArgsForCall, struct {
	}{})
	fake.recordInvocation("Tags", []interface{}{})
	fake.tagsMutex.Unlock()
	if fake.TagsStub != nil {
		return fake.TagsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.tagsReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) TagsCallCount() int {
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	return len(fake.tagsArgsForCall)
}

func (fake *FakeWorker) TagsReturns(result1 []string) {
	fake.TagsStub = nil
	fake.tagsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeWorker) TagsReturnsOnCall(i int, result1 []string) {
	fake.TagsStub = nil
	if fake.tagsReturnsOnCall == nil {
		fake.tagsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.tagsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeWorker) TeamID() int {
	fake.teamIDMutex.Lock()
	ret, specificReturn := fake.teamIDReturnsOnCall[len(fake.teamIDArgsForCall)]
	fake.teamIDArgsForCall = append(fake.teamIDArgsForCall, struct {
	}{})
	fake.recordInvocation("TeamID", []interface{}{})
	fake.teamIDMutex.Unlock()
	if fake.TeamIDStub != nil {
		return fake.TeamIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.teamIDReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) TeamIDCallCount() int {
	fake.teamIDMutex.RLock()
	defer fake.teamIDMutex.RUnlock()
	return len(fake.teamIDArgsForCall)
}

func (fake *FakeWorker) TeamIDReturns(result1 int) {
	fake.TeamIDStub = nil
	fake.teamIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorker) TeamIDReturnsOnCall(i int, result1 int) {
	fake.TeamIDStub = nil
	if fake.teamIDReturnsOnCall == nil {
		fake.teamIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.teamIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorker) TeamName() string {
	fake.teamNameMutex.Lock()
	ret, specificReturn := fake.teamNameReturnsOnCall[len(fake.teamNameArgsForCall)]
	fake.teamNameArgsForCall = append(fake.teamNameArgsForCall, struct {
	}{})
	fake.recordInvocation("TeamName", []interface{}{})
	fake.teamNameMutex.Unlock()
	if fake.TeamNameStub != nil {
		return fake.TeamNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.teamNameReturns
	return fakeReturns.result1
}

func (fake *FakeWorker) TeamNameCallCount() int {
	fake.teamNameMutex.RLock()
	defer fake.teamNameMutex.RUnlock()
	return len(fake.teamNameArgsForCall)
}

func (fake *FakeWorker) TeamNameReturns(result1 string) {
	fake.TeamNameStub = nil
	fake.teamNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) TeamNameReturnsOnCall(i int, result1 string) {
	fake.TeamNameStub = nil
	if fake.teamNameReturnsOnCall == nil {
		fake.teamNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.teamNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) Version() *string {
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

func (fake *FakeWorker) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeWorker) VersionReturns(result1 *string) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 *string
	}{result1}
}

func (fake *FakeWorker) VersionReturnsOnCall(i int, result1 *string) {
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 *string
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 *string
	}{result1}
}

func (fake *FakeWorker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.activeContainersMutex.RLock()
	defer fake.activeContainersMutex.RUnlock()
	fake.baggageclaimURLMutex.RLock()
	defer fake.baggageclaimURLMutex.RUnlock()
	fake.certsPathMutex.RLock()
	defer fake.certsPathMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.ephemeralMutex.RLock()
	defer fake.ephemeralMutex.RUnlock()
	fake.expiresAtMutex.RLock()
	defer fake.expiresAtMutex.RUnlock()
	fake.gardenAddrMutex.RLock()
	defer fake.gardenAddrMutex.RUnlock()
	fake.hTTPProxyURLMutex.RLock()
	defer fake.hTTPProxyURLMutex.RUnlock()
	fake.hTTPSProxyURLMutex.RLock()
	defer fake.hTTPSProxyURLMutex.RUnlock()
	fake.landMutex.RLock()
	defer fake.landMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.noProxyMutex.RLock()
	defer fake.noProxyMutex.RUnlock()
	fake.platformMutex.RLock()
	defer fake.platformMutex.RUnlock()
	fake.pruneMutex.RLock()
	defer fake.pruneMutex.RUnlock()
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	fake.resourceCertsMutex.RLock()
	defer fake.resourceCertsMutex.RUnlock()
	fake.resourceTypesMutex.RLock()
	defer fake.resourceTypesMutex.RUnlock()
	fake.retireMutex.RLock()
	defer fake.retireMutex.RUnlock()
	fake.startTimeMutex.RLock()
	defer fake.startTimeMutex.RUnlock()
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	fake.teamIDMutex.RLock()
	defer fake.teamIDMutex.RUnlock()
	fake.teamNameMutex.RLock()
	defer fake.teamNameMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWorker) recordInvocation(key string, args []interface{}) {
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

var _ db.Worker = new(FakeWorker)
