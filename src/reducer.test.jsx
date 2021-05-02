import reducer from "./reducer"
import { actionTypes } from "./actions.jsx"

describe("reducer", () => {
  it("should return the initial state when receiving an undefined action", () => {
    const expected = {}
    const store = undefined
    const action = "UNKNOWN"
    expect(reducer(store, action)).toEqual(expected)
  })

  it("should handle GET_DB_PATH_SUCCESS", () => {
    const databasePath = "/tmp/.kobo/KoboReader.sqlite"
    const expected = { databasePath }
    const store = {}
    const action = {
      type: actionTypes.GET_DB_PATH_SUCCESS,
      databasePath,
    }
    expect(reducer(store, action)).toEqual(expected)
  })

  it("should handle GET_DB_PATH_FAILURE", () => {
    const errorMessage = "is your hard drive running?"
    const expected = { errorMessage }
    const store = {}
    const action = {
      type: actionTypes.GET_DB_PATH_FAILURE,
      errorMessage,
    }
    expect(reducer(store, action)).toEqual(expected)
  })
})
