package config

import (
	"fmt"
	cfssl_config "github.com/cloudslit/cfssl/config"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/koding/multiconfig"

	"github.com/cloudslit/newca/pkg/util/json"
)

var (
	// C Global configuration (Must Load first, otherwise the configuration will not be available)
	C    = new(Config)
	once sync.Once
)

// MustLoad load config
func MustLoad(fpaths ...string) error {
	once.Do(func() {
		loaders := []multiconfig.Loader{
			&multiconfig.TagLoader{},
			&multiconfig.EnvironmentLoader{},
		}

		for _, fpath := range fpaths {
			if strings.HasSuffix(fpath, "toml") {
				loaders = append(loaders, &multiconfig.TOMLLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "json") {
				loaders = append(loaders, &multiconfig.JSONLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "yaml") {
				loaders = append(loaders, &multiconfig.YAMLLoader{Path: fpath})
			}
		}
		m := multiconfig.DefaultLoader{
			Loader:    multiconfig.MultiLoader(loaders...),
			Validator: multiconfig.MultiValidator(&multiconfig.RequiredValidator{}),
		}
		m.MustLoad(C)
	})
	PrintWithJSON()
	return ParseConfigByEnv()
}

func ParseConfigByEnv() error {
	// APP
	if v := os.Getenv("CA_APP_HOST"); v != "" {
		C.App.Host = v
	}
	if v := os.Getenv("CA_APP_PORT"); v != "" {
		p, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("environment variable [%s] parsing error:%v", "CA_APP_PORT", err)
		}
		C.App.Port = p
	}
	if v := os.Getenv("CA_APP_CERT_FILE"); v != "" {
		C.App.CertFile = v
	}
	if v := os.Getenv("CA_APP_KEY_FILE"); v != "" {
		C.App.KeyFile = v
	}
	// Log
	if v := os.Getenv("CA_LOG_HOOK_ENABLED"); v == "true" {
		C.Log.EnableHook = true
	}
	if v := os.Getenv("CA_LOG_REDIS_ADDR"); v != "" {
		C.LogRedisHook.Addr = v
	}
	if v := os.Getenv("CA_LOG_REDIS_KEY"); v != "" {
		C.LogRedisHook.Key = v
	}
	// Storage
	if v := os.Getenv("CA_STORAGE_TYPE"); v != "" {
		C.Storage.Type = v
	}
	// IPFS
	if v := os.Getenv("CA_IPFS_HOST"); v != "" {
		C.Ipfs.Host = v
	}
	if v := os.Getenv("CA_IPFS_PORT"); v != "" {
		p, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("environment variable [%s] parsing error:%v", "CA_IPFS_PORT", err)
		}
		C.Ipfs.Port = p
	}

	// Cfssl
	cfg, err := cfssl_config.LoadFile(C.Cfssl.ConfigFile)
	if err != nil {
		return err
	}
	C.Cfssl.Config = cfg
	return nil
}

func PrintWithJSON() {
	if C.PrintConfig {
		b, err := json.MarshalIndent(C, "", " ")
		if err != nil {
			os.Stdout.WriteString("[CONFIG] JSON marshal error: " + err.Error())
			return
		}
		os.Stdout.WriteString(string(b) + "\n")
	}
}

type Config struct {
	RunMode      string
	PrintConfig  bool
	App          App
	Log          Log
	LogRedisHook LogRedisHook
	Cfssl        Cfssl
	Storage      Storage
	Ipfs         Ipfs
	CrdtKv       CrdtKv
}

func (c *Config) IsDebugMode() bool {
	return c.RunMode == "debug"
}

func (c *Config) IsReleaseMode() bool {
	return c.RunMode == "release"
}

type LogHook string

func (h LogHook) IsRedis() bool {
	return h == "redis"
}

type Log struct {
	Level         int
	Format        string
	Output        string
	OutputFile    string
	EnableHook    bool
	HookLevels    []string
	Hook          LogHook
	HookMaxThread int
	HookMaxBuffer int
}

type LogGormHook struct {
	DBType       string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	Table        string
}

type LogRedisHook struct {
	Addr string
	Key  string
}

type Redis struct {
	Addr     string
	Password string
}

// App Configuration parameters
type App struct {
	Host               string
	Port               int
	CertFile           string
	KeyFile            string
	Cert               []byte
	Key                []byte
	ShutdownTimeout    int
	MaxContentLength   int64
	MaxReqLoggerLength int `default:"1024"`
	MaxResLoggerLength int
}

type Cfssl struct {
	ConfigFile string
	Config     *cfssl_config.Config
}

type Storage struct {
	Type string
}

func (h Storage) IsIpfs() bool {
	return h.Type == "ipfs"
}

type Ipfs struct {
	Host string
	Port int
}

type CrdtKv struct {
	NodeServiceName     string // Service Discovery Identification
	DataStorePath       string // Data storage path
	DataSyncChannel     string // Pubsub data synchronization channel
	NetDiscoveryChannel string // Node discovery channel
	Namespace           string
}
