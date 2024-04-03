package metalcloud

import (
	"fmt"
	"reflect"
)

// GetTypesThatSupportApplierInterface returns a list of types that support the applier interface
func GetTypesThatSupportApplierInterface() map[string]reflect.Type {
	typeRegistry := map[string]reflect.Type{}
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

	return typeRegistry
}

// GetObjectByKind creates an object of type <name>. Only supported on types that implement
// the Applier interface. Use GetTypesThatSupportApplierInterface to get a list of supported types
func GetObjectByKind(name string) (reflect.Value, error) {
	typeRegistry := GetTypesThatSupportApplierInterface()
	t, ok := typeRegistry[name]

	if !ok {
		return reflect.Value{}, fmt.Errorf("%s was not recongnized as a valid product", name)
	}

	v := reflect.New(t)
	return v, nil
}
