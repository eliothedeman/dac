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

func newLinuxAudioInterface(sampleRate, bufferSize, in, out int) (AudioCallback, error) {

	p := exec.Command("aplay", "-t", "raw", "-f", "S32_LE", "-r", fmt.Sprint(sampleRate), "-c", fmt.Sprint(out))
	stdin, err := p.StdinPipe()
	if err != nil {
		return nil, err
	}

	// ibuff := make([]byte, bufferSize*in*4)
	obuff := make([]byte, bufferSize*out*4)
	go p.Run()

	return func(i, o *Buffer) error {
		fillBuffer(o, obuff)
		_, err := stdin.Write(obuff)

		return err
	}, nil
}

func fillBuffer(o *Buffer, buff []byte) {
	var x int

	for i := 0; i < o.Len(); i++ {
		for c := range o.Samples {
			binary.LittleEndian.PutUint32(buff[x:x+4], uint32(o.GetSample(c, i)*math.MaxInt32))
			x += 4
		}
	}
}
