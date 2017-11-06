window.onload = () => {

    let form = document.getElementById("login-form")

    form.addEventListener("submit", (e) => {
        e.preventDefault()
        let body = {
            email: document.getElementById("email").value,
            password: document.getElementById("password").value
        }
        
        let loginInfo = {method: "POST", body: JSON.stringify(body)}
        console.log(body)
        fetch("/login", loginInfo)
        .then(body => body.text())
        .then(json => console.log(json))
        .catch(err =>  console.log(err))
    })

}