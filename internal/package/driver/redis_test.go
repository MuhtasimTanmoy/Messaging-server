package driver

import (
	"fmt"
	"github.com/nbio/st"
	"github.com/spf13/viper"
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

func TestRedis(t *testing.T) {
    driver := NewRedisDriver()

	ok, err := driver.Connect()
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	ok, err = driver.Ping()
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	count, err := driver.Del("channel")
	st.Expect(t, int(count), 0)
	st.Expect(t, err, nil)

	ok, err = driver.Set("channel", "test", 0)
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	ok, err = driver.Exists("channel")
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	value, err := driver.Get("channel")
	st.Expect(t, value, "test")
	st.Expect(t, err, nil)
}