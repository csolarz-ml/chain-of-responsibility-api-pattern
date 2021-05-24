package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/csolarz-ml/chain-of-responsibility-api-pattern/model"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := Setup()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

type cases struct {
	input   string
	amount  int
	code    int
	message string
}

func TestLoanRoute(t *testing.T) {
	router := Setup()

	t.Run("Should be not authorized by middleware", func(t *testing.T) {
		body := strings.NewReader(`{"amount": 500}`)

		request, _ := http.NewRequest("POST", fmt.Sprintf("/loan"), body)
		request.Header.Add("Content-Type", "application/json")

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})

	t.Run("Should be fail on amount validation by middleware", func(t *testing.T) {
		body := strings.NewReader(`{"amount": "none"}`)

		request, _ := http.NewRequest("POST", fmt.Sprintf("/loan"), body)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Secret-Api-Key", "whatever")

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should be ok on account executive", func(t *testing.T) {
		cases := []cases{
			{
				input:   `{"amount": 5}`,
				amount:  5,
				code:    200,
				message: "loan ok!",
			},
			{
				input:   `{"amount": 5000}`,
				amount:  5000,
				code:    200,
				message: "loan ok!",
			},
			{
				input:   `{"amount": 50000}`,
				amount:  50000,
				code:    200,
				message: "loan ok!",
			},
		}

		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				body := strings.NewReader(c.input)

				request, _ := http.NewRequest("POST", fmt.Sprintf("/loan"), body)
				request.Header.Add("Content-Type", "application/json")
				request.Header.Add("Secret-Api-Key", "whatever")

				response := httptest.NewRecorder()
				router.ServeHTTP(response, request)

				loanJson, _ := ioutil.ReadAll(response.Body)

				var loan model.Loan
				if err := json.Unmarshal([]byte(loanJson), &loan); err != nil {
					t.Fail()
				}

				assert.Equal(t, c.code, response.Code)
				assert.Equal(t, c.amount, loan.Form.Amount)
				assert.Equal(t, c.message, loan.Message)
			})
		}
	})
}
