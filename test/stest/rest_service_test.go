//go:build servicetest
// +build servicetest

package integration

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/num30/golang-service-test/pkg/router"

	"github.com/go-resty/resty/v2"

	"github.com/stretchr/testify/assert"
)

var serviceHost = "http://localhost:8080"

// Configure service host from env variable
func init() {
	if h := os.Getenv("SERVICEHOST"); len(h) > 0 {
		serviceHost = h
	}
	log.Println("testing service: ", serviceHost)
}

// Test ping endpoint
func TestPingEndpoint(t *testing.T) {
	rest := resty.New().EnableTrace().SetDebug(false)
	t.Run("Ping", func(t *testing.T) {
		r, err := rest.R().Get(serviceHost + "/ping")
		assert.NoError(t, err)
		assert.Equal(t, r.StatusCode(), http.StatusOK)
	})
}

func Test_BoxesEndpoint(t *testing.T) {
	// make sure that test data is distinguishable from prod data and unique for test run
	testBoxId := fmt.Sprintf("test_box_id_%d", time.Now().Unix())
	testContent := "test box content"

	// it's a good practice to prepare tests data before test run and clean up after test run
	// however it's not always possible so it's ok to rely on test data that was generated earlier
	// For example you may have a test user that is used in integration tests to authenticat your requests

	rest := resty.New().EnableTrace().SetDebug(false)

	// test adding new box
	t.Run("AddBox", func(tt *testing.T) {
		r, err := rest.R().SetBody(router.Box{
			Content: testContent,
		}).Put(fmt.Sprintf("%s/boxes/%s", serviceHost, testBoxId))

		if assert.NoError(tt, err) {
			assert.Equal(tt, r.StatusCode(), http.StatusCreated)
		} else {
			// fail whole test because other steps depends on this one
			t.FailNow()
		}
	})

	// test getting box content
	t.Run("Get", func(tt *testing.T) {
		res := router.Box{}
		r, err := rest.R().SetResult(&res).Get(fmt.Sprintf("%s/boxes/%s", serviceHost, testBoxId))

		if assert.NoError(tt, err) {
			assert.Equal(tt, r.StatusCode(), http.StatusOK)
			assert.Equal(tt, testContent, res.Content)
		}
	})

	// test missing box
	t.Run("GetBoxFails", func(tt *testing.T) {
		r, err := rest.R().Get(fmt.Sprintf("%s/boxes/%s", serviceHost, "missing_box_id"))
		if assert.NoError(tt, err) {
			assert.Equal(tt, r.StatusCode(), http.StatusNotFound)
		}
	})
}
