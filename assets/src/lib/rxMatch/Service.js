// @flow
import { BehaviorSubject } from 'rxjs'

class Service <A> {
    // 内部保存的数据
    state: any
    // 提供给外部的订阅数据，通过this.state$.next(state)通知外部更新数据
    state$: Observable

    constructor({ state }: {
        state: A
    }) {
        this.state = {}
        // 使用BehaviorSubject，使得stae$可以被各处订阅
        this.state$ = new BehaviorSubject()
        this._updateState(state)
    }

    getState(): A {
        return this.state
    }

    setState(state: any) {
        const newState = {
            ...this.getState(),
            ...state
        }
        this._updateState(newState)
    }

    subscribe(...args) {
        return this.state$.subscribe(...args)
    }

    _updateState(newState: A) {
        this.state = newState
        this.state$.next(newState)
    }
}

export default Service
