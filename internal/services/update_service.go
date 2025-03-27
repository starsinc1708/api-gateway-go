package services

import (
	"api-gateway/internal/models"
	"api-gateway/internal/models/telegram"
	"api-gateway/internal/utils"
	"reflect"
)

func ExtractUpdateType(update telegram.Update) string {
	v := reflect.ValueOf(update)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() != reflect.Pointer {
			continue
		}
		if !field.IsNil() {
			fieldName := v.Type().Field(i).Name
			if fieldName != "" {
				return utils.ToSnakeCase(fieldName)
			}
		}
	}
	return "unknown"
}

func ExtractUpdateSource(update telegram.Update, updateType string) (models.UpdateSource, int64) {
	if updateType == "business_connection" {
		return models.BusinessAccount, update.BusinessConnection.UserChatID
	}
	if updateType == "edited_business_message" {
		return models.BusinessAccount, update.EditedBusinessMessage.From.ID
	}
	if updateType == "deleted_business_messages" {
		return models.BusinessAccount, update.DeletedBusinessMessages.Chat.ID
	}
	if updateType == "callback_query" {
		var chatType string
		if msg := update.CallbackQuery.Message; msg != nil {
			chatType = msg.Chat.Type
			return GetSourceFromChatType(chatType), msg.Chat.ID
		}
		return models.Unknown, int64(0)
	}
	if updateType == "channel_post" {
		return models.Channel, update.ChannelPost.Chat.ID
	}
	if updateType == "edited_channel_post" {
		return models.Channel, update.EditedChannelPost.Chat.ID
	}
	if updateType == "chat_boost" {
		return GetSourceFromChatType(update.ChatBoost.Chat.Type), update.ChatBoost.Chat.ID
	}
	if updateType == "removed_chat_boost" {
		return GetSourceFromChatType(update.RemovedChatBoost.Chat.Type), update.RemovedChatBoost.Chat.ID
	}
	if updateType == "chat_member" {
		return GetSourceFromChatType(update.ChatMember.Chat.Type), update.ChatMember.Chat.ID
	}
	if updateType == "chat_join_request" {
		return GetSourceFromChatType(update.ChatJoinRequest.Chat.Type), update.ChatJoinRequest.Chat.ID
	}
	if updateType == "my_chat_member" {
		return GetSourceFromChatType(update.MyChatMember.Chat.Type), update.MyChatMember.Chat.ID
	}
	if updateType == "chosen_inline_result" {
		return models.InlineMode, update.ChosenInlineResult.From.ID
	}
	if updateType == "inline_query" {
		return models.InlineMode, update.InlineQuery.From.ID
	}
	if updateType == "message" {
		return GetSourceFromChatType(update.Message.Chat.Type), update.Message.Chat.ID
	}
	if updateType == "edited_message" {
		return GetSourceFromChatType(update.EditedMessage.Chat.Type), update.EditedMessage.Chat.ID
	}
	if updateType == "poll_answer" {
		return models.Poll, update.PollAnswer.User.ID
	}
	if updateType == "poll" {
		return models.Poll, int64(0)
	}
	if updateType == "pre_checkout_query" {
		return models.Payment, update.PreCheckoutQuery.From.ID
	}
	if updateType == "shipping_query" {
		return models.Payment, update.ShippingQuery.From.ID
	}
	if updateType == "message_reaction" {
		return GetSourceFromChatType(update.MessageReaction.Chat.Type), update.MessageReaction.Chat.ID
	}
	if updateType == "message_reaction_count" {
		return GetSourceFromChatType(update.MessageReactionCount.Chat.Type), update.MessageReactionCount.Chat.ID
	}
	return models.Unknown, int64(0)
}

func GetSourceFromChatType(chatType string) models.UpdateSource {
	switch chatType {
	case "group":
		return models.Group
	case "supergroup":
		return models.SuperGroup
	case "private":
		return models.PrivateChat
	case "channel":
		return models.Channel
	default:
		return models.Unknown
	}
}
