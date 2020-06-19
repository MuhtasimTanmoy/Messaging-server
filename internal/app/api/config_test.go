package api

import (
	"fmt"
	// "github.com/gin-gonic/gin"
	"github.com/nbio/st"
	"github.com/spf13/viper"
	// "net/http"
	// "net/http/httptest"
	"os"
	"strconv"
	"testing"
)

const BasePath = "/Users/tanmoy/memory/projects/messaging_server"

// init setup stuff
func init() {
	configFile := fmt.Sprintf("%s/%s",BasePath, "config.yml")

	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf(
			"Error while loading config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	os.Setenv("BasePath", fmt.Sprintf("%s/", BasePath))
	os.Setenv("PORT", strconv.Itoa(viper.GetInt("app.port")))
}


// TestHealthCheckController test case
func TestHealthCheckController(t *testing.T) {

	configResult := ConfigResult{Key: "key", Value: "value"}
	jsonValue, err := configResult.ConvertToJSON()
	st.Expect(t, jsonValue, `{"key":"key","value":"value"}`)
	st.Expect(t, err, nil)

	ok, err := configResult.LoadFromJSON([]byte(jsonValue))
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)
	st.Expect(t, configResult.Key, "key")
	st.Expect(t, configResult.Value, "value")
}
