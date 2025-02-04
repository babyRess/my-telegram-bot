package main

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HandleUserMessage(ctx context.Context, b *bot.Bot, update *models.Update, user User) {

	if update.Message.Text == "/start" {
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
	}
	if update.Message.Text == "/dev" {
		joinGroupButton := models.InlineKeyboardButton{
			Text: "Join Telegram Group",
			URL:  "https://t.me/mkt_royalgame",
		}

		openMiniAppButton := models.InlineKeyboardButton{
			Text: "(Rocket) Try Your Luck!(DEV)",
			WebApp: &models.WebAppInfo{
				URL: "https://api.g1388.makethatold.com/promobot/lobby/?gameIdentifier=rocket&gameEnv=local",
			},
		}
		openMiniAppButton2 := models.InlineKeyboardButton{
			Text: "(Bank) Try Your Luck!(DEV)",
			WebApp: &models.WebAppInfo{
				URL: "https://api.g1388.makethatold.com/promobot/lobby/?gameIdentifier=bank&gameEnv=local",
			},
		}
		openMiniAppButton3 := models.InlineKeyboardButton{
			Text: "(Money Tree) Try Your Luck!(DEV)",
			WebApp: &models.WebAppInfo{
				URL: "https://api.g1388.makethatold.com/promobot/lobby/?gameIdentifier=money_tree&gameEnv=local",
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
	}

	if update.Message.Text == "/slot" {
		openMiniAppButton := models.InlineKeyboardButton{
			Text: "New AnimalSlot!(DEV)",
			WebApp: &models.WebAppInfo{
				URL: "https://cdn.loteegames.com/ace/1738635245364/build/web-mobile-003/index.html",
			},
		}
		inlineKeyboard := &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{openMiniAppButton},
			},
		}
		// send to all users

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Welcome to the game! Please join the telegram group for news and click the miniapp to start playing.",
			ReplyMarkup: inlineKeyboard,
		})
	}
}
