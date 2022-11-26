import axios from "axios";
import config from "./config";

const ApiClient = axios.create({ baseURL: config.API_URL });

export default ApiClient;