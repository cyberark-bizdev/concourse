// Code generated by counterfeiter. DO NOT EDIT.
package tokenfakes

import (
	sync "sync"

	token "github.com/concourse/concourse/skymarshal/token"
	oauth2 "golang.org/x/oauth2"
)

type FakeIssuer struct {
	IssueStub        func(*token.VerifiedClaims) (*oauth2.Token, error)
	issueMutex       sync.RWMutex
	issueArgsForCall []struct {
		arg1 *token.VerifiedClaims
	}
	issueReturns struct {
		result1 *oauth2.Token
		result2 error
	}
	issueReturnsOnCall map[int]struct {
		result1 *oauth2.Token
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIssuer) Issue(arg1 *token.VerifiedClaims) (*oauth2.Token, error) {
	fake.issueMutex.Lock()
	ret, specificReturn := fake.issueReturnsOnCall[len(fake.issueArgsForCall)]
	fake.issueArgsForCall = append(fake.issueArgsForCall, struct {
		arg1 *token.VerifiedClaims
	}{arg1})
	fake.recordInvocation("Issue", []interface{}{arg1})
	fake.issueMutex.Unlock()
	if fake.IssueStub != nil {
		return fake.IssueStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.issueReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIssuer) IssueCallCount() int {
	fake.issueMutex.RLock()
	defer fake.issueMutex.RUnlock()
	return len(fake.issueArgsForCall)
}

func (fake *FakeIssuer) IssueArgsForCall(i int) *token.VerifiedClaims {
	fake.issueMutex.RLock()
	defer fake.issueMutex.RUnlock()
	argsForCall := fake.issueArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIssuer) IssueReturns(result1 *oauth2.Token, result2 error) {
	fake.IssueStub = nil
	fake.issueReturns = struct {
		result1 *oauth2.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeIssuer) IssueReturnsOnCall(i int, result1 *oauth2.Token, result2 error) {
	fake.IssueStub = nil
	if fake.issueReturnsOnCall == nil {
		fake.issueReturnsOnCall = make(map[int]struct {
			result1 *oauth2.Token
			result2 error
		})
	}
	fake.issueReturnsOnCall[i] = struct {
		result1 *oauth2.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeIssuer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.issueMutex.RLock()
	defer fake.issueMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIssuer) recordInvocation(key string, args []interface{}) {
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

var _ token.Issuer = new(FakeIssuer)
