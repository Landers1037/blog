const httprequest = axios.create({
    // baseURL: 'https://blog.renj.io',
    baseURL: "http://10.211.55.4:5000",
    timeout: 10000
});

export default httprequest