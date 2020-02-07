// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeUnsetSpaceQuotaActor struct {
	UnsetSpaceQuotaStub        func(string, string, string) (v7action.Warnings, error)
	unsetSpaceQuotaMutex       sync.RWMutex
	unsetSpaceQuotaArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	unsetSpaceQuotaReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	unsetSpaceQuotaReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUnsetSpaceQuotaActor) UnsetSpaceQuota(arg1 string, arg2 string, arg3 string) (v7action.Warnings, error) {
	fake.unsetSpaceQuotaMutex.Lock()
	ret, specificReturn := fake.unsetSpaceQuotaReturnsOnCall[len(fake.unsetSpaceQuotaArgsForCall)]
	fake.unsetSpaceQuotaArgsForCall = append(fake.unsetSpaceQuotaArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("UnsetSpaceQuota", []interface{}{arg1, arg2, arg3})
	fake.unsetSpaceQuotaMutex.Unlock()
	if fake.UnsetSpaceQuotaStub != nil {
		return fake.UnsetSpaceQuotaStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.unsetSpaceQuotaReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUnsetSpaceQuotaActor) UnsetSpaceQuotaCallCount() int {
	fake.unsetSpaceQuotaMutex.RLock()
	defer fake.unsetSpaceQuotaMutex.RUnlock()
	return len(fake.unsetSpaceQuotaArgsForCall)
}

func (fake *FakeUnsetSpaceQuotaActor) UnsetSpaceQuotaCalls(stub func(string, string, string) (v7action.Warnings, error)) {
	fake.unsetSpaceQuotaMutex.Lock()
	defer fake.unsetSpaceQuotaMutex.Unlock()
	fake.UnsetSpaceQuotaStub = stub
}

func (fake *FakeUnsetSpaceQuotaActor) UnsetSpaceQuotaArgsForCall(i int) (string, string, string) {
	fake.unsetSpaceQuotaMutex.RLock()
	defer fake.unsetSpaceQuotaMutex.RUnlock()
	argsForCall := fake.unsetSpaceQuotaArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUnsetSpaceQuotaActor) UnsetSpaceQuotaReturns(result1 v7action.Warnings, result2 error) {
	fake.unsetSpaceQuotaMutex.Lock()
	defer fake.unsetSpaceQuotaMutex.Unlock()
	fake.UnsetSpaceQuotaStub = nil
	fake.unsetSpaceQuotaReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUnsetSpaceQuotaActor) UnsetSpaceQuotaReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.unsetSpaceQuotaMutex.Lock()
	defer fake.unsetSpaceQuotaMutex.Unlock()
	fake.UnsetSpaceQuotaStub = nil
	if fake.unsetSpaceQuotaReturnsOnCall == nil {
		fake.unsetSpaceQuotaReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.unsetSpaceQuotaReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUnsetSpaceQuotaActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.unsetSpaceQuotaMutex.RLock()
	defer fake.unsetSpaceQuotaMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUnsetSpaceQuotaActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.UnsetSpaceQuotaActor = new(FakeUnsetSpaceQuotaActor)
