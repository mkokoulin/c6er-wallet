import axios from 'axios';

const ax = axios.create({
    baseURL: `/api/v1/`,
});

ax.interceptors.request.use(function (config) {
    // Do something before request is sent
    return config;
  }, function (error) {
    // Do something with request error
    return Promise.reject(error);
  });

// Add a response interceptor
ax.interceptors.response.use(function (response) {
    // Do something with response data
    return response;
  }, function (error) {
    if (error.response.status === 401) {
      window.location = window.location.protocol + "//" + window.location.host + "/login"
    }
    // Do something with response error
    return Promise.reject(error.response ? error.response.data : error);
  });

  export default ax