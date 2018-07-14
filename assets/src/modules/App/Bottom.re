let str = ReasonReact.string
let component = ReasonReact.statelessComponent("Bottom");

let make = (_children) => {
    ...component,
    render: _self => <section
            id="bottom"
            className="text-center"
        >
            (str("@CopyRight 2017 * Designed & Wrote By Jackson Liu"))
    </section>,
}
