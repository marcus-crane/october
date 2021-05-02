import React from "react"
import { render } from "react-dom"
import { HashRouter, Route, Switch } from "react-router-dom"
import { Provider } from "react-redux"
import { createStore, applyMiddleware, compose } from "redux"
import { createLogger } from "redux-logger"
import thunkMiddleware from "redux-thunk"

import reducer from "./reducer.jsx"

import DeviceSelection from "./pages/DeviceSelection.jsx"
import Bookshelf from "./pages/Bookshelf.jsx"

import "./index.css"
import "tailwindcss/tailwind.css"

const ROOT_EL = document.getElementById("root")
const initialState = {}
const loggerMiddleware = createLogger()

const store = createStore(
  reducer,
  initialState,
  compose(
    applyMiddleware(thunkMiddleware, loggerMiddleware),
    window.devToolsExtension ? window.devToolsExtension() : (f) => f
  )
)

render(
  <Provider store={store}>
    <HashRouter>
      <Route exact path="/" component={DeviceSelection} />
      <Route exact path="/books" component={Bookshelf} />
    </HashRouter>
  </Provider>,
  ROOT_EL
)
