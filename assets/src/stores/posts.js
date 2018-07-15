/**@flow */
import Rx from 'rxjs'

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

const posts = new Module({
    state: {}
})

posts.$$state.subscribe(posts => console.log(posts))
posts.save([{
    id: 1,
    title: '123',
}])
posts.remove([1])
window.posts = posts