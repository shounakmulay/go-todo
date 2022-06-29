package mocks

import (
	"go-todo/server/model/dbmodel"

	"github.com/stretchr/testify/mock"
)

type MockRoleDao struct {
	mock.Mock
}

func (m *MockRoleDao) CreateRole(role dbmodel.Role) (int, error) {
	args := m.Called(role)
	return args.Int(0), args.Error(1)
}

func (m *MockRoleDao) FindRoleByID(id int) (dbmodel.Role, error) {
	args := m.Called(id)
	return args.Get(0).(dbmodel.Role), args.Error(1)
}
