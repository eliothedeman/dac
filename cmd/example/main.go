package main

import (
	"log"

	"github.com/eliothedeman/dac"
	"github.com/eliothedeman/dac/gen"
)

func main() {
	s := &gen.Sine{Freq: 440}
	s1 := &gen.Sine{Freq: 442}
	d, err := dac.NewDAC(44100, 512, 2, 2)
	if err != nil {
		log.Fatal(err)
	}
	c := dac.GetBuffer(2)
	c1 := dac.GetBuffer(2)
	scale := dac.GetBuffer(2)
	gen.Val(0.5).Fill(scale)

	d.OnAudio(func(i, o *dac.Buffer) error {
		o.Clear()
		s.Fill(c)
		s1.Fill(c1)
		c.Multiply(c1)
		c.Multiply(scale)
		o.Add(c)
		return nil
	})

	d.MainLoop()

}
