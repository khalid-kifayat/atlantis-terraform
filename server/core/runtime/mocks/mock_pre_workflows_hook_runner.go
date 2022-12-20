// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/runtime (interfaces: PreWorkflowHookRunner)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockPreWorkflowHookRunner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockPreWorkflowHookRunner(options ...pegomock.Option) *MockPreWorkflowHookRunner {
	mock := &MockPreWorkflowHookRunner{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockPreWorkflowHookRunner) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockPreWorkflowHookRunner) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockPreWorkflowHookRunner) Run(ctx models.PreWorkflowHookCommandContext, command string, path string) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockPreWorkflowHookRunner().")
	}
	params := []pegomock.Param{ctx, command, path}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Run", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockPreWorkflowHookRunner) VerifyWasCalledOnce() *VerifierMockPreWorkflowHookRunner {
	return &VerifierMockPreWorkflowHookRunner{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockPreWorkflowHookRunner) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockPreWorkflowHookRunner {
	return &VerifierMockPreWorkflowHookRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockPreWorkflowHookRunner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockPreWorkflowHookRunner {
	return &VerifierMockPreWorkflowHookRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockPreWorkflowHookRunner) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockPreWorkflowHookRunner {
	return &VerifierMockPreWorkflowHookRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockPreWorkflowHookRunner struct {
	mock                   *MockPreWorkflowHookRunner
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockPreWorkflowHookRunner) Run(ctx models.PreWorkflowHookCommandContext, command string, path string) *MockPreWorkflowHookRunner_Run_OngoingVerification {
	params := []pegomock.Param{ctx, command, path}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Run", params, verifier.timeout)
	return &MockPreWorkflowHookRunner_Run_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockPreWorkflowHookRunner_Run_OngoingVerification struct {
	mock              *MockPreWorkflowHookRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockPreWorkflowHookRunner_Run_OngoingVerification) GetCapturedArguments() (models.PreWorkflowHookCommandContext, string, string) {
	ctx, command, path := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], command[len(command)-1], path[len(path)-1]
}

func (c *MockPreWorkflowHookRunner_Run_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PreWorkflowHookCommandContext, _param1 []string, _param2 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PreWorkflowHookCommandContext, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PreWorkflowHookCommandContext)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
	}
	return
}
