package checkHealthDelivery

import (
	"clean-code-unit-test/model/dto/json"
	"clean-code-unit-test/src/checkHealth"

	"github.com/gin-gonic/gin"
)

type checkHealthDelivery struct {
	checkHealthUC checkHealth.CheckHealthUsecase
}

func NewCheckHealthDelivery(v1Group *gin.RouterGroup, checkHealthUC checkHealth.CheckHealthUsecase) {
	handler := checkHealthDelivery{
		checkHealthUC: checkHealthUC,
	}
	checkHealthGroup := v1Group.Group("/checkHealth")
	{
		checkHealthGroup.GET("/version", handler.getVersion)
	}
}

// GetVersion is a function to get version
func (c *checkHealthDelivery) getVersion(ctx *gin.Context) {
	version, err := c.checkHealthUC.GetVersion()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}
	json.NewResponseSuccess(ctx, version, "success", "01", "01")
}
