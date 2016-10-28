import axios from "axios";

const API_BASE = "http://localhost:8000/api/";


var Api = axios.create({
    baseURL: API_BASE
});

export default Api;