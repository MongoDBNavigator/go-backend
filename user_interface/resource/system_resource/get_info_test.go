package system_resource

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"fmt"

	"errors"

	"github.com/MongoDBNavigator/go-backend/domain/system/model"
	"github.com/MongoDBNavigator/go-backend/tests/helper"
	"github.com/MongoDBNavigator/go-backend/tests/mock"
	"github.com/MongoDBNavigator/go-backend/user_interface/resource/middleware"
	"github.com/emicklei/go-restful"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetInfoSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wsContainer := restful.NewContainer()
	systemInfoReader := mock.NewMockSystemInfoReader(ctrl)

	systemInfo := model.NewSystemInfo("4.0.0", 64, "localhost")

	systemInfoReader.
		EXPECT().
		Reade().
		Return(systemInfo, nil)

	NewSystemResource(systemInfoReader, middleware.NewJwtMiddleware(helper.PASSWORD)).Register(wsContainer)

	httpRequest, _ := http.NewRequest("GET", "http://localhost/api/v1/system/info", nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	wsContainer.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusOK, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"url":"localhost","version":"4.0.0","processorArchitecture":64}`)
}

func TestGetInfoUnauthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wsContainer := restful.NewContainer()
	systemInfoReader := mock.NewMockSystemInfoReader(ctrl)

	NewSystemResource(systemInfoReader, middleware.NewJwtMiddleware(helper.PASSWORD)).Register(wsContainer)

	httpRequest, _ := http.NewRequest("GET", "http://localhost/api/v1/system/info", nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpWriter := httptest.NewRecorder()

	wsContainer.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusUnauthorized, httpWriter.Code)
}

func TestGetInfoForbidden(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wsContainer := restful.NewContainer()
	systemInfoReader := mock.NewMockSystemInfoReader(ctrl)

	NewSystemResource(systemInfoReader, middleware.NewJwtMiddleware(helper.PASSWORD)).Register(wsContainer)

	httpRequest, _ := http.NewRequest("GET", "http://localhost/api/v1/system/info", nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", "Bearer test")
	httpWriter := httptest.NewRecorder()

	wsContainer.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusForbidden, httpWriter.Code)
}

func TestGetInfoInternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wsContainer := restful.NewContainer()
	systemInfoReader := mock.NewMockSystemInfoReader(ctrl)

	systemInfoReader.
		EXPECT().
		Reade().
		Return(nil, errors.New("internal_server_error"))

	NewSystemResource(systemInfoReader, middleware.NewJwtMiddleware(helper.PASSWORD)).Register(wsContainer)

	httpRequest, _ := http.NewRequest("GET", "http://localhost/api/v1/system/info", nil)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", helper.GenerateJwtToken()))
	httpWriter := httptest.NewRecorder()

	wsContainer.Dispatch(httpWriter, httpRequest)

	assert.Equal(t, http.StatusInternalServerError, httpWriter.Code)
	space := regexp.MustCompile(`\s+`)
	assert.Equal(t, space.ReplaceAllString(httpWriter.Body.String(), ""), `{"message":"internal_server_error"}`)
}
