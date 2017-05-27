package gen

import (
	"math"

	"github.com/eliothedeman/dac"
)

// Sine is a sine wave osilator
type Sine struct {
	Freq        float64
	Phase       float64
	sampleIndex int
}

type Val float64

func (v Val) Fill(b *dac.Buffer) {

	for i := 0; i < b.Channels(); i++ {
		for j := 0; j < b.Len(); j++ {
			b.SetSample(i, j, float64(v))
		}
	}
}

// Fill with the current phase of the oscilator
func (s *Sine) Fill(b *dac.Buffer) {
	start := s.sampleIndex
	end := start
	sampelIndex := start

	for i := 0; i < b.Channels(); i++ {
		for j := 0; j < b.Len(); j++ {
			b.SetSample(i, j, math.Sin((float64(sampelIndex)/float64(b.SampleRate))*math.Pi*2*s.Freq))
			sampelIndex++
			if sampelIndex > b.SampleRate {
				sampelIndex = 0
			}
		}
		end = sampelIndex
		sampelIndex = start
	}
	s.sampleIndex = end
}
