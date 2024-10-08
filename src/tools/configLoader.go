package tools

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	_ "github.com/golang/protobuf/proto" //To force go.sum to be generated
)

type ConfigLoader struct {
	configs map[string]string
	folder  string
}

const (
	//baseFilePath = "../../configs/"
	configType = ".cfg"
)

var (
	SharedConfig *ConfigLoader
)

func (config *ConfigLoader) InitEmptyConfig() {
	config.configs = make(map[string]string)
}

func (config *ConfigLoader) LoadConfigs(folder string) {
	SharedConfig = config
	config.configs = make(map[string]string)
	if folder[len(folder)-1] != '/' {
		//config.folder = baseFilePath + folder + "/"
		config.folder = folder + "/"
	} else {
		//config.folder = baseFilePath + folder
		config.folder = folder
	}
	filesToRead := config.getConfigFiles()
	filesRead := 0
	for _, filePath := range filesToRead {
		config.readConfigFile(filePath)
		filesRead++
	}
	fmt.Printf("Finished reading %d config files from %s.\n", filesRead, folder)
}

func (config *ConfigLoader) HasConfig(key string) (has bool) {
	_, has = config.configs[key]
	return
}

func (config *ConfigLoader) GetConfig(key string) (value string) {
	return config.configs[key]
}

func (config *ConfigLoader) GetAndHasConfig(key string) (value string, has bool) {
	value, has = config.configs[key]
	return
}

func (config *ConfigLoader) GetOrDefault(key string, def string) (value string) {
	value, has := config.configs[key]
	if !has {
		return def
	}
	return value
}

func (config *ConfigLoader) GetBoolConfig(key string, def bool) bool {
	value, has := config.configs[key]
	if !has {
		return def
	}
	result, _ := strconv.ParseBool(value)
	return result
}

func (config *ConfigLoader) GetIntConfig(key string, def int) int {
	value, has := config.configs[key]
	if !has {
		return def
	}
	result, _ := strconv.ParseInt(value, 10, 64)
	return int(result)
}

func (config *ConfigLoader) GetInt64Config(key string, def int64) int64 {
	value, has := config.configs[key]
	if !has {
		return def
	}
	result, _ := strconv.ParseInt(value, 10, 64)
	return result
}

func (config *ConfigLoader) GetInt32Config(key string, def int32) int32 {
	value, has := config.configs[key]
	if !has {
		return def
	}
	result, _ := strconv.ParseInt(value, 10, 32)
	return int32(result)
}

func (config *ConfigLoader) GetFloatConfig(key string, def float64) float64 {
	value, has := config.configs[key]
	if !has {
		return def
	}
	result, _ := strconv.ParseFloat(value, 64)
	return result
}

// Uses a whitespace to slice the string into different substrings
func (config *ConfigLoader) GetStringSliceConfig(key string, def string) []string {
	value, has := config.configs[key]
	if has {
		return strings.Split(value, " ")
	}
	return strings.Split(def, " ")
}

// Splits by comma instead of space.
func (config *ConfigLoader) GetStringSliceCommaConfig(key string, def string) []string {
	value, has := config.configs[key]
	if has {
		return strings.Split(value, ",")
	}
	return strings.Split(def, ",")
}

func (config *ConfigLoader) ReplaceConfig(key string, value string) {
	config.configs[key] = value
}

func (config *ConfigLoader) getConfigFiles() (fileNames []string) {
	files, _ := ioutil.ReadDir(config.folder)
	fileNames = make([]string, 0, 10)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), configType) {
			fileNames = append(fileNames, config.folder+file.Name())
		}
	}
	return
}

func (config *ConfigLoader) readConfigFile(fileName string) {
	file, _ := os.Open(fileName)
	defer file.Close()
	in := bufio.NewReader(file)
	var err error
	var str string
	for err == nil {
		str, err = in.ReadString('\n')
		if err == nil && !config.isComment(str) {
			splitI := strings.Index(str, "=")
			parts := []string{str[0:splitI], str[splitI+1:]}
			//parts := strings.Split(str, "=")
			//fmt.Println("Read:", parts)
			config.configs[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			//fmt.Println(strings.TrimSpace(parts[0]) + "=" + strings.TrimSpace(parts[1]))
		}
	}
}

func (config *ConfigLoader) isComment(line string) (isComment bool) {
	return line == "\n" || strings.HasPrefix("//", line) || line[0] == '#'
}
