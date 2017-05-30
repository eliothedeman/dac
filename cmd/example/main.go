package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/eliothedeman/dac"
)

var cpuprofile = flag.String("cpu", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		go func() {
			<-time.After(10 * time.Second)
			pprof.StopCPUProfile()
			f.Sync()
			os.Exit(0)

		}()
	}

	d, err := dac.NewDAC(44100, 128, 2, 2)
	if err != nil {
		log.Fatal(err)
	}

	d.OnAudio(func(i, o *dac.Buffer) error {
		o.Clear()
		o.Add(i)
		return nil
	})

	d.MainLoop()
}
