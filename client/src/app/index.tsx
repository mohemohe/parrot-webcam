import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "mobx-react";
import Router, { RouteInfo } from "./containers/Router";
import Store from "./stores";
import Top from "./containers/Top";
import Login from "./containers/Login";
import Image from "./containers/Image";

const routes: RouteInfo[] = [
    {
        path: "/",
        component: Top,
    },
    {
        path: "/login",
        component: Login,
    },
    {
        path: "/image",
        component: Image,
    },
];

export function render() {
    ReactDOM.render(
        <Provider {...Store}>
            <Router routes={routes} />
        </Provider>,
        document.querySelector("#app")
    );
}
