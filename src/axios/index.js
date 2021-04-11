import customData from "../custom/custom";
import {Message} from "element-ui";

const httprequest = axios.create({
    baseURL: customData.api_url,
    withCredentials: false,
    timeout: 10000
});

httprequest.interceptors.request.use(
    config => {
        if(localStorage.getItem("token")) {
            config.headers.admin_token = localStorage.getItem("token");
        }
        return config;
    }
);

httprequest.interceptors.response.use(
    response => {
        if (response.status === 403){
                console.log("无权限访问接口");
                Message.error("无权限访问接口");
        }
        return response;
    },
    error => {
        if (error.response.status === 403 ){
            console.log("无权限访问接口");
            Message.error("无权限访问接口");
        }
    }
)
export default httprequest;