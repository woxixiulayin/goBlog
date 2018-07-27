/**@flow */
import * as React from 'react'
import assert from 'src/lib/assert'

type EntityType = {
    id: string,
    [key: string]: any
}

type EntityStateType = {
    [id: string]: EntityType
}

type StateOrPropsType = {
    [key: string]: any
}

const setEntities = (entities: Array<EntityType>): EntityStateType => {
    assert(Array.isArray(entities), 'entities should be an array')
    return entities.reduce((state, item) => {
    if (typeof item.id !== 'undefined') {
        state[item.id] = item
    }
    return state
}, {})
}

const saveEntities = (state: EntityStateType, entities: Array<EntityType> | EntityType): EntityStateType => {

    assert(Array.isArray(entities) || typeof entities !== 'undefined', 'entities should be an array or not undefined')

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

const deleteEntities = (state: EntityStateType, entityIds: Array<string> | string): EntityStateType => {

    assert(Array.isArray(entityIds) || typeof entityIds !== 'undefined', 'entityIds should be an array or not undefined')

    let list = Array.isArray(entityIds) ? entityIds : [entityIds]

    return list.reduce((state, id) => {
        state[id] && delete state[id]
        return state
    }, {...state})
}

export default {
    setEntities,
    saveEntities,
    deleteEntities,
}
