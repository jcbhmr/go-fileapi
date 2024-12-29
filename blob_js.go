package fileapi

type Blob interface {
	Size() uint64
	Type() string
	Slice(start *int64, end *int64, contentType *string) Blob
	Stream() streams.ReadableStream
	Text() string
	ArrayBuffer() []byte
	Bytes() []byte
}