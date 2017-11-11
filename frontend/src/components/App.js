import React, { Component } from 'react'
import '../css/App.css'
import LoginForm from "./LoginForm"
import Navbar from "./Navbar";
import { MuiThemeProvider } from "material-ui/styles"
import getFormValues from "../reducers"
import { connect } from "react-redux";

const Body = () => {
  return [
    <Navbar key={1}/>,
    <LoginForm key={2}/>
  ]
}

class App extends Component {

  render() {
    return (
      <MuiThemeProvider>
        <Body />
      </MuiThemeProvider>
    )
  }
}


export default App;
