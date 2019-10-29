package main

import (
	"fmt"
	. "github.com/cynic89/go-differ/differ"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//srcFilePath := "/Users/pivotal/workspace/pks-releng-ci-locks/aws-flannel-om25-terraform/unclaimed/apple-death.yml"
	//destFilePath := "/Users/pivotal/workspace/pks-locks/aws-flannel-om25-terraform/unclaimed/gentle-devourer"

	srcFilePath := os.Args[1]
	destFilePath := os.Args[2]

	diff, err := diffFromFiles(srcFilePath, destFilePath)
	if err != nil {
		fmt.Printf("Failed with error %v", err)
		os.Exit(1)
	}
	fmt.Println()
	fmt.Printf("Following differences were identified. ie. These are extra properties present in %v", srcFilePath)
	fmt.Println()
	fmt.Println()
	printResults(diff)

}

func diffFromFiles(srcFile string, destFile string) (diffMap map[string]interface{}, err error) {

	data, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed while reading file from %s", srcFile))
	}

	srcParseable, err := getParseable(srcFile, data)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed while making a parseable file from %s", srcFile))
	}

	dataJson, err := ioutil.ReadFile(destFile)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed while reading file from %s", destFile))
	}

	destParseable, err := getParseable(destFile, dataJson)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Failed while making a parseable file from %s", destFile))
	}

	return Diff(srcParseable, destParseable)

}

func getParseable(filename string, data []byte) (Parseable, error) {

	if strings.Contains(filename, ".yml") || strings.Contains(filename, ".yaml") {
		return Yaml{data}, nil
	}

	//a hack check to see if it's json since some json files in pks locks are not named properly
	if strings.Contains(filename, ".json") || data[0]==123  {
		return Json{data}, nil
	}


	return nil, errors.New("Only files of type json and yaml/yml are supported")
}

func printResults(diff map[string]interface{})  {
	for k, v := range diff{
		fmt.Printf("%v => %v", k, v)
		fmt.Println()
	}
}