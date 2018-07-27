import React, { Component } from 'react'

class Todos extends Component {
    render() {
        const { post = {} } = this.props
        console.log('rerender', post)
        return <div>
        <div>{''}</div>
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
