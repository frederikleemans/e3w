package conf

import (
	"gopkg.in/ini.v1"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Config struct {
	Port          string
	Auth          bool
	EtcdRootKey   string
	DirValue      string
	EtcdEndPoints []string
	EtcdUsername  string
	EtcdPassword  string
	CertFile      string
	KeyFile       string
	CAFile        string
}

func Init(filepath string) (*Config, error) {
	cfg, err := ini.Load(filepath)
	if err != nil {
		return nil, err
	}

	c := &Config{}

	appSec := cfg.Section("app")
	c.Port = appSec.Key("port").Value()
	c.Auth = appSec.Key("auth").MustBool()

	etcdSec := cfg.Section("etcd")
	c.EtcdRootKey = etcdSec.Key("root_key").Value()
	c.DirValue = etcdSec.Key("dir_value").Value()
	c.EtcdEndPoints = etcdSec.Key("addr").Strings(",")
	c.EtcdUsername = etcdSec.Key("username").Value()
	c.EtcdPassword = etcdSec.Key("password").Value()
	c.CertFile = etcdSec.Key("cert_file").Value()
	c.KeyFile = etcdSec.Key("key_file").Value()
	c.CAFile = etcdSec.Key("ca_file").Value()

	return c, nil
}

func OverrideByEnv(c *Config) {
	c.Port = getEnv("E3W_PORT", c.Port)
	if value, ok := os.LookupEnv("E3W_AUTH"); ok {
		ret, err := strconv.ParseBool(value)
		if err == nil {
		   c.Auth = ret
		}
	}
	c.EtcdRootKey = getEnv("E3W_ETCDROOTKEY", c.EtcdRootKey)
	c.DirValue = getEnv("E3W_DIRVALUE", c.DirValue)
	c.EtcdEndPoints = strings.Split(getEnv("E3W_ETCDENDPOINTS", strings.Join(c.EtcdEndPoints,",")),",")
	c.EtcdUsername = getEnv("E3W_ETCDUSERNAME", c.EtcdUsername)
  c.EtcdPassword = getEnv("E3W_ETCDPASSWORD", c.EtcdPassword)
	c.CertFile = getEnv("E3W_CERTFILE", c.CertFile)
	c.KeyFile = getEnv("E3W_KEYFILE", c.KeyFile)
	c.CAFile = getEnv("E3W_CAFILE", c.CAFile)
}

func Print(c Config) {
	fmt.Println("e3w config:")
	fmt.Println("Port: ",c.Port)
	fmt.Println("Auth: ", c.Auth)
	fmt.Println("EtcdRootKey: ", c.EtcdRootKey)
	fmt.Println("DirValue:", c.DirValue)
	fmt.Println("EtcdEndPoints: ", c.EtcdEndPoints)
	fmt.Println("EtcdUsername: ", c.EtcdUsername)
	fmt.Println("EtcdPassword: ", c.EtcdPassword)
	fmt.Println("CertFile: ", c.CertFile)
	fmt.Println("KeyFile: ", c.KeyFile)
	fmt.Println("CAFile: ", c.CAFile)
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
