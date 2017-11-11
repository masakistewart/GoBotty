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
        console.log(this.props)
        return (
            <FlatButton style={this.props.style} onClick={this.props.toggleLoginForm} label="Login" />
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

/**
 * This example is taking advantage of the composability of the `AppBar`
 * to render different components depending on the application state.
 */
class AppBarExampleComposition extends Component {
    render() {
        console.log(this.props.toggleLoginForm)
        return (
            <div>
                <AppBar
                    title="PWP"
                    iconElementRight={!this.props.toggleLogin ? <Login toggleLoginForm={this.props.toggleLoginForm} /> : <Logged  />}
                />
            </div>
        );
    }
}

export default AppBarExampleComposition;