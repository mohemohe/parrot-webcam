import { RouterStore } from "mobx-react-router";
import {AuthStore} from "./AuthStore";
import {ImageStore} from "./ImageStore";

const stores = {
    RouterStore: new RouterStore(),
    AuthStore: new AuthStore(),
    ImageStore: new ImageStore(),
};

export default stores;
