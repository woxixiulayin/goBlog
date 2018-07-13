package assets

import (
    "io/ioutil"
    "encoding/json"
    "goBlog/modules/log"
)

var jsPath = "./assets/assets.json"

// 从assets.json中获取js地址
func GetJsFiles() map[string]string {

    // js文件解析后的数据
    var assetsMap map[string]interface
    var jsFiles map[string]string

    data, err := ioutil.ReadFile(jsPath)

    log.Debugf("js files data is %v", string(data))

    if err != nil {
       log.Debugf("read file %v error: %v", jsPath, err) 
       return nil
    }
    
    if err := json.Unmarshal(data, &assetsMap); err != nil {
        log.Debugf("Unmarshal file %v error: %v", jsPath, err) 
        return nil
    }

    return assetsMap
}
