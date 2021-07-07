package gclient

import (
	"git.hortorgames.com/orange/mandarin/v3/bon"
)


func SendAuth() {
	//url := "http://10.1.1.58:10101"
	//http.Post()
}

func SerializeToBytes(v interface{}) []byte {
	bonData := SerializeToBon(v)
	bytes := bonData.ToBonBytes()
	bon.Release(bonData)
	return bytes
}

func SerializeToBon(v interface{}) bon.IValue {
	return bon.FromObject(v)
}
