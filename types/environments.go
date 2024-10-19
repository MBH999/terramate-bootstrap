package types

type Config struct {
	Config Configuration `yaml:"config"`
}

type Configuration struct {
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
