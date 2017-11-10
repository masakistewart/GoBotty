import React from 'react';
import ReactDOM from 'react-dom';
import './css/index.css';
import App from './components/App';
import MuiThemeProvider from "material-ui/styles/MuiThemeProvider";
import registerServiceWorker from './registerServiceWorker';

const Test = () => (
    <MuiThemeProvider>
        <App />
    </MuiThemeProvider>
)
ReactDOM.render(<Test />, document.getElementById('root'));
registerServiceWorker();
