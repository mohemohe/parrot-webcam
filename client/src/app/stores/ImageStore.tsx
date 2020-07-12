
import {action, computed, observable} from "mobx";
import StoreBase, {IModel, Mode, State} from "./StoreBase";

export class ImageStore extends StoreBase {
    @observable
    public buffer: string;

    constructor() {
        super();

        this.buffer = "";
    }

    @action
    public async fetch() {
        this.setMode(Mode.GET);
        this.setState(State.RUNNING);

        const buffer = this.buffer;

        try {
            const response = await fetch(this.apiBasePath + "v1/image", {
                headers: this.generateFetchHeader(),

            });

            if (response.status !== 200) {
                throw new Error();
            }
            const result = await response.blob();
            this.buffer = URL.createObjectURL(result);
            if (buffer.length > 0) {
                URL.revokeObjectURL(buffer);
            }

            this.setState(State.DONE);
        } catch (e) {
            this.tryShowToast("画像取得に失敗しました");
            console.error(e);
            this.setState(State.ERROR);
        }
    }
}
