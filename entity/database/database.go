package database

import "time"

type TblUserProfile struct {
	Username         string    `json:"username"`
	Name             string    `json:"name"`
	BornPlace        string    `json:"born_place"`
	BornDate         time.Time `json:"born_date"`
	MotherMaidenName string    `json:"mother_maiden_name"`
	Address          string    `json:"address"`
	CellphoneNumber  string    `json:"cellphone_number"`
	EmailAddress     string    `json:"email_address"`
	Cif              string    `json:"cif"`
	ResponseId       string    `json:"response_id"`
}

type TblUserAccount struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	Account         string `json:"account"`
	TypeAccount     string `json:"type_account"`
	ProductName     string `json:"product_name"`
	Currency        string `json:"currency"`
	CardNumber      string `json:"card_number"`
	Status          int    `json:"status"`
	FinansialStatus int    `json:"finansial_status"`
	Default         int    `json:"default"`
	ScCode          string `json:"sc_code"`
}

type User struct {
	Client          string `json:"client"`
	AccountNumber   string `json:"account_number"`
	Timestamp       string `json:"timestamp"`
	ResponseRefnum  string `json:"response_refnum"`
	ResponseId      string `json:"response_id"`
	CellphoneNumber string `json:"cellphone_number"`
	EmailAddress    string `json:"email_address"`
	Refnum          string `json:"refnum"`
}
