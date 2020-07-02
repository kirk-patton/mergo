package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/imdario/mergo"
	"gopkg.in/yaml.v2"
)

type Foo map[interface{}]interface{}

func main() {
	// 3.9 error on different types to merge
	// interesting 3.8 merged maxMemory where one was a string and the other was an int
	config := `
---
image: docker.ouroath.com:4443/yahoo-cloud/redis-docker:20200228-202415
appName: systemregistry-redis
athensDomain: stat.k8s
masters: 2
replicas: 1
ykkName: systemregistry-redis-auth
ykkGroup: systemregistry.development
ykkEnv: aws
slaveServeStaleData: "no"
maxMemory: "5Gi"
port: 8021
cpu: "1.5"
environmentVariables:
  yamasNamespace: stat
jobs:
  component:
    colo: corp.bf1
    environment: development
  prod-corp-bf1:
    colo: corp.bf1
    environment: development
`

	defaultConfig := `
replicas: 0
port: 6379
maxMemory:  3221225472
maxMemoryPolicy: "allkeys-lru"
cpu: "500m"
timeoutSeconds: 60
clusterNodeTimeoutMillis: 5000
save: ""
slaveServeStaleData: "yes"
ykkEnv: "prod"
maxclients: 10000
`

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

	err = mergo.Merge(&configMap, &defaultMap)
	if err != nil {
		panic(err)
	}

	spew.Dump(configMap)

}
