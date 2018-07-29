import React from 'react'
import { Provider, connect } from '@rematch/core'
import ReactDOM from 'react-dom'
import Router from './routes'
import store from './store'

import 'css/index.scss'

ReactDOM.render(
    <Router />,
    document.querySelector('#app')
)



