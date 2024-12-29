package fileapi

import "syscall/js"

var url2 = js.Global().Get("URL")

func URLCreateObjectURL(obj any) string {
	switch obj := obj.(type) {
		case Blob:
			return url2.Call("createObjectURL", obj.value).String()
		// case mediasource2.MediaSource:
		default:
			panic("obj must be Blob | mediasource2.MediaSource")
	}
}

func URLRevokeObjectURL(url string) {
	url2.Call("revokeObjectURL", url)
}
