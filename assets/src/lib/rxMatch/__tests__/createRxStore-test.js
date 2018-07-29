import creteRxStore from '../createRxStore'

describe('test createRxStore', () => {

    let store

    beforeEach(() => {
        store = creteRxStore()
    })

    it('test methods', () => {
        const methods = ['dispatch', 'getState', 'subscribe']
        const store = creteRxStore()

        methods.forEach(item => expect(typeof store[item] !== 'undefined').toBe(true))
    })

})