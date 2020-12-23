package settings

import (
	"fmt"
	"os"
	"github.com/appointment-octopus/auth/utils"
	"strings"
)

var expirationDelta = map[string]int{
	"production": 72,
	"preproduction": 36,
	"tests": 1,
}

type Settings struct {
	PrivateKeyPath     string
	PublicKeyPath      string
	JWTExpirationDelta int
}

var settings Settings = Settings{}
var env = "preproduction"

func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting preproduction environment due to lack of GO_ENV value")
		env = "preproduction"
	}
	LoadSettingsByEnv(env)
}

func checkFileExists(filepath string) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		panic(err)
	}
}

func buildPath(path string, env string) string {
	rootDir := utils.RootDir()

	if env == "tests" {
		rootDir = strings.Split(rootDir, "/tests")[0]
	}

	var sb strings.Builder
	sb.WriteString(rootDir)
	sb.WriteString("/")
	sb.WriteString(path)
	fmt.Println(sb.String())
	checkFileExists(sb.String())
	return sb.String()
}

func LoadSettingsByEnv(env string) {
	settings = Settings{
		JWTExpirationDelta: expirationDelta[env],
		PrivateKeyPath: buildPath("settings/keys/private_key", env),
		PublicKeyPath: buildPath("settings/keys/public_key.pub", env),
	}
}

func GetEnvironment() string {
	return env
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}

func IsTestEnvironment() bool {
	return env == "tests"
}
