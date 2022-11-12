import React, { useState, useEffect } from 'react';
import { Tab, Tabs } from 'react-bootstrap';
import { useNavigate } from "react-router-dom";
import { AuthContext } from '../contexts/authContext';
import LoginForm from '../components/LoginForm';
import RegistrationForm from '../components/RegistrationForm';
import './LoginPage.css'


function LoginPage() {
  const { pending, logged, error, login } = React.useContext(AuthContext)
  const navigate = useNavigate();

  const [values, setValues] = useState({
    login: "",
    password: "",
  });

  const submit = (e) => {
    e.preventDefault();
    login({...values})
  }

  useEffect(() => {
    if (logged) {
      navigate('/main', { replace: true });
    }
  }, [logged]);

  const handleChange = (event) => {
    setValues((form) => ({
      ...form,
      [event.target.name]: event.target.value,
    }));
   };

  return (
    <div className="LoginPage">
      <div className="LoginPage-wrapper">
        <Tabs
          defaultActiveKey="login"
          id="uncontrolled-tab-example"
          className="mb-3"
        >
          <Tab eventKey="login" title="Вход">
            <LoginForm />
          </Tab>
          <Tab eventKey="signup" title="Регистрация">
            <RegistrationForm />
          </Tab>
        </Tabs>
      </div>
    </div>
  );
}

export default LoginPage;
