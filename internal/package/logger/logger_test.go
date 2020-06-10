package logger

import (
	"fmt"
	"github.com/nbio/st"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
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

// TestLogging test cases
func TestLogging(t *testing.T) {

	currentTime := time.Now().Local()

	logFile := fmt.Sprintf(
		"%s%s/%s.log",
		os.Getenv("BasePath"),
		viper.GetString("log.path"),
		currentTime.Format("2006-01-02"),
	)

	// Start Test Cases
	Info("Info")
	Infoln("Infoln")
	Infof("Infof")
	Warning("Warning")
	Warningln("Warningln")
	Warningf("Warningf")

	data, err := ioutil.ReadFile(logFile)

	if err != nil {
		panic(err.Error())
	}

	st.Expect(t, strings.Contains(string(data), "Info\n"), true)
	st.Expect(t, strings.Contains(string(data), "Infoln\n"), true)
	st.Expect(t, strings.Contains(string(data), "Infof\n"), true)
	st.Expect(t, strings.Contains(string(data), "Warning\n"), true)
	st.Expect(t, strings.Contains(string(data), "Warningln\n"), true)
	st.Expect(t, strings.Contains(string(data), "Warningf\n"), true)

	os.Remove(logFile)
}
