package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/minaandrawos/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/Chapter08/backend/src/dblayer"
	"github.com/minaandrawos/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/Chapter08/backend/src/models"
)

func TestHandler_GetProducts(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)
	mockdbLayer := dblayer.NewMockDBLayerWithData()
	h := NewHandlerWithDB(mockdbLayer)
	const productsURL string = "/products"
	type errMSG struct {
		Error string `json:"error"`
	}
	tests := []struct {
		name             string
		inErr            error
		outStatusCode    int
		expectedRespBody interface{}
	}{
		{
			"getproductsnoerrors",
			nil,
			http.StatusOK,
			mockdbLayer.GetMockProductData(),
		},
		{
			"getproductswitherror",
			errors.New("get products error"),
			http.StatusInternalServerError,
			errMSG{Error: "get products error"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//set the input error
			mockdbLayer.SetError(tt.inErr)
			//Create a test request
			req := httptest.NewRequest(http.MethodGet, productsURL, nil)
			//create an http response recorder
			w := httptest.NewRecorder()
			//create a fresh gin context and gin engine object from the response recorder
			_, engine := gin.CreateTestContext(w)

			//configure the get request
			engine.GET(productsURL, h.GetProducts)
			//serve the request
			engine.ServeHTTP(w, req)

			//test the output
			response := w.Result()
			if response.StatusCode != tt.outStatusCode {
				t.Errorf("Received Status code %d does not match expected status code %d", response.StatusCode, tt.outStatusCode)
			}

			var respBody interface{}
			if tt.inErr != nil {
				var errmsg errMSG
				json.NewDecoder(response.Body).Decode(&errmsg)
				respBody = errmsg
			} else {
				products := []models.Product{}
				json.NewDecoder(response.Body).Decode(&products)
				respBody = products
			}

			if !reflect.DeepEqual(respBody, tt.expectedRespBody) {
				t.Logf("%+v , %+v", respBody, tt.expectedRespBody)
				t.Errorf("Received HTTP response body %+v does not match expected HTTP response Body %+v", respBody, tt.expectedRespBody)
			}
		})
	}
}
