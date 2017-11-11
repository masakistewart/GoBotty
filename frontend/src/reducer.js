import * as actions from './actions'

const initState = {
    email: "",
    password: "",
    user: false,
    toggleLogin: false
}

export default (state = initState, action) => {
    switch (action.type) {
        case actions.EMAIL_CHANGE:
            return {...state, login: state.email.concat(action.payload)}
        case actions.PASSWORD_CHANGE:
            return {...state, login: state.password.concat(action.payload)}
        case actions.TOGGLE_FORM:
            return {...state, toggleLogin: !state.toggleLogin}
        default:
           return state
    }
}