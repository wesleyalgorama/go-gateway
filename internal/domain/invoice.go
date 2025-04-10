package domain

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	Pending  Status = "pending"
	Approved Status = "approved"
	Rejected Status = "rejected"
)

type Invoice struct {
	Id             string
	AccountId      string
	Status         Status
	Amount         float64
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreditCard struct {
	Number         string
	CVV            string
	ExpiryMonth    int
	ExpiryYear     int
	CardHolderName string
}

func NewInvoice(accountId string, amount float64, description string, paymentType string, card CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	lastDigits := card.Number[len(card.Number)-4:]

	return &Invoice{
		Id:             uuid.New().String(),
		AccountId:      accountId,
		Amount:         amount,
		Status:         Pending,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func (i *Invoice) Process() error {
	if i.Amount > 10000 {
		return nil
	}

	randomSource := rand.New(rand.NewSource(time.Now().Unix()))

	var newStatus Status

	if randomSource.Float64() <= 0.7 {
		newStatus = Approved
	} else {
		newStatus = Rejected
	}

	i.Status = newStatus
	return nil
}

func (i *Invoice) UpdateStatus(newStatus Status) error {
	if i.Status != Pending {
		return ErrInvalidStatus
	}

	i.Status = newStatus
	i.UpdatedAt = time.Now()
	return nil
}
