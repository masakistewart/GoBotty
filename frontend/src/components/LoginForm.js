import React, { Component } from "react"
import { TextField } from "material-ui"
import Paper from "material-ui/Paper"
import RaisedButton from 'material-ui/RaisedButton'
import { connect } from "react-redux";
import { GET_FORM_VALUES, SEND_LOGIN_INFO, getFormValues } from "../actions"



class LoginForm extends Component{
    setState(email, password) {
        return {
            formValues: {
                email: email,
                password: password,
            },
            disableButton: false
        }
    }
    
    setState(email, password) {
        return {
            formValues: {
                email: email,
                password: password,
            },
            disableButton: false
        }
    }

    sendLoginInfo(formValues) {
        fetch("http://localhost:8080", {
            method: "GET",
            body: formValues,
            headers: {
                "content-type": "application/json"
            }
        }).catch(err => console.log(err))
    }


    render() {
        let email, password
        return (
            <div className="container-center" >
                <Paper style={style} zDepth={1}>
                    <form
                        onSubmit={(e) => {
                            e.preventDefault()
                            let emailAddr = email.input.value.trim()
                            let pass = password.input.value.trim()
                            let formVals = this.setState(emailAddr, pass)
                            if (emailAddr && pass) {
                                this.props.dispatch(getFormValues(formVals))
                                this.sendLoginInfo(this.props.formValues)  
                            } else {
                                return
                            }
                            // console.log(store.getState())
                        }}>
                        <div>
                            <TextField ref={node => {email = node }} style={textStyle} type="email" floatingLabelText="Email" hintText="Email Field" />
                        </div>
                        <div>
                            <TextField ref={node => { password = node }} style={textStyle} type="password" hintText="Password Field" floatingLabelText="Password" />
                        </div>
                        <div>
                            <RaisedButton label="Primary" type="submit" primary={true} style={buttonStyle} />
                        </div>
                    </form>
                </Paper>
            </div>
        )
    }
}

const mapStateToProps = state => {
    return { formValues: state.formValues }
}

const mapDispatchToProps = dispatch => {
    console.log("DISPATCH, ", dispatch)
    return { dispatch }
}

const style = {
    width: "80%",
    height: "auto",
    textAlign: 'center',
    padding: "30px 30px 30px 30px",
    margin: "0 auto",
}

const textStyle = {
    width: "100%"
}

const buttonStyle = {
    margin: "30px 0 0 0"
}


const ConnectedForm = connect(
    mapStateToProps,
    mapDispatchToProps
)(LoginForm)

export default ConnectedForm


