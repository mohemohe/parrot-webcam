import "regenerator-runtime/runtime";
import "core-js/stable";
import "whatwg-fetch";

import("./app").then((app) => {
    app.render();
});
