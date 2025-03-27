package telegram

// Poll represents a poll
type Poll struct {
	ID                    string          `json:"id"`
	Question              string          `json:"question"`
	Options               []PollOption    `json:"options"`
	TotalVoterCount       int             `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CorrectOptionID       int             `json:"correct_option_id,omitempty"`
	Explanation           string          `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int             `json:"open_period,omitempty"`
	CloseDate             int64           `json:"close_date,omitempty"`
}

// PollOption represents an answer option in a poll
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

// MessageAutoDeleteTimerChanged represents a service message about a change in auto-delete timer settings
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// Invoice represents an invoice
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int64  `json:"total_amount"`
}

// SuccessfulPayment represents a successful payment
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int64      `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionID        string     `json:"shipping_option_id,omitempty"`
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
}

// OrderInfo represents information about an order
type OrderInfo struct {
	Name            string           `json:"name,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	Email           string           `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// ShippingAddress represents a shipping address
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2,omitempty"`
	PostCode    string `json:"post_code"`
}

// UserShared represents a service message about a user sharing something with their chat
type UserShared struct {
	RequestID int   `json:"request_id"`
	UserID    int64 `json:"user_id"`
}

// ChatShared represents a service message about a chat being shared with the bot
type ChatShared struct {
	RequestID int   `json:"request_id"`
	ChatID    int64 `json:"chat_id"`
}

// WriteAccessAllowed represents a service message about a user being allowed to write messages
type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request,omitempty"`
	WebAppName         string `json:"web_app_name,omitempty"`
	FromAttachmentMenu bool   `json:"from_attachment_menu,omitempty"`
}

// PassportData represents Telegram Passport data shared with the bot by the user
type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials       `json:"credentials"`
}

// EncryptedPassportElement represents information about documents or other Telegram Passport elements shared with the bot
type EncryptedPassportElement struct {
	Type        string         `json:"type"`
	Data        string         `json:"data,omitempty"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	Email       string         `json:"email,omitempty"`
	Files       []PassportFile `json:"files,omitempty"`
	FrontSide   *PassportFile  `json:"front_side,omitempty"`
	ReverseSide *PassportFile  `json:"reverse_side,omitempty"`
	Selfie      *PassportFile  `json:"selfie,omitempty"`
	Translation []PassportFile `json:"translation,omitempty"`
	Hash        string         `json:"hash"`
}

// PassportFile represents a file uploaded to Telegram Passport
type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	FileDate     int64  `json:"file_date"`
}

// EncryptedCredentials represents data required for decrypting and authenticating EncryptedPassportElement
type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

// ProximityAlertTriggered represents a service message about a user triggering a proximity alert
type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"`
	Watcher  User `json:"watcher"`
	Distance int  `json:"distance"`
}

// ForumTopicCreated represents a service message about a new forum topic created in the chat
type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicEdited represents a service message about an edited forum topic
type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicClosed represents a service message about a forum topic closed in the chat
type ForumTopicClosed struct{}

// ForumTopicReopened represents a service message about a forum topic reopened in the chat
type ForumTopicReopened struct{}

// GeneralForumTopicHidden represents a service message about the General forum topic hidden in the chat
type GeneralForumTopicHidden struct{}

// GeneralForumTopicUnhidden represents a service message about the General forum topic unhidden in the chat
type GeneralForumTopicUnhidden struct{}

// VideoChatScheduled represents a service message about a video chat scheduled in the chat
type VideoChatScheduled struct {
	StartDate int64 `json:"start_date"`
}

// VideoChatStarted represents a service message about a video chat started in the chat
type VideoChatStarted struct{}

// VideoChatEnded represents a service message about a video chat ended in the chat
type VideoChatEnded struct {
	Duration int `json:"duration"`
}

// VideoChatParticipantsInvited represents a service message about new members invited to a video chat
type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}

// WebAppData represents data from a Web App
type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

// Update represents an incoming update
type Update struct {
	UpdateID                int64                    `json:"update_id"`
	Message                 *Message                 `json:"message,omitempty"`
	EditedMessage           *Message                 `json:"edited_message,omitempty"`
	ChannelPost             *Message                 `json:"channel_post,omitempty"`
	EditedChannelPost       *Message                 `json:"edited_channel_post,omitempty"`
	InlineQuery             *InlineQuery             `json:"inline_query,omitempty"`
	ChosenInlineResult      *ChosenInlineResult      `json:"chosen_inline_result,omitempty"`
	CallbackQuery           *CallbackQuery           `json:"callback_query,omitempty"`
	ShippingQuery           *ShippingQuery           `json:"shipping_query,omitempty"`
	PreCheckoutQuery        *PreCheckoutQuery        `json:"pre_checkout_query,omitempty"`
	Poll                    *Poll                    `json:"poll,omitempty"`
	PollAnswer              *PollAnswer              `json:"poll_answer,omitempty"`
	MyChatMember            *ChatMember              `json:"my_chat_member,omitempty"`
	ChatMember              *ChatMember              `json:"chat_member,omitempty"`
	ChatJoinRequest         *ChatJoinRequest         `json:"chat_join_request,omitempty"`
	ChatBoost               *ChatBoost               `json:"chat_boost,omitempty"`
	RemovedChatBoost        *RemovedChatBoost        `json:"removed_chat_boost,omitempty"`
	BusinessConnection      *BusinessConnection      `json:"business_connection,omitempty"`
	EditedBusinessMessage   *Message                 `json:"edited_business_message,omitempty"`
	DeletedBusinessMessages *DeletedBusinessMessages `json:"deleted_business_messages,omitempty"`
	MessageReaction         *MessageReaction         `json:"message_reaction,omitempty"`
	MessageReactionCount    *MessageReactionCount    `json:"message_reaction_count,omitempty"`
}

// InlineQuery represents an incoming inline query
type InlineQuery struct {
	ID       string    `json:"id"`
	From     User      `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType string    `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}

// ChosenInlineResult represents a result of an inline query that was chosen by the user and sent to their chat partner
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            User      `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            User     `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
}

// ShippingQuery represents an incoming shipping query
type ShippingQuery struct {
	ID              string          `json:"id"`
	From            User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

// PreCheckoutQuery represents an incoming pre-checkout query
type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             User       `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int64      `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}

// PollAnswer represents an answer of a user in a non-anonymous poll
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	User      User   `json:"user"`
	OptionIDs []int  `json:"option_ids"`
}

// ChatJoinRequest represents a join request sent to a chat
type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`
	From       User            `json:"from"`
	Date       int64           `json:"date"`
	Bio        string          `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

// ChatInviteLink represents an invite link for a chat
type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`
	Creator                 User   `json:"creator"`
	IsPrimary               bool   `json:"is_primary"`
	IsRevoked               bool   `json:"is_revoked"`
	Name                    string `json:"name,omitempty"`
	ExpireDate              int64  `json:"expire_date,omitempty"`
	MemberLimit             int    `json:"member_limit,omitempty"`
	PendingJoinRequestCount int    `json:"pending_join_request_count,omitempty"`
	IsApproved              bool   `json:"is_approved,omitempty"`
	IsRequestNeeded         bool   `json:"is_request_needed,omitempty"`
}

// ChatMember represents a chat member
type ChatMember struct {
	User                  User   `json:"user"`
	Chat                  Chat   `json:"chat"`
	Status                string `json:"status"`
	CustomTitle           string `json:"custom_title,omitempty"`
	IsAnonymous           bool   `json:"is_anonymous,omitempty"`
	CanBeEdited           bool   `json:"can_be_edited,omitempty"`
	CanManageChat         bool   `json:"can_manage_chat,omitempty"`
	CanDeleteMessages     bool   `json:"can_delete_messages,omitempty"`
	CanManageVideoChats   bool   `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`
	CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`
	CanChangeInfo         bool   `json:"can_change_info,omitempty"`
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`
	CanPostMessages       bool   `json:"can_post_messages,omitempty"`
	CanEditMessages       bool   `json:"can_edit_messages,omitempty"`
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`
	CanPostStories        bool   `json:"can_post_stories,omitempty"`
	CanEditStories        bool   `json:"can_edit_stories,omitempty"`
	CanDeleteStories      bool   `json:"can_delete_stories,omitempty"`
	IsMember              bool   `json:"is_member,omitempty"`
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`
	CanSendAudios         bool   `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool   `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool   `json:"can_send_photos,omitempty"`
	CanSendVideos         bool   `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     bool   `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool   `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"`
	UntilDate             int64  `json:"until_date,omitempty"`
	CanManageTopics       bool   `json:"can_manage_topics,omitempty"`
}

// ChatBoost represents a boost added to a chat or changed
type ChatBoost struct {
	Chat           Chat             `json:"chat"`
	BoostID        string           `json:"boost_id"`
	AddDate        int64            `json:"add_date"`
	ExpirationDate int64            `json:"expiration_date,omitempty"`
	Source         *ChatBoostSource `json:"source"`
}

// ChatBoostSource represents the source of a chat boost
type ChatBoostSource struct {
	Source      string `json:"source"`
	User        *User  `json:"user,omitempty"`
	IsUnclaimed bool   `json:"is_unclaimed,omitempty"`
}

// RemovedChatBoost represents a boost removed from a chat
type RemovedChatBoost struct {
	Chat       Chat             `json:"chat"`
	BoostID    string           `json:"boost_id"`
	RemoveDate int64            `json:"remove_date"`
	Source     *ChatBoostSource `json:"source"`
}

// BusinessConnection represents a business connection of the bot
type BusinessConnection struct {
	ID         string `json:"id"`
	UserChatID int64  `json:"user_chat_id"`
	Username   string `json:"username,omitempty"`
	CanReply   bool   `json:"can_reply"`
	IsEnabled  bool   `json:"is_enabled"`
}

// DeletedBusinessMessages represents messages deleted from a connected business account
type DeletedBusinessMessages struct {
	Chat       Chat    `json:"chat"`
	MessageIDs []int64 `json:"message_ids"`
}

// MessageReaction represents a change in a message reaction
type MessageReaction struct {
	Chat        Chat           `json:"chat"`
	MessageID   int64          `json:"message_id"`
	User        *User          `json:"user,omitempty"`
	ActorChat   *Chat          `json:"actor_chat,omitempty"`
	Date        int64          `json:"date"`
	OldReaction []ReactionType `json:"old_reaction,omitempty"`
	NewReaction []ReactionType `json:"new_reaction,omitempty"`
}

// MessageReactionCount represents a change in the number of reactions on a message
type MessageReactionCount struct {
	Chat      Chat            `json:"chat"`
	MessageID int64           `json:"message_id"`
	Date      int64           `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}

// ReactionType represents a type of reaction
type ReactionType struct {
	Type          string `json:"type"`
	Emoji         string `json:"emoji,omitempty"`
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

// ReactionCount represents a count of reactions of a specific type
type ReactionCount struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}
