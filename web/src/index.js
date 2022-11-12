import React from 'react';
import ReactDOM from 'react-dom/client';
import reportWebVitals from './reportWebVitals';
import { BrowserRouter } from "react-router-dom";
import App from './App';
import combineComponents from './utils/combineComponents';
import { AuthProvider } from './contexts/authContext';
import { AlertProvider } from './contexts/alertContext';
import Alert from './components/Alert';
import './index.css';
import 'bootstrap/dist/css/bootstrap.min.css';

const root = ReactDOM.createRoot(document.getElementById('root'));

const providers = [AuthProvider, AlertProvider]

export const AppContextProvider = combineComponents(...providers);

root.render(
  <React.StrictMode>
    <BrowserRouter>
      <AppContextProvider>
        <Alert />
        <App/>
      </AppContextProvider>
    </BrowserRouter>
  </React.StrictMode>
);

reportWebVitals();
