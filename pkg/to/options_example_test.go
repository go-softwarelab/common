package to_test

import (
	"fmt"

	"github.com/go-softwarelab/common/pkg/to"
)

func ExampleOptions() {
	// Prepare definition of config
	type Config struct {
		Host  string
		Port  int
		Debug bool
	}

	// Define reusable option functions
	withHost := func(host string) func(*Config) {
		return func(c *Config) {
			c.Host = host
		}
	}

	withPort := func(port int) func(*Config) {
		return func(c *Config) {
			c.Port = port
		}
	}

	withDebug := func(c *Config) {
		c.Debug = true
	}

	// Handle opts
	config := to.Options[Config](
		withHost("example.com"),
		withPort(443),
		withDebug,
	)

	fmt.Printf("Host: %s\n", config.Host)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug: %v\n", config.Debug)

	// Output:
	// Host: example.com
	// Port: 443
	// Debug: true
}

func ExampleOptionsWithDefault() {
	type Config struct {
		Host  string
		Port  int
		Debug bool
	}

	// Default configuration
	defaultConfig := Config{
		Host:  "127.0.0.1",
		Port:  3000,
		Debug: false,
	}

	// Override some default values using OptionsWithDefault
	config := to.OptionsWithDefault(defaultConfig,
		func(c *Config) {
			c.Port = 8080
		},
		func(c *Config) {
			c.Debug = true
		},
	)

	fmt.Printf("Host: %s\n", config.Host)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug: %v\n", config.Debug)

	// Output:
	// Host: 127.0.0.1
	// Port: 8080
	// Debug: true
}
