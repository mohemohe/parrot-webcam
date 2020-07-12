import React from "react";
import {inject, observer} from "mobx-react";
import {AuthStatus, AuthStore} from "../stores/AuthStore";
import {Redirect} from "react-router";
import {ImageStore} from "../stores/ImageStore";

interface Props {
    AuthStore?: AuthStore;
    ImageStore?: ImageStore;
}

@inject("AuthStore", "ImageStore")
@observer
export default class Image extends React.Component<Props, {}> {
    private authHandler?: any;

    public async componentDidMount() {
        await this.props.AuthStore?.checkAuth();
        if (this.props.AuthStore?.authStatus === AuthStatus.Unauthorized) {
            return;
        }

        this.props.ImageStore?.fetch();

        this.authHandler = setInterval(() => {
            this.props.AuthStore?.checkAuth();
        }, 60 * 60 * 1000); // NOTE: 1時間
        this.authHandler = setInterval(() => {
            this.props.ImageStore?.fetch();
        }, 1000);
    }

    public componentWillUnmount() {
        clearInterval(this.authHandler);
    }

    public render() {
        if (this.props.AuthStore?.authStatus === AuthStatus.Unauthorized) {
            return (
                <Redirect to={"/login"} />
            );
        }

        return (
            <div>
                <img src={this.props.ImageStore?.buffer} />
            </div>
        );
    }
}
