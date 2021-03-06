// Code generated by counterfeiter. DO NOT EDIT.
package bulkfakes

import (
	"os"
	"sync"

	"github.com/smartatransit/scrapedumper/pkg/bulk"
)

type FakeFile struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
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
	ReadStub        func([]byte) (int, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		arg1 []byte
	}
	readReturns struct {
		result1 int
		result2 error
	}
	readReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	ReadAtStub        func([]byte, int64) (int, error)
	readAtMutex       sync.RWMutex
	readAtArgsForCall []struct {
		arg1 []byte
		arg2 int64
	}
	readAtReturns struct {
		result1 int
		result2 error
	}
	readAtReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	ReaddirStub        func(int) ([]os.FileInfo, error)
	readdirMutex       sync.RWMutex
	readdirArgsForCall []struct {
		arg1 int
	}
	readdirReturns struct {
		result1 []os.FileInfo
		result2 error
	}
	readdirReturnsOnCall map[int]struct {
		result1 []os.FileInfo
		result2 error
	}
	ReaddirnamesStub        func(int) ([]string, error)
	readdirnamesMutex       sync.RWMutex
	readdirnamesArgsForCall []struct {
		arg1 int
	}
	readdirnamesReturns struct {
		result1 []string
		result2 error
	}
	readdirnamesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	SeekStub        func(int64, int) (int64, error)
	seekMutex       sync.RWMutex
	seekArgsForCall []struct {
		arg1 int64
		arg2 int
	}
	seekReturns struct {
		result1 int64
		result2 error
	}
	seekReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	StatStub        func() (os.FileInfo, error)
	statMutex       sync.RWMutex
	statArgsForCall []struct {
	}
	statReturns struct {
		result1 os.FileInfo
		result2 error
	}
	statReturnsOnCall map[int]struct {
		result1 os.FileInfo
		result2 error
	}
	SyncStub        func() error
	syncMutex       sync.RWMutex
	syncArgsForCall []struct {
	}
	syncReturns struct {
		result1 error
	}
	syncReturnsOnCall map[int]struct {
		result1 error
	}
	TruncateStub        func(int64) error
	truncateMutex       sync.RWMutex
	truncateArgsForCall []struct {
		arg1 int64
	}
	truncateReturns struct {
		result1 error
	}
	truncateReturnsOnCall map[int]struct {
		result1 error
	}
	WriteStub        func([]byte) (int, error)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		arg1 []byte
	}
	writeReturns struct {
		result1 int
		result2 error
	}
	writeReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	WriteAtStub        func([]byte, int64) (int, error)
	writeAtMutex       sync.RWMutex
	writeAtArgsForCall []struct {
		arg1 []byte
		arg2 int64
	}
	writeAtReturns struct {
		result1 int
		result2 error
	}
	writeAtReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	WriteStringStub        func(string) (int, error)
	writeStringMutex       sync.RWMutex
	writeStringArgsForCall []struct {
		arg1 string
	}
	writeStringReturns struct {
		result1 int
		result2 error
	}
	writeStringReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFile) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFile) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeFile) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeFile) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFile) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFile) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	stub := fake.NameStub
	fakeReturns := fake.nameReturns
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFile) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeFile) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeFile) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFile) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
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

func (fake *FakeFile) Read(arg1 []byte) (int, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.ReadStub
	fakeReturns := fake.readReturns
	fake.recordInvocation("Read", []interface{}{arg1Copy})
	fake.readMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFile) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeFile) ReadCalls(stub func([]byte) (int, error)) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = stub
}

func (fake *FakeFile) ReadArgsForCall(i int) []byte {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	argsForCall := fake.readArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFile) ReadReturns(result1 int, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) ReadReturnsOnCall(i int, result1 int, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) ReadAt(arg1 []byte, arg2 int64) (int, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.readAtMutex.Lock()
	ret, specificReturn := fake.readAtReturnsOnCall[len(fake.readAtArgsForCall)]
	fake.readAtArgsForCall = append(fake.readAtArgsForCall, struct {
		arg1 []byte
		arg2 int64
	}{arg1Copy, arg2})
	stub := fake.ReadAtStub
	fakeReturns := fake.readAtReturns
	fake.recordInvocation("ReadAt", []interface{}{arg1Copy, arg2})
	fake.readAtMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFile) ReadAtCallCount() int {
	fake.readAtMutex.RLock()
	defer fake.readAtMutex.RUnlock()
	return len(fake.readAtArgsForCall)
}

func (fake *FakeFile) ReadAtCalls(stub func([]byte, int64) (int, error)) {
	fake.readAtMutex.Lock()
	defer fake.readAtMutex.Unlock()
	fake.ReadAtStub = stub
}

func (fake *FakeFile) ReadAtArgsForCall(i int) ([]byte, int64) {
	fake.readAtMutex.RLock()
	defer fake.readAtMutex.RUnlock()
	argsForCall := fake.readAtArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFile) ReadAtReturns(result1 int, result2 error) {
	fake.readAtMutex.Lock()
	defer fake.readAtMutex.Unlock()
	fake.ReadAtStub = nil
	fake.readAtReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) ReadAtReturnsOnCall(i int, result1 int, result2 error) {
	fake.readAtMutex.Lock()
	defer fake.readAtMutex.Unlock()
	fake.ReadAtStub = nil
	if fake.readAtReturnsOnCall == nil {
		fake.readAtReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.readAtReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) Readdir(arg1 int) ([]os.FileInfo, error) {
	fake.readdirMutex.Lock()
	ret, specificReturn := fake.readdirReturnsOnCall[len(fake.readdirArgsForCall)]
	fake.readdirArgsForCall = append(fake.readdirArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.ReaddirStub
	fakeReturns := fake.readdirReturns
	fake.recordInvocation("Readdir", []interface{}{arg1})
	fake.readdirMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFile) ReaddirCallCount() int {
	fake.readdirMutex.RLock()
	defer fake.readdirMutex.RUnlock()
	return len(fake.readdirArgsForCall)
}

func (fake *FakeFile) ReaddirCalls(stub func(int) ([]os.FileInfo, error)) {
	fake.readdirMutex.Lock()
	defer fake.readdirMutex.Unlock()
	fake.ReaddirStub = stub
}

func (fake *FakeFile) ReaddirArgsForCall(i int) int {
	fake.readdirMutex.RLock()
	defer fake.readdirMutex.RUnlock()
	argsForCall := fake.readdirArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFile) ReaddirReturns(result1 []os.FileInfo, result2 error) {
	fake.readdirMutex.Lock()
	defer fake.readdirMutex.Unlock()
	fake.ReaddirStub = nil
	fake.readdirReturns = struct {
		result1 []os.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) ReaddirReturnsOnCall(i int, result1 []os.FileInfo, result2 error) {
	fake.readdirMutex.Lock()
	defer fake.readdirMutex.Unlock()
	fake.ReaddirStub = nil
	if fake.readdirReturnsOnCall == nil {
		fake.readdirReturnsOnCall = make(map[int]struct {
			result1 []os.FileInfo
			result2 error
		})
	}
	fake.readdirReturnsOnCall[i] = struct {
		result1 []os.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) Readdirnames(arg1 int) ([]string, error) {
	fake.readdirnamesMutex.Lock()
	ret, specificReturn := fake.readdirnamesReturnsOnCall[len(fake.readdirnamesArgsForCall)]
	fake.readdirnamesArgsForCall = append(fake.readdirnamesArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.ReaddirnamesStub
	fakeReturns := fake.readdirnamesReturns
	fake.recordInvocation("Readdirnames", []interface{}{arg1})
	fake.readdirnamesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFile) ReaddirnamesCallCount() int {
	fake.readdirnamesMutex.RLock()
	defer fake.readdirnamesMutex.RUnlock()
	return len(fake.readdirnamesArgsForCall)
}

func (fake *FakeFile) ReaddirnamesCalls(stub func(int) ([]string, error)) {
	fake.readdirnamesMutex.Lock()
	defer fake.readdirnamesMutex.Unlock()
	fake.ReaddirnamesStub = stub
}

func (fake *FakeFile) ReaddirnamesArgsForCall(i int) int {
	fake.readdirnamesMutex.RLock()
	defer fake.readdirnamesMutex.RUnlock()
	argsForCall := fake.readdirnamesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFile) ReaddirnamesReturns(result1 []string, result2 error) {
	fake.readdirnamesMutex.Lock()
	defer fake.readdirnamesMutex.Unlock()
	fake.ReaddirnamesStub = nil
	fake.readdirnamesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) ReaddirnamesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.readdirnamesMutex.Lock()
	defer fake.readdirnamesMutex.Unlock()
	fake.ReaddirnamesStub = nil
	if fake.readdirnamesReturnsOnCall == nil {
		fake.readdirnamesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.readdirnamesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) Seek(arg1 int64, arg2 int) (int64, error) {
	fake.seekMutex.Lock()
	ret, specificReturn := fake.seekReturnsOnCall[len(fake.seekArgsForCall)]
	fake.seekArgsForCall = append(fake.seekArgsForCall, struct {
		arg1 int64
		arg2 int
	}{arg1, arg2})
	stub := fake.SeekStub
	fakeReturns := fake.seekReturns
	fake.recordInvocation("Seek", []interface{}{arg1, arg2})
	fake.seekMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFile) SeekCallCount() int {
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	return len(fake.seekArgsForCall)
}

func (fake *FakeFile) SeekCalls(stub func(int64, int) (int64, error)) {
	fake.seekMutex.Lock()
	defer fake.seekMutex.Unlock()
	fake.SeekStub = stub
}

func (fake *FakeFile) SeekArgsForCall(i int) (int64, int) {
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	argsForCall := fake.seekArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFile) SeekReturns(result1 int64, result2 error) {
	fake.seekMutex.Lock()
	defer fake.seekMutex.Unlock()
	fake.SeekStub = nil
	fake.seekReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) SeekReturnsOnCall(i int, result1 int64, result2 error) {
	fake.seekMutex.Lock()
	defer fake.seekMutex.Unlock()
	fake.SeekStub = nil
	if fake.seekReturnsOnCall == nil {
		fake.seekReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.seekReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) Stat() (os.FileInfo, error) {
	fake.statMutex.Lock()
	ret, specificReturn := fake.statReturnsOnCall[len(fake.statArgsForCall)]
	fake.statArgsForCall = append(fake.statArgsForCall, struct {
	}{})
	stub := fake.StatStub
	fakeReturns := fake.statReturns
	fake.recordInvocation("Stat", []interface{}{})
	fake.statMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFile) StatCallCount() int {
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	return len(fake.statArgsForCall)
}

func (fake *FakeFile) StatCalls(stub func() (os.FileInfo, error)) {
	fake.statMutex.Lock()
	defer fake.statMutex.Unlock()
	fake.StatStub = stub
}

func (fake *FakeFile) StatReturns(result1 os.FileInfo, result2 error) {
	fake.statMutex.Lock()
	defer fake.statMutex.Unlock()
	fake.StatStub = nil
	fake.statReturns = struct {
		result1 os.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) StatReturnsOnCall(i int, result1 os.FileInfo, result2 error) {
	fake.statMutex.Lock()
	defer fake.statMutex.Unlock()
	fake.StatStub = nil
	if fake.statReturnsOnCall == nil {
		fake.statReturnsOnCall = make(map[int]struct {
			result1 os.FileInfo
			result2 error
		})
	}
	fake.statReturnsOnCall[i] = struct {
		result1 os.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) Sync() error {
	fake.syncMutex.Lock()
	ret, specificReturn := fake.syncReturnsOnCall[len(fake.syncArgsForCall)]
	fake.syncArgsForCall = append(fake.syncArgsForCall, struct {
	}{})
	stub := fake.SyncStub
	fakeReturns := fake.syncReturns
	fake.recordInvocation("Sync", []interface{}{})
	fake.syncMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFile) SyncCallCount() int {
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	return len(fake.syncArgsForCall)
}

func (fake *FakeFile) SyncCalls(stub func() error) {
	fake.syncMutex.Lock()
	defer fake.syncMutex.Unlock()
	fake.SyncStub = stub
}

func (fake *FakeFile) SyncReturns(result1 error) {
	fake.syncMutex.Lock()
	defer fake.syncMutex.Unlock()
	fake.SyncStub = nil
	fake.syncReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFile) SyncReturnsOnCall(i int, result1 error) {
	fake.syncMutex.Lock()
	defer fake.syncMutex.Unlock()
	fake.SyncStub = nil
	if fake.syncReturnsOnCall == nil {
		fake.syncReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.syncReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFile) Truncate(arg1 int64) error {
	fake.truncateMutex.Lock()
	ret, specificReturn := fake.truncateReturnsOnCall[len(fake.truncateArgsForCall)]
	fake.truncateArgsForCall = append(fake.truncateArgsForCall, struct {
		arg1 int64
	}{arg1})
	stub := fake.TruncateStub
	fakeReturns := fake.truncateReturns
	fake.recordInvocation("Truncate", []interface{}{arg1})
	fake.truncateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFile) TruncateCallCount() int {
	fake.truncateMutex.RLock()
	defer fake.truncateMutex.RUnlock()
	return len(fake.truncateArgsForCall)
}

func (fake *FakeFile) TruncateCalls(stub func(int64) error) {
	fake.truncateMutex.Lock()
	defer fake.truncateMutex.Unlock()
	fake.TruncateStub = stub
}

func (fake *FakeFile) TruncateArgsForCall(i int) int64 {
	fake.truncateMutex.RLock()
	defer fake.truncateMutex.RUnlock()
	argsForCall := fake.truncateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFile) TruncateReturns(result1 error) {
	fake.truncateMutex.Lock()
	defer fake.truncateMutex.Unlock()
	fake.TruncateStub = nil
	fake.truncateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFile) TruncateReturnsOnCall(i int, result1 error) {
	fake.truncateMutex.Lock()
	defer fake.truncateMutex.Unlock()
	fake.TruncateStub = nil
	if fake.truncateReturnsOnCall == nil {
		fake.truncateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.truncateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFile) Write(arg1 []byte) (int, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.WriteStub
	fakeReturns := fake.writeReturns
	fake.recordInvocation("Write", []interface{}{arg1Copy})
	fake.writeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFile) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeFile) WriteCalls(stub func([]byte) (int, error)) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = stub
}

func (fake *FakeFile) WriteArgsForCall(i int) []byte {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	argsForCall := fake.writeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFile) WriteReturns(result1 int, result2 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) WriteReturnsOnCall(i int, result1 int, result2 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) WriteAt(arg1 []byte, arg2 int64) (int, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.writeAtMutex.Lock()
	ret, specificReturn := fake.writeAtReturnsOnCall[len(fake.writeAtArgsForCall)]
	fake.writeAtArgsForCall = append(fake.writeAtArgsForCall, struct {
		arg1 []byte
		arg2 int64
	}{arg1Copy, arg2})
	stub := fake.WriteAtStub
	fakeReturns := fake.writeAtReturns
	fake.recordInvocation("WriteAt", []interface{}{arg1Copy, arg2})
	fake.writeAtMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFile) WriteAtCallCount() int {
	fake.writeAtMutex.RLock()
	defer fake.writeAtMutex.RUnlock()
	return len(fake.writeAtArgsForCall)
}

func (fake *FakeFile) WriteAtCalls(stub func([]byte, int64) (int, error)) {
	fake.writeAtMutex.Lock()
	defer fake.writeAtMutex.Unlock()
	fake.WriteAtStub = stub
}

func (fake *FakeFile) WriteAtArgsForCall(i int) ([]byte, int64) {
	fake.writeAtMutex.RLock()
	defer fake.writeAtMutex.RUnlock()
	argsForCall := fake.writeAtArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFile) WriteAtReturns(result1 int, result2 error) {
	fake.writeAtMutex.Lock()
	defer fake.writeAtMutex.Unlock()
	fake.WriteAtStub = nil
	fake.writeAtReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) WriteAtReturnsOnCall(i int, result1 int, result2 error) {
	fake.writeAtMutex.Lock()
	defer fake.writeAtMutex.Unlock()
	fake.WriteAtStub = nil
	if fake.writeAtReturnsOnCall == nil {
		fake.writeAtReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.writeAtReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) WriteString(arg1 string) (int, error) {
	fake.writeStringMutex.Lock()
	ret, specificReturn := fake.writeStringReturnsOnCall[len(fake.writeStringArgsForCall)]
	fake.writeStringArgsForCall = append(fake.writeStringArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.WriteStringStub
	fakeReturns := fake.writeStringReturns
	fake.recordInvocation("WriteString", []interface{}{arg1})
	fake.writeStringMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFile) WriteStringCallCount() int {
	fake.writeStringMutex.RLock()
	defer fake.writeStringMutex.RUnlock()
	return len(fake.writeStringArgsForCall)
}

func (fake *FakeFile) WriteStringCalls(stub func(string) (int, error)) {
	fake.writeStringMutex.Lock()
	defer fake.writeStringMutex.Unlock()
	fake.WriteStringStub = stub
}

func (fake *FakeFile) WriteStringArgsForCall(i int) string {
	fake.writeStringMutex.RLock()
	defer fake.writeStringMutex.RUnlock()
	argsForCall := fake.writeStringArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFile) WriteStringReturns(result1 int, result2 error) {
	fake.writeStringMutex.Lock()
	defer fake.writeStringMutex.Unlock()
	fake.WriteStringStub = nil
	fake.writeStringReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) WriteStringReturnsOnCall(i int, result1 int, result2 error) {
	fake.writeStringMutex.Lock()
	defer fake.writeStringMutex.Unlock()
	fake.WriteStringStub = nil
	if fake.writeStringReturnsOnCall == nil {
		fake.writeStringReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.writeStringReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeFile) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	fake.readAtMutex.RLock()
	defer fake.readAtMutex.RUnlock()
	fake.readdirMutex.RLock()
	defer fake.readdirMutex.RUnlock()
	fake.readdirnamesMutex.RLock()
	defer fake.readdirnamesMutex.RUnlock()
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	fake.truncateMutex.RLock()
	defer fake.truncateMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	fake.writeAtMutex.RLock()
	defer fake.writeAtMutex.RUnlock()
	fake.writeStringMutex.RLock()
	defer fake.writeStringMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFile) recordInvocation(key string, args []interface{}) {
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

var _ bulk.File = new(FakeFile)
