# Services and Dameons Scraper

This Go program retrieves and displays active services, startup settings, and daemon-specific configurations.

## Features

- Lists currently active services on the system.
- Retrieves startup settings for all services (enabled, disabled).
- Reads and displays configuration settings. On the program the example is from the SSH daemon (`sshd_config`).

## Prerequisites

- Go 1.16 or later installed on your system.
- Appropriate permissions to run system commands and access configuration files.

## Usage
To run on terminal: sudo go run main.go  (Root access required to access the daemon config files)
