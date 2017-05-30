package dac

// AudioInterface constructs a new input callback
type AudioInterface func(sampleRate, bufferSize, inChannels, outChannels int) (AudioInterfaceCallback, error)

// AudioInterfaceCallback Will run on the audio interface
type AudioInterfaceCallback func(AudioCallback) error

var audioInterface AudioInterface

// AudioCallback will be executed at every sample buffer
type AudioCallback func(in, out *Buffer) error

// DAC is a digital audio converter. It provides an interface for reading and writing audio samples
type DAC struct {
	callbacks      []AudioCallback
	onSample       AudioInterfaceCallback
	sampleRate     int
	inputChannels  int
	outputChannels int
}

func (d *DAC) MainLoop() {
	var err error
	for {
		err = d.onSample(func(i, o *Buffer) error {
			for _, c := range d.callbacks {
				err := c(i, o)
				if err != nil {
					panic(err)
				}
			}
			return nil
		})
		if err != nil {
			panic(err)
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
