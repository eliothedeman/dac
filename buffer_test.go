package dac

import (
	"reflect"
	"testing"
)

func Test_initCache(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initCache()
		})
	}
}

func TestSetBufferSize(t *testing.T) {
	type args struct {
		s int
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetBufferSize(tt.args.s)
		})
	}
}

func TestSetSampleRate(t *testing.T) {
	type args struct {
		r int
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetSampleRate(tt.args.r)
		})
	}
}

func TestBuffer_Len(t *testing.T) {
	type fields struct {
		SampleRate int
		Samples    [][]float64
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				SampleRate: tt.fields.SampleRate,
				Samples:    tt.fields.Samples,
			}
			if got := b.Len(); got != tt.want {
				t.Errorf("Buffer.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_Channels(t *testing.T) {
	type fields struct {
		SampleRate int
		Samples    [][]float64
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				SampleRate: tt.fields.SampleRate,
				Samples:    tt.fields.Samples,
			}
			if got := b.Channels(); got != tt.want {
				t.Errorf("Buffer.Channels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_GetSample(t *testing.T) {
	type fields struct {
		SampleRate int
		Samples    [][]float64
	}
	type args struct {
		channel int
		index   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				SampleRate: tt.fields.SampleRate,
				Samples:    tt.fields.Samples,
			}
			if got := b.GetSample(tt.args.channel, tt.args.index); got != tt.want {
				t.Errorf("Buffer.GetSample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_SetSample(t *testing.T) {
	type fields struct {
		SampleRate int
		Samples    [][]float64
	}
	type args struct {
		channel int
		index   int
		sample  float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				SampleRate: tt.fields.SampleRate,
				Samples:    tt.fields.Samples,
			}
			b.SetSample(tt.args.channel, tt.args.index, tt.args.sample)
		})
	}
}

func TestBuffer_Copy(t *testing.T) {
	type fields struct {
		SampleRate int
		Samples    [][]float64
	}
	tests := []struct {
		name   string
		fields fields
		want   *Buffer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				SampleRate: tt.fields.SampleRate,
				Samples:    tt.fields.Samples,
			}
			if got := b.Copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Buffer.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_Clear(t *testing.T) {
	type fields struct {
		SampleRate int
		Samples    [][]float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				SampleRate: tt.fields.SampleRate,
				Samples:    tt.fields.Samples,
			}
			b.Clear()
		})
	}
}

func TestBuffer_Add(t *testing.T) {
	type fields struct {
		SampleRate int
		Samples    [][]float64
	}
	type args struct {
		n *Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				SampleRate: tt.fields.SampleRate,
				Samples:    tt.fields.Samples,
			}
			b.Add(tt.args.n)
		})
	}
}

func TestBuffer_Multiply(t *testing.T) {
	type fields struct {
		SampleRate int
		Samples    [][]float64
	}
	type args struct {
		n *Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				SampleRate: tt.fields.SampleRate,
				Samples:    tt.fields.Samples,
			}
			b.Multiply(tt.args.n)
		})
	}
}

func TestGetBuffer(t *testing.T) {
	type args struct {
		channels int
	}
	tests := []struct {
		name string
		args args
		want *Buffer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBuffer(tt.args.channels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFreeBuffer(t *testing.T) {
	type args struct {
		b *Buffer
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FreeBuffer(tt.args.b)
		})
	}
}
