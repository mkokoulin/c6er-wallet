import React, { useState } from 'react';
import api from './api';

export const AuthContext = React.createContext(null);

const initialState = {
  isLoggedIn: false,
  isLoginPending: false,
  loginError: null
}

export const AuthProvider = props => {
  const [state, setState] = useState(initialState);

  const setLoginPending = (isLoginPending) => setState({isLoginPending});
  const setLoginSuccess = (isLoggedIn) => setState({isLoggedIn});
  const setLoginError = (loginError) => setState({loginError});

  const login = (login, password) => {
    setLoginPending(true);

    return api.post('http://localhost:8080/api/v1/user/login', {
      email: login,
      password: password
    })
    .then(res => {
      setLoginSuccess(true);
    })
    .catch(err => {
      setLoginError(err);
    })
    .finally(() => {
      setLoginPending(true);
    })
  }

  const logout = () => {
    setLoginPending(false);
    setLoginSuccess(false);
    setLoginError(null);
  }

  return (
    <AuthContext.Provider
      value={{
        state,
        login,
        logout,
      }}
    >
      {props.children}
    </AuthContext.Provider>
  );
};

// fake login
const fetchLogin = (email, password, callback) => 
  setTimeout(() => {
    if (email === 'admin' && password === 'admin') {
      return callback(null);
    } else {
      return callback(new Error('Invalid email and password'));
    }
  }, 1000);