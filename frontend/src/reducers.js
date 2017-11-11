import { GET_FORM_VALUES, DISABLE_BUTTON } from "./actions"
import { combineReducers } from 'redux'

const initialState = {
    formValues: {
        email: "",
        password: ""
    },
    disableButton: true
}

export function initializeApp(state = initialState, action) {
    switch (action.type) {
        case GET_FORM_VALUES:
            return { ...state, ...action.state}
        default:
            return state
    }
}


