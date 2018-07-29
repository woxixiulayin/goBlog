//  @flow
import Service from './Service'
import { combineLatest, BehaviorSubject } from 'rxjs'
import assert from 'src/lib/assert'

type typeServices = {
    [name: string]: Service
}

type typeInitParam = {
    services: typeServices,
}

export default function createRxStore({
    services,
}: typeInitParam = {
    services: {}
}): {
    getState: Function,
    subscribe: Function,
    dispatch: Object,
} {
    let state = null
    const state$ = new BehaviorSubject()
    const serviceList = Object.values(services)
    
    const dispatch = Object.keys(services).reduce((dispatch, item) => {
        assert(typeof dispatch === 'undefined', `service name '${item}' has already been used`)

        const service = services[item]

        dispatch[item] = Object.getOwnPropertyNames(service).reduce((pureService, method) => {
            pureService[method] = service[method].bind(service)
            return pureService
        })

        return dispatch
    }, {})
    

    // 订阅所有services
    combineLatest(...serviceList)
    .subscribe(newState => {
        state = newState
        state$.next(state)
    })


    const getState = (...args: any): any => {
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