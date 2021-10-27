import axios from 'axios'

export default () => {
    var currentLocation = window.location;
    let port = currentLocation.port
    if (process.env.NODE_ENV == "development"){
        port = "8080"
    }

    let baseURL = `//`+currentLocation.hostname+`:`+port;

    var apiClient = axios.create({
        baseURL: baseURL,
        headers: {
            'Content-Type': 'application/json'
        },
    });

    apiClient.interceptors.response.use(
        function (response) {
            //do nothing if ok
            return response;
        }, function (error) {
               return Promise.reject(error);
        });

    return apiClient;
}
