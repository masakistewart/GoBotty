import actionTypes  from './actions'

const initState = {
    email: "",
    password: "",
    user: false,
    toggleLogin: false
}

export default (state = initState, action) => {
    debugger
    switch (action.type) {
        case actionTypes.EMAIL_CHANGE:
            return {...state, login: state.email.concat(action.payload)}
        case actionTypes.PASSWORD_CHANGE:
            return {...state, login: state.password.concat(action.payload)}
        case actionTypes.TOGGLE_FORM:
            return {...state, toggleLogin: !state.toggleLogin}
        default:
           return state
    }
}