import configureMockStore from "redux-mock-store"
import thunk from "redux-thunk"
import * as actions from "./actions.jsx"
import expect from "expect"

const middlewares = [thunk]
const mockStore = configureMockStore(middlewares)

const { actionTypes } = actions

describe("regular actions", () => {
  it("should create GET_DB_PATH_REQUEST when user selects a volume", () => {
    const expectedAction = { type: actionTypes.GET_DB_PATH_REQUEST }
    expect(actions.getDbPathRequest()).toEqual(expectedAction)
  })

  it("should create GET_DB_PATH_SUCCESS when a valid path is selected", () => {
    const databasePath = "/tmp/.kobo/KoboReader.sqlite"
    const expectedAction = {
      type: actionTypes.GET_DB_PATH_SUCCESS,
      databasePath,
    }
    expect(actions.getDbPathSuccess(databasePath)).toEqual(expectedAction)
  })

  it("should create GET_DB_PATH_FAILURE when volume selection is cancelled", () => {
    const errorMessage = "something broke?!"
    const expectedAction = {
      type: actionTypes.GET_DB_PATH_FAILURE,
      errorMessage,
    }
    expect(actions.getDbPathFailure(errorMessage)).toEqual(expectedAction)
  })

  it("should create READ_DB_REQUEST when looking for db at detected db path", () => {
    const databasePath = "/tmp/.kobo/KoboReader.sqlite"
    const expectedAction = {
      type: actionTypes.READ_DB_REQUEST,
      databasePath,
    }
    expect(actions.readDbRequest(databasePath)).toEqual(expectedAction)
  })

  it("should create READ_DB_SUCCESS when successfully read contents of database", () => {
    const database = { some_contents: true }
    const expectedAction = {
      type: actionTypes.READ_DB_SUCCESS,
      database,
    }
    expect(actions.readDbSuccess(database)).toEqual(expectedAction)
  })

  it("should create READ_DB_FAILURE when failed to read database contents", () => {
    const errorMessage = "database is locked up by another process"
    const expectedAction = {
      type: actionTypes.READ_DB_FAILURE,
      errorMessage,
    }
    expect(actions.readDbFailure(errorMessage)).toEqual(expectedAction)
  })
})

describe("async actions", () => {
  it("should create GET_DB_PATH_SUCCESS when building db path succeeds", () => {
    const databasePath = "/tmp/.kobo/KoboReader.sqlite"
    const expectedActions = [
      { type: actionTypes.GET_DB_PATH_REQUEST },
      {
        type: actionTypes.GET_DB_PATH_SUCCESS,
        databasePath,
      },
    ]

    const store = mockStore({ database: {} })

    const fakeRenderer = { invoke: () => Promise.resolve(["/tmp"]) }

    return store.dispatch(actions.getDbPath(fakeRenderer)).then(() => {
      expect(store.getActions()).toEqual(expectedActions)
    })
  })

  it("should create GET_DB_PATH_FAILURE when fetch database fails", () => {
    const error = { message: "something broke?!" }
    const expectedActions = [
      { type: actionTypes.GET_DB_PATH_REQUEST },
      { type: actionTypes.GET_DB_PATH_FAILURE, errorMessage: error.message },
    ]

    const store = mockStore({})

    const fakeRenderer = { invoke: () => Promise.reject(error) }

    return store.dispatch(actions.getDbPath(fakeRenderer)).then(() => {
      expect(store.getActions()).toEqual(expectedActions)
    })
  })
})
