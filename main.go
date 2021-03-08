package stringertest

import "github.com/jackc/pgx"

// TxStatus mirrors pgx.TxStatus for printing.
type TxStatus int

//go:generate stringer -type TxStatus
const (
	TxStatusInFailure       TxStatus = pgx.TxStatusInFailure
	TxStatusRollbackFailure TxStatus = pgx.TxStatusRollbackFailure
	TxStatusCommitFailure   TxStatus = pgx.TxStatusCommitFailure
	TxStatusInProgress      TxStatus = pgx.TxStatusInProgress
	TxStatusCommitSuccess   TxStatus = pgx.TxStatusCommitSuccess
	TxStatusRollbackSuccess TxStatus = pgx.TxStatusRollbackSuccess
)
