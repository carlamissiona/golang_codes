package paymentcode

 

import (
	"context"
	"fmt"
	"html/template"
	
	"net/http"
	"os"

	"github.com/braintree-go/braintree-go"
)

type BraintreeJS struct {
	Key template.HTML
}

func showForm(w http.ResponseWriter, r *http.Request) {
	config := BraintreeJS{Key: "'" + template.HTML(os.Getenv("BRAINTREE_CSE_KEY")) + "'"}
	t := template.Must(template.ParseFiles("checkout_form.html"))
	err := t.Execute(w, config)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	bt := braintree.New(
		braintree.Sandbox,
		os.Getenv("BRAINTREE_MERCH_ID"),
		os.Getenv("BRAINTREE_PUB_KEY"),
		os.Getenv("BRAINTREE_PRIV_KEY"),
	)

  

	tx := &braintree.TransactionRequest{
		Type:   "sale",
		Amount: braintree.NewDecimal(10000, 2),
		CreditCard: &braintree.CreditCard{
			Number:          r.FormValue("number"),
			CVV:             r.FormValue("cvv"),
			ExpirationMonth: r.FormValue("month"),
			ExpirationYear:  r.FormValue("year"),
		},
	}

	_, err := bt.Transaction().Create(ctx, tx)

	if err == nil {
		_, _ = fmt.Fprintf(w, "<h1>Success!</h1>")
	} else {
		_, _ = fmt.Fprintf(w, "<h1>Something went wrong: "+err.Error()+"</h1>")
	}
}

func Pay() {
	http.HandleFunc("/paypal_form", showForm)
	http.HandleFunc("/create_transaction", createTransaction)
  
}


