package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/imdario/mergo"
	"gopkg.in/yaml.v2"
)

type Foo map[interface{}]interface{}

func main() {
	// v0.3.9: error on different types to merge
	// cannot append two different types (int, string)

	// v0.3.8: and earlier allow this
	//
	// There is a mergo.WithTypeCheck(c *Config)
	// but that defaults to false, so this behavior is
	// unexpected

	config := `foo: "1"`
	defaultConfig := `foo: 1`

	var (
		configMap  map[interface{}]interface{}
		defaultMap map[interface{}]interface{}
	)

	err := yaml.Unmarshal([]byte(config), &configMap)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(defaultConfig), &defaultMap)
	if err != nil {
		panic(err)
	}

	// v0.3.9 returns an error
	// versions <= v0.3.8 do not
	err = mergo.Merge(&configMap, &defaultMap)
	if err != nil {
		panic(err)
	}

	spew.Dump(configMap)

}
