import React, { useState, useEffect } from 'react';
import { Col, Row } from 'react-bootstrap';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import { useNavigate } from "react-router-dom";
import { AlertContext } from '../contexts/alertContext';
import { AuthContext } from '../contexts/authContext';

function RegistrationForm() {
  const { logged, signup } = React.useContext(AuthContext)
  const { setShowAlert } = React.useContext(AlertContext)

  const navigate = useNavigate();

  const [values, setValues] = useState({
    login: "",
    password: "",
  });

  const submit = async (e) => {
    e.preventDefault();
    try {
      await signup({...values})
      setShowAlert({ show: true, variant: 'success', message: 'Successful authentication' })
    } catch(err) {
      setShowAlert({ show: true, variant: 'danger', message: err })
    }
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
    <Row>
      <Col>
        <Form onSubmit={submit}>
          <Form.Group className="mb-3" controlId="formLogin">
            <Form.Label>Login</Form.Label>
            <Form.Control type="text" placeholder="Логин" name="login" value={values.login} onChange={handleChange}/>
          </Form.Group>

          <Form.Group className="mb-3" controlId="formPassword">
            <Form.Label>Password</Form.Label>
            <Form.Control type="password" placeholder="Пароль" name="password"  value={values.password} onChange={handleChange}/>
          </Form.Group>
          <Button variant="primary" type="submit">
            Sign up
          </Button>
        </Form>
      </Col>
    </Row>
  );
}

export default RegistrationForm;
