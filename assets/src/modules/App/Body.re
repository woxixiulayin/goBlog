let str = ReasonReact.string
let component = ReasonReact.statelessComponent("Body")

let make = (_children) => {
    ...component,
    render: _self => <div id="body">
        <div>{str("this is a body")}</div>
    </div>,
}