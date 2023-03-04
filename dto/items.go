package dto

type Item struct {
	Name     string `db:"name"`
	User     string `db:"user_email"`
	Metadata string `db:"metadata"`
}

type LoginPassword struct {
	Item
	Login    string `db:"login"`
	Password string `db:"password"`
}

type TextInfo struct {
	Item
	Text string `db:"text"`
}

type CardInfo struct {
	Item
	CardNumber      string `db:"card_number"`
	CVV             int32  `db:"cvv"`
	ExpirationMonth int32  `db:"expiration_month"`
	ExpirationYear  int32  `db:"expiration_year"`
}

type ItemsList struct {
	Passwords []*LoginPassword
	Texts     []*TextInfo
	Cards     []*CardInfo
}
