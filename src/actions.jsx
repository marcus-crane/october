import { ipcRenderer } from "electron";

import { buildSqlitePath } from "./constants";

export const GET_DB_PATH_REQUEST = "GET_DB_PATH_REQUEST";
export const GET_DB_PATH_SUCCESS = "GET_DB_PATH_SUCCESS";
export const GET_DB_PATH_FAILURE = "GET_DB_PATH_FAILURE";

export const READ_DB_REQUEST = "READ_DB_REQUEST";
export const READ_DB_SUCCESS = "READ_DB_SUCCESS";
export const READ_DB_FAILURE = "READ_DB_FAILURE";

export const getDbPathRequest = () => {
  return { type: GET_DB_PATH_REQUEST };
};

export const getDbPathSuccess = (databasePath) => {
  return {
    type: GET_DB_PATH_SUCCESS,
    databasePath,
  };
};

export const getDbPathFailure = (error) => {
  return {
    type: GET_DB_PATH_FAILURE,
    error,
  };
};

export const getDbPath = (renderer = ipcRenderer) => {
  return (dispatch) => {
    dispatch(getDbPathRequest());
    return renderer
      .invoke("select-mounted-volume")
      .then((res) => {
        if (res === undefined) {
          return Promise.reject("something broke?!");
        }
        return buildSqlitePath(res[0]);
      })
      .then((data) => dispatch(getDbPathSuccess(data)))
      .catch((error) => dispatch(getDbPathFailure(error)));
  };
};

export const readDbRequest = (databasePath) => {
  return {
    type: READ_DB_REQUEST,
    databasePath,
  };
};

export const readDbSuccess = (database) => {
  return {
    type: READ_DB_SUCCESS,
    database,
  };
};

export const readDbFailure = (error) => {
  return {
    type: READ_DB_FAILURE,
    error,
  };
};
