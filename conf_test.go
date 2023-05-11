package sage

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"sage/conf"
	"sage/conf/file"
	"sage/config"
	"testing"
)

func TestConfig_file(t *testing.T) {

	c := conf.New(conf.WithSource(file.NewSource("./config/config.yaml")),
		conf.WithDecoder(func(kv *conf.KeyValue, i map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, i)
		}))
	er := c.Load()
	fmt.Println(er)
	b, _ := json.Marshal(c)
	fmt.Println(string(b))
	fmt.Println(c)

	var bc config.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	b, _ = json.Marshal(bc)
	fmt.Println("++++", string(b))

}
