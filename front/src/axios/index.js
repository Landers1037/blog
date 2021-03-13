import customData from "../custom/custom";
import axios from "axios";
const httprequest = axios.create({
    baseURL: customData.api_url,
    timeout: 10000
});

export default httprequest