import React, { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import { AuthContext } from './authContext';

function LoginPage() {
  const auth = React.useContext(AuthContext)

  console.log(auth)

  const [values, setValues] = useState({
    login: "",
    password: "",
  });

  const submit = (e) => {
    console.log(auth)
    e.preventDefault();
    auth.login(values)
  }

  const handleChange = (event) => {
    setValues((form) => ({
      ...form,
      [event.target.name]: event.target.value,
    }));
   };

  return (
    <Form onSubmit={submit}>
      <Form.Group className="mb-3" controlId="formLogin">
        <Form.Label>Login</Form.Label>
        <Form.Control type="text" placeholder="Enter email" name="login" value={values.login} onChange={handleChange}/>
        <Form.Text className="text-muted">
          We'll never share your email with anyone else.
        </Form.Text>
      </Form.Group>

      <Form.Group className="mb-3" controlId="formPassword">
        <Form.Label>Password</Form.Label>
        <Form.Control type="password" placeholder="Password" name="password"  value={values.password} onChange={handleChange}/>
      </Form.Group>
      <Form.Group className="mb-3" controlId="formBasicCheckbox">
        <Form.Check type="checkbox" label="Check me out" />
      </Form.Group>
      <Button variant="primary" type="submit">
        Submit
      </Button>
    </Form>
  );
}

export default LoginPage;
