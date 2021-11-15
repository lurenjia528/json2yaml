package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
)

var f string
var o string

func init() {
	flag.StringVar(&f, "f", "", "read file")
	flag.StringVar(&o, "o", "", "output json|yaml")
}

func main() {
	flag.Parse()
	b, err := ioutil.ReadFile(f)
	if err != nil {
		println("读文件失败")
	}

	if "json" == o {
		j, err := yaml.YAMLToJSON(b)
		if err != nil {
			fmt.Println("转换失败", err)
		}
		var out bytes.Buffer
		json.Indent(&out, j, "", "\t")
		fmt.Println(out.String())
		os.Exit(0)
	}
	if "yaml" == o {
		y, err := yaml.JSONToYAML(b)
		if err != nil {
			fmt.Println("转换失败", err)
		}
		fmt.Println(string(y))
		os.Exit(0)
	}
}
