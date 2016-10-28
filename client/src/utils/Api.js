import axios from 'axios';

const API_BASE = "http://api.rocket.dev/api/";


var Api = axios.create({
    baseURL: API_BASE
});

export default Api;