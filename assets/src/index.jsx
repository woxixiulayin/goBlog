import React from 'react'
import ReactDOM from 'react-dom'
import Router from './routes'
import store from './store'

import 'css/index.scss'

window.store = store

ReactDOM.render(
    <Router />,
    document.querySelector('#app')
)



