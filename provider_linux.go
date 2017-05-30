package dac

import (
	"encoding/binary"
	"fmt"
	"math"
	"os/exec"
)

func init() {
	audioInterface = newLinuxAudioInterface
}

func newLinuxAudioInterface(sampleRate, bufferSize, in, out int) (AudioInterfaceCallback, error) {

	aplay := exec.Command("aplay", "-t", "raw", "-f", "S32_LE", "-r", fmt.Sprint(sampleRate), "-c", fmt.Sprint(out))
	arecord := exec.Command("arecord", "-t", "raw", "-f", "S32_LE", "-r", fmt.Sprint(sampleRate), "-c", fmt.Sprint(in))
	stdout, err := arecord.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stdin, err := aplay.StdinPipe()
	if err != nil {
		return nil, err
	}

	// ibuff := make([]byte, bufferSize*in*4)
	obuff := make([]byte, bufferSize*out*4)
	ibuff := make([]byte, bufferSize*in*4)

	start := make(chan error)
	go func() {
		err := aplay.Start()
		start <- err
	}()
	err = <-start
	if err != nil {
		return nil, err
	}

	go func() {
		err := arecord.Start()
		start <- err
	}()
	err = <-start
	if err != nil {
		return nil, err
	}

	i := GetBuffer(in)
	o := GetBuffer(out)

	close(start)
	return func(f AudioCallback) error {
		var err error
		_, err = stdout.Read(ibuff)
		if err != nil {
			return err
		}
		fillInputBuffer(i, ibuff)
		o.Clear()
		err = f(i, o)
		if err != nil {
			return err
		}

		_, err = stdin.Write(obuff)
		fillOutputBuffer(o, obuff)
		return err
	}, nil
}

func fillInputBuffer(b *Buffer, buff []byte) {
	var x int

	for i := 0; i < b.Len(); i++ {
		for c := range b.Samples {
			b.SetSample(c, i, float64(binary.LittleEndian.Uint32(buff[x:x+4]))/math.MaxInt32)
			x += 4
		}
	}
}

func fillOutputBuffer(o *Buffer, buff []byte) {
	var x int

	for i := 0; i < o.Len(); i++ {
		for c := range o.Samples {
			binary.LittleEndian.PutUint32(buff[x:x+4], uint32(o.GetSample(c, i)*math.MaxInt32))
			x += 4
		}
	}
}
