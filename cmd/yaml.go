package cmd

import (
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var yamlCMD = &cobra.Command{
	Use:   "yaml",
	Short: "YAML config parser",
	Long:  `Enables you to generate configmap and values.yaml based on an example config formatted with YAML`,
	Run:   yamlFunc,
}

func init() {
	rootCMD.AddCommand(yamlCMD)

	yamlCMD.PersistentFlags().StringVarP(&inputPath, "example-config", "x", "config.example.yml", "Path to example yaml config file")
	yamlCMD.PersistentFlags().StringVar(&configMapPath, "configmap", "configmap.yaml", "Path to configmap file output")
	yamlCMD.PersistentFlags().StringVar(&valuesPath, "values", "values.yaml", "Path to values file output")
}

func yamlFunc(_ *cobra.Command, _ []string) {
	rawFile, err := ioutil.ReadFile(inputPath)
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

	err = ioutil.WriteFile(configMapPath, rawConfigMap, 0644)
	if err != nil {
		panic(err)
	}

	rawValues, err := yaml.Marshal(values)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(valuesPath, rawValues, 0644)
	if err != nil {
		panic(err)
	}
}

func convertName(val string) string {
	return strings.Replace(val, "_", "", -1)
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
		}
	}
}
