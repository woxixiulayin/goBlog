type state = {
    currentRouter: string,
}

let component = ReasonReact.reducerComponent("Header");

type router = {
    name: string,
    path: string
};

let routers: array(router) = [|
    {
        name: "home",
        path: "home",
    },
    {
        name: "about",
        path: "home",
    },
    {
        name: "github",
        path: "github",
    },
|]

let make = (_children) => {
    ...component,
    reducer: (action, _state) => ReasonReact.Update({ currentRouter: action }),
    initialState: () => {currentRouter: "home"},
    render: (self) => {
        let links = routers
        |> Array.map(item => {
            <a
              href={"#" ++ item.path}
              className={self.state.currentRouter == item.name ? "current" : ""}
              onClick={(_) => self.send(item.name)}
            >
                {ReasonReact.string(item.name)}
            </a>
        })
        |> ReasonReact.array;

        <header
          className="text-center relative"
        >
            <h2>{ReasonReact.string("Jackson Liu's house")}</h2>
            <nav className="nav clearfix">
           <div className="links">
            links
           </div>
           </nav>
        </header>
    }
}
