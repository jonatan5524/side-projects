// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	model "github.com/jonatan5524/side-projects-manager/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// ProjectRepository is an autogenerated mock type for the ProjectRepository type
type ProjectRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0
func (_m *ProjectRepository) Delete(_a0 model.Project) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Project) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByPath provides a mock function with given fields: _a0
func (_m *ProjectRepository) DeleteByPath(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMany provides a mock function with given fields: _a0
func (_m *ProjectRepository) DeleteMany(_a0 ...*model.Project) error {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...*model.Project) error); ok {
		r0 = rf(_a0...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0
func (_m *ProjectRepository) Get(_a0 string) (model.Project, error) {
	ret := _m.Called(_a0)

	var r0 model.Project
	if rf, ok := ret.Get(0).(func(string) model.Project); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(model.Project)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *ProjectRepository) GetAll() ([]*model.Project, error) {
	ret := _m.Called()

	var r0 []*model.Project
	if rf, ok := ret.Get(0).(func() []*model.Project); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllFilteredGit provides a mock function with given fields:
func (_m *ProjectRepository) GetAllFilteredGit() ([]*model.Project, error) {
	ret := _m.Called()

	var r0 []*model.Project
	if rf, ok := ret.Get(0).(func() []*model.Project); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecent provides a mock function with given fields: _a0
func (_m *ProjectRepository) GetRecent(_a0 int) ([]*model.Project, error) {
	ret := _m.Called(_a0)

	var r0 []*model.Project
	if rf, ok := ret.Get(0).(func(int) []*model.Project); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: _a0
func (_m *ProjectRepository) Put(_a0 model.Project) (uint64, error) {
	ret := _m.Called(_a0)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(model.Project) uint64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Project) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
