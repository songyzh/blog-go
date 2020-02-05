package configs

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	"os"
)

var DB *sqlx.DB

func init(){
	DB = sqlx.MustConnect("mysql", os.Getenv("BLOG_DATABASE_GO"))
	DB.MapperFunc(strcase.ToSnake)
	DB = DB.Unsafe()
}
