import React, { useState } from "react";

export const AlertContext = React.createContext(null);

const initialState = {
    show: false,
    variant: 'success', // success, danger, warning, info
    message: ''
}

export const AlertProvider = props => {
    const [state, setState] = useState(initialState);

    const setShowAlert = (state) => {
        setState(state)

        setTimeout(() => {
            setState(initialState)
        }, 3000)
    }

    return (
        <AlertContext.Provider
            value={{
                state,
                setShowAlert
            }}
        >
            {props.children}
        </AlertContext.Provider>
    );
};