import React from "react";
import { render } from "react-dom";
import { Router } from "@reach/router";
import { Provider } from "react-redux";
import { createStore, applyMiddleware, compose } from "redux";
import { createLogger } from "redux-logger";
import thunkMiddleware from "redux-thunk";

import reducer from "./reducers";
import App from "./app.jsx";

import "./index.css";
import "tailwindcss/tailwind.css";

const ROOT_EL = document.getElementById("root");
const initalState = {};
const loggerMiddleware = createLogger();

const store = createStore(
  reducer,
  initialState,
  compose(
    applyMiddleware(thunkMiddleware, loggerMiddleware),
    window.devToolsExtension ? window.devToolsExtension() : (f) => f
  )
);

render(
  <Provider store={store}>
    <Router>
      <App path="/" />
    </Router>
  </Provider>,
  ROOT_EL
);
