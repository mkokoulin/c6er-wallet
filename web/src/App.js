import React, { useEffect } from 'react';
import {
  Routes,
  Route,
  useLocation,
  Navigate
} from "react-router-dom";
import { Container } from 'react-bootstrap';
import { AuthContext } from './contexts/authContext';
import LoginPage from './pages/LoginPage';
import MainPage from './pages/MainPage';
import { useJwt } from "react-jwt";
import { getCookie } from "./utils/cookie";
import './App.css'

function App() {
  const auth = React.useContext(AuthContext);

  // useEffect(() => {
  //   // const a = await auth.checkAuth()

  //   // if (!a) {
  //   //   return <Navigate to="/login" replace />;
  //   // }
  // })

  return (
    <Container className="App">
      <Routes>
        <Route path="/" element={<Navigate replace to="/main" />}/>
        <Route path="/login" element={<LoginPage />} />
        <Route
          path="/main"
          element={
            <RequireAuth>
              <MainPage />
            </RequireAuth>
          }
        />
      </Routes>
    </Container>
  );
}

function RequireAuth({ children }) {
  let auth = React.useContext(AuthContext);
  let location = useLocation();

  const accessToken = getCookie('access_token');

  const { decodedToken, isExpired } = useJwt(accessToken);

  console.log(isExpired);
  console.log(isExpired);

  if (!auth.logged) {
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  return children;
}

export default App;
