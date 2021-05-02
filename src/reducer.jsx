import { actionTypes } from "./actions.jsx"

export default function reducer(store = {}, action) {
  switch (action.type) {
    case actionTypes.GET_DB_PATH_SUCCESS:
      return getDbPathSuccess(store, action)
    case actionTypes.GET_DB_PATH_FAILURE:
      return getDbPathFailure(store, action)
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
