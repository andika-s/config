# Config

A simple configuration management in Golang.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)

## Installation
Get the repository by following this command below:
```bash
go get github.com/andika-s/config
```

## Usage
Start to create new config initialization and put the configuration script using yaml extension and access the process function.
```bash
func main() {
	cfg := config.New("config.yaml")
	instance, err := cfg.Process()
	if err != nil {
		log.Fatal(err)
	}
}
```