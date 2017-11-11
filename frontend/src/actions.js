export const GET_FORM_VALUES = "GET_FORM_VALUES"
export const SEND_LOGIN_INFO = "SEND_LOGIN_INFO"
export const DISABLE_BUTTON = "DISABLE_BUTTON"

export function getFormValues(formValues) {
    debugger
    return { type: GET_FORM_VALUES, state: formValues}
}
