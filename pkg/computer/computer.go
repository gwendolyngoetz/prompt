package computer

import (
    "fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"strings"

	"github.com/go-ini/ini"
)

func GetUserName() string {
	user, err := user.Current()
fmt.Print(user.Username)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return user.Username
}

func GetHostname() string {
	// Built in os.Hostname() return fqdn
	hostname, err := exec.Command("hostname", "-s").Output()

	if err != nil {
		log.Fatalf(err.Error())
	}

	return strings.TrimSpace(string(hostname))
}

func GetOsIcon() string {
	distro := ""

	switch os := runtime.GOOS; os {
	case "darwin":
		distro = "osx"
	case "linux":
		distro = readOSRelease()["ID"]
	}

	result := ""

	switch distro {
	case "osx":
		result = ""
	case "ubuntu":
		result = ""
	case "debian":
		result = ""
	case "fedora":
		result = ""
	case "arch":
		result = ""
	case "manjaro":
		result = ""
	case "nixos":
		result = ""
	case "linux":
		result = ""
	default:
		result = ""
	}

	return result
}

func readOSRelease() map[string]string {
	cfg, err := ini.Load("/etc/os-release")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}

	ConfigParams := make(map[string]string)
	ConfigParams["ID"] = cfg.Section("").Key("ID").String()

	return ConfigParams
}

func GetPwd() string {
	// dir, err := exec.Command("dirs").Output()

	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// return string(dir)

	user, _ := user.Current()
	dir, _ := os.Getwd()

	result := strings.Replace(dir, user.HomeDir, "~", 1)

	return result
}

func IsNarrowWindow() bool {
	cols, err := exec.Command("tput", "cols").Output()

	if err != nil {
		log.Fatalf(err.Error())
	}

	count, _ := strconv.Atoi(string(cols))

	return count <= 64
}

func IsRemote() bool {
	if _, ok := os.LookupEnv("SSH_CLIENT"); ok {
		return true
	}

	return false
}

func IsSudo() bool {
	if _, ok := os.LookupEnv("SUDO_USER"); ok {
		return true
	}

	return false
	// return true
}
