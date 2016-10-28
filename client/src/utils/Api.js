import axios from "axios";

const API_BASE = "/api/";

var Api = axios.create({
    baseURL: API_BASE
});

export default Api;