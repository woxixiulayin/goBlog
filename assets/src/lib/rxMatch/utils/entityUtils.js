// @flow
import assert from 'src/lib/assert'
import { EntityType, EntityStateType } from '../types'

export const setEntities = (entities: Array<EntityType>): EntityStateType => {
    assert(Array.isArray(entities), 'entities should be an array')
    return entities.reduce((state, item) => {
        const newState = { ...state }
        if (typeof item.id !== 'undefined') {
            newState[item.id] = item
        }
        return newState
    }, {})
}

export const saveEntities = (state: EntityStateType, entities: Array<EntityType> | EntityType): EntityStateType => {

    assert(Array.isArray(entities) || typeof entities !== 'undefined', 'entities should be an array or not undefined')

    let list = Array.isArray(entities) ? entities : [entities]

    return list.reduce((state, item) => {
        const newState = { ...state }
        if (typeof item.id !== 'undefined') {
            newState[item.id] = state[item.id]
            ? { ...state[item.id], ...item }
            : item
        }
        return newState
    }, { ...state })
}

export const deleteEntities = (state: EntityStateType, entityIds: Array<number> | number): EntityStateType => {

    assert(Array.isArray(entityIds) || typeof entityIds !== 'undefined', 'entityIds should be an array or not undefined')

    let list = Array.isArray(entityIds) ? entityIds : [entityIds]

    return list.reduce((state, id) => {
        const newState = { ...state }
        newState[id] && delete newState[id]
        return newState
    }, { ...state })
}
