package metalcloud

import (
	"fmt"
	"reflect"
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
		&SwitchDevice{},
		&Variable{},
	}

	for _, v := range myTypes {
		t := reflect.ValueOf(v).Elem()
		u := reflect.TypeOf(v).Elem()
		typeRegistry[u.Name()] = t.Type()
	}
}

//GetObjectByKind creates an object of type <name>
func GetObjectByKind(name string) (reflect.Value, error) {
	initTypeRegistry()
	t, ok := typeRegistry[name]

	if !ok {
		return reflect.Value{}, fmt.Errorf("%s was not recongnized as a valid product", name)
	}

	v := reflect.New(t)
	return v, nil
}
