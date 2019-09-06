// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"context"
	"io"
	"sync"

	"github.com/bipol/scrapedumper/pkg/worker"
)

type FakeScraper struct {
	PrefixStub        func() string
	prefixMutex       sync.RWMutex
	prefixArgsForCall []struct {
	}
	prefixReturns struct {
		result1 string
	}
	prefixReturnsOnCall map[int]struct {
		result1 string
	}
	ScrapeStub        func(context.Context) (io.ReadCloser, error)
	scrapeMutex       sync.RWMutex
	scrapeArgsForCall []struct {
		arg1 context.Context
	}
	scrapeReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	scrapeReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeScraper) Prefix() string {
	fake.prefixMutex.Lock()
	ret, specificReturn := fake.prefixReturnsOnCall[len(fake.prefixArgsForCall)]
	fake.prefixArgsForCall = append(fake.prefixArgsForCall, struct {
	}{})
	fake.recordInvocation("Prefix", []interface{}{})
	fake.prefixMutex.Unlock()
	if fake.PrefixStub != nil {
		return fake.PrefixStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.prefixReturns
	return fakeReturns.result1
}

func (fake *FakeScraper) PrefixCallCount() int {
	fake.prefixMutex.RLock()
	defer fake.prefixMutex.RUnlock()
	return len(fake.prefixArgsForCall)
}

func (fake *FakeScraper) PrefixCalls(stub func() string) {
	fake.prefixMutex.Lock()
	defer fake.prefixMutex.Unlock()
	fake.PrefixStub = stub
}

func (fake *FakeScraper) PrefixReturns(result1 string) {
	fake.prefixMutex.Lock()
	defer fake.prefixMutex.Unlock()
	fake.PrefixStub = nil
	fake.prefixReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeScraper) PrefixReturnsOnCall(i int, result1 string) {
	fake.prefixMutex.Lock()
	defer fake.prefixMutex.Unlock()
	fake.PrefixStub = nil
	if fake.prefixReturnsOnCall == nil {
		fake.prefixReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.prefixReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeScraper) Scrape(arg1 context.Context) (io.ReadCloser, error) {
	fake.scrapeMutex.Lock()
	ret, specificReturn := fake.scrapeReturnsOnCall[len(fake.scrapeArgsForCall)]
	fake.scrapeArgsForCall = append(fake.scrapeArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Scrape", []interface{}{arg1})
	fake.scrapeMutex.Unlock()
	if fake.ScrapeStub != nil {
		return fake.ScrapeStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.scrapeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeScraper) ScrapeCallCount() int {
	fake.scrapeMutex.RLock()
	defer fake.scrapeMutex.RUnlock()
	return len(fake.scrapeArgsForCall)
}

func (fake *FakeScraper) ScrapeCalls(stub func(context.Context) (io.ReadCloser, error)) {
	fake.scrapeMutex.Lock()
	defer fake.scrapeMutex.Unlock()
	fake.ScrapeStub = stub
}

func (fake *FakeScraper) ScrapeArgsForCall(i int) context.Context {
	fake.scrapeMutex.RLock()
	defer fake.scrapeMutex.RUnlock()
	argsForCall := fake.scrapeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeScraper) ScrapeReturns(result1 io.ReadCloser, result2 error) {
	fake.scrapeMutex.Lock()
	defer fake.scrapeMutex.Unlock()
	fake.ScrapeStub = nil
	fake.scrapeReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeScraper) ScrapeReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.scrapeMutex.Lock()
	defer fake.scrapeMutex.Unlock()
	fake.ScrapeStub = nil
	if fake.scrapeReturnsOnCall == nil {
		fake.scrapeReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.scrapeReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeScraper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.prefixMutex.RLock()
	defer fake.prefixMutex.RUnlock()
	fake.scrapeMutex.RLock()
	defer fake.scrapeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeScraper) recordInvocation(key string, args []interface{}) {
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

var _ worker.Scraper = new(FakeScraper)