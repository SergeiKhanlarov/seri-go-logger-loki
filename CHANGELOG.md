# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial project structure
- Loki client implementation
- Logger provider for seri-go-logger integration
- Basic configuration types
- Documentation and examples

### Features
- Support for all log levels (Debug, Info, Warn, Error, Fatal)
- Asynchronous log sending
- Custom HTTP client configuration
- Field merging functionality

### Technical Details
- Interface-based architecture
- Proper error handling and wrapping
- Comprehensive test coverage
- Go module support

## [v0.1.0] - 2025-11-30

### Added
- Initial release
- Core Loki client functionality
- Integration with seri-go-logger
- Basic documentation
- Example usage

### Dependencies
- Requires Go 1.19+
- Compatible with seri-go-logger v1.x
- HTTP client for Loki API communication