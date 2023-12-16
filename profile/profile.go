package profile

import "go-scraper/action"

type Profile struct {
	version  string
	wakeUp   []map[string]string
	function action.Function
}

func NewProfile() *Profile {
	return &Profile{}
}
