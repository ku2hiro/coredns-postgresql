package coredns_postgresql

import (
	"testing"
)

func TestFindRecordTypeA(t *testing.T) {
	p := Postgresql{
		user:     "coredns",
		password: "coredns",
		database: "coredns",
	}
	addrs, err := p.findRecordTypeA("example.com.")
	if err != nil {
		t.Fail()
	}
	for _, a := range addrs {
		if a != "1.1.1.1" {
			t.Fail()
		}
	}
}
