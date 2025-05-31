# Anvil

A lightweight, fast system configuration tool for automating server setup and management tasks on local and remote Linux systems.

## Why Anvil?

Anvil was created to address common pain points with existing configuration management tools:

- **Simplicity**: Unlike Ansible's complex YAML syntax and steep learning curve, Anvil uses simple, composable actions
- **Performance**: Built in Go for fast execution and minimal resource usage
- **Learning**: A hands-on project to understand configuration management principles without the complexity of enterprise tools
- **Transparency**: Clear command execution with dry-run support to see exactly what will happen
- **Flexibility**: Works locally or over SSH with the same interface

## Features

### Core Actions
- **User Management**: Create users with optional group assignment
- **Package Management**: Install packages with optional repository updates
- **OS Detection**: Automatic detection of Linux distribution (Debian/Ubuntu vs Fedora/RedHat families)
- **Dry Run Mode**: Preview all commands before execution

### Pre-built Recipes
Quick server configuration templates:
- **LAMP Server**: Apache, MySQL, PHP setup
- **Basic Web Server**: Essential web hosting components
- **Nginx Web Server**: Modern web server configuration

### Execution Modes
- **Local Execution**: Run commands on the current system
- **SSH Execution**: Execute commands on remote systems (planned)
- **Dry Run**: Preview mode that shows commands without executing them

## Usage

### Basic Commands

```bash
# Create a user
anvil create-user john

# Create a user and add to a group
anvil create-user --group sudo john

# Install a package
anvil install-package nginx

# Install a package with repository update
anvil install-package --update docker.io

# Deploy a complete server configuration
anvil recipe lamp

# List available recipes
anvil recipe --list

# Detect current OS
anvil detect-os

# Preview commands without executing (dry run)
anvil --dry-run recipe lamp
```

### Available Recipes

- `lamp` - Complete LAMP stack (Apache, MySQL, PHP)
- `webserver` - Basic web server setup
- `nginx` - Nginx web server configuration

## Architecture

Anvil follows clean architecture principles with well-defined interfaces:

- **Actions**: Composable units of work (create user, install package, etc.)
- **Executors**: Abstraction for command execution (local, SSH, dry-run)
- **OS Interface**: Cross-distribution compatibility layer
- **Recipes**: Collections of actions for common server configurations
- **Observers**: Event system for monitoring execution progress

## Building and Testing

```bash
# Build the project
go build ./...

# Run all tests
go test ./...

# Run tests for specific packages
go test ./internal/core/...
go test ./internal/actions/...

# Tidy dependencies
go mod tidy
```

## Contributing

This project follows Go best practices and maintains high test coverage. See `CLAUDE.md` for detailed development guidelines.
