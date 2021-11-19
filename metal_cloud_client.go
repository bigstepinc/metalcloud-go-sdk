// Code generated by ifacemaker; DO NOT EDIT.

package metalcloud



// MetalCloudClient interface used for mocking and abstracting the backend
type MetalCloudClient interface {
	//Datacenters returns datacenters for all users
	Datacenters(onlyActive bool) (*map[string]Datacenter, error)
	//DatacentersByUserID returns datacenters for specific user
	DatacentersByUserID(userID int, onlyActive bool) (*map[string]Datacenter, error)
	//DatacentersByUserEmail returns datacenters by email
	DatacentersByUserEmail(userEmail string, onlyActive bool) (*map[string]Datacenter, error)
	//DatacenterGet returns details of a specific datacenter
	DatacenterGet(datacenterName string) (*Datacenter, error)
	//DatacenterGetForUserByEmail returns details of a specific datacenter
	DatacenterGetForUserByEmail(datacenterName string, userID string) (*Datacenter, error)
	//DatacenterGetForUserByID returns details of a specific datacenter
	DatacenterGetForUserByID(datacenterName string, userID int) (*Datacenter, error)
	//DatacenterConfigGet returns details of a specific datacenter
	DatacenterConfigGet(datacenterName string) (*DatacenterConfig, error)
	//DatacenterConfigUpdate Updates configuration information for a specified Datacenter.
	DatacenterConfigUpdate(datacenterName string, datacenterConfig DatacenterConfig) error
	//DatacenterCreate creates a new Datacenter
	DatacenterCreate(datacenter Datacenter, datacenterConfig DatacenterConfig) (*Datacenter, error)
	//DatacenterAgentsConfigJSONDownloadURL returns the agent url (and automatically decrypts it)
	DatacenterAgentsConfigJSONDownloadURL(datacenterName string, decrypt bool) (string, error)
	//DriveArrays retrieves the list of drives arrays of an infrastructure
	DriveArrays(infrastructureID int) (*map[string]DriveArray, error)
	//DriveArraysByLabel retrieves the list of drives arrays of an infrastructure
	DriveArraysByLabel(infrastructureLabel string) (*map[string]DriveArray, error)
	//DriveArrayGet retrieves a DriveArray object with specified ids
	DriveArrayGet(driveArrayID int) (*DriveArray, error)
	//DriveArrayGetByLabel retrieves a DriveArray object with specified ids
	DriveArrayGetByLabel(driveArrayLabel string) (*DriveArray, error)
	//DriveArrayCreate creates a drive array. Requires deploy.
	DriveArrayCreate(infrastructureID int, driveArray DriveArray) (*DriveArray, error)
	//DriveArrayCreateByLabel creates a drive array. Requires deploy.
	DriveArrayCreateByLabel(infrastructureLabel string, driveArray DriveArray) (*DriveArray, error)
	//DriveArrayEdit alters a deployed drive array. Requires deploy.
	DriveArrayEdit(driveArrayID int, driveArrayOperation DriveArrayOperation) (*DriveArray, error)
	//DriveArrayEditByLabel alters a deployed drive array. Requires deploy.
	DriveArrayEditByLabel(driveArrayLabel string, driveArrayOperation DriveArrayOperation) (*DriveArray, error)
	//DriveArrayDelete deletes a Drive Array with specified id
	DriveArrayDelete(driveArrayID int) error
	//DriveArrayDeleteByLabel deletes a Drive Array with specified id
	DriveArrayDeleteByLabel(driveArrayLabel string) error
	//DriveArrayDrives returns the drives of a drive array
	DriveArrayDrives(driveArray int) (*map[string]Drive, error)
	//DriveArrayDrivesByLabel returns the drives of a drive array
	DriveArrayDrivesByLabel(driveArrLabel string) (*map[string]Drive, error)
	//DriveSnapshotCreate creates a drive snapshot
	DriveSnapshotCreate(driveID int) (*Snapshot, error)
	//DriveSnapshotDelete creates a drive snapshot
	DriveSnapshotDelete(driveSnapshotID int) error
	//DriveSnapshotRollback rolls a Drive back to a specified DriveSnapshot. The specified snapshot is not destroyed and can be reused.
	DriveSnapshotRollback(driveSnapshotID int) error
	//DriveSnapshotGet gets a drive snapshot
	DriveSnapshotGet(driveSnapshotID int) (*Snapshot, error)
	//DriveSnapshots retrieves a list of all the snapshot objects
	DriveSnapshots(driveID int) (*map[string]Snapshot, error)
	//ExternalConnections returns a list of external connections for the specified datacenter
	ExternalConnections(datacenterName string) (*[]ExternalConnection, error)
	//ExternalConnectionCreate creates an external connection.
	ExternalConnectionCreate(externalConnection ExternalConnection) (*ExternalConnection, error)
	//ExternalConnectionGet returns an external connection with specified id
	ExternalConnectionGet(externalConnectionID int) (*ExternalConnection, error)
	//ExternalConnectionGetByLabel returns an external connection with specified id
	ExternalConnectionGetByLabel(externalConnectionLabel string) (*ExternalConnection, error)
	//ExternalConnectionEdit updates an external connection.
	ExternalConnectionEdit(externalConnectionID int, externalConnection ExternalConnection) (*ExternalConnection, error)
	//ExternalConnectionEditByLabel updates an external connection.
	ExternalConnectionEditByLabel(externalConnectionLabel string, externalConnection ExternalConnection) (*ExternalConnection, error)
	//ExternalConnectionDelete deletes an external connection.
	ExternalConnectionDelete(externalConnectionID int) error
	//ExternalConnectionDeleteByLabel deletes an external connection.
	ExternalConnectionDeleteByLabel(externalConnectionLabel string) error
	//InfrastructureCreate creates an infrastructure
	InfrastructureCreate(infrastructure Infrastructure) (*Infrastructure, error)
	//Infrastructures returns a list of infrastructures
	Infrastructures() (*map[string]Infrastructure, error)
	//InfrastructureEdit alters an infrastructure
	InfrastructureEdit(infrastructureID int, infrastructureOperation InfrastructureOperation) (*Infrastructure, error)
	//InfrastructureEditByLabel alters an infrastructure
	InfrastructureEditByLabel(infrastructureLabel string, infrastructureOperation InfrastructureOperation) (*Infrastructure, error)
	//InfrastructureDelete deletes an infrastructure and all associated elements. Requires deploy
	InfrastructureDelete(infrastructureID int) error
	//InfrastructureDeleteByLabel deletes an infrastructure and all associated elements. Requires deploy
	InfrastructureDeleteByLabel(infrastructureLabel string) error
	//InfrastructureOperationCancel reverts (undos) alterations done before deploy
	InfrastructureOperationCancel(infrastructureID int) error
	//InfrastructureOperationCancelByLabel reverts (undos) alterations done before deploy
	InfrastructureOperationCancelByLabel(infrastructureLabel string) error
	//InfrastructureDeploy initiates a deploy operation that will apply all registered changes for the respective infrastructure
	InfrastructureDeploy(infrastructureID int, shutdownOptions ShutdownOptions, allowDataLoss bool, skipAnsible bool) error
	//InfrastructureDeployByLabel initiates a deploy operation that will apply all registered changes for the respective infrastructure
	InfrastructureDeployByLabel(infrastructureLabel string, shutdownOptions ShutdownOptions, allowDataLoss bool, skipAnsible bool) error
	//InfrastructureDeployWithOptions initiates a deploy operation that will apply all registered changes for the respective infrastructure. With options.
	InfrastructureDeployWithOptions(infrastructureID int, shutdownOptions ShutdownOptions, deployOptions *DeployOptions, allowDataLoss bool, skipAnsible bool) error
	//InfrastructureDeployWithOptionsByLabel initiates a deploy operation that will apply all registered changes for the respective infrastructure. With options.
	InfrastructureDeployWithOptionsByLabel(infrastructureLabel string, shutdownOptions ShutdownOptions, deployOptions *DeployOptions, allowDataLoss bool, skipAnsible bool) error
	//InfrastructureGet returns a specific infrastructure by id
	InfrastructureGet(infrastructureID int) (*Infrastructure, error)
	//InfrastructureGetByLabel returns a specific infrastructure by id
	InfrastructureGetByLabel(infrastructureLabel string) (*Infrastructure, error)
	//InfrastructureUserLimits returns user metadata
	InfrastructureUserLimits(infrastructureID int) (*map[string]interface{}, error)
	//InfrastructureUserLimitsByLabel returns user metadata
	InfrastructureUserLimitsByLabel(infrastructureLabel string) (*map[string]interface{}, error)
	//InstanceArrayInterfaceAttachNetwork attaches an InstanceArrayInterface to a Network
	InstanceArrayInterfaceAttachNetwork(instanceArrayID int, instanceArrayInterfaceIndex int, networkID int) (*InstanceArray, error)
	//InstanceArrayInterfaceDetach detaches an InstanceArrayInterface from any Network element that is attached to.
	InstanceArrayInterfaceDetach(instanceArrayID int, instanceArrayInterfaceIndex int) (*InstanceArray, error)
	//InstanceArrayGet returns an InstanceArray with specified id
	InstanceArrayGet(instanceArrayID int) (*InstanceArray, error)
	//InstanceArrayGetByLabel returns an InstanceArray with specified id
	InstanceArrayGetByLabel(instanceArrayLabel string) (*InstanceArray, error)
	//InstanceArrays returns list of instance arrays of specified infrastructure
	InstanceArrays(infrastructureID int) (*map[string]InstanceArray, error)
	//InstanceArraysByLabel returns list of instance arrays of specified infrastructure
	InstanceArraysByLabel(infrastructureLabel string) (*map[string]InstanceArray, error)
	//InstanceArrayCreate creates an instance array (colletion of identical instances). Requires Deploy.
	InstanceArrayCreate(infrastructureID int, instanceArray InstanceArray) (*InstanceArray, error)
	//InstanceArrayCreateByLabel creates an instance array (colletion of identical instances). Requires Deploy.
	InstanceArrayCreateByLabel(infrastructureLabel string, instanceArray InstanceArray) (*InstanceArray, error)
	//InstanceArrayEdit alterns a deployed instance array. Requires deploy.
	InstanceArrayEdit(instanceArrayID int, instanceArrayOperation InstanceArrayOperation, bSwapExistingInstancesHardware *bool, bKeepDetachingDrives *bool, objServerTypeMatches *ServerTypeMatches, arrInstancesToBeDeleted *[]int) (*InstanceArray, error)
	//InstanceArrayEditByLabel alterns a deployed instance array. Requires deploy.
	InstanceArrayEditByLabel(instanceArrayLabel string, instanceArrayOperation InstanceArrayOperation, bSwapExistingInstancesHardware *bool, bKeepDetachingDrives *bool, objServerTypeMatches *ServerTypeMatches, arrInstancesToBeDeleted *[]int) (*InstanceArray, error)
	//InstanceArrayDelete deletes an instance array. Requires deploy.
	InstanceArrayDelete(instanceArrayID int) error
	//InstanceArrayDeleteByLabel deletes an instance array. Requires deploy.
	InstanceArrayDeleteByLabel(instanceArrayLabel string) error
	//InstanceArrayStop stops a specified InstanceArray.
	InstanceArrayStop(instanceArrayID int) (*InstanceArray, error)
	//InstanceArrayStopByLabel stops a specified InstanceArray.
	InstanceArrayStopByLabel(instanceArrayLabel string) (*InstanceArray, error)
	//InstanceArrayStart starts a specified InstanceArray.
	InstanceArrayStart(instanceArrayID int) (*InstanceArray, error)
	//InstanceArrayStartByLabel starts a specified InstanceArray.
	InstanceArrayStartByLabel(instanceArrayLabel string) (*InstanceArray, error)
	//InstanceEdit edits an instance. Requires deploy
	InstanceEdit(instanceID int, instanceOperation InstanceOperation) (*Instance, error)
	//InstanceEditByLabel edits an instance. Requires deploy
	InstanceEditByLabel(instanceLabel string, instanceOperation InstanceOperation) (*Instance, error)
	//InstanceArrayInstances retrieves a list of all the Instance objects associated with a specified InstanceArray.
	InstanceArrayInstances(instanceArrayID int) (*map[string]Instance, error)
	//InstanceArrayInstancesByLabel retrieves a list of all the Instance objects associated with a specified InstanceArray.
	InstanceArrayInstancesByLabel(instanceArrayLabel string) (*map[string]Instance, error)
	//InstanceGet returns a specific instance by id
	InstanceGet(instanceID int) (*Instance, error)
	//InstanceGetByLabel returns a specific instance by id
	InstanceGetByLabel(instanceLabel string) (*Instance, error)
	//InstanceServerPowerSet reboots or powers on an instance
	InstanceServerPowerSet(instanceID int, operation string) error
	//InstanceServerPowerSetByLabel reboots or powers on an instance
	InstanceServerPowerSetByLabel(instanceLabel string, operation string) error
	//InstanceServerPowerGet returns the power status of an instance
	InstanceServerPowerGet(instanceID int) (*string, error)
	//InstanceServerPowerGetByLabel returns the power status of an instance
	InstanceServerPowerGetByLabel(instanceLabel string) (*string, error)
	//InstanceServerPowerGetBatch returns the power status of multiple instances
	InstanceServerPowerGetBatch(infrastructureID int, instanceIDs []int) (*map[string]string, error)
	//InstanceServerPowerGetBatchByLabel returns the power status of multiple instances
	InstanceServerPowerGetBatchByLabel(infrastructureLabel string, instanceIDs []int) (*map[string]string, error)
	//GetUserEmail returns the user configured for this connection
	GetUserEmail() string
	//GetEndpoint returns the endpoint configured for this connection
	GetEndpoint() string
	//GetUserID returns the ID of the user extracted from the API key
	GetUserID() int
	//NetworkGet retrieves a network object
	NetworkGet(networkID int) (*Network, error)
	//NetworkGetByLabel retrieves a network object
	NetworkGetByLabel(networkLabel string) (*Network, error)
	//Networks returns a list of all network objects of an infrastructure
	Networks(infrastructureID int) (*map[string]Network, error)
	//NetworksByLabel returns a list of all network objects of an infrastructure
	NetworksByLabel(infrastructureLabel string) (*map[string]Network, error)
	//NetworkCreate creates a network
	NetworkCreate(infrastructureID int, network Network) (*Network, error)
	//NetworkCreateByLabel creates a network
	NetworkCreateByLabel(infrastructureLabel string, network Network) (*Network, error)
	//NetworkEdit applies a change to an existing network
	NetworkEdit(networkID int, networkOperation NetworkOperation) (*Network, error)
	//NetworkEditByLabel applies a change to an existing network
	NetworkEditByLabel(networkLabel string, networkOperation NetworkOperation) (*Network, error)
	//NetworkDelete deletes a network.
	NetworkDelete(networkID int) error
	//NetworkDeleteByLabel deletes a network.
	NetworkDeleteByLabel(networkLabel string) error
	//NetworkJoin merges two specified Network objects.
	NetworkJoin(networkID int, networkToBeDeletedID int) error
	//NetworkJoinByLabel merges two specified Network objects.
	NetworkJoinByLabel(networkLabel string, networkToBeDeletedID int) error
	//NetworkProfiles returns a list of network profiles for the specified datacenter
	NetworkProfiles(datacenterName string) (*map[int]NetworkProfile, error)
	//NetworkProfileCreate creates a network profile.
	NetworkProfileCreate(datacenterName string, networkProfile NetworkProfile) (*NetworkProfile, error)
	NetworkProfileListByInstanceArray(instanceArrayID id) (*map[int]int, error)
	//NetworkProfileGet returns a NetworkProfile with specified id
	NetworkProfileGet(networkProfileID int) (*NetworkProfile, error)
	//NetworkProfileGetByLabel returns a NetworkProfile with specified id
	NetworkProfileGetByLabel(networkProfileLabel string) (*NetworkProfile, error)
	//NetworkProfileUpdate updates a network profile.
	NetworkProfileUpdate(networkProfileID int, networkProfile NetworkProfile) (*NetworkProfile, error)
	//NetworkProfileUpdateByLabel updates a network profile.
	NetworkProfileUpdateByLabel(networkProfileLabel string, networkProfile NetworkProfile) (*NetworkProfile, error)
	//NetworkProfileDelete deletes a network profile.
	NetworkProfileDelete(networkProfileID int) error
	//NetworkProfileDeleteByLabel deletes a network profile.
	NetworkProfileDeleteByLabel(networkProfileLabel string) error
	//OSAssetCreate creates a osAsset object
	OSAssetCreate(osAsset OSAsset) (*OSAsset, error)
	//OSAssetDelete permanently destroys a OSAsset.
	OSAssetDelete(osAssetID int) error
	//OSAssetUpdate updates a osAsset
	OSAssetUpdate(osAssetID int, osAsset OSAsset) (*OSAsset, error)
	//OSAssetGet returns a OSAsset specified by nOSAssetID. The OSAsset's protected value is never returned.
	OSAssetGet(osAssetID int) (*OSAsset, error)
	//OSAssets retrieves a list of all the OSAsset objects which a specified User is allowed to see through ownership or delegation. The OSAsset objects never return the actual protected OSAsset value.
	OSAssets() (*map[string]OSAsset, error)
	//OSAssetMakePublic makes an OS Asset public
	OSAssetMakePublic(osAssetID int) (*OSAsset, error)
	//OSAssetMakePrivate makes an OS Asset private and owned by the current user
	OSAssetMakePrivate(osAssetID int, userID int) (*OSAsset, error)
	//OSTemplateCreate creates a osTemplate object
	OSTemplateCreate(osTemplate OSTemplate) (*OSTemplate, error)
	//OSTemplateDelete permanently destroys a OSTemplate.
	OSTemplateDelete(osTemplateID int) error
	//OSTemplateUpdate updates a osTemplate
	OSTemplateUpdate(osTemplateID int, osTemplate OSTemplate) (*OSTemplate, error)
	//OSTemplateGet returns a OSTemplate specified by nOSTemplateID. The OSTemplate's protected value is never returned.
	OSTemplateGet(osTemplateID int, decryptPasswd bool) (*OSTemplate, error)
	//OSTemplates retrieves a list of all the OSTemplate objects which a specified User is allowed to see through ownership or delegation. The OSTemplate objects never return the actual protected OSTemplate value.
	OSTemplates() (*map[string]OSTemplate, error)
	//OSTemplateOSAssets returns the OSAssets assigned to an OSTemplate.
	OSTemplateOSAssets(osTemplateID int) (*map[string]OSTemplateOSAssetData, error)
	//OSTemplateAddOSAsset adds an asset to a template
	OSTemplateAddOSAsset(osTemplateID int, osAssetID int, path string, variablesJSON string) error
	//OSTemplateRemoveOSAsset removes an asset from a template
	OSTemplateRemoveOSAsset(osTemplateID int, osAssetID int) error
	//OSTemplateUpdateOSAssetPath updates an asset mapping
	OSTemplateUpdateOSAssetPath(osTemplateID int, osAssetID int, path string) error
	//OSTemplateUpdateOSAssetVariables updates an asset variable
	OSTemplateUpdateOSAssetVariables(osTemplateID int, osAssetID int, variablesJSON string) error
	//OSTemplateMakePublic makes a template public
	OSTemplateMakePublic(osTemplateID int) error
	//OSTemplateMakePrivate makes a template private
	OSTemplateMakePrivate(osTemplateID int, userID int) error
	//SecretCreate creates a secret
	SecretCreate(secret Secret) (*Secret, error)
	//SecretDelete Permanently destroys a Secret.
	SecretDelete(secretID int) error
	//SecretUpdate This function allows updating the secret_usage, secret_label and secret_base64 of a Secret
	SecretUpdate(secretID int, secret Secret) (*Secret, error)
	//SecretGet returns a Secret specified by nSecretID. The secret's protected value is never returned.
	SecretGet(secretID int) (*Secret, error)
	//Secrets retrieves a list of all the Secret objects which a specified User is allowed to see through ownership or delegation. The secret objects never return the actual protected secret value.
	Secrets(usage string) (*map[string]Secret, error)
	//ServersSearch searches for servers matching certain filter
	ServersSearch(filter string) (*[]ServerSearchResult, error)
	//ServerGetByUUID retrieves information about a specified Server by using the server's UUID
	ServerGetByUUID(serverUUID string, decryptPasswd bool) (*Server, error)
	//ServerGet returns a server's details
	ServerGet(serverID int, decryptPasswd bool) (*Server, error)
	//ServerCreate manually creates a server record
	ServerCreate(server Server, autoGenerate bool) (int, error)
	//ServerEditComplete - perform a complete edit
	ServerEditComplete(serverID int, server Server) (*Server, error)
	//ServerEditIPMI - edit only IPMI settings
	ServerEditIPMI(serverID int, server Server) (*Server, error)
	//ServerEditAvailability - edit only server availability settings
	ServerEditAvailability(serverID int, server Server) (*Server, error)
	//ServerEdit edits a server record
	ServerEdit(serverID int, serverEditType string, server Server) (*Server, error)
	//ServerDelete deletes all the information about a specified Server.
	ServerDelete(serverID int, skipIPMI bool) error
	//ServerDecomission decomissions the server row and deletes all child rows.
	ServerDecomission(serverID int, skipIPMI bool) error
	//ServerFirmwareComponentUpgrade Creates a firmware upgrading session for the specified component.
	//If no strServerComponentFirmwareNewVersion or strFirmwareBinaryURL are provided the system will use the values from the database which should have been previously added
	ServerFirmwareComponentUpgrade(serverID int, serverComponentID int, serverComponentFirmwareNewVersion string, firmwareBinaryURL string) error
	//ServerFirmwareUpgrade creates a firmware upgrading session that affects all components from the specified server that have a target version set and are updatable.
	ServerFirmwareUpgrade(serverID int) error
	//ServerFirmwareComponentTargetVersionSet Sets a firmware target version for the upgrading process. The system will apply the upgrade at the next upgrading session.
	ServerFirmwareComponentTargetVersionSet(serverComponentID int, serverComponentFirmwareNewVersion string) error
	//ServerFirmwareComponentTargetVersionUpdate Updates for every component of the specified server the available firmware versions that can be used as target by the firmware upgrading process. The available versions are extracted from a vendor specific catalog.
	ServerFirmwareComponentTargetVersionUpdate(serverComponentID int) error
	//ServerFirmwareComponentTargetVersionAdd Adds a new available firmware version for a server component along with the url of the binary. If the version already exists the old url will be overwritten.
	ServerFirmwareComponentTargetVersionAdd(serverComponentID int, version string, firmareBinaryURL string) error
	//ServerComponentGet returns a server's component's details
	ServerComponentGet(serverComponentID int) (*ServerComponent, error)
	//ServerComponents searches for servers matching certain filter
	ServerComponents(serverID int, filter string) (*[]ServerComponent, error)
	//ServerPowerSet reboots or powers on a server
	ServerPowerSet(serverID int, operation string) error
	//ServerFirmwarePolicyGet returns a server policy's details
	ServerFirmwarePolicyGet(serverFirmwarePolicyID int) (*ServerFirmwareUpgradePolicy, error)
	//ServerFirmwareUpgradePolicyCreate creates a server firmware policy.
	ServerFirmwareUpgradePolicyCreate(serverFirmwarePolicy *ServerFirmwareUpgradePolicy) (*ServerFirmwareUpgradePolicy, error)
	//ServerFirmwarePolicyAddRule add a new rule for a policy.
	ServerFirmwarePolicyAddRule(serverFirmwarePolicyID int, serverRule *ServerFirmwareUpgradePolicyRule) (*ServerFirmwareUpgradePolicy, error)
	//ServerFirmwarePolicyDeleteRule deletes a rule from a policy.
	ServerFirmwarePolicyDeleteRule(serverFirmwarePolicyID int, serverRule *ServerFirmwareUpgradePolicyRule) error
	//ServerFirmwareUpgradePolicyDelete deletes all the information about a specified ServerFirmwareUpgradePolicy.
	ServerFirmwareUpgradePolicyDelete(serverFirmwarePolicyID int) error
	ServerFirmwareUgradePolicyInstanceArraySet(serverFirmwarePolicyID int, instanceArrayList []int) error
	//ServerTypesMatchHardwareConfiguration Retrieves a list of server types that match the provided hardware configuration. The function does not check for availability, only compatibility, so physical servers associated with the returned server types might be unavailable.
	ServerTypesMatchHardwareConfiguration(datacenterName string, hardwareConfiguration HardwareConfiguration) (*map[int]ServerType, error)
	//ServerTypeDatacenter retrieves all the server type IDs for servers found in a specified Datacenter
	ServerTypeDatacenter(datacenterName string) (*[]int, error)
	//ServerTypes retrieves all ServerType objects from the database.
	ServerTypes(bOnlyAvailable bool) (*map[int]ServerType, error)
	//ServerTypesForDatacenter retrieves all ServerType objects from the database.
	ServerTypesForDatacenter(datacenterName string, bOnlyAvailable bool) (*map[int]ServerType, error)
	//ServerTypeGet retrieves a server type by id
	ServerTypeGet(serverTypeID int) (*ServerType, error)
	//ServerTypeGetByLabel retrieves a server type by id
	ServerTypeGetByLabel(serverTypeLabel string) (*ServerType, error)
	//ServerTypesMatches matches available servers with a certain Instance&#39;s configuration, using the properties specified in the objHardwareConfiguration object, and returns the number of compatible servers for each server_type_id.
	ServerTypesMatches(infrastructureID int, hardwareConfiguration HardwareConfiguration, instanceArrayID *int, bAllowServerSwap bool) (*map[string]ServerType, error)
	//ServerTypesMatchesByLabel matches available servers with a certain Instance&#39;s configuration, using the properties specified in the objHardwareConfiguration object, and returns the number of compatible servers for each server_type_id.
	ServerTypesMatchesByLabel(infrastructureLabel string, hardwareConfiguration HardwareConfiguration, instanceArrayID *int, bAllowServerSwap bool) (*map[string]ServerType, error)
	SharedDriveAttachInstanceArray(sharedDriveID int, instanceArrayID int) (*SharedDrive, error)
	SharedDriveDetachInstanceArray(sharedDriveID int, instanceArrayID int) (*SharedDrive, error)
	//SharedDrives retrieves the list of shared drives of an infrastructure
	SharedDrives(infrastructureID int) (*map[string]SharedDrive, error)
	//SharedDriveCreate creates a shared drive array. Requires deploy.
	SharedDriveCreate(infrastructureID int, sharedDrive SharedDrive) (*SharedDrive, error)
	//SharedDriveCreateByLabel creates a shared drive array. Requires deploy.
	SharedDriveCreateByLabel(infrastructureLabel string, sharedDrive SharedDrive) (*SharedDrive, error)
	//SharedDriveGet Retrieves a shared drive
	SharedDriveGet(sharedDriveID int) (*SharedDrive, error)
	//SharedDriveGetByLabel Retrieves a shared drive
	SharedDriveGetByLabel(sharedDriveLabel string) (*SharedDrive, error)
	//SharedDriveEdit alters a deployed drive array. Requires deploy.
	SharedDriveEdit(sharedDriveID int, sharedDriveOperation SharedDriveOperation) (*SharedDrive, error)
	//SharedDriveEditByLabel alters a deployed drive array. Requires deploy.
	SharedDriveEditByLabel(sharedDriveLabel string, sharedDriveOperation SharedDriveOperation) (*SharedDrive, error)
	//SharedDriveDelete deletes a shared drive.
	SharedDriveDelete(sharedDriveID int) error
	//SharedDriveDeleteByLabel deletes a shared drive.
	SharedDriveDeleteByLabel(sharedDriveLabel string) error
	//StageDefinitionCreate creates a stageDefinition
	StageDefinitionCreate(stageDefinition StageDefinition) (*StageDefinition, error)
	//StageDefinitionDelete Permanently destroys a StageDefinition.
	StageDefinitionDelete(stageDefinitionID int) error
	//StageDefinitionUpdate This function allows updating the stageDefinition_usage, stageDefinition_label and stageDefinition_base64 of a StageDefinition
	StageDefinitionUpdate(stageDefinitionID int, stageDefinition StageDefinition) (*StageDefinition, error)
	//StageDefinitionGet returns a StageDefinition specified by nStageDefinitionID. The stageDefinition's protected value is never returned.
	StageDefinitionGet(stageDefinitionID int) (*StageDefinition, error)
	//StageDefinitions retrieves a list of all the StageDefinition objects which a specified User is allowed to see through ownership or delegation. The stageDefinition objects never return the actual protected stageDefinition value.
	StageDefinitions() (*map[string]StageDefinition, error)
	//SubnetPoolCreate creates a new SubnetPool.
	SubnetPoolCreate(subnetPool SubnetPool) (*SubnetPool, error)
	//SubnetPoolGet retrieves information regarding a specified SubnetPool.
	SubnetPoolGet(subnetPoolID int) (*SubnetPool, error)
	//SubnetPoolPrefixSizesStats retrieves information regarding the utilization of a specified SubnetPool.
	SubnetPoolPrefixSizesStats(subnetPoolID int) (*SubnetPoolUtilization, error)
	//SubnetPoolDelete deletes the specified SubnetPool
	SubnetPoolDelete(subnetPoolID int) error
	//SubnetPools retrieves all switch devices registered in the database.
	SubnetPools() (*[]SubnetPool, error)
	//SubnetPoolSearch retrieves all switch devices registered in the database with the specified filter
	SubnetPoolSearch(filter string) (*[]SubnetPool, error)
	//SwitchDeviceGet Retrieves information regarding a specified SwitchDevice.
	SwitchDeviceGet(networkEquipmentID int, decryptPasswd bool) (*SwitchDevice, error)
	//SwitchDeviceGetByIdentifierString Retrieves information regarding a specified SwitchDevice by identifier string.
	SwitchDeviceGetByIdentifierString(networkEquipmentIdentifierString string, decryptPasswd bool) (*SwitchDevice, error)
	//SwitchDeviceCreate Creates a record for a new SwitchDevice.
	SwitchDeviceCreate(switchDevice SwitchDevice, bOverwriteWithHostnameFromFetchedSwitch bool) (*SwitchDevice, error)
	//SwitchDeviceDelete deletes a specified switch device and its registered interfaces.
	SwitchDeviceDelete(networkEquipmentID int) error
	//SwitchDevices retrieves all switch devices registered in the database.
	SwitchDevices(datacenter string, switchType string) (*map[string]SwitchDevice, error)
	//SwitchDevicesInDatacenter retrieves all switch devices in a datacenter
	SwitchDevicesInDatacenter(datacenter string) (*map[string]SwitchDevice, error)
	//SwitchDeviceUpdate updates an existing switch configuration
	SwitchDeviceUpdate(networkEquipmentID int, switchDevice SwitchDevice, bOverwriteWithHostnameFromFetchedSwitch bool) (*SwitchDevice, error)
	//UserGet describes returns user account specifications.
	UserGet(userID int) (*User, error)
	//UserGetByEmail describes returns user account specifications.
	UserGetByEmail(userLabel string) (*User, error)
	//UserEmailToUserID returns the user id of an user given an email
	UserEmailToUserID(userEmail string) (*int, error)
	//VariableCreate creates a variable object
	VariableCreate(variable Variable) (*Variable, error)
	//VariableDelete permanently destroys a Variable.
	VariableDelete(variableID int) error
	//VariableUpdate updates a variable
	VariableUpdate(variableID int, variable Variable) (*Variable, error)
	//VariableGet returns a Variable specified by nVariableID. The Variable's protected value is never returned.
	VariableGet(variableID int) (*Variable, error)
	//Variables retrieves a list of all the Variable objects which a specified User is allowed to see through ownership or delegation. The Variable objects never return the actual protected Variable value.
	Variables(usage string) (*map[string]Variable, error)
	//VolumeTemplates retrives the list of available templates
	VolumeTemplates() (*map[string]VolumeTemplate, error)
	//VolumeTemplateMakePublic makes a template public
	VolumeTemplateMakePublic(volumeTemplateID int) error
	//VolumeTemplateMakePrivate makes a template private
	VolumeTemplateMakePrivate(volumeTemplateID int, userID int) error
	//VolumeTemplateGet returns the specified volume template
	VolumeTemplateGet(volumeTemplateID int) (*VolumeTemplate, error)
	//VolumeTemplateGetByLabel returns the specified volume template
	VolumeTemplateGetByLabel(volumeTemplateLabel string) (*VolumeTemplate, error)
	//VolumeTemplateCreateFromDrive creates a private volume template from a drive
	VolumeTemplateCreateFromDrive(driveID int, objVolumeTemplate VolumeTemplate) (*VolumeTemplate, error)
	//VolumeTemplateCreateFromDriveByLabel creates a private volume template from a drive
	VolumeTemplateCreateFromDriveByLabel(driveLabel string, objVolumeTemplate VolumeTemplate) (*VolumeTemplate, error)
	//WorkflowCreate creates a workflow
	WorkflowCreate(workflow Workflow) (*Workflow, error)
	//WorkflowDelete Permanently destroys a Workflow.
	WorkflowDelete(workflowID int) error
	//WorkflowUpdate This function allows updating the workflow_usage, workflow_label and workflow_base64 of a Workflow
	WorkflowUpdate(workflowID int, workflow Workflow) (*Workflow, error)
	//WorkflowGet returns a Workflow specified by nWorkflowID. The workflow's protected value is never returned.
	WorkflowGet(workflowID int) (*Workflow, error)
	//Workflows retrieves a list of all the Workflow objects which a specified User is allowed to see through ownership or delegation.
	Workflows() (*map[string]Workflow, error)
	//WorkflowsWithUsage retrieves a list of all the Workflow objects which the current User is allowed to see through ownership or delegation with a specific usage.
	WorkflowsWithUsage(usage string) (*map[string]Workflow, error)
	//WorkflowStages retrieves a list of all the StageDefinitions objects in this workflow
	WorkflowStages(workflowID int) (*[]WorkflowStageDefinitionReference, error)
	//WorkflowStageGet returns a StageDefinition specified by workflowStageID.
	WorkflowStageGet(workflowStageID int) (*WorkflowStageDefinitionReference, error)
	//WorkflowStageAddAsNewRunLevel adds a new stage in this workflow
	WorkflowStageAddAsNewRunLevel(workflowID int, stageDefinitionID int, destinationRunLevel int) error
	//WorkflowStageAddIntoRunLevel adds a new stage in this workflow
	WorkflowStageAddIntoRunLevel(workflowID int, stageDefinitionID int, destinationRunLevel int) error
	//WorkflowMoveAsNewRunLevel moves a stage in this workflow from a runlevel to another
	WorkflowMoveAsNewRunLevel(workflowID int, stageDefinitionID int, sourceRunLevel int, destinationRunLevel int) error
	//WorkflowMoveIntoRunLevel moves a stage in this workflow from a runlevel to another
	WorkflowMoveIntoRunLevel(workflowID int, stageDefinitionID int, sourceRunLevel int, destinationRunLevel int) error
	//WorkflowStageDelete deletes a stage from a workflow entirelly
	WorkflowStageDelete(workflowStageID int) error
	//InfrastructureDeployCustomStageAddIntoRunlevel adds a stage into a runlevel
	InfrastructureDeployCustomStageAddIntoRunlevel(infraID int, stageID int, runLevel int, stageRunMoment string) error
	//InfrastructureDeployCustomStageDeleteIntoRunlevel delete a stage into a runlevel
	InfrastructureDeployCustomStageDeleteIntoRunlevel(infraID int, stageID int, runLevel int, stageRunMoment string) error
	//InfrastructureDeployCustomStages retrieves a list of all the StageDefinition objects which a specified User is allowed to see through ownership or delegation. The stageDefinition objects never return the actual protected stageDefinition value.
	InfrastructureDeployCustomStages(infraID int, stageDefinitionType string) (*[]WorkflowStageAssociation, error)
}
