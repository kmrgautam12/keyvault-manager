package database

import (
	utils "Authentication-Go/Utils"
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DBManager struct {
	dbpool *pgxpool.Pool
}

func InitPgDBConnection() *DBManager {

	connectionString, err := pgxpool.ParseConfig(UserDbConnStr)
	if err != nil {
		panic(err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), connectionString)
	if err != nil {
		panic(err)
	}
	return &DBManager{
		dbpool: pool,
	}
}

func (db *DBManager) InsertToDb(ctx *gin.Context, stmt string) error {
	p, err := db.dbpool.Exec(context.Background(), stmt)
	if err != nil {
		utils.SentErrorResponse500(ctx, err)
		return nil
	}
	if p.RowsAffected() == 0 {
		utils.SentErrorResponse500(ctx, errors.New("no matching row"))
		return nil
	}
	return nil
}

func (db *DBManager) GetUserFromDB(ctx *gin.Context, stmt string) (bool, error) {
	p, err := db.dbpool.Exec(context.Background(), stmt)
	if err != nil {
		return false, err
	}
	if p.RowsAffected() == 0 {
		return false, err
	}
	return true, nil

}
