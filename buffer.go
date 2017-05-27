package dac

import (
	"sync"
)

// bufferSize defines the number of sampels a buffer will hold
var gBufferSize = 512
var gSampleRate = 44100
var bufferCache sync.Pool

func init() {
	initCache()
}

func initCache() {
	bufferCache = sync.Pool{
		New: func() interface{} {
			s := make([]float64, gBufferSize)
			return s
		},
	}

	for i := 0; i < 32; i++ {
		bufferCache.Put(bufferCache.New())

	}
}

// SetBufferSize initializes the buffer cache with sample size
func SetBufferSize(s int) {
	gBufferSize = s
	initCache()
}

// SetSampleRate sets the sample rate of buffers and re initilizes the cache
func SetSampleRate(r int) {
	gSampleRate = r
	initCache()
}

// Buffer holds audio samples
type Buffer struct {
	SampleRate int         // Samples per second
	Samples    [][]float64 // PCM data
}

// Len returns the number of samples in the bufffer, per channel
func (b *Buffer) Len() int {
	if len(b.Samples) == 0 {
		return 0
	}

	return len(b.Samples[0])
}

// Channels returns the number of channels
func (b *Buffer) Channels() int {
	return len(b.Samples)
}

// GetSample returns the sample at the given index
func (b *Buffer) GetSample(channel, index int) float64 {
	return b.Samples[channel][index]
}

// SetSample returns the sample at the given index
func (b *Buffer) SetSample(channel, index int, sample float64) {
	b.Samples[channel][index] = sample
}

// Copy the buffer
func (b *Buffer) Copy() *Buffer {
	x := GetBuffer(b.Channels())
	for i := 0; i < x.Channels(); i++ {
		for j := 0; j < x.Len(); j++ {
			x.SetSample(i, j, b.GetSample(i, j))
		}
	}
	x.SampleRate = b.SampleRate
	return x
}

// Clear the buffer to all 0s
func (b *Buffer) Clear() {
	for i := 0; i < b.Channels(); i++ {
		for j := 0; j < b.Len(); j++ {
			b.SetSample(i, j, 0)
		}
	}
}

// Add two buffers
func (b *Buffer) Add(n *Buffer) {
	for i := 0; i < b.Channels(); i++ {
		for j := 0; j < b.Len(); j++ {
			b.SetSample(i, j, b.GetSample(i, j)+n.GetSample(i, j))
		}
	}
}

// Multiply two buffers
func (b *Buffer) Multiply(n *Buffer) {
	for i := 0; i < b.Channels(); i++ {
		for j := 0; j < b.Len(); j++ {
			b.SetSample(i, j, b.GetSample(i, j)*n.GetSample(i, j))
		}
	}
}

// GetBuffer will pull a new buffer from the cache
func GetBuffer(channels int) *Buffer {
	b := &Buffer{
		SampleRate: gSampleRate,
	}
	b.Samples = make([][]float64, channels)
	for i := range b.Samples {
		b.Samples[i] = bufferCache.Get().([]float64)
	}

	return b
}

// FreeBuffer will pull a new buffer from the cache
func FreeBuffer(b *Buffer) {
	bufferCache.Put(b)
}
