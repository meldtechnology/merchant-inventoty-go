package util

import (
	"errors"
)

type TRANSACTION_TYPE int
type STATUS_TYPE int

const (
	PURCHASE_ORDER TRANSACTION_TYPE = iota + 1
	SALES_ORDER
)

const (
	PENDING STATUS_TYPE = iota + 1
	APPROVED
	REJECTED
	PARTIALLY_FULFILLED
	FULFILLED
)

var TransactionName = map[TRANSACTION_TYPE]string{
	PURCHASE_ORDER: "PURCHASE_ORDER",
	SALES_ORDER:    "SALES_ORDER",
}

var StatusName = map[STATUS_TYPE]string{
	PENDING:             "PENDING",
	APPROVED:            "APPROVED",
	REJECTED:            "REJECTED",
	PARTIALLY_FULFILLED: "PARTIALLY_FULFILLED",
	FULFILLED:           "FULFILLED",
}

var GetStatusName = map[string]STATUS_TYPE{
	"PENDING":             PENDING,
	"APPROVED":            APPROVED,
	"REJECTED":            REJECTED,
	"PARTIALLY_FULFILLED": PARTIALLY_FULFILLED,
	"FULFILLED":           FULFILLED,
}

func ValidateStatus(value string) error {
	FOUND := GetStatusName[value]

	if FOUND == 0 {
		return errors.New(value + " is not a valid Status")
	}
	return nil
}
