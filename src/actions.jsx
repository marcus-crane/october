import { ipcRenderer } from 'electron'

import { GET_KOBO_SQLITE_PATH } from './constants'

export const READ_KOBO_DB_REQUEST = 'READ_KOBO_DB_REQUEST'
export const READ_KOBO_DB_SUCCESS = 'READ_KOBO_DB_SUCCESS'
export const READ_KOBO_DB_FAILURE = 'READ_KOBO_DB_FAILURE'

export const readKoboDbRequest = () => {
  return { type: READ_KOBO_DB_REQUEST }
}

export const readKoboDbSuccess = (databasePath) => {
  return {
    type: READ_KOBO_DB_SUCCESS,
    databasePath
  }
}

export const readKoboDbFailure = (error) => {
  return {
    type: READ_KOBO_DB_FAILURE,
    error
  }
}

export const readKoboDb = (renderer = ipcRenderer) => {
  return dispatch => {
    dispatch(readKoboDbRequest())
    return renderer.invoke('select-mounted-volume')
      .then(res => {
        if (res === undefined) {
          return Promise.reject('something broke?!')
        }
        return GET_KOBO_SQLITE_PATH(res[0])
      })
      .then(data => dispatch(readKoboDbSuccess(data)))
      .catch(error => dispatch(readKoboDbFailure(error)))
  }
}
