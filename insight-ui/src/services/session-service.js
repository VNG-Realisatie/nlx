import axios from 'axios';

function getSessionData(baseUrl) {
  return axios
    .get(`${baseUrl}/get-session`, {
      withCredentials: true,
    })
    .then(response => response.data);
}

function deauthenticate(baseUrl) {
  return axios
    .get(`${baseUrl}/deauthenticate`, {
      withCredentials: true,
    });
}

const service = {
  getSessionData,
  deauthenticate,
};

export default service;