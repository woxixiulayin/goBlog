package assets

import (
    "io/ioutil"
    "encoding/json"
    "goBlog/modules/log"
)

var jsPath = "./assets/assets.json"

// 对应assets.home.js，
type JsFile struct {
    Js string `json:"js"`
}

// 使用tag指明键值对应的jsonkey
// struct中的键值要大写，外部才能使用
type AssetsMapType struct {
    Home JsFile `json:"home"`
}

type JsFiles struct {
    Home string
}

// 从assets.json中获取js地址
func GetJsFiles() JsFiles {

    // json文件格式如下
    // {\"home\":{\"js\":\"/Users/Jackson/go/src/goBlog/assets/build/index.427672d296ac31fd9608.js\"}}"}
    var assetsMap AssetsMapType
    var jsFiles JsFiles

    data, err := ioutil.ReadFile(jsPath)

    log.Debugf("js files data is %v", string(data))

    if err != nil {
       log.Debugf("read file %v error: %v", jsPath, err) 
       return jsFiles
    }
    
    if err := json.Unmarshal(data, &assetsMap); err != nil {
        log.Debugf("Unmarshal file %v error: %v", jsPath, err) 
        return jsFiles
    }

    jsFiles.Home = assetsMap.Home.Js

    return jsFiles
}
