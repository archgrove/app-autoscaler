// This file was generated by counterfeiter
package fakes

import (
	"autoscaler/metricscollector/noaa"
	"sync"

	"github.com/cloudfoundry/sonde-go/events"
)

type FakeNoaaConsumer struct {
	ContainerEnvelopesStub        func(appGuid string, authToken string) ([]*events.Envelope, error)
	containerEnvelopesMutex       sync.RWMutex
	containerEnvelopesArgsForCall []struct {
		appGuid   string
		authToken string
	}
	containerEnvelopesReturns struct {
		result1 []*events.Envelope
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNoaaConsumer) ContainerEnvelopes(appGuid string, authToken string) ([]*events.Envelope, error) {
	fake.containerEnvelopesMutex.Lock()
	fake.containerEnvelopesArgsForCall = append(fake.containerEnvelopesArgsForCall, struct {
		appGuid   string
		authToken string
	}{appGuid, authToken})
	fake.recordInvocation("ContainerEnvelopes", []interface{}{appGuid, authToken})
	fake.containerEnvelopesMutex.Unlock()
	if fake.ContainerEnvelopesStub != nil {
		return fake.ContainerEnvelopesStub(appGuid, authToken)
	} else {
		return fake.containerEnvelopesReturns.result1, fake.containerEnvelopesReturns.result2
	}
}

func (fake *FakeNoaaConsumer) ContainerEnvelopesCallCount() int {
	fake.containerEnvelopesMutex.RLock()
	defer fake.containerEnvelopesMutex.RUnlock()
	return len(fake.containerEnvelopesArgsForCall)
}

func (fake *FakeNoaaConsumer) ContainerEnvelopesArgsForCall(i int) (string, string) {
	fake.containerEnvelopesMutex.RLock()
	defer fake.containerEnvelopesMutex.RUnlock()
	return fake.containerEnvelopesArgsForCall[i].appGuid, fake.containerEnvelopesArgsForCall[i].authToken
}

func (fake *FakeNoaaConsumer) ContainerEnvelopesReturns(result1 []*events.Envelope, result2 error) {
	fake.ContainerEnvelopesStub = nil
	fake.containerEnvelopesReturns = struct {
		result1 []*events.Envelope
		result2 error
	}{result1, result2}
}

func (fake *FakeNoaaConsumer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.containerEnvelopesMutex.RLock()
	defer fake.containerEnvelopesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeNoaaConsumer) recordInvocation(key string, args []interface{}) {
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

var _ noaa.NoaaConsumer = new(FakeNoaaConsumer)