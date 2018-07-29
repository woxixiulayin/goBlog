import creteRxStore from '../createRxStore'
import Service from '../Service'

describe('test createRxStore', () => {

    let store,
        serviceA,
        serviceB

    beforeEach(() => {
        
    })

    it.skip('test methods', () => {
        const methods = ['dispatch', 'getState', 'subscribe']
        const store = creteRxStore()

        methods.forEach(item => expect(typeof store[item] !== 'undefined').toBe(true))
    })

    it('test getState', () => {
        serviceA = new Service({ state: 'a'})
        serviceB = new Service({ state: 'b'})
        store = creteRxStore({
            services: {
                serviceA,
                serviceB
            }
        })
        const initState = store.getState()
        expect(initState.serviceA).toBe('a')
        expect(initState.serviceB).toBe('b')
    })
})