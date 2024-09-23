# Services and Dameons Scraper

This Go program retrieves and displays active services, startup settings, and daemon-specific configurations.

## Features

- Lists currently active services on the system.
- Retrieves startup settings for all services.
- Reads and displays configuration settings. Default is the SSH daemon (`sshd_config`).

## Prerequisites

- Go 1.16 or later installed on your system.
- Appropriate permissions to run system commands and access configuration files.

## Usage
To run on terminal: 
sudo go run main.go  (Root access required to access the daemon config files)

use -config flag to give the path for the daemon config file (By Default SSH Dameon Config File)
sudo go run main.go -config enter-path-here

Note: Its set to not read the comments and empty lines of the config file, can be changed in code easily. 