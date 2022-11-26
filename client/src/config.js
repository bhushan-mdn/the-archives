// Config.js

// import prod from "./prod.config.json";
// import staging from "./staging.config.json";
import dev from "./dev.config.json";
import localhost from "./localhost.config.json";

let config = {};

switch (window.location.hostname) {
    case "www.myapp.com":
        config = prod;
        break;
    case "staging.myapp.com":
        config = staging;
        break;
    case "192.168.1.61":
        config = dev;
        break;
    case "localhost":
        config = localhost;
        break;
    default:
        config = {
            "API_URL": `http://${window.location.hostname}/api`
        };
}

export default config;
