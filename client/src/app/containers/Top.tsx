import React from "react";
import {inject, observer} from "mobx-react";
import {RouterStore} from "mobx-react-router";
import {AuthStatus, AuthStore} from "../stores/AuthStore";
import {Redirect} from "react-router";

interface Props {
    RouterStore?: RouterStore;
    AuthStore?: AuthStore;
}

@inject("RouterStore", "AuthStore")
@observer
export default class Top extends React.Component<Props, {}> {
    public componentDidMount() {
        this.props.AuthStore?.checkAuth();
    }

    public render() {
        if (this.props.AuthStore?.authStatus === AuthStatus.Authorized) {
            return (
                <Redirect to={"/image"} />
            )
        }

        if (this.props.AuthStore?.authStatus === AuthStatus.Unauthorized) {
            return (
                <Redirect to={"/login"} />
            );
        }

        return (
            <div>
                Loading...
            </div>
        );
    }
}
