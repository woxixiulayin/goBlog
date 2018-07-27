// @flow
import qs from 'qs'
import axios from 'axios'

const log = console.log
const defaultTimeout = 30000

/**
 * 设置ajax请求的header
 */
const instance = axios.create({
    headers: {
        // 'X-Requested-With': 'XMLHttpRequest',
        'Content-Type': 'application/x-www-form-urlencoded',
    },
    timeout: defaultTimeout,
})

/**
 * 1.定义底层ajax请求使用的库，可以用jQuery、axios、fetch等，在此处封装
 * 2.全局拦截
 * @param {object} ajax请求库
 * @return {object} ajax请求代理
 */
const request = (($) => {

    let alerting = false

       // 身份未验证时的默认操作
       const onNotVerified = param => {
        if (!alerting) {
            alerting = true
            const msg = '身份过期，请刷新页面重新登录'
            alert(msg)
            log.error(param)
            window.location.reload()
        }
    }

    function ajax(method, url, data = {}, timeout = defaultTimeout) {

        log(`${method} "${url}" with data ==>`, data)

        const postData = { ...data }

        return $({
            url,
            method,
            // 'application/x-www-form-urlencoded'需要使用qs
            data: qs.stringify(postData),
            timeout
        })
            .then((response) => {
                log("response is ", response, response.data)
                const code = Number(response.data.code)
                if (code) {  //  存在code码表示是符合有课规定的请求
                    switch (code) {
                    case 200:
                        return response.data.data
                    // 身份未验证
                    case 401:
                        return onNotVerified({
                            url,
                            method,
                            code: 401,
                            param: JSON.stringify(postData)
                        })
                    default: {
                        const err = new Error(`${method} ${url} 请求失败: ${JSON.stringify(response.data)}`)
                        // 将错误码和信息加入err中
                        err.msg = response.data.msg
                        err.code = Number(response.data.code)
                        throw err
                    }
                    }
                } else { // 不存在code码表示不属于有课规定的请求，直接返回数据
                    log('response is --->', response)
                    return response.data
                }
            })
            .catch((error) => {
                error.url = url
                error.method = method
                error.param = JSON.stringify(postData)
                throw error
            })

    }

    // notVerifiedUrl:当身份未验证时的跳转地址
    function get(url, postData, timeout) {
        if (postData && Object.keys(postData).length !== 0) {
            log.error('注意：get方法无法通过data对象传递请求参数！', url)
        }
        return ajax('get', url, postData, timeout)
    }

    function post(url, postData, timeout) {
        return ajax('post', url, postData, timeout)
    }

    return {
        post,
        get,
    }
})(instance)

export default request
