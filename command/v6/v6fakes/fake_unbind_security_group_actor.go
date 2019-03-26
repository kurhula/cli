// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/constant"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeUnbindSecurityGroupActor struct {
	UnbindSecurityGroupByNameAndSpaceStub        func(string, string, constant.SecurityGroupLifecycle) (v2action.Warnings, error)
	unbindSecurityGroupByNameAndSpaceMutex       sync.RWMutex
	unbindSecurityGroupByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 constant.SecurityGroupLifecycle
	}
	unbindSecurityGroupByNameAndSpaceReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	unbindSecurityGroupByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	UnbindSecurityGroupByNameOrganizationNameAndSpaceNameStub        func(string, string, string, constant.SecurityGroupLifecycle) (v2action.Warnings, error)
	unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex       sync.RWMutex
	unbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 constant.SecurityGroupLifecycle
	}
	unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameAndSpace(arg1 string, arg2 string, arg3 constant.SecurityGroupLifecycle) (v2action.Warnings, error) {
	fake.unbindSecurityGroupByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.unbindSecurityGroupByNameAndSpaceReturnsOnCall[len(fake.unbindSecurityGroupByNameAndSpaceArgsForCall)]
	fake.unbindSecurityGroupByNameAndSpaceArgsForCall = append(fake.unbindSecurityGroupByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 constant.SecurityGroupLifecycle
	}{arg1, arg2, arg3})
	fake.recordInvocation("UnbindSecurityGroupByNameAndSpace", []interface{}{arg1, arg2, arg3})
	fake.unbindSecurityGroupByNameAndSpaceMutex.Unlock()
	if fake.UnbindSecurityGroupByNameAndSpaceStub != nil {
		return fake.UnbindSecurityGroupByNameAndSpaceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.unbindSecurityGroupByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameAndSpaceCallCount() int {
	fake.unbindSecurityGroupByNameAndSpaceMutex.RLock()
	defer fake.unbindSecurityGroupByNameAndSpaceMutex.RUnlock()
	return len(fake.unbindSecurityGroupByNameAndSpaceArgsForCall)
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameAndSpaceCalls(stub func(string, string, constant.SecurityGroupLifecycle) (v2action.Warnings, error)) {
	fake.unbindSecurityGroupByNameAndSpaceMutex.Lock()
	defer fake.unbindSecurityGroupByNameAndSpaceMutex.Unlock()
	fake.UnbindSecurityGroupByNameAndSpaceStub = stub
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameAndSpaceArgsForCall(i int) (string, string, constant.SecurityGroupLifecycle) {
	fake.unbindSecurityGroupByNameAndSpaceMutex.RLock()
	defer fake.unbindSecurityGroupByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.unbindSecurityGroupByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameAndSpaceReturns(result1 v2action.Warnings, result2 error) {
	fake.unbindSecurityGroupByNameAndSpaceMutex.Lock()
	defer fake.unbindSecurityGroupByNameAndSpaceMutex.Unlock()
	fake.UnbindSecurityGroupByNameAndSpaceStub = nil
	fake.unbindSecurityGroupByNameAndSpaceReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameAndSpaceReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.unbindSecurityGroupByNameAndSpaceMutex.Lock()
	defer fake.unbindSecurityGroupByNameAndSpaceMutex.Unlock()
	fake.UnbindSecurityGroupByNameAndSpaceStub = nil
	if fake.unbindSecurityGroupByNameAndSpaceReturnsOnCall == nil {
		fake.unbindSecurityGroupByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.unbindSecurityGroupByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameOrganizationNameAndSpaceName(arg1 string, arg2 string, arg3 string, arg4 constant.SecurityGroupLifecycle) (v2action.Warnings, error) {
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.Lock()
	ret, specificReturn := fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturnsOnCall[len(fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall)]
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall = append(fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 constant.SecurityGroupLifecycle
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("UnbindSecurityGroupByNameOrganizationNameAndSpaceName", []interface{}{arg1, arg2, arg3, arg4})
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.Unlock()
	if fake.UnbindSecurityGroupByNameOrganizationNameAndSpaceNameStub != nil {
		return fake.UnbindSecurityGroupByNameOrganizationNameAndSpaceNameStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameOrganizationNameAndSpaceNameCallCount() int {
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.RLock()
	defer fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.RUnlock()
	return len(fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall)
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameOrganizationNameAndSpaceNameCalls(stub func(string, string, string, constant.SecurityGroupLifecycle) (v2action.Warnings, error)) {
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.Lock()
	defer fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.Unlock()
	fake.UnbindSecurityGroupByNameOrganizationNameAndSpaceNameStub = stub
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall(i int) (string, string, string, constant.SecurityGroupLifecycle) {
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.RLock()
	defer fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.RUnlock()
	argsForCall := fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameOrganizationNameAndSpaceNameReturns(result1 v2action.Warnings, result2 error) {
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.Lock()
	defer fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.Unlock()
	fake.UnbindSecurityGroupByNameOrganizationNameAndSpaceNameStub = nil
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUnbindSecurityGroupActor) UnbindSecurityGroupByNameOrganizationNameAndSpaceNameReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.Lock()
	defer fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.Unlock()
	fake.UnbindSecurityGroupByNameOrganizationNameAndSpaceNameStub = nil
	if fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturnsOnCall == nil {
		fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUnbindSecurityGroupActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.unbindSecurityGroupByNameAndSpaceMutex.RLock()
	defer fake.unbindSecurityGroupByNameAndSpaceMutex.RUnlock()
	fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.RLock()
	defer fake.unbindSecurityGroupByNameOrganizationNameAndSpaceNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUnbindSecurityGroupActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.UnbindSecurityGroupActor = new(FakeUnbindSecurityGroupActor)
