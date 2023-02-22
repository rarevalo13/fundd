package main

import (
	fiberserver "fundd.com/fundd/fiberServer"
)

// type  publicTokenResp struct {
// 	PublicToken string `json:"public_token"`
// 	RequestID string `json:"request_id"`
// }

// type publicTokenReq struct {
// 	ClientID string `json:"client_id"`
// 	Secret string `json:"secret"`
// 	InstitutionID string `json:"institution_id"`
// 	InitialProducts []string `json:"initial_products"'`
// 	Options options `json:"options"`
// }

// type options struct {
// 	Webhook string `json:"webhook,omitempty"`
// 	OverrideUserName string `json:"override_username,omitempty"`
// 	OverridePassword string `json:"override_password,omitempty"`
// 	Transactions txns `json:"transactions,omitempty"`
// }

// type txns struct {
// 	StartDate string `json:"start_date"`
// 	EndDate string `json:"end_date"`
// }

// type accessToken struct {
// 	AccessToken string `json:"access_token"`
// 	ItemID string `json:"item_id"`
// 	RequestID string `json:"request_id"`
// }

// type accessToekenReq struct {
// 	ClientID string `json:"client_id"`
// 	Secret string `json:"secret"`
// 	PublicToken string `json:"public_token"`
// }

// var client_id = os.Getenv("PLAID_CLIENT_ID")
// var clientSecret = os.Getenv("PLAID_SECRET")
// var baseURL = "https://sandbox.plaid.com"
// var txnEndpoint = "/transactions/get"
// var sandboxTokenCreate = "/sandbox/public_token/create"
// var publicTokenExchange = "/item/public_token/exchange"
// var CapitalOne = "ins_9"
// var Chase = "ins_3"
// var Tinker = "ins_120019"

func main() {
	// publicToken := createPublicToken()
	// fmt.Println(exchangePublicToken(publicToken))
	fiberserver.Server()

	// client, err := plaid.NewClient(plaid.ClientOptions{"5ff63628ca7ffe0012a506fc", "a89428559e9fc9dd3f36c61cacf458", plaid.Sandbox})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// linkToken, err := client.CreatePublicToken()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Print(linkToken)
}

// func createPublicToken() string {
// 	var newpublicTokenReq publicTokenResp
// 	getNewPublicToken := publicTokenReq{
// 		ClientID:        client_id,
// 		Secret:          clientSecret,
// 		InstitutionID:   Chase,
// 		InitialProducts: []string{"transactions"},
// 		Options:         options{Transactions: txns{StartDate: "2020-01-01", EndDate: "2020-01-08"}},
// 	}
// 	requestBody, err := json.Marshal(getNewPublicToken)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	req, err := http.NewRequest("POST", baseURL+sandboxTokenCreate, bytes.NewBufferString(string(requestBody)))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	req.Header.Add("Content-Type", "application/json")
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	body, err := ioutil.ReadAll(resp.Body)
// 	err = json.Unmarshal(body, &newpublicTokenReq)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()
// 	return newpublicTokenReq.PublicToken
// }

// func exchangePublicToken(publicToken string) string {
// 	var newAccessToken accessToken
// 	getNewAccessToken := accessToekenReq{
// 		ClientID:    client_id,
// 		Secret:      clientSecret,
// 		PublicToken: publicToken}

// 	requestBody, err := json.Marshal(getNewAccessToken)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	req, err := http.NewRequest("POST", baseURL+publicTokenExchange, bytes.NewBufferString(string(requestBody)))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	req.Header.Add("Content-Type", "application/json")
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	body, err := ioutil.ReadAll(resp.Body)
// 	err = json.Unmarshal(body, &newAccessToken)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()
// 	return newAccessToken.AccessToken
// }
