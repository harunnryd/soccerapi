package team

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/harunnryd/skeltun/internal/app/driver/db"
	"github.com/harunnryd/skeltun/internal/app/handler/team/param"
	"github.com/harunnryd/skeltun/internal/app/handler/team/transporter"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

// Suite ...
type Suite struct {
	suite.Suite
	pgsqlConn *gorm.DB
	mock      sqlmock.Sqlmock

	team ITeam
	helper
	response
}

type helper struct {
	db  *sql.DB
	err error
}

type response struct {
	doCreateResp transporter.DoCreate
	getTeamsResp []transporter.GetTeams
	getTeamResp  transporter.GetTeam
	doUpdateResp transporter.DoUpdate
	doDeleteResp transporter.DoDelete
}

// SetupSuite ...
func (suite *Suite) SetupSuite() {
	suite.helper.db, suite.mock, suite.helper.err = sqlmock.New()
	require.NoError(suite.T(), suite.helper.err)

	suite.pgsqlConn, suite.helper.err = gorm.Open(postgres.New(postgres.Config{
		Conn:                 suite.helper.db,
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	require.NoError(suite.T(), suite.helper.err)

	require.NotNil(suite.T(), suite.pgsqlConn)

	suite.pgsqlConn.Debug()

	suite.team = New(
		WithDatabase(db.PgsqlDialectParam, suite.pgsqlConn),
	)
}

// TestDoCreate ...
func (suite *Suite) TestDoCreate() {
	params := param.DoCreate{
		Team: param.Team{
			Name: "Arsenal",
		},
	}

	suite.mock.MatchExpectationsInOrder(false)
	suite.mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "teams" ("created_at","updated_at","deleted_at","name") VALUES ($1,$2,$3,$4) RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), params.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(uuid.NewV4()))

	suite.response.doCreateResp, suite.helper.err = suite.team.DoCreate(context.Background(), params)

	require.NoError(suite.T(), suite.helper.err)

	require.Equal(suite.T(), "Arsenal", suite.response.doCreateResp.Name)
}

// TestGetTeams ...
func (suite *Suite) TestGetTeams() {
	suite.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "teams" LIMIT 10 OFFSET 1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(uuid.NewV4(), "Juventus"))

	suite.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "players" WHERE "players"."team_id" = $1`)).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(uuid.NewV4(), "John Doe"))

	suite.response.getTeamsResp, suite.helper.err = suite.team.GetTeams(context.Background(), param.GetTeams{Pagination: param.Pagination{
		Limit:  "10",
		Offset: "1",
	}})

	require.NoError(suite.T(), suite.helper.err)
}

// TestGetTeam ...
func (suite *Suite) TestGetTeam() {
	params := param.GetTeam{ID: uuid.NewV4()}

	suite.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "teams" WHERE id = $1 LIMIT 1`)).
		WithArgs(params.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(uuid.NewV4(), "Liverpool"))

	suite.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "players" WHERE "players"."team_id" = $1`)).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(uuid.NewV4(), "John Doe"))

	suite.response.getTeamResp, suite.helper.err = suite.team.GetTeam(context.Background(), params)

	require.NoError(suite.T(), suite.helper.err)
}

// TestDoUpdate ...
func (suite *Suite) TestDoUpdate() {
	params := param.DoUpdate{
		Team: param.Team{
			ID:   uuid.NewV4(),
			Name: "Real Madrid",
		},
	}

	suite.mock.MatchExpectationsInOrder(false)

	suite.mock.
		ExpectExec(regexp.QuoteMeta(`UPDATE "teams" SET "updated_at"=$1,"name"=$2 WHERE id = $3`)).
		WithArgs(sqlmock.AnyArg(), params.Name, params.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	suite.response.doUpdateResp, suite.helper.err = suite.team.DoUpdate(context.Background(), params)

	require.NoError(suite.T(), suite.helper.err)
}

func (suite *Suite) TestDoDelete() {
	params := param.DoDelete{
		ID: uuid.NewV4(),
	}

	suite.mock.MatchExpectationsInOrder(false)

	suite.mock.
		ExpectExec(regexp.QuoteMeta(`DELETE FROM "teams" WHERE id = $1`)).
		WithArgs(params.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	suite.response.doDeleteResp, suite.helper.err = suite.team.DoDelete(context.Background(), params)

	require.NoError(suite.T(), suite.helper.err)
}

// AfterTest ...
func (suite *Suite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
	suite.helper.err = nil
}

// Suite We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
