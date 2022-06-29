package controller_test

import (
	// "errors"
	"errors"
	"go-todo/server/controller"
	"go-todo/server/model/dbmodel"
	"go-todo/server/model/reqmodel"
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
