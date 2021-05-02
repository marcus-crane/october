import React from "react";
import { render } from "react-dom";
import {
  createMemorySource,
  createHistory,
  LocationProvider,
  Router,
} from "@reach/router";
import { Provider } from "react-redux";
import { createStore, applyMiddleware, compose } from "redux";
import { createLogger } from "redux-logger";
import thunkMiddleware from "redux-thunk";

import reducer from "./reducer.jsx";
import App from "./app.jsx";

import "./index.css";
import "tailwindcss/tailwind.css";

const ROOT_EL = document.getElementById("root");
const initialState = {};
const loggerMiddleware = createLogger();

const source = createMemorySource("/");
const history = createHistory(source);

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
    <LocationProvider history={history}>
      <Router>
        <App path="/" />
      </Router>
    </LocationProvider>
  </Provider>,
  ROOT_EL
);
