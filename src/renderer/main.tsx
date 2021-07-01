import React from "react"
import { render } from "react-dom"
import App from './App'

const ROOT_EL = document.getElementById("root")

render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  ROOT_EL
)
