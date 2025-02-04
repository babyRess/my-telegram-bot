package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// HandleAdminMessage processes messages from admin users
func HandleAdminMessage(ctx context.Context, b *bot.Bot, update *models.Update, user User) {
	// add all default command when as a admin

	if strings.HasPrefix(update.Message.Text, "/updateUrl") {
		// Extract the URL from the message text
		parts := strings.Split(update.Message.Text, " ")
		if len(parts) < 3 {
			// Handle the case where no URL is provided
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Please provide a URL after the command. example: /updateUrl <url> <app name>",
			})
			return
		}

		urlMiniApp = parts[1]
		// app name is from parts[2] to end
		appName = strings.Join(parts[2:], " ")
		allUsers := getUsers()
		fmt.Println(allUsers)

		joinGroupButton := models.InlineKeyboardButton{
			Text: "Join Telegram Group",
			URL:  "https://t.me/mkt_royalgame",
		}

		openMiniAppButton := models.InlineKeyboardButton{
			Text: "(Rocket) Try Your Luck!",
			WebApp: &models.WebAppInfo{
				URL: "https://api.g1388.makethatold.com/promobot/lobby/?gameIdentifier=rocket",
			},
		}
		openMiniAppButton2 := models.InlineKeyboardButton{
			Text: "(Bank) Try Your Luck!",
			WebApp: &models.WebAppInfo{
				URL: "https://api.g1388.makethatold.com/promobot/lobby/?gameIdentifier=bank",
			},
		}
		openMiniAppButton3 := models.InlineKeyboardButton{
			Text: "(Money Tree) Try Your Luck!",
			WebApp: &models.WebAppInfo{
				URL: "https://api.g1388.makethatold.com/promobot/lobby/?gameIdentifier=money_tree",
			},
		}
		inlineKeyboard := &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{joinGroupButton},
				{openMiniAppButton},
				{openMiniAppButton2},
				{openMiniAppButton3},
			},
		}
		// send to all users

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Welcome to the game! Please join the telegram group for news and click the miniapp to start playing.",
			ReplyMarkup: inlineKeyboard,
		})

		b.SetChatMenuButton(
			ctx, &bot.SetChatMenuButtonParams{
				ChatID: update.Message.Chat.ID,
				MenuButton: &models.MenuButtonWebApp{
					Type: "web_app",
					Text: appName,
					WebApp: models.WebAppInfo{
						URL: urlMiniApp,
					},
				}})
		for _, user := range allUsers {
			// Set the chat menu button with the extracted URL
			b.SetChatMenuButton(
				ctx, &bot.SetChatMenuButtonParams{
					ChatID: user,
					MenuButton: &models.MenuButtonWebApp{
						Type: "web_app",
						Text: "Try Your Luck!🍀",
						WebApp: models.WebAppInfo{
							URL: "https://api.g1388.makethatold.com/promobot/lobby/?gameIdentifier=bank",
						},
					}})
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:      user,
				Text:        "Welcome to the game! Please join the telegram group for news and click the miniapp to start playing.",
				ReplyMarkup: inlineKeyboard,
			})
		}

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Updated",
		})
	}
}
