import axios from 'axios';

const API_BASE = "http://rocket.dev/api/";

var Api = axios.create({
    baseURL: API_BASE
});

export default Api;