/**@flow */
import Rx from 'rxjs'
import * as React from 'react'

type EntityType = {
    id: number
}

type StateType = {
    [id: number]: EntityType
}

const setEntities = (entities: Array<EntityType>): StateType => entities.reduce((state, item) => {
    if (typeof item.id !== 'undefined') {
        state[item.id] = item
    }
    return state
}, {})

const saveEntities = (state: StateType, entities: Array<EntityType> | EntityType): StateType => {

    let list = Array.isArray(entities) ? entities : [entities]

    return list.reduce((state, item) => {
        if (typeof item.id !== 'undefined') {
            state[item.id] = state[item.id]
            ? { ...state[item.id], ...item }
            : item
        }
        return state
    }, {...state})
}

const deleteEntities = (state: StateType, entityIds: Array<number> | number): StateType => {
    let list = Array.isArray(entityIds) ? entityIds : [entityIds]

    return list.reduce((state, id) => {
        state[id] && delete state[id]
        return state
    }, {...state})
}
    const bindRxObservable = (
        $$data: Rx.BehaviorSubject,
        mapState: {[stateKey: string]: Function}
    ) => WrappedComponent => {

    return class extends React.Component<any, {[string]: any}> {

        constructor() {
            super()
            this.state = {}
        }

        componentDidMount() {
            // 记录组件的props，后面可以通过组件的props来选择$data中需要订阅的数据
            const owmProps = this.props
        
            const stateKeys = Object.keys(mapState)
        
            const $$subject = $$data
            .map(data => {
                console.log('get next data', data)
                return data
            })
            .map(data =>
                // 从data中生成对应的组件state
                stateKeys.reduce((state, key) => {
                    const keyData = mapState[key](data, owmProps)
                    if (typeof keyData !== 'undefined') {
                        state[key] = keyData
                    }
                    return state
                }
                , {})
            )
            // 根据本次state对比上次state计算出需要更新的state
            .scan((oldState, newState) => 
                stateKeys.reduce((state, key) => {
                    if (
                        newState[key] !== oldState[key]) {
                        state[key] = newState[key]
                    }
                    return state
                }, {})
            , {})

            // 监听绑定生成的state
            $$subject.subscribe(state => {
                console.log(`component get subscribe state`, state)
                if (Object.getOwnPropertyNames(state).length > 0) {
                    this.setState(state)
                }
            })
        }

        render() {
            return <WrappedComponent {...this.props} {...this.state} />
        }
    }
}

class Module {

    // constructor({
    //     state = {}
    // }) {
    // }
    state: StateType
    $$state: Rx.BehaviorSubject

    constructor({ state = {} } : {
        state : StateType
    }) {
        this.state = {}
        this.$$state = new Rx.BehaviorSubject()
        this._updateState(state)
    }
    
    getState(): StateType {
        return this.state
    }

    _updateState(newState: StateType) {
        this.state = newState
        this.$$state.next(newState)
    }
    
    set(entities: Array<EntityType>) {
        this._updateState(setEntities(entities))
    }
    
    save(entities: Array<EntityType>) {
        this._updateState(saveEntities(this.state, entities))
    }

    remove(entities: Array<number>) {
        this._updateState(deleteEntities(this.state, entities))
    }
}

export default Module
export {
    bindRxObservable
}
