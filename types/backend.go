package types

// type Backend struct {
// 	AzureRM *AzureRMBackend `yaml:"azurerm,omitempty"`
// 	AWS     *AWSBackend     `yaml:"aws,omitempty"`
// }
//
// type AzureRMBackend struct {
// 	ResourceGroupName  string `yaml:"resource_group_name"`
// 	StorageAccountName string `yaml:"storage_account_name"`
// 	ContainerName      string `yaml:"container_name"`
// }
//
// type AWSBackend struct {
// 	BucketName    string `yaml:"bucket_name"`
// 	Region        string `yaml:"region"`
// 	DynamoDBTable string `yaml:"dynamo_db_table"`
// }
//
// Validate checks if either AzureRM or AWS is present and validates the required fields
// func (b *Backend) Validate() error {
// 	if b.AzureRM != nil && b.AWS == nil {
// 		if b.AzureRM.ResourceGroupName == "" || b.AzureRM.StorageAccountName == "" || b.AzureRM.ContainerName == "" {
// 			return errors.New("all fields in AzureRMBackend are required when using 'azurerm' backend")
// 		}
// 	} else if b.AWS != nil && b.AzureRM == nil {
// 		if b.AWS.BucketName == "" || b.AWS.Region == "" || b.AWS.DynamoDBTable == "" {
// 			return errors.New("all fields in AWSBackend are required when using 'aws' backend")
// 		}
// 	} else {
// 		return errors.New("exactly one of 'azurerm' or 'aws' backend configuration must be provided")
// 	}
// 	return nil
// }
