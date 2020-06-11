package auth

import (
    "fmt"
    "os"
	"strconv"
	"testing"
	"time"
    "github.com/spf13/viper"
    "github.com/nbio/st"
    "github.com/MuhtasimTanmoy/messaging_server/internal/package/utils"
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
func TestAuth(t *testing.T) {

    now := time.Now().Unix()
	id := utils.GenerateUUID4()

	_, err := GenerateJWTToken(
		fmt.Sprintf("%s@%d", id, now),
		now,
		viper.GetString("app.secret"),
	)

    st.Expect(t, err, nil)
}