//go:build gofuzz

package coredns_postgresql

import (
	"github.com/coredns/coredns/plugin/pkg/fuzz"
)

// Fuzz fuzzes cache.
func Fuzz(data []byte) int {
	w := Whoami{}
	return fuzz.Do(w, data)
}
