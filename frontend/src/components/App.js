import { MuiThemeProvider } from "material-ui/styles"
import React, { Component } from 'react'

import Login from "./Login"
import Navbar from "./Navbar";
import '../css/App.css'

const Body = (props) => {
  return [
    <Navbar key={1} {...props}/>,
    props.toggleLoginForm ? <Login {...props} key={2}/> : null
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
