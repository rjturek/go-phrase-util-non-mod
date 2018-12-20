package rjtphrasenonmod

import "github.com/golang/example/stringutil"

var phrase = "Far away from night-clubs and from saxophones."

func GetPhrase() string {
	return stringutil.Reverse(stringutil.Reverse(phrase))
}
