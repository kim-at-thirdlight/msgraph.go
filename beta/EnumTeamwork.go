// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TeamworkActivityTopicSource undocumented
type TeamworkActivityTopicSource string

const (
	// TeamworkActivityTopicSourceVEntityURL undocumented
	TeamworkActivityTopicSourceVEntityURL TeamworkActivityTopicSource = "entityUrl"
	// TeamworkActivityTopicSourceVText undocumented
	TeamworkActivityTopicSourceVText TeamworkActivityTopicSource = "text"
)

var (
	// TeamworkActivityTopicSourcePEntityURL is a pointer to TeamworkActivityTopicSourceVEntityURL
	TeamworkActivityTopicSourcePEntityURL = &_TeamworkActivityTopicSourcePEntityURL
	// TeamworkActivityTopicSourcePText is a pointer to TeamworkActivityTopicSourceVText
	TeamworkActivityTopicSourcePText = &_TeamworkActivityTopicSourcePText
)

var (
	_TeamworkActivityTopicSourcePEntityURL = TeamworkActivityTopicSourceVEntityURL
	_TeamworkActivityTopicSourcePText      = TeamworkActivityTopicSourceVText
)

// TeamworkApplicationIdentityType undocumented
type TeamworkApplicationIdentityType string

const (
	// TeamworkApplicationIdentityTypeVAadApplication undocumented
	TeamworkApplicationIdentityTypeVAadApplication TeamworkApplicationIdentityType = "aadApplication"
	// TeamworkApplicationIdentityTypeVBot undocumented
	TeamworkApplicationIdentityTypeVBot TeamworkApplicationIdentityType = "bot"
	// TeamworkApplicationIdentityTypeVTenantBot undocumented
	TeamworkApplicationIdentityTypeVTenantBot TeamworkApplicationIdentityType = "tenantBot"
	// TeamworkApplicationIdentityTypeVOffice365Connector undocumented
	TeamworkApplicationIdentityTypeVOffice365Connector TeamworkApplicationIdentityType = "office365Connector"
	// TeamworkApplicationIdentityTypeVOutgoingWebhook undocumented
	TeamworkApplicationIdentityTypeVOutgoingWebhook TeamworkApplicationIdentityType = "outgoingWebhook"
	// TeamworkApplicationIdentityTypeVUnknownFutureValue undocumented
	TeamworkApplicationIdentityTypeVUnknownFutureValue TeamworkApplicationIdentityType = "unknownFutureValue"
)

var (
	// TeamworkApplicationIdentityTypePAadApplication is a pointer to TeamworkApplicationIdentityTypeVAadApplication
	TeamworkApplicationIdentityTypePAadApplication = &_TeamworkApplicationIdentityTypePAadApplication
	// TeamworkApplicationIdentityTypePBot is a pointer to TeamworkApplicationIdentityTypeVBot
	TeamworkApplicationIdentityTypePBot = &_TeamworkApplicationIdentityTypePBot
	// TeamworkApplicationIdentityTypePTenantBot is a pointer to TeamworkApplicationIdentityTypeVTenantBot
	TeamworkApplicationIdentityTypePTenantBot = &_TeamworkApplicationIdentityTypePTenantBot
	// TeamworkApplicationIdentityTypePOffice365Connector is a pointer to TeamworkApplicationIdentityTypeVOffice365Connector
	TeamworkApplicationIdentityTypePOffice365Connector = &_TeamworkApplicationIdentityTypePOffice365Connector
	// TeamworkApplicationIdentityTypePOutgoingWebhook is a pointer to TeamworkApplicationIdentityTypeVOutgoingWebhook
	TeamworkApplicationIdentityTypePOutgoingWebhook = &_TeamworkApplicationIdentityTypePOutgoingWebhook
	// TeamworkApplicationIdentityTypePUnknownFutureValue is a pointer to TeamworkApplicationIdentityTypeVUnknownFutureValue
	TeamworkApplicationIdentityTypePUnknownFutureValue = &_TeamworkApplicationIdentityTypePUnknownFutureValue
)

var (
	_TeamworkApplicationIdentityTypePAadApplication     = TeamworkApplicationIdentityTypeVAadApplication
	_TeamworkApplicationIdentityTypePBot                = TeamworkApplicationIdentityTypeVBot
	_TeamworkApplicationIdentityTypePTenantBot          = TeamworkApplicationIdentityTypeVTenantBot
	_TeamworkApplicationIdentityTypePOffice365Connector = TeamworkApplicationIdentityTypeVOffice365Connector
	_TeamworkApplicationIdentityTypePOutgoingWebhook    = TeamworkApplicationIdentityTypeVOutgoingWebhook
	_TeamworkApplicationIdentityTypePUnknownFutureValue = TeamworkApplicationIdentityTypeVUnknownFutureValue
)

// TeamworkConversationIdentityType undocumented
type TeamworkConversationIdentityType string

const (
	// TeamworkConversationIdentityTypeVTeam undocumented
	TeamworkConversationIdentityTypeVTeam TeamworkConversationIdentityType = "team"
	// TeamworkConversationIdentityTypeVChannel undocumented
	TeamworkConversationIdentityTypeVChannel TeamworkConversationIdentityType = "channel"
	// TeamworkConversationIdentityTypeVChat undocumented
	TeamworkConversationIdentityTypeVChat TeamworkConversationIdentityType = "chat"
	// TeamworkConversationIdentityTypeVUnknownFutureValue undocumented
	TeamworkConversationIdentityTypeVUnknownFutureValue TeamworkConversationIdentityType = "unknownFutureValue"
)

var (
	// TeamworkConversationIdentityTypePTeam is a pointer to TeamworkConversationIdentityTypeVTeam
	TeamworkConversationIdentityTypePTeam = &_TeamworkConversationIdentityTypePTeam
	// TeamworkConversationIdentityTypePChannel is a pointer to TeamworkConversationIdentityTypeVChannel
	TeamworkConversationIdentityTypePChannel = &_TeamworkConversationIdentityTypePChannel
	// TeamworkConversationIdentityTypePChat is a pointer to TeamworkConversationIdentityTypeVChat
	TeamworkConversationIdentityTypePChat = &_TeamworkConversationIdentityTypePChat
	// TeamworkConversationIdentityTypePUnknownFutureValue is a pointer to TeamworkConversationIdentityTypeVUnknownFutureValue
	TeamworkConversationIdentityTypePUnknownFutureValue = &_TeamworkConversationIdentityTypePUnknownFutureValue
)

var (
	_TeamworkConversationIdentityTypePTeam               = TeamworkConversationIdentityTypeVTeam
	_TeamworkConversationIdentityTypePChannel            = TeamworkConversationIdentityTypeVChannel
	_TeamworkConversationIdentityTypePChat               = TeamworkConversationIdentityTypeVChat
	_TeamworkConversationIdentityTypePUnknownFutureValue = TeamworkConversationIdentityTypeVUnknownFutureValue
)

// TeamworkTagType undocumented
type TeamworkTagType string

const (
	// TeamworkTagTypeVStandard undocumented
	TeamworkTagTypeVStandard TeamworkTagType = "standard"
)

var (
	// TeamworkTagTypePStandard is a pointer to TeamworkTagTypeVStandard
	TeamworkTagTypePStandard = &_TeamworkTagTypePStandard
)

var (
	_TeamworkTagTypePStandard = TeamworkTagTypeVStandard
)

// TeamworkUserIdentityType undocumented
type TeamworkUserIdentityType string

const (
	// TeamworkUserIdentityTypeVAadUser undocumented
	TeamworkUserIdentityTypeVAadUser TeamworkUserIdentityType = "aadUser"
	// TeamworkUserIdentityTypeVOnPremiseAadUser undocumented
	TeamworkUserIdentityTypeVOnPremiseAadUser TeamworkUserIdentityType = "onPremiseAadUser"
	// TeamworkUserIdentityTypeVAnonymousGuest undocumented
	TeamworkUserIdentityTypeVAnonymousGuest TeamworkUserIdentityType = "anonymousGuest"
	// TeamworkUserIdentityTypeVFederatedUser undocumented
	TeamworkUserIdentityTypeVFederatedUser TeamworkUserIdentityType = "federatedUser"
	// TeamworkUserIdentityTypeVPersonalMicrosoftAccountUser undocumented
	TeamworkUserIdentityTypeVPersonalMicrosoftAccountUser TeamworkUserIdentityType = "personalMicrosoftAccountUser"
	// TeamworkUserIdentityTypeVSkypeUser undocumented
	TeamworkUserIdentityTypeVSkypeUser TeamworkUserIdentityType = "skypeUser"
	// TeamworkUserIdentityTypeVPhoneUser undocumented
	TeamworkUserIdentityTypeVPhoneUser TeamworkUserIdentityType = "phoneUser"
	// TeamworkUserIdentityTypeVUnknownFutureValue undocumented
	TeamworkUserIdentityTypeVUnknownFutureValue TeamworkUserIdentityType = "unknownFutureValue"
)

var (
	// TeamworkUserIdentityTypePAadUser is a pointer to TeamworkUserIdentityTypeVAadUser
	TeamworkUserIdentityTypePAadUser = &_TeamworkUserIdentityTypePAadUser
	// TeamworkUserIdentityTypePOnPremiseAadUser is a pointer to TeamworkUserIdentityTypeVOnPremiseAadUser
	TeamworkUserIdentityTypePOnPremiseAadUser = &_TeamworkUserIdentityTypePOnPremiseAadUser
	// TeamworkUserIdentityTypePAnonymousGuest is a pointer to TeamworkUserIdentityTypeVAnonymousGuest
	TeamworkUserIdentityTypePAnonymousGuest = &_TeamworkUserIdentityTypePAnonymousGuest
	// TeamworkUserIdentityTypePFederatedUser is a pointer to TeamworkUserIdentityTypeVFederatedUser
	TeamworkUserIdentityTypePFederatedUser = &_TeamworkUserIdentityTypePFederatedUser
	// TeamworkUserIdentityTypePPersonalMicrosoftAccountUser is a pointer to TeamworkUserIdentityTypeVPersonalMicrosoftAccountUser
	TeamworkUserIdentityTypePPersonalMicrosoftAccountUser = &_TeamworkUserIdentityTypePPersonalMicrosoftAccountUser
	// TeamworkUserIdentityTypePSkypeUser is a pointer to TeamworkUserIdentityTypeVSkypeUser
	TeamworkUserIdentityTypePSkypeUser = &_TeamworkUserIdentityTypePSkypeUser
	// TeamworkUserIdentityTypePPhoneUser is a pointer to TeamworkUserIdentityTypeVPhoneUser
	TeamworkUserIdentityTypePPhoneUser = &_TeamworkUserIdentityTypePPhoneUser
	// TeamworkUserIdentityTypePUnknownFutureValue is a pointer to TeamworkUserIdentityTypeVUnknownFutureValue
	TeamworkUserIdentityTypePUnknownFutureValue = &_TeamworkUserIdentityTypePUnknownFutureValue
)

var (
	_TeamworkUserIdentityTypePAadUser                      = TeamworkUserIdentityTypeVAadUser
	_TeamworkUserIdentityTypePOnPremiseAadUser             = TeamworkUserIdentityTypeVOnPremiseAadUser
	_TeamworkUserIdentityTypePAnonymousGuest               = TeamworkUserIdentityTypeVAnonymousGuest
	_TeamworkUserIdentityTypePFederatedUser                = TeamworkUserIdentityTypeVFederatedUser
	_TeamworkUserIdentityTypePPersonalMicrosoftAccountUser = TeamworkUserIdentityTypeVPersonalMicrosoftAccountUser
	_TeamworkUserIdentityTypePSkypeUser                    = TeamworkUserIdentityTypeVSkypeUser
	_TeamworkUserIdentityTypePPhoneUser                    = TeamworkUserIdentityTypeVPhoneUser
	_TeamworkUserIdentityTypePUnknownFutureValue           = TeamworkUserIdentityTypeVUnknownFutureValue
)
