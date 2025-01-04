package database

import (
	"context"
	"fmt"

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

func (db *DBManager) InsertToDb(ctx *gin.Context, stmt string, user string, hash_pass string) error {
	p, err := db.dbpool.Exec(context.Background(), stmt, user, hash_pass)
	if err != nil {
		return fmt.Errorf("error creating user: error : %s", err.Error())
	}
	if p.RowsAffected() == 0 {
		return fmt.Errorf("error creating user: error : %s", err)
	}
	return nil
}

func (db *DBManager) GetUserFromDB(ctx *gin.Context, stmt string) (bool, error) {
	p, err := db.dbpool.Query(context.Background(), stmt)
	if err != nil {
		return false, err
	}
	defer p.Close()
	if p.Err() != nil {
		return false, err
	}
	if p.Next() {
		return true, nil
	}
	return false, nil

}

func (db *DBManager) GetUserFromDBService(ctx *gin.Context, stmt string) (bool, error) {
	p, err := db.dbpool.Query(context.Background(), stmt)
	if err != nil {
		return false, err
	}
	defer p.Close()
	if p.Err() != nil {
		return false, err
	}
	if p.Next() {
		p.Values()
		return true, nil
	}
	return false, nil

}
