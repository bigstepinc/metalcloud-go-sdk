package metalcloud

//Applier should create or update an object.
type Applier interface {
	CreateOrUpdate(MetalCloudClient) error
	Delete(MetalCloudClient) error
	Validate() error
}
