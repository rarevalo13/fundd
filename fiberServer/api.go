package fiberserver

import (
	"fmt"
	"log"
	"time"

	"fundd.com/fundd/db"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/plaid/plaid-go/plaid"
	"gopkg.in/ini.v1"
)

func initClient() plaid.ClientOptions {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal(err)
	}

	var CLIENT_ID = cfg.Section("plaid").Key("CLIENT_ID").String()
	var SECRET = cfg.Section("plaidsandbox").Key("CLIENT_SECRET").String()
	var environment = plaid.Sandbox

	var clientOptions = plaid.ClientOptions{
		ClientID:    CLIENT_ID,
		Secret:      SECRET,
		Environment: environment,
	}
	return clientOptions
}

var client, err = plaid.NewClient(initClient())

//This endpoint should do the following:
//1. Make a call to clinet.GetLinkToken
//2. Pass that to client.ExchangePublicToken
//That should generate an access_token that can be saved
// and used to make requests
//var client, err = plaid.NewClient(clientOptions)

func GetToken(ctx *fiber.Ctx) error {

	newToken := &plaid.LinkTokenConfigs{
		User:         &plaid.LinkTokenUser{ClientUserID: uuid.NewString()},
		ClientName:   "Fund\\'d",
		Language:     "en",
		CountryCodes: []string{"US"},
		Products:     []string{"transactions"},
	}

	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"Message": err,
		})
	}

	linkTokenResp, err := client.CreateLinkToken(*newToken)
	if err != nil {
		log.Panic(err)
	}

	return ctx.JSON(fiber.Map{
		"link_token": linkTokenResp.LinkToken,
	})

}

func GetAccessToken(ctx *fiber.Ctx) error {
	type PublicToken struct {
		PublicToken string `json:"public_token"`
	}

	createAccessToken := PublicToken{}

	ctx.BodyParser(&createAccessToken)
	//log.Println(createAccessToken.PublicToken)
	accessToken, err := client.ExchangePublicToken(createAccessToken.PublicToken)

	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"Message": err,
		})
	}
	return ctx.JSON(fiber.Map{
		"access_token": accessToken.AccessToken,
		"item_id":      accessToken.ItemID,
	})

}

//This endpoing should get recent transactions for last two weeks.

func GetRecentTxns(ctx *fiber.Ctx) error {

	type AccessToken struct {
		AccessToken string `json:"access_token"`
		ItemID      string `json:"item_id,omitempty"`
	}
	useAccessToken := AccessToken{}
	ctx.BodyParser(&useAccessToken)
	const iso8601TimeFormat = "2006-01-02"
	startDate := time.Now().Add(-30 * 24 * time.Hour).Format(iso8601TimeFormat)
	endDate := time.Now().Format(iso8601TimeFormat)

	transactionsResp, err := client.GetTransactions(useAccessToken.AccessToken, startDate, endDate)
	if err != nil {
		log.Fatal(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"Message": err,
		})
	}
	return ctx.JSON(fiber.Map{
		"recent_transactions": transactionsResp.Transactions,
	})
}

// func sampleAccessTest(ctx *fiber.Ctx) error {
// 	type AccessToken struct {
// 		AccessToken string `json:"access_token"`
// 		ItemID      string `json:"item_id,omitempty"`
// 	}
// 	useAccessToken := AccessToken{}
// 	ctx.BodyParser(&useAccessToken)
// 	fmt.Println("this is the accesstoken from the post body ", useAccessToken.AccessToken)
// 	if err != nil {
// 		log.Fatal(err)
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"success": false,
// 			"Message": err,
// 		})
// 	}
// 	return ctx.JSON(fiber.Map{
// 		"token": useAccessToken.AccessToken,
// 	})

// }

func GetAllTxns(ctx *fiber.Ctx) {
	ctx.SendString("All Txns")
}

// func GetToken(ctx *fiber.Ctx) {
// 	ctx.Send("TOKEN")
// }

// func registerUser() {

// }

func SimpleGet(ctx *fiber.Ctx) error {
	table := "users"
	db.ReadFromTable(table)

	return ctx.SendString("read from the DB")
}
