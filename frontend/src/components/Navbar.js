import React, { Component } from 'react';
import AppBar from 'material-ui/AppBar';
import IconButton from 'material-ui/IconButton';
import IconMenu from 'material-ui/IconMenu';
import MenuItem from 'material-ui/MenuItem';
import FlatButton from 'material-ui/FlatButton';
import Toggle from 'material-ui/Toggle';
import MoreVertIcon from 'material-ui/svg-icons/navigation/more-vert';
import NavigationClose from 'material-ui/svg-icons/navigation/close';
import "../css/index.css"

class Login extends Component {
    static muiName = 'FlatButton';

    render() {
        return (
            <FlatButton style={this.props.style} onClick={this.props.toggleLoginForm} label={this.props.toggleLogin ? "Login" : "Signup"} />
        );
    }
}

const Logged = (props) => (
    <IconMenu
        {...props}
        iconButtonElement={
            <IconButton><MoreVertIcon /></IconButton>
        }
        targetOrigin={{ horizontal: 'right', vertical: 'top' }}
        anchorOrigin={{ horizontal: 'right', vertical: 'top' }}
    >
        <MenuItem primaryText="Refresh" />
        <MenuItem primaryText="Help" />
        <MenuItem primaryText="Sign out" />
    </IconMenu>
);

Logged.muiName = 'IconMenu';

class AppBarExampleComposition extends Component {
    render() {
        return (
            <div>
                <AppBar
                    title="PWP"
                    iconElementRight={
                        <Login
                            style={this.props.style}
                            toggleLoginForm={this.props.toggleLoginForm}
                            toggleLogin={this.props.toggleLogin}
                        />
                    }
                />
            </div>
        );
    }
}

export default AppBarExampleComposition;