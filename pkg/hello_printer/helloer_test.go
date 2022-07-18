package hello_printer_test

import (
	"testing"

	"github.com/yngvark.com/go-cmd-template/pkg/hello_printer"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		expect string
	}{
		{
			name:   "Should work",
			expect: "Hello",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expect, hello_printer.Hello())
		})
	}
}
