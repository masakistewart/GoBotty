import { MuiThemeProvider } from "material-ui/styles"
import React, { Component } from 'react'

import Form from "./Form"
import Navbar from "./Navbar"
import '../css/App.css'

const Body = (props) => {
  return [
    <Navbar key={1} {...props}/>,
    <Form {...props} key={2}/>
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
