import { actionTypes } from "../actions/database"

export default function reducer(store = {}, action) {
  switch (action.type) {
    case actionTypes.SELECT_DEVICE_PATH_SUCCESS:
      return selectDevicePathSuccess(store, action)
    case actionTypes.SELECT_DEVICE_PATH_FAILURE:
      return selectDevicePathFailure(store, action)
    case actionTypes.READ_DATABASE_SUCCESS:
      return readDatabaseSuccess(store, action)
    case actionTypes.READ_DATABASE_FAILURE:
      return readDatabaseFailure(store, action)
    default:
      return store
  }
}

const selectDevicePathSuccess = (store, action) => {
  const { databasePath } = action
  return {
    ...store,
    databasePath,
  }
}

const selectDevicePathFailure = (store, action) => {
  const { errorMessage } = action
  return {
    ...store,
    errorMessage,
  }
}

const readDatabaseSuccess = (store, action) => {
  const { database } = action
  return {
    ...store,
    database
  }
}

const readDatabaseFailure = (store, action) => {
  const { errorMessage } = action
  return {
    ...store,
    errorMessage
  }
}