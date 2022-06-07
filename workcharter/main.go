package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"gopkg.in/olahol/melody.v1"
	"io/ioutil"
	"log"
	"net/http"
	_ "os"
	_ "os/exec"
)

type PaymentJSON struct {
	Amount       int64  `json:"amount"`
	ReceiptEmail string `json:"receiptEmail"`
}

func main() {

	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "echo.html")
	})

	r.GET("/api/charges", func(c *gin.Context) {
		var json PaymentJSON
		c.BindJSON(&json)

		// apiKey := os.Getenv("StripeKey")
	 
apiKey := os.Getenv("API_KEY")
		
    log.Println("===What Is Token===")
    log.Println(stripe.String("tok_visa"))
    log.Println("===What Is Token===")
		stripe.Key = apiKey
	 
    response_stripe, err := charge.New(&stripe.ChargeParams{
      Amount:       stripe.Int64(json.Amount),
      Currency:     stripe.String(string(stripe.CurrencyAUD)),
      Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")}, // this should come from clientside
      ReceiptEmail: stripe.String(json.ReceiptEmail)})
    
    log.Println("===What Is Response_stripe===")
    log.Println(response_stripe)
    log.Println("===What Is Response_stripe===")
		if err != nil {
			// Handle any errors from attempt to charge
			c.String(http.StatusBadRequest, "Request failed")
			return
		}
		c.String(http.StatusCreated, "Successfully charged")

		
	})

	r.GET("/app", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Println("Message==== ")
		fmt.Println(string(msg))
		sbresp := "respstr"
		api_str := "https://wrapapi.com/use/worktrack/worklifearticles/getblogs/latest?wrapAPIKey=LFmRhQ131KpgnsRwKpeHMNNnl9a6vhvJ"
		if string(msg) == "wwr" {
			api_str = "https://wrapapi.com/use/worktrack/weworkremotely/getblogs/latest?wrapAPIKey=LFmRhQ131KpgnsRwKpeHMNNnl9a6vhvJ"
		}

		resp, err := http.Get(api_str)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		sb := string(body)
		sbresp = sb
		// fmt.Println(curl) )
		fmt.Println("curl==== ")
		fmt.Println(sb)
		fmt.Println("curl==== ")

		fmt.Println(string(msg))

		var buf []byte
		tomsg := []byte(sbresp)
		buf = append(msg, tomsg...)
		m.Broadcast(buf)
	})
	r.GET("/ws/list-articles", func(c *gin.Context) {
		// get articles websocket
		fmt.Println("\n  Get WS ")
		fmt.Println(m)
		m.HandleRequest(c.Writer, c.Request)
	})

	r.Run(":8000")
}
