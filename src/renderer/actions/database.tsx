import { ipcRenderer } from "electron"

import { buildSqlitePath } from "../constants"

export const actionTypes = {
  SELECT_DEVICE_PATH_REQUEST: "SELECT_DEVICE_PATH_REQUEST",
  SELECT_DEVICE_PATH_SUCCESS: "SELECT_DEVICE_PATH_SUCCESS",
  SELECT_DEVICE_PATH_FAILURE: "SELECT_DEVICE_PATH_FAILURE",
  READ_DATABASE_REQUEST: "READ_DATABASE_REQUEST",
  READ_DATABASE_SUCCESS: "READ_DATABASE_SUCCESS",
  READ_DATABASE_FAILURE: "READ_DATABASE_FAILURE",
}

export const selectDevicePathRequest = () => {
  return { type: actionTypes.SELECT_DEVICE_PATH_REQUEST }
}

export const selectDevicePathSuccess = (databasePath) => {
  return {
    type: actionTypes.SELECT_DEVICE_PATH_SUCCESS,
    databasePath,
  }
}

export const selectDevicePathFailure = (errorMessage) => {
  return {
    type: actionTypes.SELECT_DEVICE_PATH_FAILURE,
    errorMessage,
  }
}

export const selectDevicePath = (renderer = ipcRenderer) => {
  return (dispatch) => {
    dispatch(selectDevicePathRequest())
    return renderer
      .invoke("select-mounted-volume")
      .then((res) => {
        if (res === undefined) {
          return Promise.reject("Device selection was aborted.")
        }
        return buildSqlitePath(res[0])
      })
      .then((data) => dispatch(selectDevicePathSuccess(data)))
      .catch((err) => dispatch(selectDevicePathFailure(err.message)))
  }
}

export const readDatabaseRequest = (databasePath) => {
  return {
    type: actionTypes.READ_DATABASE_REQUEST,
    databasePath,
  }
}

export const readDatabaseSuccess = (database) => {
  return {
    type: actionTypes.READ_DATABASE_SUCCESS,
    database,
  }
}

export const readDatabaseFailure = (errorMessage) => {
  return {
    type: actionTypes.READ_DATABASE_FAILURE,
    errorMessage,
  }
}

export const readDatabase = (path, renderer = ipcRenderer) => {
  return (dispatch) => {
    dispatch(readDatabaseRequest(path))
    return renderer
      .invoke("read-database", { path })
      .then((database) => {
        if (Object.keys(database).length === 0) return {}
        if (!Object.keys(database).includes("books")) return {}
        return {
          books: database.books,
        }
      })
      .then((data) => dispatch(readDatabaseSuccess(data)))
      .catch((error) => dispatch(readDatabaseFailure(error.message)))
  }
}
