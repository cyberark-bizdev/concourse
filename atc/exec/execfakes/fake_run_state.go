// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	io "io"
	sync "sync"

	atc "github.com/concourse/concourse/atc"
	exec "github.com/concourse/concourse/atc/exec"
	worker "github.com/concourse/concourse/atc/worker"
)

type FakeRunState struct {
	ArtifactsStub        func() *worker.ArtifactRepository
	artifactsMutex       sync.RWMutex
	artifactsArgsForCall []struct {
	}
	artifactsReturns struct {
		result1 *worker.ArtifactRepository
	}
	artifactsReturnsOnCall map[int]struct {
		result1 *worker.ArtifactRepository
	}
	ReadPlanOutputStub        func(atc.PlanID, io.Writer)
	readPlanOutputMutex       sync.RWMutex
	readPlanOutputArgsForCall []struct {
		arg1 atc.PlanID
		arg2 io.Writer
	}
	ReadUserInputStub        func(atc.PlanID, exec.InputHandler) error
	readUserInputMutex       sync.RWMutex
	readUserInputArgsForCall []struct {
		arg1 atc.PlanID
		arg2 exec.InputHandler
	}
	readUserInputReturns struct {
		result1 error
	}
	readUserInputReturnsOnCall map[int]struct {
		result1 error
	}
	ResultStub        func(atc.PlanID, interface{}) bool
	resultMutex       sync.RWMutex
	resultArgsForCall []struct {
		arg1 atc.PlanID
		arg2 interface{}
	}
	resultReturns struct {
		result1 bool
	}
	resultReturnsOnCall map[int]struct {
		result1 bool
	}
	SendPlanOutputStub        func(atc.PlanID, exec.OutputHandler) error
	sendPlanOutputMutex       sync.RWMutex
	sendPlanOutputArgsForCall []struct {
		arg1 atc.PlanID
		arg2 exec.OutputHandler
	}
	sendPlanOutputReturns struct {
		result1 error
	}
	sendPlanOutputReturnsOnCall map[int]struct {
		result1 error
	}
	SendUserInputStub        func(atc.PlanID, io.ReadCloser)
	sendUserInputMutex       sync.RWMutex
	sendUserInputArgsForCall []struct {
		arg1 atc.PlanID
		arg2 io.ReadCloser
	}
	StoreResultStub        func(atc.PlanID, interface{})
	storeResultMutex       sync.RWMutex
	storeResultArgsForCall []struct {
		arg1 atc.PlanID
		arg2 interface{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRunState) Artifacts() *worker.ArtifactRepository {
	fake.artifactsMutex.Lock()
	ret, specificReturn := fake.artifactsReturnsOnCall[len(fake.artifactsArgsForCall)]
	fake.artifactsArgsForCall = append(fake.artifactsArgsForCall, struct {
	}{})
	fake.recordInvocation("Artifacts", []interface{}{})
	fake.artifactsMutex.Unlock()
	if fake.ArtifactsStub != nil {
		return fake.ArtifactsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.artifactsReturns
	return fakeReturns.result1
}

func (fake *FakeRunState) ArtifactsCallCount() int {
	fake.artifactsMutex.RLock()
	defer fake.artifactsMutex.RUnlock()
	return len(fake.artifactsArgsForCall)
}

func (fake *FakeRunState) ArtifactsReturns(result1 *worker.ArtifactRepository) {
	fake.ArtifactsStub = nil
	fake.artifactsReturns = struct {
		result1 *worker.ArtifactRepository
	}{result1}
}

func (fake *FakeRunState) ArtifactsReturnsOnCall(i int, result1 *worker.ArtifactRepository) {
	fake.ArtifactsStub = nil
	if fake.artifactsReturnsOnCall == nil {
		fake.artifactsReturnsOnCall = make(map[int]struct {
			result1 *worker.ArtifactRepository
		})
	}
	fake.artifactsReturnsOnCall[i] = struct {
		result1 *worker.ArtifactRepository
	}{result1}
}

func (fake *FakeRunState) ReadPlanOutput(arg1 atc.PlanID, arg2 io.Writer) {
	fake.readPlanOutputMutex.Lock()
	fake.readPlanOutputArgsForCall = append(fake.readPlanOutputArgsForCall, struct {
		arg1 atc.PlanID
		arg2 io.Writer
	}{arg1, arg2})
	fake.recordInvocation("ReadPlanOutput", []interface{}{arg1, arg2})
	fake.readPlanOutputMutex.Unlock()
	if fake.ReadPlanOutputStub != nil {
		fake.ReadPlanOutputStub(arg1, arg2)
	}
}

func (fake *FakeRunState) ReadPlanOutputCallCount() int {
	fake.readPlanOutputMutex.RLock()
	defer fake.readPlanOutputMutex.RUnlock()
	return len(fake.readPlanOutputArgsForCall)
}

func (fake *FakeRunState) ReadPlanOutputArgsForCall(i int) (atc.PlanID, io.Writer) {
	fake.readPlanOutputMutex.RLock()
	defer fake.readPlanOutputMutex.RUnlock()
	argsForCall := fake.readPlanOutputArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRunState) ReadUserInput(arg1 atc.PlanID, arg2 exec.InputHandler) error {
	fake.readUserInputMutex.Lock()
	ret, specificReturn := fake.readUserInputReturnsOnCall[len(fake.readUserInputArgsForCall)]
	fake.readUserInputArgsForCall = append(fake.readUserInputArgsForCall, struct {
		arg1 atc.PlanID
		arg2 exec.InputHandler
	}{arg1, arg2})
	fake.recordInvocation("ReadUserInput", []interface{}{arg1, arg2})
	fake.readUserInputMutex.Unlock()
	if fake.ReadUserInputStub != nil {
		return fake.ReadUserInputStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.readUserInputReturns
	return fakeReturns.result1
}

func (fake *FakeRunState) ReadUserInputCallCount() int {
	fake.readUserInputMutex.RLock()
	defer fake.readUserInputMutex.RUnlock()
	return len(fake.readUserInputArgsForCall)
}

func (fake *FakeRunState) ReadUserInputArgsForCall(i int) (atc.PlanID, exec.InputHandler) {
	fake.readUserInputMutex.RLock()
	defer fake.readUserInputMutex.RUnlock()
	argsForCall := fake.readUserInputArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRunState) ReadUserInputReturns(result1 error) {
	fake.ReadUserInputStub = nil
	fake.readUserInputReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRunState) ReadUserInputReturnsOnCall(i int, result1 error) {
	fake.ReadUserInputStub = nil
	if fake.readUserInputReturnsOnCall == nil {
		fake.readUserInputReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.readUserInputReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRunState) Result(arg1 atc.PlanID, arg2 interface{}) bool {
	fake.resultMutex.Lock()
	ret, specificReturn := fake.resultReturnsOnCall[len(fake.resultArgsForCall)]
	fake.resultArgsForCall = append(fake.resultArgsForCall, struct {
		arg1 atc.PlanID
		arg2 interface{}
	}{arg1, arg2})
	fake.recordInvocation("Result", []interface{}{arg1, arg2})
	fake.resultMutex.Unlock()
	if fake.ResultStub != nil {
		return fake.ResultStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.resultReturns
	return fakeReturns.result1
}

func (fake *FakeRunState) ResultCallCount() int {
	fake.resultMutex.RLock()
	defer fake.resultMutex.RUnlock()
	return len(fake.resultArgsForCall)
}

func (fake *FakeRunState) ResultArgsForCall(i int) (atc.PlanID, interface{}) {
	fake.resultMutex.RLock()
	defer fake.resultMutex.RUnlock()
	argsForCall := fake.resultArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRunState) ResultReturns(result1 bool) {
	fake.ResultStub = nil
	fake.resultReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRunState) ResultReturnsOnCall(i int, result1 bool) {
	fake.ResultStub = nil
	if fake.resultReturnsOnCall == nil {
		fake.resultReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.resultReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRunState) SendPlanOutput(arg1 atc.PlanID, arg2 exec.OutputHandler) error {
	fake.sendPlanOutputMutex.Lock()
	ret, specificReturn := fake.sendPlanOutputReturnsOnCall[len(fake.sendPlanOutputArgsForCall)]
	fake.sendPlanOutputArgsForCall = append(fake.sendPlanOutputArgsForCall, struct {
		arg1 atc.PlanID
		arg2 exec.OutputHandler
	}{arg1, arg2})
	fake.recordInvocation("SendPlanOutput", []interface{}{arg1, arg2})
	fake.sendPlanOutputMutex.Unlock()
	if fake.SendPlanOutputStub != nil {
		return fake.SendPlanOutputStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sendPlanOutputReturns
	return fakeReturns.result1
}

func (fake *FakeRunState) SendPlanOutputCallCount() int {
	fake.sendPlanOutputMutex.RLock()
	defer fake.sendPlanOutputMutex.RUnlock()
	return len(fake.sendPlanOutputArgsForCall)
}

func (fake *FakeRunState) SendPlanOutputArgsForCall(i int) (atc.PlanID, exec.OutputHandler) {
	fake.sendPlanOutputMutex.RLock()
	defer fake.sendPlanOutputMutex.RUnlock()
	argsForCall := fake.sendPlanOutputArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRunState) SendPlanOutputReturns(result1 error) {
	fake.SendPlanOutputStub = nil
	fake.sendPlanOutputReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRunState) SendPlanOutputReturnsOnCall(i int, result1 error) {
	fake.SendPlanOutputStub = nil
	if fake.sendPlanOutputReturnsOnCall == nil {
		fake.sendPlanOutputReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendPlanOutputReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRunState) SendUserInput(arg1 atc.PlanID, arg2 io.ReadCloser) {
	fake.sendUserInputMutex.Lock()
	fake.sendUserInputArgsForCall = append(fake.sendUserInputArgsForCall, struct {
		arg1 atc.PlanID
		arg2 io.ReadCloser
	}{arg1, arg2})
	fake.recordInvocation("SendUserInput", []interface{}{arg1, arg2})
	fake.sendUserInputMutex.Unlock()
	if fake.SendUserInputStub != nil {
		fake.SendUserInputStub(arg1, arg2)
	}
}

func (fake *FakeRunState) SendUserInputCallCount() int {
	fake.sendUserInputMutex.RLock()
	defer fake.sendUserInputMutex.RUnlock()
	return len(fake.sendUserInputArgsForCall)
}

func (fake *FakeRunState) SendUserInputArgsForCall(i int) (atc.PlanID, io.ReadCloser) {
	fake.sendUserInputMutex.RLock()
	defer fake.sendUserInputMutex.RUnlock()
	argsForCall := fake.sendUserInputArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRunState) StoreResult(arg1 atc.PlanID, arg2 interface{}) {
	fake.storeResultMutex.Lock()
	fake.storeResultArgsForCall = append(fake.storeResultArgsForCall, struct {
		arg1 atc.PlanID
		arg2 interface{}
	}{arg1, arg2})
	fake.recordInvocation("StoreResult", []interface{}{arg1, arg2})
	fake.storeResultMutex.Unlock()
	if fake.StoreResultStub != nil {
		fake.StoreResultStub(arg1, arg2)
	}
}

func (fake *FakeRunState) StoreResultCallCount() int {
	fake.storeResultMutex.RLock()
	defer fake.storeResultMutex.RUnlock()
	return len(fake.storeResultArgsForCall)
}

func (fake *FakeRunState) StoreResultArgsForCall(i int) (atc.PlanID, interface{}) {
	fake.storeResultMutex.RLock()
	defer fake.storeResultMutex.RUnlock()
	argsForCall := fake.storeResultArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRunState) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.artifactsMutex.RLock()
	defer fake.artifactsMutex.RUnlock()
	fake.readPlanOutputMutex.RLock()
	defer fake.readPlanOutputMutex.RUnlock()
	fake.readUserInputMutex.RLock()
	defer fake.readUserInputMutex.RUnlock()
	fake.resultMutex.RLock()
	defer fake.resultMutex.RUnlock()
	fake.sendPlanOutputMutex.RLock()
	defer fake.sendPlanOutputMutex.RUnlock()
	fake.sendUserInputMutex.RLock()
	defer fake.sendUserInputMutex.RUnlock()
	fake.storeResultMutex.RLock()
	defer fake.storeResultMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRunState) recordInvocation(key string, args []interface{}) {
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

var _ exec.RunState = new(FakeRunState)
