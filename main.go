package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogData struct {
	UserInfo *models.User     `bson:"user_info"`
	Date     time.Time        `bson:"date"`
	Location *models.Location `bson:"location"`
}

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Username  string
	IsAdmin   bool
}

var urlMiniApp string
var appName string

var adminUserIDs = []int64{}
var mongoClient *mongo.Client

func main() {
	// Get the bot token from environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables.")
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}
	ctx := context.Background()

	mongoClient = connectToMongo(ctx)
	// loop update adminUserEvery 15 min
	go func() {
		for {
			adminUserIDs = getAdminUserIDs()
			time.Sleep(15 * time.Minute)
		}
	}()
	// Set up bot configuration
	opts := []bot.Option{
		bot.WithDefaultHandler(handleMessage),
	}

	// Create the bot instance
	b, err := bot.New(botToken, opts...)
	if err != nil {
		log.Fatal(err)
	}
	sendMessageToChannel(ctx, b, nil)
	// get url and app name
	// defaultApp, err := b.GetChatMenuButton(ctx, &bot.GetChatMenuButtonParams{
	// 	ChatID: 5038436839,
	// })
	// appName = defaultApp.WebApp.Text
	// urlMiniApp = defaultApp.WebApp.WebApp.URL

	// Create a context with an increased timeout
	// Start the bot
	fmt.Println("Bot start")

	b.Start(ctx)

	// make this bot running forever
}

// update mini app url and app name when urlMiniApp change

// handleMessage is a simple handler function for received messages
func handleMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	user := User{
		ID:        update.Message.From.ID,
		FirstName: update.Message.From.FirstName,
		LastName:  update.Message.From.LastName,
		Username:  update.Message.From.Username,
		IsAdmin:   isAdmin(update.Message.From.ID),
	}
	// if update.Message.Text == "/notice" {
	// 	// send message to channel
	// 	sendMessageToChannel(ctx, b, update)
	// }
	if update.Message.Text == "/start" {
		// add to mongo db all info about user
		logData := LogData{
			UserInfo: update.Message.From,
			Date:     time.Unix(int64(update.Message.Date), 0),
			Location: update.Message.Location,
		}
		addDataToMongo(logData)
	}

	if update.Message.Text == "/help" {
		helpMessage := "Available Commands:\n" +
			"/help - Show this help message"

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   helpMessage,
		})
		return
	}

	if user.IsAdmin {
		HandleAdminMessage(ctx, b, update, user)
	}
	HandleUserMessage(ctx, b, update, user)

}

func isAdmin(userID int64) bool {
	for _, id := range adminUserIDs {
		if id == userID {
			return true
		}
	}
	return false
}

func addDataToMongo(logData LogData) {
	collection := mongoClient.Database("telegram_bot").Collection("logs")
	//check if user already exists
	filter := bson.D{{"user_info.id", logData.UserInfo.ID}}
	var result User
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// User does not exist, insert new user
		_, err := collection.InsertOne(context.Background(), logData)
		fmt.Println(err, logData.UserInfo.ID)
	} else {
		fmt.Println("User already exists")
	}

}

func connectToMongo(ctx context.Context) *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://datnt13022:asdasd11@cluster0.juehd.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(ctx, bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Successfully connected and pinged MongoDB.")
	return client
}

// get admin user id from mongo db
func getAdminUserIDs() []int64 {
	collection := mongoClient.Database("telegram_bot").Collection("admins")
	var adminIDs []int64

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatalf("Failed to find admin user IDs: %v", err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var document bson.M
		if err := cursor.Decode(&document); err != nil {
			log.Fatalf("Failed to decode admin user ID: %v", err)
		}
		userId, err := strconv.ParseInt(document["userId"].(string), 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse admin user ID: %v", err)
		}
		adminIDs = append(adminIDs, userId)
	}
	fmt.Println(adminIDs)
	return adminIDs
}

func sendMessageToChannel(ctx context.Context, b *bot.Bot, update *models.Update) {
	// send message to channel
	fmt.Println("send message to channel")

	// create inlineButton
	//send link open mini app telegram
	inlineKeyboard := [][]models.InlineKeyboardButton{
		{{Text: "Try Your Luck!🍀", URL: "https://t.me/tetleleksmgv_bot?startapp=gameapp&gameIdentifier=bank"}},
	}
	// https: //t.me/Catizenbot/gameapp?startapp=open_3
	_, err := b.SendPhoto(ctx, &bot.SendPhotoParams{
		ChatID: -1002254268367,
		Caption: `🚀 NO WINNING? NO PROBLEM! 🚀
💰💰FREE MONEY JUST FOR PLAYING!💰💰

🐹 Dear Royaler,

👉 We have added many NEW games to our gamification services check the Game list [here](https://t.me/tetleleksmgv_bot)

😉 If you decide to keep your tokens instead of CLAIMING them, you’ll be (very) pleasantly surprised!

👍 Our game will boost the prizes and continue to give you FREE REWARDS!

🧡 Stay tuned! We’ve prepared more exciting features!!!`,
		ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboard},
		ParseMode:   "Markdown", // or "HTML"
		Photo: &models.InputFileString{
			Data: "https://staging-acegames.s3-ap-southeast-1.amazonaws.com/uploads/telebots.jpeg",
		},
	})
	if err != nil {
		log.Printf("Failed to send photo: %v", err)
	}
}

func getUsers() []int64 {
	collection := mongoClient.Database("telegram_bot").Collection("logs")
	var users []int64
	uniqueUserIDs := make(map[int64]struct{}) // To avoid duplicates

	// Use an aggregation pipeline to get distinct user_info.id values
	pipeline := mongo.Pipeline{
		{
			{"$group", bson.D{
				{"_id", "$user_info.id"},
			}},
		},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Fatalf("Failed to find users: %v", err)
	}
	defer cursor.Close(context.Background())

	// Iterate through the cursor and extract user IDs
	for cursor.Next(context.Background()) {
		var document bson.M
		if err := cursor.Decode(&document); err != nil {
			log.Fatalf("Failed to decode user ID: %v", err)
		}

		// Extract the user ID from the "_id" field
		if userID, ok := document["_id"].(int64); ok {
			// Add the user ID to the map (ensuring uniqueness)
			uniqueUserIDs[userID] = struct{}{}
		}
	}

	// Convert the map to a slice of user IDs
	for userID := range uniqueUserIDs {
		users = append(users, userID)
	}

	return users
}
