package buildlog

import (
	"bytes"
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"

	"github.com/fajran/buildlog/pkg/buildlog"
	"github.com/fajran/buildlog/pkg/server"
	"github.com/fajran/buildlog/pkg/storage/disk"
)

type Server struct {
	Address string `yaml:"address"`
}

type DbConfig struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Options  string `yaml:"options"`
}

type Config struct {
	Server   Server   `yaml:"server"`
	Database DbConfig `yaml:"database"`
	DataPath string   `yaml:"data-path"`
}

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "", "Configuration file")
}

func (db DbConfig) Uri() string {
	var b bytes.Buffer

	b.WriteString("postgres://")

	if db.Username != "" {
		b.WriteString(db.Username)
		if db.Password != "" {
			b.WriteString(":")
			b.WriteString(db.Password)
		}
		b.WriteString("@")
	}

	b.WriteString(db.Host)
	b.WriteString("/")
	b.WriteString(db.Name)

	if db.Options != "" {
		b.WriteString("?")
		b.WriteString(db.Options)
	}

	return b.String()
}

func openDb(config DbConfig) (*sql.DB, error) {
	uri := config.Uri()
	log.Printf("Connecting to database %s", uri)
	return sql.Open("postgres", uri)
}

func readConfig(filename string) (Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	return config, err
}

func fillDefaultValues(config Config) {
	if config.Server.Address == "" {
		config.Server.Address = ":8080"
	}
}

func validateConfig(config Config) error {
	if config.DataPath == "" {
		return fmt.Errorf(`Missing configuration value: "data-path"`)
	}

	if config.Database.Host == "" {
		return fmt.Errorf(`Missing configuration value: "database.host"`)
	}
	if config.Database.Name == "" {
		return fmt.Errorf(`Missing configuration value: "database.name"`)
	}

	return nil
}

func Run() {
	flag.Parse()
	if configFile == "" {
		log.Fatal("Please enter configuration file")
	}

	config, err := readConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}
	fillDefaultValues(config)
	err = validateConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	db, err := openDb(config.Database)
	if err != nil {
		log.Fatal(err)
	}

	storage, err := disk.NewDiskStorage(config.DataPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Using data path at %s", config.DataPath)

	bl := buildlog.NewBuildLog(db, storage)

	log.Printf("Migrating database")
	err = bl.MigrateDb()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Running server on %s", config.Server.Address)
	s := server.NewServer(config.Server.Address, bl)
	s.Start()
}
