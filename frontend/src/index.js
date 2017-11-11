import registerServiceWorker from './registerServiceWorker'
import React from 'react'
import ReactDOM from 'react-dom'
import { Provider } from "react-redux"
import App from './components/App'
import store from "./store"
import './css/index.css'

const emailChangeHanlder = (val) => store.dispatch({ type: 'EMAIL_CHANGE', payload: val })
const passwordChangeHanlder = (val) => store.dispatch({ type: 'PASSWORD_CHANGE', payload: val })
const toggleLoginForm = () => { 
    debugger; 
    store.dispatch({type: 'TOGGLE_FORM'}) }

const render = () => {
    const state = store.getState()
    ReactDOM.render(
        <App
            emailChangeHanlder={emailChangeHanlder}
            passwordCHangeHanlder={passwordChangeHanlder}
            toggleLoginForm={toggleLoginForm}
            {...state}
        />, document.getElementById('root'))
}
render()

store.subscribe(render)

registerServiceWorker();
