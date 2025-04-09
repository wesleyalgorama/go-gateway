package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id        string
	Name      string
	Email     string
	ApiKey    string
	Balance   float64
	mu        sync.RWMutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

func generateApiKey() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func NewAccount(name string, email string) *Account {
	account := &Account{
		Id:        uuid.New().String(),
		Name:      name,
		Email:     email,
		ApiKey:    generateApiKey(),
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}

func (a *Account) AddBalance(amount float64) {
	a.mu.Lock() // lock other transactions
	a.Balance += amount
	a.UpdatedAt = time.Now()
	defer a.mu.Unlock() // defer garants that it will be called at the end
}
