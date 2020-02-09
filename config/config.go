package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	_ "strings"
	"syscall"
	"time"

	homedir "github.com/mitchellh/go-homedir"
)

const (
	// ConfigFileName is the name of config file
	ConfigFileName = "config.json"
	configFileDir  = ".gic"
)

var (
	configDir = os.Getenv("GIC_CONFIG")
	p         = fmt.Println
)

// init the config file path
func init() {
	if configDir == "" {
		homePath, err := homedir.Dir()
		if err != nil {
			// just use ~ instead
			configDir = filepath.Join("~", configFileDir)
		}
		configDir = filepath.Join(homePath, configFileDir)
	}
}

// Dir returns the directory the configuration file is stored in
func Dir() string {
	return configDir
}

// load the config file into struct
func Load(configDir string) (*ConfigFile, error) {
	if configDir == "" {
		configDir = Dir()
	}

	filename := filepath.Join(configDir, ConfigFileName)
	// save the file to configfile struct
	configFile := InitConf(filename)

	if _, err := os.Stat(filename); err == nil {
		file, err := os.Open(filename)
		if err != nil {
			return configFile, errors.New("cannot open file " + filename)
		}
		defer file.Close()
		err = json.NewDecoder(file).Decode(configFile)
		return configFile, err
	} else if !os.IsNotExist(err) {
		// if file is there but we can't stat it for any reason other
		// than it doesn't exist then stop
		return configFile, errors.New("cannot stat file " + filename)
	}

	return configFile, nil
}

// check the token is created less than one hour
func CheckToken() (token string) {
	// check the token exists
	// check the token is created less than one hour
	configFile, _ := Load("")
	token = configFile.Token.Val
	if token != "" {
		now := time.Now()
		// p(now.Format(time.RFC3339))
		createdTime, err := time.Parse(time.RFC3339, configFile.Token.CreatedAt)
		// p("token created time: ", createdTime)
		createdTimePlus1 := createdTime.Add(1 * time.Hour)
		// p("created time +1 hour: ", createdTimePlus1)
		if err == nil {
			if now.Before(createdTimePlus1) {
				// p("before")
				return token
			}
		} else {
			p("created parse error: ", err)
		}
	}

	return ""
}

// ConfigFile ~/.gic/config.json file info
type ConfigFile struct {
	Url      string          `json:"url,omitempty"`
	Username string          `json:"username,omitempty"`
	Cred     string          `json:"cred,omitempty"`
	Token    ConfigFileToken `json:"token,omitempty"`
	Filename string          `json:"-"`
}

type ConfigFileToken struct {
	Val string `json:"val,omitempty"`
	// timestamp
	CreatedAt string `json:"createdAt,omitempty"`
}

func InitConf(fn string) *ConfigFile {
	return &ConfigFile{
		Url:      "",
		Username: "",
		Cred:     "",
		Token:    ConfigFileToken{Val: "", CreatedAt: ""},
		Filename: fn,
	}
}

// Save encodes and writes out all the authorization information
func (configFile *ConfigFile) Save() error {
	if configFile.Filename == "" {
		return errors.New("Can't save config with empty filename")
	}

	// create config file on disk
	dir := filepath.Dir(configFile.Filename)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	temp, err := ioutil.TempFile(dir, filepath.Base(configFile.Filename))
	if err != nil {
		return err
	}
	err = configFile.SaveToWriter(temp)
	defer temp.Close()
	if err != nil {
		os.Remove(temp.Name())
		return err
	}
	// Try copying the current config file (if any) ownership and permissions
	copyFilePermissions(configFile.Filename, temp.Name())

	return os.Rename(temp.Name(), configFile.Filename)
}

// SaveToWriter encodes and writes out all the config information to the given writer
func (configFile *ConfigFile) SaveToWriter(writer io.Writer) error {

	// InitConf(writer.Name())
	// configFile.Filename = writer.Name()
	configFile.Cred = ""
	data, err := json.MarshalIndent(configFile, "", "\t")
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	return err
}

// copyFilePermissions copies file ownership and permissions from "src" to "dst",
// ignoring any error during the process.
func copyFilePermissions(src, dst string) {
	var (
		mode     os.FileMode = 0600
		uid, gid int
	)

	fi, err := os.Stat(src)
	if err != nil {
		return
	}
	if fi.Mode().IsRegular() {
		mode = fi.Mode()
	}
	if err := os.Chmod(dst, mode); err != nil {
		return
	}

	uid = int(fi.Sys().(*syscall.Stat_t).Uid)
	gid = int(fi.Sys().(*syscall.Stat_t).Gid)

	if uid > 0 && gid > 0 {
		_ = os.Chown(dst, uid, gid)
	}
}
