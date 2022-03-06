// Code generated by mockery v2.10.0. DO NOT EDIT.

package domain

import (
	context "context"

	model "github.com/bangumi/server/model"
	mock "github.com/stretchr/testify/mock"
)

// MockPersonRepo is an autogenerated mock type for the PersonRepo type
type MockPersonRepo struct {
	mock.Mock
}

type MockPersonRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPersonRepo) EXPECT() *MockPersonRepo_Expecter {
	return &MockPersonRepo_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, id
func (_m *MockPersonRepo) Get(ctx context.Context, id uint32) (model.Person, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Person
	if rf, ok := ret.Get(0).(func(context.Context, uint32) model.Person); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Person)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPersonRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockPersonRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint32
func (_e *MockPersonRepo_Expecter) Get(ctx interface{}, id interface{}) *MockPersonRepo_Get_Call {
	return &MockPersonRepo_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *MockPersonRepo_Get_Call) Run(run func(ctx context.Context, id uint32)) *MockPersonRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *MockPersonRepo_Get_Call) Return(_a0 model.Person, _a1 error) *MockPersonRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetActors provides a mock function with given fields: ctx, subjectID, characterIDs
func (_m *MockPersonRepo) GetActors(ctx context.Context, subjectID uint32, characterIDs ...uint32) (map[uint32][]model.Person, error) {
	_va := make([]interface{}, len(characterIDs))
	for _i := range characterIDs {
		_va[_i] = characterIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subjectID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 map[uint32][]model.Person
	if rf, ok := ret.Get(0).(func(context.Context, uint32, ...uint32) map[uint32][]model.Person); ok {
		r0 = rf(ctx, subjectID, characterIDs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint32][]model.Person)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, ...uint32) error); ok {
		r1 = rf(ctx, subjectID, characterIDs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPersonRepo_GetActors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetActors'
type MockPersonRepo_GetActors_Call struct {
	*mock.Call
}

// GetActors is a helper method to define mock.On call
//  - ctx context.Context
//  - subjectID uint32
//  - characterIDs ...uint32
func (_e *MockPersonRepo_Expecter) GetActors(ctx interface{}, subjectID interface{}, characterIDs ...interface{}) *MockPersonRepo_GetActors_Call {
	return &MockPersonRepo_GetActors_Call{Call: _e.mock.On("GetActors",
		append([]interface{}{ctx, subjectID}, characterIDs...)...)}
}

func (_c *MockPersonRepo_GetActors_Call) Run(run func(ctx context.Context, subjectID uint32, characterIDs ...uint32)) *MockPersonRepo_GetActors_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]uint32, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(uint32)
			}
		}
		run(args[0].(context.Context), args[1].(uint32), variadicArgs...)
	})
	return _c
}

func (_c *MockPersonRepo_GetActors_Call) Return(_a0 map[uint32][]model.Person, _a1 error) *MockPersonRepo_GetActors_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetCharacterRelated provides a mock function with given fields: ctx, subjectID
func (_m *MockPersonRepo) GetCharacterRelated(ctx context.Context, subjectID uint32) ([]CharacterCast, error) {
	ret := _m.Called(ctx, subjectID)

	var r0 []CharacterCast
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []CharacterCast); ok {
		r0 = rf(ctx, subjectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]CharacterCast)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, subjectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPersonRepo_GetCharacterRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCharacterRelated'
type MockPersonRepo_GetCharacterRelated_Call struct {
	*mock.Call
}

// GetCharacterRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - subjectID uint32
func (_e *MockPersonRepo_Expecter) GetCharacterRelated(ctx interface{}, subjectID interface{}) *MockPersonRepo_GetCharacterRelated_Call {
	return &MockPersonRepo_GetCharacterRelated_Call{Call: _e.mock.On("GetCharacterRelated", ctx, subjectID)}
}

func (_c *MockPersonRepo_GetCharacterRelated_Call) Run(run func(ctx context.Context, subjectID uint32)) *MockPersonRepo_GetCharacterRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *MockPersonRepo_GetCharacterRelated_Call) Return(_a0 []CharacterCast, _a1 error) *MockPersonRepo_GetCharacterRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetSubjectRelated provides a mock function with given fields: ctx, subjectID
func (_m *MockPersonRepo) GetSubjectRelated(ctx context.Context, subjectID uint32) ([]model.Person, []model.PersonSubjectRelation, error) {
	ret := _m.Called(ctx, subjectID)

	var r0 []model.Person
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []model.Person); ok {
		r0 = rf(ctx, subjectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Person)
		}
	}

	var r1 []model.PersonSubjectRelation
	if rf, ok := ret.Get(1).(func(context.Context, uint32) []model.PersonSubjectRelation); ok {
		r1 = rf(ctx, subjectID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]model.PersonSubjectRelation)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, uint32) error); ok {
		r2 = rf(ctx, subjectID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockPersonRepo_GetSubjectRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubjectRelated'
type MockPersonRepo_GetSubjectRelated_Call struct {
	*mock.Call
}

// GetSubjectRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - subjectID uint32
func (_e *MockPersonRepo_Expecter) GetSubjectRelated(ctx interface{}, subjectID interface{}) *MockPersonRepo_GetSubjectRelated_Call {
	return &MockPersonRepo_GetSubjectRelated_Call{Call: _e.mock.On("GetSubjectRelated", ctx, subjectID)}
}

func (_c *MockPersonRepo_GetSubjectRelated_Call) Run(run func(ctx context.Context, subjectID uint32)) *MockPersonRepo_GetSubjectRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *MockPersonRepo_GetSubjectRelated_Call) Return(_a0 []model.Person, _a1 []model.PersonSubjectRelation, _a2 error) *MockPersonRepo_GetSubjectRelated_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}