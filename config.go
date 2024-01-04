package config

import (
	"os"
	"sync"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

// Package description: This Go package provides a simple configuration management system
// with a singleton pattern, utilizing YAML files for configuration data.

// Struct instance: The 'instance' struct represents a configuration instance,
type instance struct {
	// Put custom struct or elements with instance struct for specific configuration.
	// Should like this 'Service service `yaml:"service"`'.
}

// The example struct testing the service atributte
// type service struct {
// 	Name    string `yaml:"name" default:"default"`
// 	Host    string `yaml:"host" default:"0.0.0.0"`
// 	Port    string `yaml:"port" default:"3000"`
// 	Version string `yaml:"version" default:"1.0.0"`
// }

// Struct config: The 'config' struct is responsible for managing the configuration instance,
// handling its initialization, loading from a YAML file, and providing a singleton pattern.
type config struct {
	once     sync.Once // Ensures that the configuration is loaded only once
	instance *instance // The singleton instance of the configuration
	path     string    // The file path for loading the configuration from a YAML file

}

// Function New: The 'New' function creates and returns a new configuration instance.
// It takes a 'path' parameter, which represents the file path for loading the configuration.
func New(path string) *config {
	return &config{path: path, instance: &instance{}}
}

// Method Process: The 'Process' method ensures that the configuration is loaded and processed.
// It uses the 'sync.Once' construct to guarantee that the loading operation is performed only once.
// If loading fails, default values are set using the 'defaults.Set' function.
func (c *config) Process() (*instance, error) {
	var err error
	c.once.Do(func() {
		err = c.load()
		if err != nil {
			defaults.Set(c.instance)
		}
	})
	return c.instance, err
}

// Method load: The 'load' method reads the configuration from a YAML file specified by the 'path'.
// If the file doesn't exist, it returns an error. The loaded configuration is stored in the 'instance'
// field of the 'config' struct, and a message is printed indicating a successful load.
func (c *config) load() error {
	_, err := os.Stat(c.path)
	if err != nil && os.IsNotExist(err) {
		return err
	}
	file, err := os.ReadFile(c.path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, &c.instance)
	if err != nil {
		return err
	}
	return nil
}
