package Configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

type configuration struct {
	Database database
}

type database struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Sslmode  string
}

var once sync.Once
var c *configuration

func initConfiguration() {
	text, err := ioutil.ReadFile("configuration.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(text, &c)
	if err != nil {
		panic(err)
	}
}

func ConnStringFromConf() string {
	once.Do(initConfiguration)
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host, c.Database.User, c.Database.Password, c.Database.Dbname, c.Database.Sslmode)
}
