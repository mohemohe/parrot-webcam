import React from "react";
import { inject, observer } from "mobx-react";
import {
    Route,
    Router as ReactRouter,
    Switch,
    RouteComponentProps,
} from "react-router";
import { createBrowserHistory } from "history";
import MobxReactRouter, {
    RouterStore,
    syncHistoryWithStore,
} from "mobx-react-router";

class NotFound extends React.Component<{}, {}> {
    public render() {
        return <div>404 not found</div>;
    }
}

export interface RouteInfo {
    path?: string;
    component?: any;
    children?: RouteInfo[];
}

export interface Props {
    RouterStore?: RouterStore;

    routes: RouteInfo[];
    notFound?:
        | React.ComponentType<RouteComponentProps<any>>
        | React.ComponentType<any>;
}

interface State {}

@inject("RouterStore")
@observer
export default class Router extends React.Component<Props, State> {
    constructor(props: Props, state: State) {
        super(props, state);

        const history = createBrowserHistory();
        this.history = syncHistoryWithStore(history, this.props.RouterStore!);
        this.pathname = this.history.location.pathname;
        this.history.subscribe((location) => {
            if (location.pathname != this.pathname) {
                this.pathname = location.pathname;
                window.scrollTo(0, 0);
            }
        });

        this.routeArray = [];

        this.generateRoute(props, props.routes);
    }

    public UNSAFE_componentWillReceiveProps(nextProps: Readonly<Props>) {
        this.generateRoute(nextProps, nextProps.routes);
    }

    private history: MobxReactRouter.SynchronizedHistory;
    private pathname: string;
    private routeArray: JSX.Element[];

    protected generateRoute(props: Props, routes: RouteInfo[]) {
        this.routeArray = [];
        this.parseRoute(props, routes);
        this.routeArray.push(
            <Route
                component={props.notFound || NotFound}
                key={this.routeArray.length}
            />
        );
    }

    protected parseRoute(props: Props, routes: RouteInfo[]) {
        routes.forEach((route) => {
            if (route.path) {
                this.routeArray.push(
                    <Route
                        exact
                        path={route.path}
                        component={
                            route.component
                                ? route.component
                                : props.notFound || NotFound
                        }
                        key={this.routeArray.length}
                    />
                );
            }
            if (route.children) {
                this.parseRoute(props, route.children);
            }
        });
    }

    public render() {
        return (
            <ReactRouter {...this.props} history={this.history}>
                <Switch>{this.routeArray}</Switch>
            </ReactRouter>
        );
    }
}
