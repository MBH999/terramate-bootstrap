package types

type Terramate struct {
	Backend BackendConfig `yaml:"backend"`
	Stacks  StacksConfig  `yaml:"stacks"`
}

type BackendConfig struct {
	AzureRM AzureRMBackend `yaml:"azurerm,omitempty"`
	// AWS     *AWSBackend     `yaml:"aws,omitempty"`
}

type AzureRMBackend struct {
	ResourceGroupName  string `yaml:"resource_group_name"`
	StorageAccountName string `yaml:"storage_account_name"`
	Location           string `yaml:"location"`
}

//
// type AWSBackend struct {
// 	BucketName    string `yaml:"bucket_name"`
// 	Region        string `yaml:"region"`
// 	DynamoDBTable string `yaml:"dynamo_db_table"`
// }

type StacksConfig struct {
	Environments  map[string]Environment  `yaml:"environments"`
	Regions       map[string]Region       `yaml:"regions"`
	ResourceTypes map[string]ResourceType `yaml:"resource_types,omitempty"`
}

type Environment struct {
	Tags []string `yaml:"tags,omitempty"`
}

type Region struct {
	Tags []string `yaml:"tags,omitempty"`
}

type ResourceType struct {
	Tags                []string `yaml:"tags,omitempty"`
	ExcludeEnvironments []string `yaml:"exclude_environments,omitempty"`
}
