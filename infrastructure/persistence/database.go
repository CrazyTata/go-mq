package persistence

import (
	"go-mq/domain"

	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// NewDatabase 创建数据库连接
func NewDatabase(dataSource string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	// 设置连接池参数
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	return db, nil
}

func RowBuilder(table string, columns ...string) squirrel.SelectBuilder {
	return squirrel.Select(columns...).From(table).Where(squirrel.Eq{"is_deleted": domain.IsNotDeleted})
}

func CountBuilder(table string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(*)").From(table).Where(squirrel.Eq{"is_deleted": domain.IsNotDeleted})
}
