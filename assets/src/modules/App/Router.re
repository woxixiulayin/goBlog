type router = {
    name: string,
    path: string
};

let routers: array(router) = [|
    {
        name: "首页",
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