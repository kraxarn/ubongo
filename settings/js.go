//go:build js && wasm

package settings

import "syscall/js"

func localStorage() js.Value {
	return js.Global().Get("localstorage")
}

func get(key string) string {
	value := localStorage().Call("getItem", key)
	if value.IsNull() {
		return ""
	}
	return value.String()
}

func set(key, value string) {
	localStorage().Call("setItem", key, value)
}
