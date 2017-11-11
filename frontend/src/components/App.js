import { MuiThemeProvider } from "material-ui/styles"
import React, { Component } from 'react'

import Login from "./Login"
import Navbar from "./Navbar";
import Signup from './Signup'
import '../css/App.css'

const Body = (props) => {
  return [
    <Navbar key={1} {...props}/>,
    props.toggleLogin ? <Login {...props} key={2}/> : <Signup />
  ]
}

class App extends Component {
  render() {
    return (
      <MuiThemeProvider>
        <Body {...this.props} />
      </MuiThemeProvider>
    )
  }
}


export default App;
