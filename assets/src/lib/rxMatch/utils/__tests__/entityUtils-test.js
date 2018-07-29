import { setEntities, deleteEntities, saveEntities } from '../entityUtils'

describe('test entity utils', () => {

    let entities = {}

    beforeEach(() => {
        entities = {
            1: { id: 1, name: 'one' }
        }
    })

   it('test set', () => {
       const entityList = [
           {id: 2, name: 'two'},
           {id: 3, name: 'three'},
       ]
       const newEntity = setEntities(entityList)
       expect(newEntity[2].name === 'two').toBe(true)
       expect(newEntity[3].name === 'three').toBe(true)
       
    })
    
    it('test save', () => {
        const four = { id: 4, name: 'four'}
        const five = { id: 5, name: 'five'}

        let newEntity = saveEntities(entities, four)
        expect(newEntity[4].name === 'four').toBe(true)
        
        newEntity = saveEntities(entities, [four, five])
        expect(newEntity[4].name === 'four').toBe(true)
        expect(newEntity[5].name === 'five').toBe(true)
    })
    
    it('test delete', () => {
        const entities = {
            1: {id: 1, name: 'one'},
            2: {id: 2, name: 'two'},
        }
        
        let newEntity = deleteEntities(entities, 1)
        expect(typeof newEntity[1] === 'undefined').toBe(true)
        
        newEntity = deleteEntities(entities, [1, 2])
        expect(typeof newEntity[1] === 'undefined').toBe(true)
        expect(typeof newEntity[2] === 'undefined').toBe(true)
    })
})
