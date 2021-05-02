import { actionTypes } from "./actions.jsx"

export default function reducer(store = {}, action) {
  switch (action.type) {
    case actionTypes.GET_DB_PATH_SUCCESS:
      return getDbPathSuccess(store, action)
    case actionTypes.GET_DB_PATH_FAILURE:
      return getDbPathFailure(store, action)
    case actionTypes.READ_DB_SUCCESS:
      return readDbSuccess(store, action)
    case actionTypes.READ_DB_FAILURE:
      return readDbFailure(store, action)
    default:
      return store
  }
}

const getDbPathSuccess = (store, action) => {
  const { databasePath } = action
  return {
    ...store,
    databasePath,
  }
}

const getDbPathFailure = (store, action) => {
  const { errorMessage } = action
  return {
    ...store,
    errorMessage,
  }
}

const readDbSuccess = (store, action) => {
  const { database } = action
  return {
    ...store,
    database,
  }
}

const readDbFailure = (store, action) => {
  const { errorMessage } = action
  return {
    ...store,
    errorMessage,
  }
}
