package controller_test

import (
	"errors"
	"go-todo/server/controller"
	"go-todo/server/model/dbmodel"
	"go-todo/server/model/reqmodel"
	"go-todo/server/model/resmodel"
	"go-todo/test/mocks"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RoleControllerTestSuite struct {
	suite.Suite
	roleController *controller.RoleController
	mockRoleDao    *mocks.MockRoleDao
}

func Test_RoleControllerTestSuite(t *testing.T) {
	suite.Run(t, new(RoleControllerTestSuite))
}

func (r *RoleControllerTestSuite) SetupTest() {
	r.mockRoleDao = new(mocks.MockRoleDao)
	r.roleController = controller.NewRoleController(r.mockRoleDao)
}

func (r *RoleControllerTestSuite) Test_CreateRole() {
	tests := []struct {
		TestName    string
		Name        string
		AccessLevel int
		returnID    int
		returnErr   error
	}{
		{
			TestName:    "creates a new role",
			Name:        "user",
			AccessLevel: 100,
			returnID:    10,
			returnErr:   nil,
		},
		{
			TestName:    "returns the error from dao",
			Name:        "admin",
			AccessLevel: 200,
			returnID:    0,
			returnErr:   errors.New("test err: error creating role"),
		},
	}

	for _, t := range tests {
		r.Run(t.TestName, func() {
			mockReqRole := reqmodel.CreateRole{
				Name:        t.Name,
				AccessLevel: t.AccessLevel,
			}
			mockDBRole := dbmodel.Role{
				Name:        t.Name,
				AccessLevel: t.AccessLevel,
			}
			r.mockRoleDao.On("CreateRole", mockDBRole).Return(t.returnID, t.returnErr)

			val, err := r.roleController.CreateRole(mockReqRole)

			if t.returnErr == nil {
				require.NoError(r.T(), err)
			} else {
				require.Error(r.T(), err)
			}

			require.Equal(r.T(), t.returnID, val)
		})
	}
	r.mockRoleDao.AssertNumberOfCalls(r.T(), "CreateRole", 2)
	r.mockRoleDao.AssertExpectations(r.T())
}

func (r *RoleControllerTestSuite) Test_FindRoleByID() {
	tests := []struct {
		TestName        string
		RoleID          int
		RoleName        string
		RoleAccessLevel int
		returnErr       error
	}{
		{
			TestName:        "returns role res model",
			RoleID:          10,
			RoleName:        "user",
			RoleAccessLevel: 100,
			returnErr:       nil,
		},
		{
			TestName:        "returns error from role dao",
			RoleID:          20,
			RoleName:        "admin",
			RoleAccessLevel: 200,
			returnErr:       errors.New("test err: error finding role"),
		},
	}

	for _, t := range tests {
		r.Run(t.TestName, func() {
			mockDbRole := dbmodel.Role{
				ID:          t.RoleID,
				Name:        t.RoleName,
				AccessLevel: t.RoleAccessLevel,
			}
			r.mockRoleDao.On("FindRoleByID", t.RoleID).Return(mockDbRole, t.returnErr)

			role, err := r.roleController.FindRoleByID(t.RoleID)

			if t.returnErr == nil {
				require.NoError(r.T(), err)
			} else {
				require.Error(r.T(), err)
			}

			expectedRole := resmodel.Role{
				ID:          t.RoleID,
				Name:        t.RoleName,
				AccessLevel: t.RoleAccessLevel,
			}
			require.Equal(r.T(), expectedRole, role)
		})
	}
	r.mockRoleDao.AssertNumberOfCalls(r.T(), "FindRoleByID", 2)
	r.mockRoleDao.AssertExpectations(r.T())
}
