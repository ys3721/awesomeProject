package config

import (
    "io/ioutil"
    "encoding/json"
)
    
var tableURL = "./"
    
var tables Configs
var maps Maps

func GetTables() *Configs {
    return &tables
}

func GetMaps() *Maps {
    return &maps
}
    

    
type Configs struct {
}

type Maps struct {
}
    
func DecodeAll(url string) (err error) {
    if(url[len(url)-1] != '/') {
        url = url + "/"
    }
    tableURL = url
    tables = Configs{}
    return 
}

func DecodeAllByPath() (err error) {
    return DecodeAll("./out/go/config")
}
    
func Load(filename string, v interface{})(err error) {
    //ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return
    }
    //读取的数据为json格式，需要进行解码
    err = json.Unmarshal(data, v)
    if err != nil {
        return
    }
    return
}


