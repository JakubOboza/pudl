package pudl

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	REFRESH_THROTTLE = 10
)

type PudlData struct {
	Config struct {
		TopCount int
	} `yaml:"config"`
	LastFetchTime time.Time
	Data          []Pudl `yaml:"db"`
}

func (pd *PudlData) ToYaml() ([]byte, error) {
	serialized, err := yaml.Marshal(&pd)
	if err != nil {
		return nil, err
	}
	return serialized, nil
}

func (pd *PudlData) SaveToFile() error {

	filePath, err := pudlConfigPath()

	if err != nil {
		return err
	}

	data, err := pd.ToYaml()

	if err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)

	return err
}

func pudlConfigPath() (string, error) {
	dirname, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return path.Join(dirname, PUDL_CONFIG_FILE), nil
}

func LoadDataFromYaml(rawYamlData []byte) (*PudlData, error) {
	pd := &PudlData{}

	err := yaml.Unmarshal(rawYamlData, pd)
	if err != nil {
		return nil, err
	}

	return pd, nil
}

func LoadConfigFromFile() (*PudlData, error) {
	filePath, err := pudlConfigPath()

	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return defaultConfig(), nil
	}

	rawContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	return LoadDataFromYaml(rawContent)
}

func defaultConfig() *PudlData {
	layout := "2006-01-02T15:04:05.000Z"
	str := "1985-01-31T22:22:13.371Z"
	defaultTime, _ := time.Parse(layout, str)
	return &PudlData{
		Config: struct{ TopCount int }{
			TopCount: 3,
		},
		LastFetchTime: defaultTime,
		Data:          []Pudl{},
	}
}
