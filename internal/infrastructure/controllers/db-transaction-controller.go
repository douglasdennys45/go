package controllers

type DBTransaction interface {
	Begin() error
	Commit() error
	Rollback() error
	Close() error
}

type DBTransactionController struct {
	Controller
	DBTransaction
}
