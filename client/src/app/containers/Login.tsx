import React from "react";
import {style} from "typestyle";
import {Button, Form, Input} from "antd";
import {inject, observer} from "mobx-react";
import {RouterStore} from "mobx-react-router";
import {AuthStatus, AuthStore} from "../stores/AuthStore";
import {Redirect} from "react-router";
import {Mode} from "../stores/StoreBase";

interface Props {
    RouterStore?: RouterStore;
    AuthStore?: AuthStore;
}

const classNames = {
    buttonContainer: style({
        display: "flex",
        justifyContent: "flex-end",
    }),
}

@inject("RouterStore", "AuthStore")
@observer
export default class Login extends React.Component<Props, {}> {
    public render() {
        if (this.props.AuthStore?.authStatus === AuthStatus.Authorized) {
            return <Redirect to={"/"} />
        }

        return (
            <div>
                <h1>Webcam:</h1>
                <Form
                    name="basic"
                    initialValues={{ remember: true }}
                    onFinish={(values: any) => {this.props.AuthStore?.login(values.id, values.password)}}
                >
                    <Form.Item
                        label="ユーザーID"
                        name="id"
                        rules={[{ required: true, message: "ユーザーIDは必須です" }]}
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        label="パスワード"
                        name="password"
                        rules={[{ required: true, message: "パスワードは必須です" }]}
                    >
                        <Input.Password />
                    </Form.Item>

                    {this.props.AuthStore?.mode === Mode.LOGIN && this.props.AuthStore?.authStatus === AuthStatus.Unauthorized && <div>ログイン情報が間違っています</div>}

                    <Form.Item className={classNames.buttonContainer}>
                        <Button type="primary" htmlType="submit">
                            ログイン
                        </Button>
                    </Form.Item>
                </Form>
            </div>
        );
    }
}
