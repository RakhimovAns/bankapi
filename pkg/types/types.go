package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д.)
type Money int64

// Currency представляет код валюты
type Currency string
type Catergory string
// Коды валют
const (
  TJS Currency = "TJS"
  RUB Currency = "RUB"
  USD Currency = "USD"
  EUR Currency = "EUR"
)

// PAN представляет номер карты
type PAN string

// Card представляет информацию о платёжной карте
type Card struct {
  ID         int
  PAN        PAN
  Balance    Money // использовали Money
  MinBalance Money // использовали Money
  Currency   Currency
  Color      string
  Name       string
  Active     bool
}
type Payment struct{
	ID int 
	Amount Money
	Catergory Catergory
}