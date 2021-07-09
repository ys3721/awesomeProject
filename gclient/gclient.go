package gclient

import (
	"fmt"
	"git.hortorgames.com/orange/mandarin/v3/bon"
	"git.hortorgames.com/orange/mandarin/v3/reflectex"
	"git.hortorgames.com/orange/mandarin/v3/utils"
	"reflect"
)

func main()  {
	SendAuth()
}

func SendAuth() {
	ul := "Login_AuthUser"
	p:= "{\"_RoleId\":\"\",\"_ServerId\":\"1\",\"platform\":\"debug\",\"info\":\"{\\\"id\\\":\\\"yaotest1\\\"}\",\"scene\":\"1\",\"referrerInfo\":\"\\\"1\\\"\"}"
	ti := reflectex.TypeInfoOfName(ul)
	cmd := reflect.New(ti.Type).Interface()
	utils.JsonDecodeString(p, cmd)
	fmt.Println(cmd)
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
