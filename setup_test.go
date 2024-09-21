package coredns_postgresql

import (
	"testing"

	"github.com/coredns/caddy"
)

func TestSetup(t *testing.T) {
	c := caddy.NewTestController("dns", `postgresql`)
	if err := setup(c); err == nil {
		t.Fatalf("Expected errors, but got: %v", err)
	}

	c = caddy.NewTestController("dns", `postgresql dsn {
  user coredns
  password coredns
  database coredns
}`)

	if err := setup(c); err != nil {
		t.Fatalf("Expected no errors, but got: %v", err)
	}
}
