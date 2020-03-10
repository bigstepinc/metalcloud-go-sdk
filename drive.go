package metalcloud

//Drive represents a drive
type Drive struct {
	DriveID               int               `json:"drive_id,omitempty"`
	DriveLabel            string            `json:"drive_label,omitempty"`
	DriveSubdomain        string            `json:"drive_subdomain,omitempty"`
	DriveArrayID          int               `json:"drive_array_id,omitempty"`
	InstanceID            int               `json:"instance_id,omitempty"`
	DriveSizeMBytes       int               `json:"drive_size_mbytes,omitempty"`
	DriveStorageType      string            `json:"drive_storage_type,omitempty"`
	InfrastructureID      int               `json:"infrastructure_id,omitempty"`
	TemplateIDOrigin      int               `json:"template_id_origin,omitempty"`
	DriveCredentials      *DriveCredentials `json:"drive_credentials,omitempty"`
	DriveServiceStatus    string            `json:"drive_service_status,omitempty"`
	DriveCreatedTimestamp string            `json:"drive_created_timestamp,omitempty"`
	DriveUpdatedTimestamp string            `json:"drive_updated_timestamp,omitempty"`
	DriveOperatingSystem  *OperatingSystem  `json:"drive_operating_system,omitempty"`
	DriveFileSystem       *DriveFileSystem  `json:"drive_filesystem,omitempty"`
}

//DriveCredentials credentials to connect to the drive
type DriveCredentials struct {
	ISCSI ISCSI `json:"iscsi,omitempty"`
}

//DriveFileSystem filesystem details
type DriveFileSystem struct {
	DriveFilesystemType           string `json:"drive_filesystem_type,omitempty"`
	DriveFilesystemBlockSizeBytes int    `json:"drive_filesystem_block_size_bytes,omitempty"`
	DriveFilesystemMountPath      string `json:"drive_filesystem_mount_path,omitempty"`
}
