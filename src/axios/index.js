const httprequest = axios.create({
    baseURL: 'https://blog.renj.io',
    timeout: 30000
});

export default httprequest