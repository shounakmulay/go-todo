package daos_test

import (
	"database/sql"
	"go-todo/internal/log"
	"go-todo/server/daos"
	"go-todo/server/model/dbmodel"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type RoleTestSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	dao daos.IRoleDao
}

func Test_RoleTestSuite(t *testing.T) {
	suite.Run(t, new(RoleTestSuite))
}

func (rs *RoleTestSuite) SetupTest() {
	var (
		db  *sql.DB
		err error
	)
	db, rs.mock, err = sqlmock.New()
	require.NoError(rs.T(), err)
	print(db)

	rs.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}))
	require.NoError(rs.T(), err)

	rs.DB.Logger = logger.New(log.NewGormLogger(), logger.Config{})

	rs.dao = daos.NewRoleDao(rs.DB)
}

func (rs *RoleTestSuite) AfterTest(_, _ string) {
	require.NoError(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *RoleTestSuite) Test_FindRoleByID_ReturnsDBRoleWhen() {
	dbTime := time.Now()

	tests := []struct {
		name          string
		expectedModel dbmodel.Role
	}{
		{
			name: "Passing id 10",
			expectedModel: dbmodel.Role{
				ID:          10,
				Name:        "user",
				AccessLevel: 100,
				CreatedAt:   dbTime,
				UpdatedAt:   dbTime,
			},
		},
		{
			name: "Passing id 123456789",
			expectedModel: dbmodel.Role{
				ID:          123456789,
				Name:        "admin",
				AccessLevel: 100,
				CreatedAt:   dbTime,
				UpdatedAt:   dbTime,
			},
		},
	}

	for _, test := range tests {
		rs.Run(test.name, func() {
			model := test.expectedModel

			// expected query
			rs.mock.
				ExpectQuery("SELECT \\* FROM `roles` WHERE `roles`\\.`id` = \\? ORDER BY `roles`.`id` LIMIT 1").
				WithArgs(model.ID).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "access_level", "created_at", "updated_at"}).
						AddRow(model.ID, model.Name, model.AccessLevel, model.CreatedAt, model.UpdatedAt),
				)

			// Call function under test
			res, err := rs.dao.FindRoleByID(model.ID)
			if err != nil {
				rs.FailNow(err.Error())
			}

			// Check output
			rs.Equal(test.expectedModel, res)

			// Check expectations
			newerr := rs.mock.ExpectationsWereMet()
			if newerr != nil {
				rs.FailNow(newerr.Error())
			}
		})
	}
}

func (rs *RoleTestSuite) Test_FindRoleByID_ReturnsErrorWhen() {
	tests := []struct {
		name          string
		modelID       int
		expectedError error
	}{
		{
			name:          "Record not found",
			modelID:       10,
			expectedError: gorm.ErrRecordNotFound,
		},
	}

	for _, test := range tests {
		rs.Run(test.name, func() {

			// expected query
			rs.mock.
				ExpectQuery("SELECT \\* FROM `roles` WHERE `roles`\\.`id` = \\? ORDER BY `roles`.`id` LIMIT 1").
				WithArgs(test.modelID).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "access_level", "created_at", "updated_at"}),
				)

			// Call function under test
			_, err := rs.dao.FindRoleByID(test.modelID)
			if rs.Error(err) {
				// Check expected error
				rs.Equal(test.expectedError, err)
			}

			// Check expectations
			newerr := rs.mock.ExpectationsWereMet()
			if newerr != nil {
				rs.FailNow(newerr.Error())
			}
		})
	}
}
