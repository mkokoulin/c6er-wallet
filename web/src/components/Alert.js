import React from 'react';
import { Alert as AlertComponent } from 'react-bootstrap';
import { AlertContext } from '../contexts/alertContext';
import './Alert.css';

function Alert() {
  const { state } = React.useContext(AlertContext)

  return (
    <div className="Alert">
      { state.show ?
        <AlertComponent key={state.variant} variant={state.variant}>
          {state.message}
        </AlertComponent> : null }
    </div>
  );
}

export default Alert;
