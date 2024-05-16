// Code generated by mockery v2.40.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/syself/hrobot-go/models"

	v1beta1 "github.com/syself/cluster-api-provider-hetzner/api/v1beta1"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// DeleteBootRescue provides a mock function with given fields: id
func (_m *Client) DeleteBootRescue(id int) (*models.Rescue, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBootRescue")
	}

	var r0 *models.Rescue
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.Rescue, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *models.Rescue); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Rescue)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_DeleteBootRescue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteBootRescue'
type Client_DeleteBootRescue_Call struct {
	*mock.Call
}

// DeleteBootRescue is a helper method to define mock.On call
//   - id int
func (_e *Client_Expecter) DeleteBootRescue(id interface{}) *Client_DeleteBootRescue_Call {
	return &Client_DeleteBootRescue_Call{Call: _e.mock.On("DeleteBootRescue", id)}
}

func (_c *Client_DeleteBootRescue_Call) Run(run func(id int)) *Client_DeleteBootRescue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Client_DeleteBootRescue_Call) Return(_a0 *models.Rescue, _a1 error) *Client_DeleteBootRescue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_DeleteBootRescue_Call) RunAndReturn(run func(int) (*models.Rescue, error)) *Client_DeleteBootRescue_Call {
	_c.Call.Return(run)
	return _c
}

// GetBMServer provides a mock function with given fields: _a0
func (_m *Client) GetBMServer(_a0 int) (*models.Server, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetBMServer")
	}

	var r0 *models.Server
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.Server, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) *models.Server); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Server)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_GetBMServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBMServer'
type Client_GetBMServer_Call struct {
	*mock.Call
}

// GetBMServer is a helper method to define mock.On call
//   - _a0 int
func (_e *Client_Expecter) GetBMServer(_a0 interface{}) *Client_GetBMServer_Call {
	return &Client_GetBMServer_Call{Call: _e.mock.On("GetBMServer", _a0)}
}

func (_c *Client_GetBMServer_Call) Run(run func(_a0 int)) *Client_GetBMServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Client_GetBMServer_Call) Return(_a0 *models.Server, _a1 error) *Client_GetBMServer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_GetBMServer_Call) RunAndReturn(run func(int) (*models.Server, error)) *Client_GetBMServer_Call {
	_c.Call.Return(run)
	return _c
}

// GetBootRescue provides a mock function with given fields: id
func (_m *Client) GetBootRescue(id int) (*models.Rescue, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetBootRescue")
	}

	var r0 *models.Rescue
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.Rescue, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *models.Rescue); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Rescue)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_GetBootRescue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBootRescue'
type Client_GetBootRescue_Call struct {
	*mock.Call
}

// GetBootRescue is a helper method to define mock.On call
//   - id int
func (_e *Client_Expecter) GetBootRescue(id interface{}) *Client_GetBootRescue_Call {
	return &Client_GetBootRescue_Call{Call: _e.mock.On("GetBootRescue", id)}
}

func (_c *Client_GetBootRescue_Call) Run(run func(id int)) *Client_GetBootRescue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Client_GetBootRescue_Call) Return(_a0 *models.Rescue, _a1 error) *Client_GetBootRescue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_GetBootRescue_Call) RunAndReturn(run func(int) (*models.Rescue, error)) *Client_GetBootRescue_Call {
	_c.Call.Return(run)
	return _c
}

// GetReboot provides a mock function with given fields: _a0
func (_m *Client) GetReboot(_a0 int) (*models.Reset, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetReboot")
	}

	var r0 *models.Reset
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.Reset, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) *models.Reset); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Reset)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_GetReboot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReboot'
type Client_GetReboot_Call struct {
	*mock.Call
}

// GetReboot is a helper method to define mock.On call
//   - _a0 int
func (_e *Client_Expecter) GetReboot(_a0 interface{}) *Client_GetReboot_Call {
	return &Client_GetReboot_Call{Call: _e.mock.On("GetReboot", _a0)}
}

func (_c *Client_GetReboot_Call) Run(run func(_a0 int)) *Client_GetReboot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Client_GetReboot_Call) Return(_a0 *models.Reset, _a1 error) *Client_GetReboot_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_GetReboot_Call) RunAndReturn(run func(int) (*models.Reset, error)) *Client_GetReboot_Call {
	_c.Call.Return(run)
	return _c
}

// ListBMServers provides a mock function with given fields:
func (_m *Client) ListBMServers() ([]models.Server, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListBMServers")
	}

	var r0 []models.Server
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Server, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Server); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Server)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_ListBMServers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListBMServers'
type Client_ListBMServers_Call struct {
	*mock.Call
}

// ListBMServers is a helper method to define mock.On call
func (_e *Client_Expecter) ListBMServers() *Client_ListBMServers_Call {
	return &Client_ListBMServers_Call{Call: _e.mock.On("ListBMServers")}
}

func (_c *Client_ListBMServers_Call) Run(run func()) *Client_ListBMServers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_ListBMServers_Call) Return(_a0 []models.Server, _a1 error) *Client_ListBMServers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_ListBMServers_Call) RunAndReturn(run func() ([]models.Server, error)) *Client_ListBMServers_Call {
	_c.Call.Return(run)
	return _c
}

// ListSSHKeys provides a mock function with given fields:
func (_m *Client) ListSSHKeys() ([]models.Key, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListSSHKeys")
	}

	var r0 []models.Key
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Key, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Key); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Key)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_ListSSHKeys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSSHKeys'
type Client_ListSSHKeys_Call struct {
	*mock.Call
}

// ListSSHKeys is a helper method to define mock.On call
func (_e *Client_Expecter) ListSSHKeys() *Client_ListSSHKeys_Call {
	return &Client_ListSSHKeys_Call{Call: _e.mock.On("ListSSHKeys")}
}

func (_c *Client_ListSSHKeys_Call) Run(run func()) *Client_ListSSHKeys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_ListSSHKeys_Call) Return(_a0 []models.Key, _a1 error) *Client_ListSSHKeys_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_ListSSHKeys_Call) RunAndReturn(run func() ([]models.Key, error)) *Client_ListSSHKeys_Call {
	_c.Call.Return(run)
	return _c
}

// RebootBMServer provides a mock function with given fields: _a0, _a1
func (_m *Client) RebootBMServer(_a0 int, _a1 v1beta1.RebootType) (*models.ResetPost, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RebootBMServer")
	}

	var r0 *models.ResetPost
	var r1 error
	if rf, ok := ret.Get(0).(func(int, v1beta1.RebootType) (*models.ResetPost, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(int, v1beta1.RebootType) *models.ResetPost); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResetPost)
		}
	}

	if rf, ok := ret.Get(1).(func(int, v1beta1.RebootType) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_RebootBMServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RebootBMServer'
type Client_RebootBMServer_Call struct {
	*mock.Call
}

// RebootBMServer is a helper method to define mock.On call
//   - _a0 int
//   - _a1 v1beta1.RebootType
func (_e *Client_Expecter) RebootBMServer(_a0 interface{}, _a1 interface{}) *Client_RebootBMServer_Call {
	return &Client_RebootBMServer_Call{Call: _e.mock.On("RebootBMServer", _a0, _a1)}
}

func (_c *Client_RebootBMServer_Call) Run(run func(_a0 int, _a1 v1beta1.RebootType)) *Client_RebootBMServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(v1beta1.RebootType))
	})
	return _c
}

func (_c *Client_RebootBMServer_Call) Return(_a0 *models.ResetPost, _a1 error) *Client_RebootBMServer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_RebootBMServer_Call) RunAndReturn(run func(int, v1beta1.RebootType) (*models.ResetPost, error)) *Client_RebootBMServer_Call {
	_c.Call.Return(run)
	return _c
}

// SetBMServerName provides a mock function with given fields: _a0, _a1
func (_m *Client) SetBMServerName(_a0 int, _a1 string) (*models.Server, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SetBMServerName")
	}

	var r0 *models.Server
	var r1 error
	if rf, ok := ret.Get(0).(func(int, string) (*models.Server, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(int, string) *models.Server); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Server)
		}
	}

	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_SetBMServerName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetBMServerName'
type Client_SetBMServerName_Call struct {
	*mock.Call
}

// SetBMServerName is a helper method to define mock.On call
//   - _a0 int
//   - _a1 string
func (_e *Client_Expecter) SetBMServerName(_a0 interface{}, _a1 interface{}) *Client_SetBMServerName_Call {
	return &Client_SetBMServerName_Call{Call: _e.mock.On("SetBMServerName", _a0, _a1)}
}

func (_c *Client_SetBMServerName_Call) Run(run func(_a0 int, _a1 string)) *Client_SetBMServerName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(string))
	})
	return _c
}

func (_c *Client_SetBMServerName_Call) Return(_a0 *models.Server, _a1 error) *Client_SetBMServerName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_SetBMServerName_Call) RunAndReturn(run func(int, string) (*models.Server, error)) *Client_SetBMServerName_Call {
	_c.Call.Return(run)
	return _c
}

// SetBootRescue provides a mock function with given fields: id, fingerprint
func (_m *Client) SetBootRescue(id int, fingerprint string) (*models.Rescue, error) {
	ret := _m.Called(id, fingerprint)

	if len(ret) == 0 {
		panic("no return value specified for SetBootRescue")
	}

	var r0 *models.Rescue
	var r1 error
	if rf, ok := ret.Get(0).(func(int, string) (*models.Rescue, error)); ok {
		return rf(id, fingerprint)
	}
	if rf, ok := ret.Get(0).(func(int, string) *models.Rescue); ok {
		r0 = rf(id, fingerprint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Rescue)
		}
	}

	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(id, fingerprint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_SetBootRescue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetBootRescue'
type Client_SetBootRescue_Call struct {
	*mock.Call
}

// SetBootRescue is a helper method to define mock.On call
//   - id int
//   - fingerprint string
func (_e *Client_Expecter) SetBootRescue(id interface{}, fingerprint interface{}) *Client_SetBootRescue_Call {
	return &Client_SetBootRescue_Call{Call: _e.mock.On("SetBootRescue", id, fingerprint)}
}

func (_c *Client_SetBootRescue_Call) Run(run func(id int, fingerprint string)) *Client_SetBootRescue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(string))
	})
	return _c
}

func (_c *Client_SetBootRescue_Call) Return(_a0 *models.Rescue, _a1 error) *Client_SetBootRescue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_SetBootRescue_Call) RunAndReturn(run func(int, string) (*models.Rescue, error)) *Client_SetBootRescue_Call {
	_c.Call.Return(run)
	return _c
}

// SetSSHKey provides a mock function with given fields: name, publickey
func (_m *Client) SetSSHKey(name string, publickey string) (*models.Key, error) {
	ret := _m.Called(name, publickey)

	if len(ret) == 0 {
		panic("no return value specified for SetSSHKey")
	}

	var r0 *models.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*models.Key, error)); ok {
		return rf(name, publickey)
	}
	if rf, ok := ret.Get(0).(func(string, string) *models.Key); ok {
		r0 = rf(name, publickey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Key)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, publickey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_SetSSHKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetSSHKey'
type Client_SetSSHKey_Call struct {
	*mock.Call
}

// SetSSHKey is a helper method to define mock.On call
//   - name string
//   - publickey string
func (_e *Client_Expecter) SetSSHKey(name interface{}, publickey interface{}) *Client_SetSSHKey_Call {
	return &Client_SetSSHKey_Call{Call: _e.mock.On("SetSSHKey", name, publickey)}
}

func (_c *Client_SetSSHKey_Call) Run(run func(name string, publickey string)) *Client_SetSSHKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Client_SetSSHKey_Call) Return(_a0 *models.Key, _a1 error) *Client_SetSSHKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_SetSSHKey_Call) RunAndReturn(run func(string, string) (*models.Key, error)) *Client_SetSSHKey_Call {
	_c.Call.Return(run)
	return _c
}

// ValidateCredentials provides a mock function with given fields:
func (_m *Client) ValidateCredentials() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ValidateCredentials")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_ValidateCredentials_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateCredentials'
type Client_ValidateCredentials_Call struct {
	*mock.Call
}

// ValidateCredentials is a helper method to define mock.On call
func (_e *Client_Expecter) ValidateCredentials() *Client_ValidateCredentials_Call {
	return &Client_ValidateCredentials_Call{Call: _e.mock.On("ValidateCredentials")}
}

func (_c *Client_ValidateCredentials_Call) Run(run func()) *Client_ValidateCredentials_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_ValidateCredentials_Call) Return(_a0 error) *Client_ValidateCredentials_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_ValidateCredentials_Call) RunAndReturn(run func() error) *Client_ValidateCredentials_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
