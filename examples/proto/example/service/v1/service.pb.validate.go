// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: example/service/v1/service.proto

package exampleservicev1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// define the regex for a UUID once up-front
var _service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on SendMessageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SendMessageRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendMessageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendMessageRequestMultiError, or nil if none found.
func (m *SendMessageRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SendMessageRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetMessage()) < 1 {
		err := SendMessageRequestValidationError{
			field:  "Message",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateUuid(m.GetRecipientId()); err != nil {
		err = SendMessageRequestValidationError{
			field:  "RecipientId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SendMessageRequestMultiError(errors)
	}

	return nil
}

func (m *SendMessageRequest) _validateUuid(uuid string) error {
	if matched := _service_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// SendMessageRequestMultiError is an error wrapping multiple validation errors
// returned by SendMessageRequest.ValidateAll() if the designated constraints
// aren't met.
type SendMessageRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendMessageRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendMessageRequestMultiError) AllErrors() []error { return m }

// SendMessageRequestValidationError is the validation error returned by
// SendMessageRequest.Validate if the designated constraints aren't met.
type SendMessageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendMessageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendMessageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendMessageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendMessageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendMessageRequestValidationError) ErrorName() string {
	return "SendMessageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SendMessageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendMessageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendMessageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendMessageRequestValidationError{}

// Validate checks the field values on SendMessageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SendMessageResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendMessageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendMessageResponseMultiError, or nil if none found.
func (m *SendMessageResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SendMessageResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = SendMessageResponseValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SendMessageResponseMultiError(errors)
	}

	return nil
}

func (m *SendMessageResponse) _validateUuid(uuid string) error {
	if matched := _service_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// SendMessageResponseMultiError is an error wrapping multiple validation
// errors returned by SendMessageResponse.ValidateAll() if the designated
// constraints aren't met.
type SendMessageResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendMessageResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendMessageResponseMultiError) AllErrors() []error { return m }

// SendMessageResponseValidationError is the validation error returned by
// SendMessageResponse.Validate if the designated constraints aren't met.
type SendMessageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendMessageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendMessageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendMessageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendMessageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendMessageResponseValidationError) ErrorName() string {
	return "SendMessageResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SendMessageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendMessageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendMessageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendMessageResponseValidationError{}

// Validate checks the field values on GetMessageRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetMessageRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMessageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetMessageRequestMultiError, or nil if none found.
func (m *GetMessageRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMessageRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = GetMessageRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetMessageRequestMultiError(errors)
	}

	return nil
}

func (m *GetMessageRequest) _validateUuid(uuid string) error {
	if matched := _service_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetMessageRequestMultiError is an error wrapping multiple validation errors
// returned by GetMessageRequest.ValidateAll() if the designated constraints
// aren't met.
type GetMessageRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMessageRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMessageRequestMultiError) AllErrors() []error { return m }

// GetMessageRequestValidationError is the validation error returned by
// GetMessageRequest.Validate if the designated constraints aren't met.
type GetMessageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMessageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMessageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMessageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMessageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMessageRequestValidationError) ErrorName() string {
	return "GetMessageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetMessageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMessageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMessageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMessageRequestValidationError{}

// Validate checks the field values on GetMessageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetMessageResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMessageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetMessageResponseMultiError, or nil if none found.
func (m *GetMessageResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMessageResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = GetMessageResponseValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetMessage()) < 1 {
		err := GetMessageResponseValidationError{
			field:  "Message",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetMessageResponseMultiError(errors)
	}

	return nil
}

func (m *GetMessageResponse) _validateUuid(uuid string) error {
	if matched := _service_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetMessageResponseMultiError is an error wrapping multiple validation errors
// returned by GetMessageResponse.ValidateAll() if the designated constraints
// aren't met.
type GetMessageResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMessageResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMessageResponseMultiError) AllErrors() []error { return m }

// GetMessageResponseValidationError is the validation error returned by
// GetMessageResponse.Validate if the designated constraints aren't met.
type GetMessageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMessageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMessageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMessageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMessageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMessageResponseValidationError) ErrorName() string {
	return "GetMessageResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetMessageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMessageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMessageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMessageResponseValidationError{}
