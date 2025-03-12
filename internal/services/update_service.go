package services

import (
	telegram_api "api-gateway/internal/generated/telegram-api"
	"api-gateway/internal/models"
	"api-gateway/internal/utils"
	"reflect"
)

func ExtractUpdateType(update telegram_api.Update) (string, bool) {
	v := reflect.ValueOf(update)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() != reflect.Pointer {
			continue
		}
		if !field.IsNil() {
			fieldName := v.Type().Field(i).Name
			if fieldName != "" {
				return utils.ToSnakeCase(fieldName), true
			}
		}
	}
	return "", false
}

func ExtractUpdateSource(update telegram_api.Update, updateType string) (models.UpdateSource, bool) {
	if updateType == "business_connection" || updateType == "deleted_business_messages" || updateType == "edited_business_message" {
		return models.BusinessAccount, true
	}
	if updateType == "callback_query" {
		chatType := update.CallbackQuery.Message.GetInaccessibleMessage().Chat.Type
		if chatType == "" {
			chatType = update.CallbackQuery.Message.GetMessage().Chat.Type
		}
		return GetSourceFromChatType(chatType), true
	}
	if updateType == "channel_post" || updateType == "edited_channel_post" {
		return models.Channel, true
	}
	if updateType == "chat_boost" {
		return GetSourceFromChatType(update.ChatBoost.Chat.Type), true
	}
	if updateType == "removed_chat_boost" {
		return GetSourceFromChatType(update.RemovedChatBoost.Chat.Type), true
	}
	if updateType == "chat_member" {
		return GetSourceFromChatType(update.ChatMember.Chat.Type), true
	}
	if updateType == "chat_join_request" {
		return GetSourceFromChatType(update.ChatJoinRequest.Chat.Type), true
	}
	if updateType == "my_chat_member" {
		return GetSourceFromChatType(update.MyChatMember.Chat.Type), true
	}
	if updateType == "chosen_inline_result" || updateType == "inline_query" {
		return models.InlineMode, true
	}
	if updateType == "message" {
		return GetSourceFromChatType(update.Message.Chat.Type), true
	}
	if updateType == "edited_message" {
		return GetSourceFromChatType(update.EditedMessage.Chat.Type), true
	}
	if updateType == "poll" || updateType == "poll_answer" {
		return models.Poll, true
	}
	if updateType == "pre_checkout_query" || updateType == "purchased_paid_media" || updateType == "shipping_query" {
		return models.Payment, true
	}
	if updateType == "message_reaction" {
		return GetSourceFromChatType(update.MessageReaction.Chat.Type), true
	}
	if updateType == "message_reaction_count" {
		return GetSourceFromChatType(update.MessageReactionCount.Chat.Type), true
	}
	return models.Unknown, false
}

func GetSourceFromChatType(chatType string) models.UpdateSource {
	switch chatType {
	case "group":
		return models.Group
	case "super_group":
		return models.SuperGroup
	case "private":
		return models.PrivateChat
	case "channel":
		return models.Channel
	default:
		return models.Unknown
	}
}
