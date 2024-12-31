package common

import goerrors "errors"

// ErrNotFinalized is returned when a transaction is not finalized
var ErrNotFinalized = goerrors.New("not finalized")

// ErrTxFailed is returned when a transaction has failed
var ErrTxFailed = goerrors.New("transaction failed")

// ErrFailedToGetTransactions is returned when a transaction is not found
var ErrFailedToGetTransactions = goerrors.New("failed to get transactions")
