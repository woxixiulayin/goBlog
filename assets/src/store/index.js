import models from 'src/models'
import { init } from '@rematch/core'

const store = init({
    models
})

window.store = store

export default store
