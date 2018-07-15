import React from 'react'
import ReactDOM from 'react-dom'
import Router from './routes'
import './stores/posts'

import 'css/index.scss'

ReactDOM.render(
    <Router />,
    document.querySelector('#app')
)



