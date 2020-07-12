import * as React from "react";
import { Link, LinkProps } from "react-router-dom";

export interface Props extends LinkProps {
    buttonProps?: React.ComponentProps<"button">;
}

export default class LinkButton extends React.Component<Props, {}> {
    public render() {
        return (
            <Link {...this.props}>
                <button {...this.props.buttonProps}>
                    {this.props.children}
                </button>
            </Link>
        );
    }
}
