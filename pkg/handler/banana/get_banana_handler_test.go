package banana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	dbb "github.com/makarski/mybanana/pkg/db/banana"
	bmock "github.com/makarski/mybanana/pkg/db/banana/mocks"
	hmock "github.com/makarski/mybanana/pkg/handler/mocks"
)

func TestGetBananaHandler(t *testing.T) {
	bananaFinder := &bmock.BananaFinder{}
	urlReader := &hmock.URLParamReader{}
	bananaHandler := NewGetBananaHandler(bananaFinder, urlReader)

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()

		expectedID := uint64(2)
		expectedDBBanana := &dbb.Banana{ID: expectedID, Name: "Favourite Banana"}

		expectedBanana := &Banana{}
		expectedBanana.fromDB(expectedDBBanana)

		var expectedJSON bytes.Buffer
		err := json.NewEncoder(&expectedJSON).Encode([]*Banana{expectedBanana})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://somehost/bananas/%d", expectedID), nil)
		assert.Nil(t, err)

		bananaFinder.On("Find", uint64(2)).Return(expectedDBBanana, nil)
		urlReader.On("Read", req, "bananaID").Return(fmt.Sprintf("%d", expectedID))

		bananaHandler.ServeHTTP(w, req)
		assert.JSONEq(t, expectedJSON.String(), w.Body.String())
	})

	t.Run("Bad Request: Invalid ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "somerequest", nil)
		assert.Nil(t, err)

		urlReader.On("Read", req, "bananaID").Return("abcd")

		var expectedResponse bytes.Buffer
		err = json.NewEncoder(&expectedResponse).Encode(map[string]string{"description": "invalid banana ID provided"})
		assert.NoError(t, err)

		bananaHandler.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, expectedResponse.String(), w.Body.String())
	})
}
