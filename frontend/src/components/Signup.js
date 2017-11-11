import React, { Component } from "react"
import { TextField } from "material-ui"
import Paper from "material-ui/Paper"
import RaisedButton from 'material-ui/RaisedButton'

class Signup extends Component {
    render() {
        console.log(this.props)
        let email, password
        return (
            <div className="container-center" >
                <Paper style={style} zDepth={1}>
                    <h2>Signup</h2>
                    <div>
                        <TextField style={textStyle} onChange={this.onChangeEmailHandler} type="email" floatingLabelText="Email" hintText="Email Field" />
                    </div>
                    <div>
                        <TextField style={textStyle} onChange={this.onChangePasswordHandler} type="password" hintText="Password Field" floatingLabelText="Password" />
                    </div>
                    <div>
                        <RaisedButton label="Login" primary={true} style={buttonStyle} />
                    </div>
                </Paper>
            </div>
        )
    }
}

const style = {
    width: "80%",
    height: "auto",
    textAlign: 'center',
    padding: "30px 30px 30px 30px",
    margin: "150px auto 0 auto",
}

const textStyle = {
    width: "100%"
}

const buttonStyle = {
    margin: "30px 0 0 0"
}

export default Signup