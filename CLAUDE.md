# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go library that provides a comprehensive client for the Backlog REST API. Backlog is a project management and bug tracking tool. The library is structured as a single-package Go module with separate files for different API resource types.

## Development Commands

All development tasks are managed through the Makefile:

### Core Development Commands
- `make test` - Run unit tests (short tests with 300s timeout)
- `make test-race` - Run race condition tests 
- `make test-integration` - Run full integration tests (600s timeout)
- `make lint` - Run golangci-lint (installs linter if needed)
- `make fmt` - Format all Go code
- `make cov` - Generate test coverage report and open in browser
- `make pr-prep` - Complete pre-PR workflow (fmt + lint + test-race + test-integration)

### Other Commands
- `make upgrade` - Update all dependencies
- `make help` - Show all available make targets

## Architecture

### Core Client Structure
- **`backlog.go`** - Main client implementation with HTTP request handling, authentication, and core utilities
- **`logger.go`** - Internal logging interfaces and implementations
- **`rate_limit.go`** - Rate limiting functionality for API calls

### API Resource Files
Each Backlog API resource type has its own dedicated file:
- `activity.go` - Activity/timeline operations  
- `category.go` - Issue categories
- `customfield.go` - Custom field management
- `file.go` - File attachments
- `git.go` - Git repository management
- `issue.go` - Issues (tickets/bugs)
- `issue_type.go` - Issue type definitions
- `misc.go` - Miscellaneous API endpoints
- `priority.go` - Issue priorities
- `project.go` - Project management
- `pullrequest.go` - Pull request operations
- `resolution.go` - Issue resolutions
- `space.go` - Backlog space operations
- `team.go` - Team management
- `time.go` - Time tracking utilities
- `user.go` - User management
- `version.go` - Project versions/milestones
- `watching.go` - Watch/notification settings
- `webhook.go` - Webhook management
- `wiki.go` - Wiki operations

### Key Design Patterns

1. **Client Options Pattern**: The main `Client` struct accepts functional options for configuration (HTTP client, debug mode, custom logger)

2. **Context Support**: All API methods have `Context` variants (e.g., `GetUser()` and `GetUserContext(ctx)`) for timeout/cancellation control

3. **Pointer Helpers**: Utility functions `Bool()`, `Int()`, `String()`, `Int64()` create pointers to primitive values for optional API fields

4. **Rate Limiting**: Built-in rate limit tracking and handling

5. **File Upload Support**: Multipart file upload functionality via `UploadMultipartFile()`

## Testing

- Each `.go` file has a corresponding `_test.go` file
- Uses `testify` for assertions (`github.com/stretchr/testify`)
- Test data stored in `testdata/` directory
- Supports both unit and integration testing modes
- Race condition testing available

## Dependencies

Core dependencies from `go.mod`:
- `github.com/google/go-querystring` - URL query parameter encoding
- `github.com/pkg/errors` - Enhanced error handling
- `github.com/stretchr/testify` - Testing framework
- `github.com/kylelemons/godebug` - Debugging utilities

## Usage Pattern

Typical client instantiation and usage:
```go
client := backlog.New("API_KEY", "BASE_URL", options...)
result, err := client.SomeAPIMethod()
// or with context
result, err := client.SomeAPIMethodContext(ctx)
```