package utils

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"sauce-service/src/app/category"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func getTestContext(method, url string, body []byte) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req
	return c, w
}

func newTestContext() *gin.Context {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	return c
}

func TestHandleError(t *testing.T) {
	t.Run("NotFoundError", func(t *testing.T) {
		c, w := getTestContext("GET", "/test", nil)
		HandleError(c, gorm.ErrRecordNotFound, "Not found")
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Not found")
	})

	t.Run("InternalServerError", func(t *testing.T) {
		c, w := getTestContext("GET", "/test", nil)
		HandleError(c, errors.New("something went wrong"), "irrelevant")
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "something went wrong")
	})
}

func TestExtractIDParam(t *testing.T) {
	t.Run("WithID", func(t *testing.T) {
		c := newTestContext()
		c.Params = gin.Params{{Key: "id", Value: "123"}}
		id := ExtractIDParam(c)
		assert.Equal(t, "123", id)
	})

	t.Run("MissingID", func(t *testing.T) {
		c := newTestContext()
		id := ExtractIDParam(c)
		assert.Equal(t, "", id)
	})
}

func TestBindAndValidateJSON(t *testing.T) {
	t.Run("ValidInput", func(t *testing.T) {
		body := []byte(`{"name":"Océane"}`)
		c, _ := getTestContext("POST", "/test", body)
		var input category.CategoryInput
		result := BindAndValidateJSON(c, &input)
		assert.True(t, result)
		assert.Equal(t, "Océane", input.Name)
	})

	t.Run("MissingRequiredField", func(t *testing.T) {
		body := []byte(`{}`)
		c, w := getTestContext("POST", "/test", body)
		var input category.CategoryInput
		result := BindAndValidateJSON(c, &input)
		assert.False(t, result)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Name")
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		body := []byte(`{invalid json}`)
		c, w := getTestContext("POST", "/test", body)
		var input category.CategoryInput
		result := BindAndValidateJSON(c, &input)
		assert.False(t, result)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "invalid character")
	})
}

func TestRespondError(t *testing.T) {
	t.Run("WithMessage", func(t *testing.T) {
		c, w := getTestContext("GET", "/test", nil)
		RespondError(c, http.StatusBadRequest, "invalid request")
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "invalid request")
	})

	t.Run("EmptyMessage", func(t *testing.T) {
		c, w := getTestContext("GET", "/test", nil)
		RespondError(c, http.StatusInternalServerError, "")
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), `"error":""`)
	})
}

func TestRespondSuccess(t *testing.T) {
	t.Run("WithData", func(t *testing.T) {
		c, w := getTestContext("GET", "/test", nil)
		data := gin.H{"message": "success"}
		RespondSuccess(c, http.StatusOK, data)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "success")
	})

	t.Run("WithNilData", func(t *testing.T) {
		c, w := getTestContext("GET", "/test", nil)
		RespondSuccess(c, http.StatusOK, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "null")
	})
}
