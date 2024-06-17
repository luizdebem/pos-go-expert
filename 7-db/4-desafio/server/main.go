package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type USDBRLResponse struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type Exchange struct {
	ID          uint `gorm:"primaryKey;"`
	Value       string
	CurrencyIn  string
	CurrencyOut string
	gorm.Model
}

func main() {
	resetDatabase()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /cotacao", getCotacao)
	http.ListenAndServe(":8080", mux)
}

func getCotacao(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println("A requisição externa teve tempo excedido.")
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	var cotacao USDBRLResponse
	if err := json.NewDecoder(res.Body).Decode(&cotacao); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = persistExchange(cotacao)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json := []byte(`{"cotacao": ` + cotacao.USDBRL.Bid + `}`)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func resetDatabase() {
	os.Remove("./exchange.db")
	db, err := gorm.Open(sqlite.Open("./exchange.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Exchange{})
}

func persistExchange(exchange USDBRLResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	db, err := gorm.Open(sqlite.Open("./exchange.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	result := db.WithContext(ctx).Create(&Exchange{
		Value:       exchange.USDBRL.Bid,
		CurrencyIn:  "USD",
		CurrencyOut: "BRL",
	})

	if result.Error != nil {
		if errors.Is(result.Error, context.DeadlineExceeded) {
			fmt.Println("A persistência no banco de dados teve tempo excedido.")
		}
		return result.Error
	}
	return nil

}
