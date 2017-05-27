package dac

import "time"

// AudioInterface constructs a new input callback
type AudioInterface func(sampleRate, bufferSize, inChannels, outChannels int) (AudioCallback, error)

var audioInterface AudioInterface

// AudioCallback will be executed at every sample buffer
type AudioCallback func(in, out *Buffer) error

// DAC is a digital audio converter. It provides an interface for reading and writing audio samples
type DAC struct {
	callbacks      []AudioCallback
	onSample       AudioCallback
	sampleRate     int
	inputChannels  int
	outputChannels int
}

func (d *DAC) MainLoop() {
	i := GetBuffer(d.inputChannels)
	o := GetBuffer(d.outputChannels)
	for range time.Tick(time.Second / time.Duration(d.sampleRate*gBufferSize)) {
		err := d.onSample(i, o)
		if err != nil {
			panic(err)
		}

		for _, c := range d.callbacks {
			err = c(i, o)
			if err != nil {
				panic(err)
			}
		}

	}
}

// OnAudio adds a new input callback
func (d *DAC) OnAudio(a AudioCallback) {
	d.callbacks = append(d.callbacks, a)
}

// NewDAC connects to the given dac and opens given chennels
func NewDAC(sampleRate, bufferSize, inputChannels, outputChannels int) (d *DAC, err error) {
	d = new(DAC)
	d.inputChannels = inputChannels
	d.outputChannels = outputChannels
	SetBufferSize(bufferSize)
	SetSampleRate(sampleRate)
	d.sampleRate = sampleRate
	d.onSample, err = audioInterface(sampleRate, bufferSize, inputChannels, outputChannels)

	return
}
