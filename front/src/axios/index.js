import axios from 'axios'
import customData from "../custom/custom";
const httprequest = axios.create({
    baseURL: customData.api_url,
    timeout: 10000
});

export default httprequest;