// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// EnrollmentState undocumented
type EnrollmentState string

const (
	// EnrollmentStateVUnknown undocumented
	EnrollmentStateVUnknown EnrollmentState = "unknown"
	// EnrollmentStateVEnrolled undocumented
	EnrollmentStateVEnrolled EnrollmentState = "enrolled"
	// EnrollmentStateVPendingReset undocumented
	EnrollmentStateVPendingReset EnrollmentState = "pendingReset"
	// EnrollmentStateVFailed undocumented
	EnrollmentStateVFailed EnrollmentState = "failed"
	// EnrollmentStateVNotContacted undocumented
	EnrollmentStateVNotContacted EnrollmentState = "notContacted"
)

var (
	// EnrollmentStatePUnknown is a pointer to EnrollmentStateVUnknown
	EnrollmentStatePUnknown = &_EnrollmentStatePUnknown
	// EnrollmentStatePEnrolled is a pointer to EnrollmentStateVEnrolled
	EnrollmentStatePEnrolled = &_EnrollmentStatePEnrolled
	// EnrollmentStatePPendingReset is a pointer to EnrollmentStateVPendingReset
	EnrollmentStatePPendingReset = &_EnrollmentStatePPendingReset
	// EnrollmentStatePFailed is a pointer to EnrollmentStateVFailed
	EnrollmentStatePFailed = &_EnrollmentStatePFailed
	// EnrollmentStatePNotContacted is a pointer to EnrollmentStateVNotContacted
	EnrollmentStatePNotContacted = &_EnrollmentStatePNotContacted
)

var (
	_EnrollmentStatePUnknown      = EnrollmentStateVUnknown
	_EnrollmentStatePEnrolled     = EnrollmentStateVEnrolled
	_EnrollmentStatePPendingReset = EnrollmentStateVPendingReset
	_EnrollmentStatePFailed       = EnrollmentStateVFailed
	_EnrollmentStatePNotContacted = EnrollmentStateVNotContacted
)
