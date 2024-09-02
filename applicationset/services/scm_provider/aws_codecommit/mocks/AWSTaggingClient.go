// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	request "github.com/aws/aws-sdk-go/aws/request"
	mock "github.com/stretchr/testify/mock"

	resourcegroupstaggingapi "github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
)

// AWSTaggingClient is an autogenerated mock type for the AWSTaggingClient type
type AWSTaggingClient struct {
	mock.Mock
}

// GetResourcesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *AWSTaggingClient) GetResourcesWithContext(_a0 context.Context, _a1 *resourcegroupstaggingapi.GetResourcesInput, _a2 ...request.Option) (*resourcegroupstaggingapi.GetResourcesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetResourcesWithContext")
	}

	var r0 *resourcegroupstaggingapi.GetResourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.GetResourcesInput, ...request.Option) (*resourcegroupstaggingapi.GetResourcesOutput, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *resourcegroupstaggingapi.GetResourcesInput, ...request.Option) *resourcegroupstaggingapi.GetResourcesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resourcegroupstaggingapi.GetResourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *resourcegroupstaggingapi.GetResourcesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAWSTaggingClient creates a new instance of AWSTaggingClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAWSTaggingClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *AWSTaggingClient {
	mock := &AWSTaggingClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}