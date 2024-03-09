package mocks

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// TxMock ...
type TxMock struct {
	pgxpool.Tx
}

// Commit ...
func (t *TxMock) Commit(_ context.Context) error {
	return nil
}

// Rollback ...
func (t *TxMock) Rollback(_ context.Context) error {
	return nil
}

// BeginTx ...
func (t *TxMock) BeginTx(_ context.Context) error {
	return nil
}
