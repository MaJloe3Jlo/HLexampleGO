
import axios from 'axios';

const HOST = process.env.NODE_ENV === 'development' ? 'http://localhost:7777' : 'http://www.host:8080';

export default {
  get() {
    return axios({
      method: 'GET',
      url: HOST + '/'
    })
  },
  save(data) {
    return axios({
      method: 'POST',
      url: HOST + '/post',
      data
    })
  }
};