import axios from '@/libs/api.request'

export const login = ({ userName, password }) => {
  const data = {
    userName,
    password
  }
  return axios.request({
    url: '/users/login',
    data,
    method: 'post'
  })
}
export const addUser = ({ userName, passWord,age,gender,occupation }) => {
  const data = {
    userName,
    passWord,
    age,
    gender,
    occupation
  }
  return axios.request({
    url: '/users/addUser',
    data,
    method: 'post'
  })
}



export const getUserInfo = (token) => {
  return axios.request({
    url: '/users/get_info',
    params: {
      token
    },
    method: 'get'
  })
}


export const getPageMovie =({ token,pageNo,pageSize }) => {
  return axios.request({
    url: 'movie/PageMovie',
    method: 'get',
    headers: {'Authorization': token},
    params: {
      pageNo,
      pageSize
    }
  })
}

export const movieRate =({ token,movieId,rating }) => {
  return axios.request({
    url: 'movie/movieRate',
    method: 'get',
    headers: {'Authorization': token},
    params: {
      movieId,
      rating
    }
  })
}

export const logout = (token) => {
  return axios.request({
    url: '/users/logout',
    method: 'get',
    headers: {'Authorization': token}
  })
}


export const getMessage = () => {
  return axios.request({
    url: 'message/init',
    method: 'get'
  })
}

export const getContentByMsgId = msg_id => {
  return axios.request({
    url: 'message/content',
    method: 'get',
    params: {
      msg_id
    }
  })
}

export const hasRead = msg_id => {
  return axios.request({
    url: 'message/has_read',
    method: 'post',
    data: {
      msg_id
    }
  })
}

export const removeReaded = msg_id => {
  return axios.request({
    url: 'message/remove_readed',
    method: 'post',
    data: {
      msg_id
    }
  })
}

export const restoreTrash = msg_id => {
  return axios.request({
    url: 'message/restore',
    method: 'post',
    data: {
      msg_id
    }
  })
}
