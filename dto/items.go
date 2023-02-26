package dto

type Item struct {
	Name     string
	User     string
	Metadata string
}

type LoginPassword struct {
	Item
	Login    string
	Password string
}

type TextInfo struct {
	Item
	Text string
}

type ItemsList struct {
	Passwords []*LoginPassword
	Texts     []*TextInfo
}
