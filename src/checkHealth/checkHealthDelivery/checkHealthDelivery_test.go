package checkHealthDelivery

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockCheckHealthUsecase struct {
	mock.Mock
}

func (m *mockCheckHealthUsecase) GetVersion() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

type CheckHealthDeliveryTestSuite struct {
	suite.Suite
	router  *gin.Engine
	usecase *mockCheckHealthUsecase
}

func (suite *CheckHealthDeliveryTestSuite) SetupTest() {
	suite.router = gin.Default()
	suite.usecase = new(mockCheckHealthUsecase)

	v1Group := suite.router.Group("/api/v1")
	NewCheckHealthDelivery(v1Group, suite.usecase)
}

func (suite *CheckHealthDeliveryTestSuite) TestGetVersion_Success() {
	expected := "1.0.0"
	suite.usecase.On("GetVersion").Return(expected, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/checkHealth/version", nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(`{"responseCode":"2000101","responseMessage":"success","data":"1.0.0"}`, w.Body.String())
}

func (suite *CheckHealthDeliveryTestSuite) TestGetVersion_Error() {
	expected := ""
	suite.usecase.On("GetVersion").Return(expected, sql.ErrConnDone)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/checkHealth/version", nil)

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.JSONEq(`{"responseCode":"5000101","responseMessage":"internal server error","error":"sql: connection is already closed"}`, w.Body.String())
}

func (suite *CheckHealthDeliveryTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
}

func TestCheckHealthDeliveryTestSuite(t *testing.T) {
	suite.Run(t, new(CheckHealthDeliveryTestSuite))
}
