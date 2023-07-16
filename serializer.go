package serializer

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/yukselcodingwithyou/stringbyteconverter"
	"gopkg.in/yaml.v3"
	"log"
)

// FileType is the supported file type in to serialize
type FileType string

const (
	YAML FileType = "YAML"
	JSON FileType = "JSON"
)

var converter stringbyteconverter.ToStringConverter

func init() {
	converter = stringbyteconverter.NewToStringConverter()
}

type Serializer interface {
	Serialize(value interface{}) (string, error)
}

// New creates a new serializer
func New(fileType FileType) Serializer {
	switch fileType {
	case YAML:
		return &yamlSerializer{}
	case JSON:
		return &jsonSerializer{}
	default:
		log.Fatalln("there is no such serializer implementation")
	}
	return &jsonSerializer{}
}

type jsonSerializer struct{}

// Serialize serializes an interface to string json
func (j jsonSerializer) Serialize(value interface{}) (string, error) {

	data, err := jsoniter.Marshal(value)
	return *converter.Convert(data), err
}

type yamlSerializer struct{}

// Serialize serializes an interface to string yml
func (y yamlSerializer) Serialize(value interface{}) (string, error) {

	data, err := yaml.Marshal(value)
	return *converter.Convert(data), err
}
