package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"go.elastic.co/apm/module/apmsql/v2"
	"log"
	"net/http"
)

type User struct {
	bun.BaseModel `bun:"table:users""`
	Name          string `bun:"name"`
}

type UserDto struct {
	Name string `json:"name"`
}

var users []UserDto

func ConnectDB() *bun.DB {
	apmsql.Register("postgres", &stdlib.Driver{})
	sqlDb, err := apmsql.Open("postgres", "postgresql://root:root@cockroachdb:26257?defaultdb")

	if err != nil {
		log.Fatalf("DB connection error -> %v", err)
	}

	sqlDb.SetMaxOpenConns(10)

	err = sqlDb.Ping()
	if err != nil {
		log.Fatalf("DB connection error -> %v", err)
	}

	db := bun.NewDB(sqlDb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook())

	return db
}

func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var users = &[]User{}

		log.Print("aaa")

		//conn := ConnectDB()
		//defer conn.Close()

		//err := conn.NewSelect().Model(users).Scan(context.Background())

		//if err != nil {
		//	c.IndentedJSON(http.StatusInternalServerError, "failed to get users")
		//	return
		//}

		c.IndentedJSON(http.StatusOK, users)
	}
}

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser UserDto
		if err := c.BindJSON(&newUser); err != nil {
			c.IndentedJSON(http.StatusBadRequest, "failed to bind req body")
			return
		}

		//conn := ConnectDB()
		//defer conn.Close()

		//var user = User{
		//	Name: newUser.Name,
		//}
		//_, err := conn.NewInsert().Model(&user).Exec(context.Background())
		//if err != nil {
		//	c.IndentedJSON(http.StatusInternalServerError, err)
		//	return
		//}

		users = append(users, newUser)

		c.IndentedJSON(http.StatusOK, newUser)
	}
}

func Empty() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, UserDto{})
	}
}

func UserRoutes(router *gin.Engine) {
	//conn := ConnectDB()
	router.GET("/users", GetAll())
	router.GET("/users/empty", Empty())
	router.POST("/users", Create())
}
