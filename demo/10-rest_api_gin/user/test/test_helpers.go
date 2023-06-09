package test

import (
	"demo/internal/mongodb"
	"demo/internal/web"
	"demo/user"
	"demo/user/repository"
	"flag"

	"github.com/gin-gonic/gin"
)

var (
	httpAddr   = flag.String("http", ":8000", "Http address")
	mongoDbUrl = flag.String("mongodb", "mongodb://mongoadmin:secret@localhost:27017", "MongoDb address")
)

func NewTestRouter() *gin.Engine {
	flag.Parse()
	router := web.NewRouter()

	client := mongodb.MongoDB(*mongoDbUrl)
	repository := repository.NewMongoDb(client)
	accountService := user.NewAccountService(repository)

	router.POST("users", user.CreateAccountHandler(accountService))

	return router
}
