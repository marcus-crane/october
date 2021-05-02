import { ipcRenderer } from "electron"

import { buildSqlitePath } from "./constants"

export const actionTypes = {
  GET_DB_PATH_REQUEST: "GET_DB_PATH_REQUEST",
  GET_DB_PATH_SUCCESS: "GET_DB_PATH_SUCCESS",
  GET_DB_PATH_FAILURE: "GET_DB_PATH_FAILURE",
  READ_DB_REQUEST: "READ_DB_REQUEST",
  READ_DB_SUCCESS: "READ_DB_SUCCESS",
  READ_DB_FAILURE: "READ_DB_FAILURE",
}

export const getDbPathRequest = () => {
  return { type: actionTypes.GET_DB_PATH_REQUEST }
}

export const getDbPathSuccess = (databasePath) => {
  return {
    type: actionTypes.GET_DB_PATH_SUCCESS,
    databasePath,
  }
}

export const getDbPathFailure = (errorMessage) => {
  return {
    type: actionTypes.GET_DB_PATH_FAILURE,
    errorMessage,
  }
}

export const getDbPath = (renderer = ipcRenderer) => {
  return (dispatch) => {
    dispatch(getDbPathRequest())
    return renderer
      .invoke("select-mounted-volume")
      .then((res) => {
        if (res === undefined) {
          return Promise.reject("something broke?!")
        }
        return buildSqlitePath(res[0])
      })
      .then((data) => dispatch(getDbPathSuccess(data)))
      .catch((error) => dispatch(getDbPathFailure(error)))
  }
}

export const readDbRequest = (databasePath) => {
  return {
    type: actionTypes.READ_DB_REQUEST,
    databasePath,
  }
}

export const readDbSuccess = (database) => {
  return {
    type: actionTypes.READ_DB_SUCCESS,
    database,
  }
}

export const readDbFailure = (errorMessage) => {
  return {
    type: actionTypes.READ_DB_FAILURE,
    errorMessage,
  }
}

export const readDb = (path, renderer = ipcRenderer) => {
  return (dispatch) => {
    dispatch(readDbRequest())
    return renderer
      .invoke("read-database", { path })
      .then((database) => {
        console.log(database)
        return true
      })
      .catch((error) => dispatch(readDbFailure(error)))
  }
}
