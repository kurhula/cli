// Code generated by counterfeiter. DO NOT EDIT.
package cfnetworkingactionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/cfnetworkingaction"
	"code.cloudfoundry.org/cli/actor/v3action"
)

type FakeV3Actor struct {
	GetApplicationByNameAndSpaceStub        func(string, string) (v3action.Application, v3action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	GetApplicationsByGUIDsStub        func(...string) ([]v3action.Application, v3action.Warnings, error)
	getApplicationsByGUIDsMutex       sync.RWMutex
	getApplicationsByGUIDsArgsForCall []struct {
		arg1 []string
	}
	getApplicationsByGUIDsReturns struct {
		result1 []v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	getApplicationsByGUIDsReturnsOnCall map[int]struct {
		result1 []v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	GetApplicationsBySpaceStub        func(string) ([]v3action.Application, v3action.Warnings, error)
	getApplicationsBySpaceMutex       sync.RWMutex
	getApplicationsBySpaceArgsForCall []struct {
		arg1 string
	}
	getApplicationsBySpaceReturns struct {
		result1 []v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	getApplicationsBySpaceReturnsOnCall map[int]struct {
		result1 []v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	GetOrganizationByNameStub        func(string) (v3action.Organization, v3action.Warnings, error)
	getOrganizationByNameMutex       sync.RWMutex
	getOrganizationByNameArgsForCall []struct {
		arg1 string
	}
	getOrganizationByNameReturns struct {
		result1 v3action.Organization
		result2 v3action.Warnings
		result3 error
	}
	getOrganizationByNameReturnsOnCall map[int]struct {
		result1 v3action.Organization
		result2 v3action.Warnings
		result3 error
	}
	GetOrganizationsByGUIDsStub        func(...string) ([]v3action.Organization, v3action.Warnings, error)
	getOrganizationsByGUIDsMutex       sync.RWMutex
	getOrganizationsByGUIDsArgsForCall []struct {
		arg1 []string
	}
	getOrganizationsByGUIDsReturns struct {
		result1 []v3action.Organization
		result2 v3action.Warnings
		result3 error
	}
	getOrganizationsByGUIDsReturnsOnCall map[int]struct {
		result1 []v3action.Organization
		result2 v3action.Warnings
		result3 error
	}
	GetSpaceByNameAndOrganizationStub        func(string, string) (v3action.Space, v3action.Warnings, error)
	getSpaceByNameAndOrganizationMutex       sync.RWMutex
	getSpaceByNameAndOrganizationArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getSpaceByNameAndOrganizationReturns struct {
		result1 v3action.Space
		result2 v3action.Warnings
		result3 error
	}
	getSpaceByNameAndOrganizationReturnsOnCall map[int]struct {
		result1 v3action.Space
		result2 v3action.Warnings
		result3 error
	}
	GetSpacesByGUIDsStub        func(...string) ([]v3action.Space, v3action.Warnings, error)
	getSpacesByGUIDsMutex       sync.RWMutex
	getSpacesByGUIDsArgsForCall []struct {
		arg1 []string
	}
	getSpacesByGUIDsReturns struct {
		result1 []v3action.Space
		result2 v3action.Warnings
		result3 error
	}
	getSpacesByGUIDsReturnsOnCall map[int]struct {
		result1 []v3action.Space
		result2 v3action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpace(arg1 string, arg2 string) (v3action.Application, v3action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{arg1, arg2})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceCalls(stub func(string, string) (v3action.Application, v3action.Warnings, error)) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = stub
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceReturns(result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetApplicationsByGUIDs(arg1 ...string) ([]v3action.Application, v3action.Warnings, error) {
	fake.getApplicationsByGUIDsMutex.Lock()
	ret, specificReturn := fake.getApplicationsByGUIDsReturnsOnCall[len(fake.getApplicationsByGUIDsArgsForCall)]
	fake.getApplicationsByGUIDsArgsForCall = append(fake.getApplicationsByGUIDsArgsForCall, struct {
		arg1 []string
	}{arg1})
	fake.recordInvocation("GetApplicationsByGUIDs", []interface{}{arg1})
	fake.getApplicationsByGUIDsMutex.Unlock()
	if fake.GetApplicationsByGUIDsStub != nil {
		return fake.GetApplicationsByGUIDsStub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationsByGUIDsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV3Actor) GetApplicationsByGUIDsCallCount() int {
	fake.getApplicationsByGUIDsMutex.RLock()
	defer fake.getApplicationsByGUIDsMutex.RUnlock()
	return len(fake.getApplicationsByGUIDsArgsForCall)
}

func (fake *FakeV3Actor) GetApplicationsByGUIDsCalls(stub func(...string) ([]v3action.Application, v3action.Warnings, error)) {
	fake.getApplicationsByGUIDsMutex.Lock()
	defer fake.getApplicationsByGUIDsMutex.Unlock()
	fake.GetApplicationsByGUIDsStub = stub
}

func (fake *FakeV3Actor) GetApplicationsByGUIDsArgsForCall(i int) []string {
	fake.getApplicationsByGUIDsMutex.RLock()
	defer fake.getApplicationsByGUIDsMutex.RUnlock()
	argsForCall := fake.getApplicationsByGUIDsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeV3Actor) GetApplicationsByGUIDsReturns(result1 []v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.getApplicationsByGUIDsMutex.Lock()
	defer fake.getApplicationsByGUIDsMutex.Unlock()
	fake.GetApplicationsByGUIDsStub = nil
	fake.getApplicationsByGUIDsReturns = struct {
		result1 []v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetApplicationsByGUIDsReturnsOnCall(i int, result1 []v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.getApplicationsByGUIDsMutex.Lock()
	defer fake.getApplicationsByGUIDsMutex.Unlock()
	fake.GetApplicationsByGUIDsStub = nil
	if fake.getApplicationsByGUIDsReturnsOnCall == nil {
		fake.getApplicationsByGUIDsReturnsOnCall = make(map[int]struct {
			result1 []v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getApplicationsByGUIDsReturnsOnCall[i] = struct {
		result1 []v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetApplicationsBySpace(arg1 string) ([]v3action.Application, v3action.Warnings, error) {
	fake.getApplicationsBySpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationsBySpaceReturnsOnCall[len(fake.getApplicationsBySpaceArgsForCall)]
	fake.getApplicationsBySpaceArgsForCall = append(fake.getApplicationsBySpaceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetApplicationsBySpace", []interface{}{arg1})
	fake.getApplicationsBySpaceMutex.Unlock()
	if fake.GetApplicationsBySpaceStub != nil {
		return fake.GetApplicationsBySpaceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationsBySpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV3Actor) GetApplicationsBySpaceCallCount() int {
	fake.getApplicationsBySpaceMutex.RLock()
	defer fake.getApplicationsBySpaceMutex.RUnlock()
	return len(fake.getApplicationsBySpaceArgsForCall)
}

func (fake *FakeV3Actor) GetApplicationsBySpaceCalls(stub func(string) ([]v3action.Application, v3action.Warnings, error)) {
	fake.getApplicationsBySpaceMutex.Lock()
	defer fake.getApplicationsBySpaceMutex.Unlock()
	fake.GetApplicationsBySpaceStub = stub
}

func (fake *FakeV3Actor) GetApplicationsBySpaceArgsForCall(i int) string {
	fake.getApplicationsBySpaceMutex.RLock()
	defer fake.getApplicationsBySpaceMutex.RUnlock()
	argsForCall := fake.getApplicationsBySpaceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeV3Actor) GetApplicationsBySpaceReturns(result1 []v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.getApplicationsBySpaceMutex.Lock()
	defer fake.getApplicationsBySpaceMutex.Unlock()
	fake.GetApplicationsBySpaceStub = nil
	fake.getApplicationsBySpaceReturns = struct {
		result1 []v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetApplicationsBySpaceReturnsOnCall(i int, result1 []v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.getApplicationsBySpaceMutex.Lock()
	defer fake.getApplicationsBySpaceMutex.Unlock()
	fake.GetApplicationsBySpaceStub = nil
	if fake.getApplicationsBySpaceReturnsOnCall == nil {
		fake.getApplicationsBySpaceReturnsOnCall = make(map[int]struct {
			result1 []v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getApplicationsBySpaceReturnsOnCall[i] = struct {
		result1 []v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetOrganizationByName(arg1 string) (v3action.Organization, v3action.Warnings, error) {
	fake.getOrganizationByNameMutex.Lock()
	ret, specificReturn := fake.getOrganizationByNameReturnsOnCall[len(fake.getOrganizationByNameArgsForCall)]
	fake.getOrganizationByNameArgsForCall = append(fake.getOrganizationByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrganizationByName", []interface{}{arg1})
	fake.getOrganizationByNameMutex.Unlock()
	if fake.GetOrganizationByNameStub != nil {
		return fake.GetOrganizationByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationByNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV3Actor) GetOrganizationByNameCallCount() int {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return len(fake.getOrganizationByNameArgsForCall)
}

func (fake *FakeV3Actor) GetOrganizationByNameCalls(stub func(string) (v3action.Organization, v3action.Warnings, error)) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = stub
}

func (fake *FakeV3Actor) GetOrganizationByNameArgsForCall(i int) string {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	argsForCall := fake.getOrganizationByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeV3Actor) GetOrganizationByNameReturns(result1 v3action.Organization, result2 v3action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	fake.getOrganizationByNameReturns = struct {
		result1 v3action.Organization
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetOrganizationByNameReturnsOnCall(i int, result1 v3action.Organization, result2 v3action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	if fake.getOrganizationByNameReturnsOnCall == nil {
		fake.getOrganizationByNameReturnsOnCall = make(map[int]struct {
			result1 v3action.Organization
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getOrganizationByNameReturnsOnCall[i] = struct {
		result1 v3action.Organization
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetOrganizationsByGUIDs(arg1 ...string) ([]v3action.Organization, v3action.Warnings, error) {
	fake.getOrganizationsByGUIDsMutex.Lock()
	ret, specificReturn := fake.getOrganizationsByGUIDsReturnsOnCall[len(fake.getOrganizationsByGUIDsArgsForCall)]
	fake.getOrganizationsByGUIDsArgsForCall = append(fake.getOrganizationsByGUIDsArgsForCall, struct {
		arg1 []string
	}{arg1})
	fake.recordInvocation("GetOrganizationsByGUIDs", []interface{}{arg1})
	fake.getOrganizationsByGUIDsMutex.Unlock()
	if fake.GetOrganizationsByGUIDsStub != nil {
		return fake.GetOrganizationsByGUIDsStub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationsByGUIDsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV3Actor) GetOrganizationsByGUIDsCallCount() int {
	fake.getOrganizationsByGUIDsMutex.RLock()
	defer fake.getOrganizationsByGUIDsMutex.RUnlock()
	return len(fake.getOrganizationsByGUIDsArgsForCall)
}

func (fake *FakeV3Actor) GetOrganizationsByGUIDsCalls(stub func(...string) ([]v3action.Organization, v3action.Warnings, error)) {
	fake.getOrganizationsByGUIDsMutex.Lock()
	defer fake.getOrganizationsByGUIDsMutex.Unlock()
	fake.GetOrganizationsByGUIDsStub = stub
}

func (fake *FakeV3Actor) GetOrganizationsByGUIDsArgsForCall(i int) []string {
	fake.getOrganizationsByGUIDsMutex.RLock()
	defer fake.getOrganizationsByGUIDsMutex.RUnlock()
	argsForCall := fake.getOrganizationsByGUIDsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeV3Actor) GetOrganizationsByGUIDsReturns(result1 []v3action.Organization, result2 v3action.Warnings, result3 error) {
	fake.getOrganizationsByGUIDsMutex.Lock()
	defer fake.getOrganizationsByGUIDsMutex.Unlock()
	fake.GetOrganizationsByGUIDsStub = nil
	fake.getOrganizationsByGUIDsReturns = struct {
		result1 []v3action.Organization
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetOrganizationsByGUIDsReturnsOnCall(i int, result1 []v3action.Organization, result2 v3action.Warnings, result3 error) {
	fake.getOrganizationsByGUIDsMutex.Lock()
	defer fake.getOrganizationsByGUIDsMutex.Unlock()
	fake.GetOrganizationsByGUIDsStub = nil
	if fake.getOrganizationsByGUIDsReturnsOnCall == nil {
		fake.getOrganizationsByGUIDsReturnsOnCall = make(map[int]struct {
			result1 []v3action.Organization
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getOrganizationsByGUIDsReturnsOnCall[i] = struct {
		result1 []v3action.Organization
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetSpaceByNameAndOrganization(arg1 string, arg2 string) (v3action.Space, v3action.Warnings, error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	ret, specificReturn := fake.getSpaceByNameAndOrganizationReturnsOnCall[len(fake.getSpaceByNameAndOrganizationArgsForCall)]
	fake.getSpaceByNameAndOrganizationArgsForCall = append(fake.getSpaceByNameAndOrganizationArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetSpaceByNameAndOrganization", []interface{}{arg1, arg2})
	fake.getSpaceByNameAndOrganizationMutex.Unlock()
	if fake.GetSpaceByNameAndOrganizationStub != nil {
		return fake.GetSpaceByNameAndOrganizationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getSpaceByNameAndOrganizationReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV3Actor) GetSpaceByNameAndOrganizationCallCount() int {
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	return len(fake.getSpaceByNameAndOrganizationArgsForCall)
}

func (fake *FakeV3Actor) GetSpaceByNameAndOrganizationCalls(stub func(string, string) (v3action.Space, v3action.Warnings, error)) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = stub
}

func (fake *FakeV3Actor) GetSpaceByNameAndOrganizationArgsForCall(i int) (string, string) {
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	argsForCall := fake.getSpaceByNameAndOrganizationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV3Actor) GetSpaceByNameAndOrganizationReturns(result1 v3action.Space, result2 v3action.Warnings, result3 error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = nil
	fake.getSpaceByNameAndOrganizationReturns = struct {
		result1 v3action.Space
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetSpaceByNameAndOrganizationReturnsOnCall(i int, result1 v3action.Space, result2 v3action.Warnings, result3 error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = nil
	if fake.getSpaceByNameAndOrganizationReturnsOnCall == nil {
		fake.getSpaceByNameAndOrganizationReturnsOnCall = make(map[int]struct {
			result1 v3action.Space
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getSpaceByNameAndOrganizationReturnsOnCall[i] = struct {
		result1 v3action.Space
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetSpacesByGUIDs(arg1 ...string) ([]v3action.Space, v3action.Warnings, error) {
	fake.getSpacesByGUIDsMutex.Lock()
	ret, specificReturn := fake.getSpacesByGUIDsReturnsOnCall[len(fake.getSpacesByGUIDsArgsForCall)]
	fake.getSpacesByGUIDsArgsForCall = append(fake.getSpacesByGUIDsArgsForCall, struct {
		arg1 []string
	}{arg1})
	fake.recordInvocation("GetSpacesByGUIDs", []interface{}{arg1})
	fake.getSpacesByGUIDsMutex.Unlock()
	if fake.GetSpacesByGUIDsStub != nil {
		return fake.GetSpacesByGUIDsStub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getSpacesByGUIDsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV3Actor) GetSpacesByGUIDsCallCount() int {
	fake.getSpacesByGUIDsMutex.RLock()
	defer fake.getSpacesByGUIDsMutex.RUnlock()
	return len(fake.getSpacesByGUIDsArgsForCall)
}

func (fake *FakeV3Actor) GetSpacesByGUIDsCalls(stub func(...string) ([]v3action.Space, v3action.Warnings, error)) {
	fake.getSpacesByGUIDsMutex.Lock()
	defer fake.getSpacesByGUIDsMutex.Unlock()
	fake.GetSpacesByGUIDsStub = stub
}

func (fake *FakeV3Actor) GetSpacesByGUIDsArgsForCall(i int) []string {
	fake.getSpacesByGUIDsMutex.RLock()
	defer fake.getSpacesByGUIDsMutex.RUnlock()
	argsForCall := fake.getSpacesByGUIDsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeV3Actor) GetSpacesByGUIDsReturns(result1 []v3action.Space, result2 v3action.Warnings, result3 error) {
	fake.getSpacesByGUIDsMutex.Lock()
	defer fake.getSpacesByGUIDsMutex.Unlock()
	fake.GetSpacesByGUIDsStub = nil
	fake.getSpacesByGUIDsReturns = struct {
		result1 []v3action.Space
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetSpacesByGUIDsReturnsOnCall(i int, result1 []v3action.Space, result2 v3action.Warnings, result3 error) {
	fake.getSpacesByGUIDsMutex.Lock()
	defer fake.getSpacesByGUIDsMutex.Unlock()
	fake.GetSpacesByGUIDsStub = nil
	if fake.getSpacesByGUIDsReturnsOnCall == nil {
		fake.getSpacesByGUIDsReturnsOnCall = make(map[int]struct {
			result1 []v3action.Space
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getSpacesByGUIDsReturnsOnCall[i] = struct {
		result1 []v3action.Space
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getApplicationsByGUIDsMutex.RLock()
	defer fake.getApplicationsByGUIDsMutex.RUnlock()
	fake.getApplicationsBySpaceMutex.RLock()
	defer fake.getApplicationsBySpaceMutex.RUnlock()
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	fake.getOrganizationsByGUIDsMutex.RLock()
	defer fake.getOrganizationsByGUIDsMutex.RUnlock()
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	fake.getSpacesByGUIDsMutex.RLock()
	defer fake.getSpacesByGUIDsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3Actor) recordInvocation(key string, args []interface{}) {
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

var _ cfnetworkingaction.V3Actor = new(FakeV3Actor)
