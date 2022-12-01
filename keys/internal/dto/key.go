package dto

type RegisterKey struct {
	Holder      string
	Affiliation string
	Value       string
}

type Key struct {
	Holder      string
	Affiliation string
	Value       string
}

type Keys struct {
	Keys []*Key
}
