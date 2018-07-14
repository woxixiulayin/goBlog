let component = ReasonReact.statelessComponent("Header")
let str = ReasonReact.string

let make = (_children) => {
    ...component,
    render: _self => <header id="Header">
    (str("this is header"))
    </header>
}