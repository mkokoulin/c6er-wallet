import React, { useState } from 'react';
import api from '../api';
import { getCookie } from '../utils/cookie';

export const AuthContext = React.createContext(null);

export const AuthProvider = props => {
  const [pending, setPending] = useState(false);
  const [logged, setlogged] = useState(false);
  const [error, setError] = useState(null);

  const login = (payload) => {
    setPending(true);

    return api.post('/user/login', {
      login: payload.login,
      password: payload.password,
    })
    .then(res => {
      setlogged(true);
      return res
    })
    .catch(err => {
      setError(err);
      throw err;
    })
    .finally(() => {
      setPending(false);
    })
  }

  const signup = (payload) => {
    const accessToken = getCookie('access_token');

    setPending(true);

    return api.post('/user/signup', {
      login: payload.login,
      password: payload.password,
    })
    .then(res => {
      setlogged(true);
      return res
    })
    .catch(err => {
      setError(err);
      throw err;
    })
    .finally(() => {
      setPending(false);
    })
  }

  const checkAuth = async () => {
    const accessToken = getCookie('access_token');

    try {
      await api.get('/user/auth', {
        headers: {'Authorization': `Bearer ${accessToken}`}
      });
      return true
    } catch(err) {
      return false
    }
  }

  const logout = () => {
    setPending(false);
    setlogged(false);
    setError(null);
  }

  return (
    <AuthContext.Provider
      value={{
        pending,
        logged,
        error,
        login,
        signup,
        logout,
        checkAuth
      }}
    >
      {props.children}
    </AuthContext.Provider>
  );
};

// fake login
const fetchLogin = (login, password, callback) => 
  setTimeout(() => {
    if (login === 'admin' && password === 'admin') {
      return callback(null);
    } else {
      return callback(new Error('Invalid login and password'));
    }
  }, 1000);