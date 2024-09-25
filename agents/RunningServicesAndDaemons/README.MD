# Services and Daemons Scrapper

This program fetches service statuses, configurations, and startup settings.

## Features

- Lists currently active services on the system
- Retrieves startup settings for all services
- Provides detailed information about a specific service
- Attempts to locate and display the configuration file for a given service

## Prerequisites

- Appropriate permissions to run system commands and access configuration files (root access may be required for some operations)

## Usage


```
sudo go run main.go [flags]
```

- `-service <name>`: Specifies a service name to query for detailed information
- `-show-config`:  Displays config file. ( `-service` is required to specify the service for which the config is to be shown.)

### Examples

1. Display information about all services (active services and startup settings:
   ```
   sudo go run main.go
   ```

2. Get information about a specific service:
   ```
   sudo go run main.go -service crond
   ```

3. Show the configuration of a specific service:
   ```
   sudo go run main.go -service crond -show-config
   ```


## Note

Root access (sudo) is required to access most daemon configuration files and to fetch certain service information. 
