package utils

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestParsePageLimit(t *testing.T) {
	var testCase gin.Context
	request, err := http.NewRequest("GET", "localhost:8080/api/v1/batches?page=2&limit=10", nil)

	if err != nil {
		t.Errorf("Fail to test func ParsePageLimit")
	}

	testCase.Request = request
	pageResult, limitResult := ParsePageLimit(&testCase)
	pageExp, limitExp := 2, 10

	if !assert.EqualValues(t, pageExp, pageResult) || !assert.EqualValues(t, limitExp, limitResult) {
		t.Errorf("expected limit %v but got %v and page %v but got %v", limitExp, limitResult, pageExp, pageResult)
	}
}
