import React from "react"
import { AppBar } from "material-ui";

const style = {
    marginBottom: "150px"
}

const Navbar = () => {
    return (
        <AppBar style={style} title="Gophers Galore!!" />
    )
}

export default Navbar