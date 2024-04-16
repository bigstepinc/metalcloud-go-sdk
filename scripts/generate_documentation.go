package main

import (
	"fmt"
	"os"

	metalcloud "github.com/metalsoft-io/metal-cloud-sdk-go/v3"
	"github.com/projectdiscovery/yamldoc-go/encoder"
)

type docFunc func() *encoder.FileDoc

// add here your other objects
var OBJECT_DOC_FUNCS = []docFunc{
	metalcloud.GetInstanceArrayDoc,
	metalcloud.GetSubnetOOBDoc,
	metalcloud.GetDatacenterWithConfigDoc,
	metalcloud.GetVariableDoc,
}

// This will generate the markdown documentation for all the objects
// do not forget to include all of the objects here and regenerate
func main() {
	if len(os.Args) != 2 {
		panic(fmt.Errorf("syntax: %s <output_file_name>", os.Args[0]))
	}

	fpath := os.Args[1]
	f, err := os.OpenFile(fpath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	for _, object_doc_func := range OBJECT_DOC_FUNCS {
		fileDoc := object_doc_func()

		data, err := fileDoc.Encode()
		if err != nil {
			panic(err)
		}
		_, err = f.Write(data)
		if err != nil {
			panic(err)
		}
	}
	/*
		data, err := metalcloud.GetSubnetOOBDoc().Encode()
		if err != nil {
			panic(err)
		}
		_, err = f.Write(data)
		if err != nil {
			panic(err)
		}

		data, err = metalcloud.GetVariableDoc().Encode()
		if err != nil {
			panic(err)
		}
		_, err = f.Write(data)
		if err != nil {
			panic(err)
		}
	*/

	f.Close()

}
