package router

import (
	"clean-code-unit-test/src/checkHealth/checkHealthDelivery"
	"clean-code-unit-test/src/checkHealth/checkHealthRepository"
	"clean-code-unit-test/src/checkHealth/checkHealthUsecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	checkHealthRepo := checkHealthRepository.NewCheckHealthRepository(db)
	checkHealthUC := checkHealthUsecase.NewCheckHealthUsecase(checkHealthRepo)
	checkHealthDelivery.NewCheckHealthDelivery(v1Group, checkHealthUC)
}
