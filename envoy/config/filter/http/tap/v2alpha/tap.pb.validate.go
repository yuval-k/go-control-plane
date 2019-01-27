// Code generated by protoc-gen-validate
// source: envoy/config/filter/http/tap/v2alpha/tap.proto
// DO NOT EDIT!!!

package envoy_config_filter_http_tap_v2alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on Tap with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Tap) Validate() error {
	if m == nil {
		return nil
	}

	switch m.ConfigType.(type) {

	case *Tap_AdminConfig:

		if v, ok := interface{}(m.GetAdminConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TapValidationError{
					Field:  "AdminConfig",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return TapValidationError{
			Field:  "ConfigType",
			Reason: "value is required",
		}

	}

	return nil
}

// TapValidationError is the validation error returned by Tap.Validate if the
// designated constraints aren't met.
type TapValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e TapValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTap.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = TapValidationError{}

// Validate checks the field values on AdminConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AdminConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetConfigId()) < 1 {
		return AdminConfigValidationError{
			Field:  "ConfigId",
			Reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// AdminConfigValidationError is the validation error returned by
// AdminConfig.Validate if the designated constraints aren't met.
type AdminConfigValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AdminConfigValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdminConfig.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AdminConfigValidationError{}
