[%bs.raw {|require('./css/app.scss')|}]

let component = ReasonReact.statelessComponent("App")

let make = (_children) => {
    ...component,
    render: _self => <div className="app">
        <Header />
        <Body />
        <Bottom />
    </div>
}