# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

### Build
```bash
go build ./...
```

### Run Tests
```bash
# Run all tests
go test ./...

# Run tests in a specific package
go test ./internal/core/...
go test ./internal/actions/...

# Run a specific test
go test -run TestName ./internal/core
```

### Module Operations
```bash
# Download dependencies
go mod download

# Tidy dependencies
go mod tidy
```

## Architecture

Anvil is a Go-based system configuration tool that executes actions on local or remote systems. The codebase follows a clean architecture pattern with well-defined interfaces.

### Core Components

1. **Action Interface** (`internal/core/action.go`): Defines the contract for all system actions. Actions implement `Handle()` method that takes an Executor, OS, and optional ActionObserver.

2. **Executor Interface** (`internal/core/executor.go`): Abstracts command execution with implementations for:
   - `LocalExecutor`: Executes commands locally
   - `SshExecutor`: Executes commands over SSH
   - `FakeExecutor`: Test double that records command history and returns predefined responses

3. **OS Interface** (`internal/core/os.go`): Abstracts OS-specific commands for different Linux distributions:
   - `DebianFamily` (Ubuntu, Debian): Uses `adduser` for user creation
   - `FedoraFamily` (Fedora, RedHat): Uses `useradd` for user creation
   - Provides methods: `CreateUser()`, `CheckUser()`, `GroupUser()`

4. **OS Detection** (`internal/core/osdetect.go`): Automatically detects the target OS by parsing `/etc/os-release` and returns the appropriate OS implementation.

5. **Observer Pattern**: The codebase uses observers for monitoring execution:
   - `ExecutionObserver`: Interface for monitoring command execution (methods: OnExecutionStart, OnExecutionEnd, OnExecutionOutput)
   - `ActionObserver`: Extends ExecutionObserver with action-level monitoring (OnActionStart, OnActionEnd)

### Action Implementation Pattern

Actions follow a consistent pattern as seen in `internal/actions/create_user.go`:
- Use options pattern for configuration (e.g., `CreateUserOpts`, `CreateUserOptsFunc`)
- Check preconditions before executing (e.g., check if user exists before creating)
- Support optional features through options (e.g., `WithGroup()`)
- Notify observers at appropriate points

### Testing Strategy

Tests use the `FakeExecutor` to verify:
- Correct commands are executed in the right order
- Actions handle different scenarios (user exists vs doesn't exist)
- Optional features work correctly (e.g., adding user to group)