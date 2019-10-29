package differ

import (
	"encoding/json"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Parseable interface {
	Parse() (map[string]interface{}, error)
}

type Json struct {
	Val []byte

}

type Yaml struct {
	Val []byte

}

func (j Json) Parse() (parsedVal map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(j.Val), &parsedVal)
	if err != nil {
		return nil, errors.Wrap(err, "Error while parsing json")
	}

	return parsedVal, nil
}

func (y Yaml) Parse() (parsedVal map[string]interface{}, err error) {
	err = yaml.Unmarshal([]byte(y.Val), &parsedVal)
	if err != nil {
		return nil, errors.Wrap(err, "Error while parsing yaml")
	}

	return parsedVal, nil
}


func Diff(src Parseable, dest Parseable) (diffMap map[string]interface{}, err error){

	srcMap, err := src.Parse()
	if err != nil {
		return nil, err
	}

	destMap, err := dest.Parse()
	if err != nil {
		return nil, err
	}

	diffMap = findDiffKeys(srcMap, destMap)

	return diffMap, nil
}

func findDiffKeys(src map[string]interface{}, dest map[string]interface{}) (diffMap map[string]interface{}) {

	diffMap = make(map[string]interface{})

	for key, val := range src {

		if dest[key] == nil{
			diffMap[key] = val
		}
	}

	return diffMap
}