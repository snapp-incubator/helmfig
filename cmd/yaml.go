package cmd

import (
	"os"
	"reflect"

	"github.com/iancoleman/strcase"
	log "github.com/sirupsen/logrus"
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
	rawFile, err := os.ReadFile(inputPath)
	if err != nil {
		log.WithError(err).WithField("path", inputPath).Fatal("error in reading the input config example file")
	}

	var parsedConfig map[interface{}]interface{}
	err = yaml.Unmarshal(rawFile, &parsedConfig)
	if err != nil {
		log.WithError(err).Fatal("error in unmarshalling the file with YAML format")
	}

	var configMap = map[interface{}]interface{}{}
	var values = map[interface{}]interface{}{}
	traverse(parsedConfig, configMap, values, "")

	rawConfigMap, err := yaml.Marshal(configMap)
	if err != nil {
		log.WithError(err).Fatal("error in marshalling the configmap file with YAML format")
	}

	err = os.WriteFile(configMapPath, rawConfigMap, 0644)
	if err != nil {
		log.WithError(err).WithField("path", configMapPath).Fatal("error in storing the configmap file")
	}

	log.WithField("path", configMapPath).Info("ConfigMap object content created successfully")

	rawValues, err := yaml.Marshal(values)
	if err != nil {
		log.WithError(err).Fatal("error in marshalling the values file with YAML format")
	}

	err = os.WriteFile(valuesPath, rawValues, 0644)
	if err != nil {
		log.WithError(err).WithField("path", valuesPath).Fatal("error in storing the values file")
	}

	log.WithField("path", valuesPath).Info("Values parameters created successfully")
}

func traverse(m, configMap, values map[interface{}]interface{}, valuesPath string) {
	for k, v := range m {
		lowerCamelKey := strcase.ToLowerCamel(k.(string))
		if reflect.TypeOf(v).Kind() == reflect.Map {
			var localConfigMap = map[interface{}]interface{}{}
			var localValues = map[interface{}]interface{}{}
			traverse(v.(map[interface{}]interface{}), localConfigMap, localValues, valuesPath+lowerCamelKey+".")
			configMap[k] = localConfigMap
			values[lowerCamelKey] = localValues
		} else {
			configMap[k] = "{{ " + valuesPath + lowerCamelKey + " }}"
			values[lowerCamelKey] = v
		}
	}
}
