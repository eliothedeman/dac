package gen

import (
	"github.com/eliothedeman/dac"
)

// Gen is used to fill a buffer. They should not be shared.
type Gen interface {
	Fill(b *dac.Buffer)
}
