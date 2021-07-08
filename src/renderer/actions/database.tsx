import { ipcRenderer } from "electron"

import { buildSqlitePath } from "../constants"

export const actionTypes = {
  SELECT_DEVICE_PATH_REQUEST: "SELECT_DEVICE_PATH_REQUEST",
  SELECT_DEVICE_PATH_SUCCESS: "SELECT_DEVICE_PATH_SUCCESS",
  SELECT_DEVICE_PATH_FAILURE: "SELECT_DEVICE_PATH_FAILURE",
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
