import React from "react"
import { render } from "react-dom"
import { Provider } from "react-redux"
import { createStore, applyMiddleware, compose } from "redux"
import { createLogger } from "redux-logger"
import thunkMiddleware from "redux-thunk"

import "./index.css"
import "tailwindcss/tailwind.css"

import reducer from "./reducers/index.tsx"

import Router from "./router.tsx"

const ROOT_EL = document.getElementById("root")
const initialState = {}
const loggerMiddleware = createLogger()

const store = createStore(
  reducer,
  initialState,
  compose(
    applyMiddleware(thunkMiddleware, loggerMiddleware),
    window.devToolsExtension ? window.devToolsExtension() : f => f
  )
)

render(
  <React.StrictMode>
    <Provider store={store}>
      <Router />
    </Provider>
  </React.StrictMode>,
  ROOT_EL
)
