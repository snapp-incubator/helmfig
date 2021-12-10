package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
)

func convertName(val string) string {
	return strings.Replace(val, "_", "", -1)
}

func main() {
	rawFile, err := ioutil.ReadFile("config.example.yml")
	if err != nil {
		panic(err)
	}

	var parsedConfig map[interface{}]interface{}
	err = yaml.Unmarshal(rawFile, &parsedConfig)
	if err != nil {
		panic(err)
	}

	var configMap = map[interface{}]interface{}{}
	var values = map[interface{}]interface{}{}
	f(parsedConfig, configMap, values, "")

	rawConfigMap, err := yaml.Marshal(configMap)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("configmap.yaml", rawConfigMap, 0644)
	if err != nil {
		panic(err)
	}

	rawValues, err := yaml.Marshal(values)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("values.yaml", rawValues, 0644)
	if err != nil {
		panic(err)
	}
}

func f(m, configMap, values map[interface{}]interface{}, valuesPath string) {
	for k, v := range m {
		if reflect.TypeOf(v).Kind() == reflect.Map {
			var localConfigMap = map[interface{}]interface{}{}
			var localValues = map[interface{}]interface{}{}
			f(v.(map[interface{}]interface{}), localConfigMap, localValues, valuesPath+convertName(k.(string))+".")
			configMap[k] = localConfigMap
			values[convertName(k.(string))] = localValues
		} else {
			configMap[k] = "{{ " + valuesPath + convertName(k.(string)) + " }}"
			values[convertName(k.(string))] = v
			fmt.Println(k, v, reflect.TypeOf(v))
		}
	}
}
