import React from 'react';
import {
  Routes,
  Route,
  useLocation,
  Navigate
} from "react-router-dom";
import { Container } from 'react-bootstrap';
import { AuthProvider, AuthContext } from './authContext';
import LoginPage from './LoginPage';
import MainPage from './MainPage';

function App() {
  return (
    <Container>
      <AuthProvider>
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
      </AuthProvider>
    </Container>
  );
}

function RequireAuth({ children }) {
  let auth = React.useContext(AuthContext);;
  let location = useLocation();

  if (!auth.state.isLoggedIn) {
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  return children;
}

export default App;
