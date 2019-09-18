package main

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	_cache "github.com/faridtriwicaksono/employee_api/infrastructure/cache"
	_db "github.com/faridtriwicaksono/employee_api/infrastructure/db"
	_userPresenter "github.com/faridtriwicaksono/employee_api/services/employee/presenter"
	_userRepo "github.com/faridtriwicaksono/employee_api/services/employee/repository"
	_userUseCase "github.com/faridtriwicaksono/employee_api/services/employee/usecase"

)

func auth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"tongseng": "kambingenak",
	})
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Panic in Employee API, recovered: ", err)
			debug.PrintStack()
		}
	}()

	DbClient := _db.GetDBAccess()
	RedisClient := _cache.GetRedisAccess()

	router := gin.Default()

	router.Use(auth())

	router.GET("/test/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	userRepositoryDB := _userRepo.NewMysqlEmployeeRepository(DbClient)
	userRepositoryCache := _userRepo.NewRedisEmployeeRepository(RedisClient)
	userUseCase := _userUseCase.NewEmployeeUsecase(userRepositoryDB, userRepositoryCache)
	_userPresenter.NewEmployeeHTTPHandler(router, userUseCase)


	s := &http.Server{
		Addr:           ":9090",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
