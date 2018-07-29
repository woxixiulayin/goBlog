//  @flow
import Service from './Service'
import { combineLatest, BehaviorSubject, of } from 'rxjs'
import assert from 'src/lib/assert'

type typeServices = {
    [name: string]: Service<any>
}

type typeInitParam = {
    services: typeServices,
}

export default function createRxStore({
    services,
}: typeInitParam = {
    services: {}
}) {
    let state = null
    const state$ = new BehaviorSubject()
    const serviceState$List = Object.values(services).map(service => service.state$)
    const serviceKeys = Object.keys(services)

    log('serviceList is', serviceState$List, '\nserviceKeys is', serviceKeys)
    const dispatch = serviceKeys.reduce((dispatch, serviceName) => {
        assert(typeof dispatch[serviceName] === 'undefined', `service name '${serviceName}' has already been used`)

        const service = services[serviceName]

        dispatch[serviceName] = Object.getOwnPropertyNames(service).reduce((pureService, method) => {

            if (typeof service[method] !== 'function') {
                return pureService
            }

            log('pureService is {', pureService, `} \nmethod is ${method}`)

            pureService[method] = function (...args) {
                log(`dispatch service [${serviceName}] width param`, ...args)
                return service[method].call(service, ...args)
            }

            return pureService
        }, {})

        return dispatch
    }, {})

    const initService = of('@@INIT')

    // 订阅所有services
    combineLatest(...serviceState$List, initService)
    .subscribe(stateLists => {
        log('new state Lists is ', stateLists)

        const newState = serviceKeys.reduce((state, key, index) => {
            state[key] = stateLists[index]
            return state
        }, {})

        state = newState
        // 触发订阅更新
        log('new state$ is ', newState)
        state$.next(state)
    })

    

    const getState = (): any => {
        return state
    }

    const subscribe = (...args: any): Function => {
        return state$.subscribe(...args)
    }

    return {
        getState,
        subscribe,
        dispatch,
    }
}