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

  it("should handle READ_DB_SUCCESS", () => {
    const databaseContents = { something: true }
    const expected = {
      database: databaseContents,
    }
    const store = {}
    const action = {
      type: actionTypes.READ_DB_SUCCESS,
      database: databaseContents,
    }
    expect(reducer(store, action)).toEqual(expected)
  })

  it("should handle READ_DB_FAILURE", () => {
    const errorMessage = "some of undefined is undefined"
    const expected = { errorMessage }
    const store = {}
    const action = {
      type: actionTypes.READ_DB_FAILURE,
      errorMessage,
    }
    expect(reducer(store, action)).toEqual(expected)
  })
})
