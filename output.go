package dac

// Output is an analog to digital converter
type Output struct {
	channels []*Buffer
}

// Read audio from a given channel
func (o *Output) Read(channel int) *Buffer {
	return o.channels[channel].Copy()
}
