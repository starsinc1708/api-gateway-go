package models

type UpdateSource string

const (
	BusinessAccount UpdateSource = "business_account"
	Channel         UpdateSource = "channel"
	Group           UpdateSource = "group"
	SuperGroup      UpdateSource = "supergroup"
	PrivateChat     UpdateSource = "private_chat"
	InlineMode      UpdateSource = "inline_mode"
	Payment         UpdateSource = "payment"
	Poll            UpdateSource = "poll"
	Unknown         UpdateSource = "unknown"
)

type ModuleInfo struct {
	ModuleName    string
	TransportType string
}
