import React from 'react'
import {render} from 'react-dom'
import './css/index.css'
import AppComponent from './components/App'
import registerServiceWorker from './registerServiceWorker'
import Navbar from "./components/Navbar"
import { createStore } from 'redux'
import {initializeApp} from './reducers'
import { GET_FORM_VALUES } from "./actions"
import { Provider } from "react-redux"



let store = createStore(initializeApp)

store.dispatch({
    type: "GET_FORM_VALUES", state: store.getState()
})


console.log(store.getState())
const App = () => (
    <Provider store={store} >
        <AppComponent/>
    </Provider>
)

render(
    <App />,
    document.getElementById('root')
);
registerServiceWorker();
