
import axios from 'axios';

export default {
  get() {
    return axios({
      method: 'GET',
      url: '/api/get'
    })
  },
  save(data) {
    return axios({
      method: 'POST',
      url: '/api/post',
      data
    })
  },
  list() {
    return axios({
      method: 'GET',
      url: '/api/list'
    })
  }
};