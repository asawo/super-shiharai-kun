package errors

import (
	"github.com/morikuni/failure"
)

// Transient errors
var (
	TemporaryUnavailable failure.StringCode = "TemporaryUnavailable"
	Canceled             failure.StringCode = "Canceled"
	Timeout              failure.StringCode = "Timeout"
)

// User errors
var (
	InvalidRequest   failure.StringCode = "InvalidRequest"
	NotFound         failure.StringCode = "NotFound"
	Unauthorized     failure.StringCode = "Unauthorized"
	PermissionDenied failure.StringCode = "PermissionDenied"
	MethodNotAllowed failure.StringCode = "MethodNotAllowed"
)

// Internal errors
var (
	Internal      failure.StringCode = "Internal"
	Unknown       failure.StringCode = "Unknown"
	AlreadyExists failure.StringCode = "AlreadyExists"
)

// Service specific errors
var (
	EventNotPublished failure.StringCode = "EventNotPublished"
)

var validErrors = map[failure.Code]bool{
	TemporaryUnavailable: true,
	Canceled:             true,
	Timeout:              true,
	InvalidRequest:       true,
	NotFound:             true,
	Unauthorized:         true,
	PermissionDenied:     true,
	MethodNotAllowed:     true,
	Internal:             true,
	Unknown:              true,
	AlreadyExists:        true,
	EventNotPublished:    true,
}
