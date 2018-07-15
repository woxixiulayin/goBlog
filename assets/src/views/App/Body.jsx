import React, { Component } from 'react'
import { EntityModule, bindRxObservable } from 'src/lib/Module'

const todos = new EntityModule({
    state: {
    }
})

window.todos = todos
@bindRxObservable(todos.$$state, {
    todo: (state, ownProps) => state[ownProps.id] || {}
})
class Todos extends Component {
    render() {
        const { todo = {} } = this.props
        return <div>
        {todo.title}
        </div>
    }
}

const Body = props => <section id="body">
    <div className="content">
        <Todos id={1} />
        {props.children || ''}
    </div>
</section>

export default Body
