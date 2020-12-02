package metalcloud

//Applier should create or update an object.
type Applier interface {
	CreateOrUpdate(interface{}) error
	Delete(interface{}) error
}
