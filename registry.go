package metalcloud

import (
	"fmt"
	"reflect"
	"strings"
)

var typeRegistry = make(map[string]reflect.Type)

func initTypeRegistry() {
	myTypes := []Applier{
		&InstanceArray{},
		&Datacenter{},
		&DriveArray{},
		&Infrastructure{},
		&Network{},
		&OSAsset{},
		&OSTemplate{},
		&Secret{},
		&Server{},
		&SharedDrive{},
		&StageDefinition{},
		&Workflow{},
		&SubnetPool{},
		&SubnetOOB{},
		&SwitchDevice{},
		&Variable{},
	}

	for _, v := range myTypes {
		t := reflect.ValueOf(v).Elem()
		u := reflect.TypeOf(v).Elem()
		typeRegistry[u.Name()] = t.Type()
	}
}

// GetObjectByKind creates an object of type <name>
func GetObjectByKind(name string) (reflect.Value, error) {
	initTypeRegistry()
	t, ok := typeRegistry[name]
	typesList := []string{}
	for k, _ := range typeRegistry {
		typesList = append(typesList, k)
	}

	if !ok {
		return reflect.Value{}, fmt.Errorf("%s was not supported by the apply method. Only the following types are supported: %s", name, strings.Join(typesList, ","))
	}

	v := reflect.New(t)
	return v, nil
}
