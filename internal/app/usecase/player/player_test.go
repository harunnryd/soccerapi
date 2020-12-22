package player

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/harunnryd/skeltun/internal/app/driver/db"
	"github.com/harunnryd/skeltun/internal/app/handler/player/param"
	"github.com/harunnryd/skeltun/internal/app/handler/player/transporter"
	"github.com/harunnryd/skeltun/internal/app/repo"
	iPlayerRepo "github.com/harunnryd/skeltun/internal/app/repo/player"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Suite ...
type Suite struct {
	suite.Suite
	mock      sqlmock.Sqlmock
	pgsqlConn *gorm.DB

	iPlayerRepo iPlayerRepo.IPlayer
	iRepo       repo.IRepo
	player      IPlayer
	helper
	response
}

type helper struct {
	err error
	db  *sql.DB
}

type response struct {
	doCreateResp   transporter.DoCreate
	getPlayersResp []transporter.GetPlayers
	getPlayerResp  transporter.GetPlayer
	doDeleteResp   transporter.DoDelete
	doUpdateResp   transporter.DoUpdate
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

	suite.iPlayerRepo = iPlayerRepo.New(
		iPlayerRepo.WithDatabase(db.PgsqlDialectParam, suite.pgsqlConn),
	)

	suite.iRepo = repo.New()
	suite.iRepo.SetPlayer(suite.iPlayerRepo)

	suite.player = New(WithRepo(suite.iRepo))
}

// TestDoCreate ...
func (suite *Suite) TestDoCreate() {
	params := param.DoCreate{
		Player: param.Player{
			TeamID: uuid.NewV4(),
			Name:   "John Doe",
		},
	}

	suite.mock.MatchExpectationsInOrder(false)

	suite.mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "players" ("created_at","updated_at","deleted_at","team_id","name") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), params.TeamID, params.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(uuid.NewV4()))

	suite.response.doCreateResp, suite.helper.err = suite.player.DoCreate(context.Background(), params)

	require.NoError(suite.T(), suite.helper.err)
}

// TestGetPlayers ...
func (suite *Suite) TestGetPlayers() {
	suite.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "players" LIMIT 10 OFFSET 1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(uuid.NewV4(), "John Doe"))

	suite.response.getPlayersResp, suite.helper.err = suite.player.GetPlayers(context.Background(), param.GetPlayers{Pagination: param.Pagination{
		Limit:  "10",
		Offset: "1",
	}})

	require.NoError(suite.T(), suite.helper.err)
}

// TestGetPlayer ...
func (suite *Suite) TestGetPlayer() {
	params := param.GetPlayer{ID: uuid.NewV4()}

	suite.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "players" WHERE id = $1 LIMIT 1`)).
		WithArgs(params.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(uuid.NewV4(), "John Doe"))

	suite.response.getPlayerResp, suite.helper.err = suite.player.GetPlayer(context.Background(), params)

	require.NoError(suite.T(), suite.helper.err)
}

// TestDoUpdate ...
func (suite *Suite) TestDoUpdate() {
	params := param.DoUpdate{
		Player: param.Player{
			ID:     uuid.NewV4(),
			Name:   "John Wick",
			TeamID: uuid.NewV4(),
		},
	}

	suite.mock.MatchExpectationsInOrder(false)

	suite.mock.
		ExpectExec(regexp.QuoteMeta(`UPDATE "players" SET "updated_at"=$1,"team_id"=$2,"name"=$3 WHERE id = $4`)).
		WithArgs(sqlmock.AnyArg(), params.TeamID, params.Name, params.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	suite.response.doUpdateResp, suite.helper.err = suite.player.DoUpdate(context.Background(), params)

	require.NoError(suite.T(), suite.helper.err)
}

func (suite *Suite) TestDoDelete() {
	params := param.DoDelete{
		ID: uuid.NewV4(),
	}

	suite.mock.MatchExpectationsInOrder(false)

	suite.mock.
		ExpectExec(regexp.QuoteMeta(`DELETE FROM "players" WHERE id = $1`)).
		WithArgs(params.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	suite.response.doDeleteResp, suite.helper.err = suite.player.DoDelete(context.Background(), params)

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
