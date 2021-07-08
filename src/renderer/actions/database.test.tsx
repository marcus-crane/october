import configureMockStore from "redux-mock-store"
import thunk from "redux-thunk"
import expect from "expect"

import * as actions from "./database"

const middlewares = [thunk]
const mockStore = configureMockStore(middlewares)

const { actionTypes } = actions

describe("regular actions", () => {
  it("should create SELECT_DEVICE_PATH_REQUEST when user invokes device volume selection", () => {
    const expectedAction = { type: actionTypes.SELECT_DEVICE_PATH_REQUEST }
    expect(actions.selectDevicePathRequest()).toEqual(expectedAction)
  })

  it("should create SELECT_DEVICE_PATH_SUCCESS when a path is selected", () => {
    const databasePath = "/tmp/.kobo/KoboReader.sqlite"
    const expectedAction = {
      type: actionTypes.SELECT_DEVICE_PATH_SUCCESS,
      databasePath,
    }
    expect(actions.selectDevicePathSuccess(databasePath)).toEqual(
      expectedAction
    )
  })

  it("should create SELECT_DEVICE_PATH_FAILURE when user back out of volume selection", () => {
    const errorMessage = "something broke?!"
    const expectedAction = {
      type: actionTypes.SELECT_DEVICE_PATH_FAILURE,
      errorMessage,
    }
    expect(actions.selectDevicePathFailure(errorMessage)).toEqual(
      expectedAction
    )
  })

  it("should create READ_DATABASE_REQUEST when looking for a db on the selected volume", () => {
    const databasePath = "/tmp/.kobo/KoboReader.sqlite"
    const expectedAction = {
      type: actionTypes.READ_DATABASE_REQUEST,
      databasePath,
    }
    expect(actions.readDatabaseRequest(databasePath)).toEqual(expectedAction)
  })

  it("should create READ_DATABASE_SUCCESS when successfully read contents of database", () => {
    const database = { some_contents: true }
    const expectedAction = {
      type: actionTypes.READ_DATABASE_SUCCESS,
      database,
    }
    expect(actions.readDatabaseSuccess(database)).toEqual(expectedAction)
  })

  it("should create READ_DATABASE_FAILURE when failed to read database contents", () => {
    const errorMessage = "database is locked"
    const expectedAction = {
      type: actionTypes.READ_DATABASE_FAILURE,
      errorMessage,
    }
    expect(actions.readDatabaseFailure(errorMessage)).toEqual(expectedAction)
  })
})

describe("async actions", () => {
  it("should create SELECT_DEVICE_PATH_SUCCESS when a mounted device is selected", () => {
    const databasePath = "/tmp/.kobo/KoboReader.sqlite"
    const expectedActions = [
      { type: actionTypes.SELECT_DEVICE_PATH_REQUEST },
      {
        type: actionTypes.SELECT_DEVICE_PATH_SUCCESS,
        databasePath,
      },
    ]

    const store = mockStore({})

    const fakeIpcRenderer = { invoke: () => Promise.resolve(["/tmp"]) }

    return store
      .dispatch(actions.selectDevicePath(fakeIpcRenderer))
      .then(() => {
        expect(store.getActions()).toEqual(expectedActions)
      })
  })

  it("should create SELECT_DEVICE_PATH_FAILURE when a user backs out or something goes wrong", () => {
    const error = { message: "something broke?!" }
    const expectedActions = [
      { type: actionTypes.SELECT_DEVICE_PATH_REQUEST },
      {
        type: actionTypes.SELECT_DEVICE_PATH_FAILURE,
        errorMessage: error.message,
      },
    ]

    const store = mockStore({})

    const fakeIpcRenderer = { invoke: () => Promise.reject(error) }

    return store
      .dispatch(actions.selectDevicePath(fakeIpcRenderer))
      .then(() => {
        expect(store.getActions()).toEqual(expectedActions)
      })
  })

  it("should create READ_DATABASE_SUCCESS when database is empty", () => {
    const databasePath = "/tmp/validPath"
    const database = {}
    const expectedActions = [
      { type: actionTypes.READ_DATABASE_REQUEST, databasePath },
      {
        type: actionTypes.READ_DATABASE_SUCCESS,
        database,
      },
    ]

    const store = mockStore({})

    const fakeIpcRenderer = { invoke: () => Promise.resolve(database) }

    return store
      .dispatch(actions.readDatabase(databasePath, fakeIpcRenderer))
      .then(() => {
        expect(store.getActions()).toEqual(expectedActions)
      })
  })

  it("should create READ_DATABASE_SUCCESS when database has book entries", () => {
    const databasePath = "/tmp/validPath"
    const database = {
      books: [
        {
          Attribution: "Friedrich Nietzsche",
          Title: "Beyond Good and Evil",
          VolumeID:
            "file:///mnt/onboard/Nietzsche, Friedrich/Beyond Good and Evil - Friedrich Nietzsche.kepub.epub",
          ___PercentRead: 50,
        },
        {
          Attribution: "Voltaire",
          Title: "Candide",
          VolumeID: "file:///mnt/onboard/Voltaire/Candide.kepub.epub",
          ___PercentRead: 100,
        },
      ],
    }
    const expectedActions = [
      { type: actionTypes.READ_DATABASE_REQUEST, databasePath },
      {
        type: actionTypes.READ_DATABASE_SUCCESS,
        database,
      },
    ]

    const store = mockStore({})

    const fakeIpcRenderer = { invoke: () => Promise.resolve(database) }

    return store
      .dispatch(actions.readDatabase(databasePath, fakeIpcRenderer))
      .then(() => {
        expect(store.getActions()).toEqual(expectedActions)
      })
  })

  it("should create READ_DB_FAILURE when fetching database fails", () => {
    const databasePath = "/tmp/invalidPath"
    const error = { message: "database is non existent" }
    const expectedActions = [
      { type: actionTypes.READ_DATABASE_REQUEST, databasePath },
      { type: actionTypes.READ_DATABASE_FAILURE, errorMessage: error.message },
    ]

    const store = mockStore({})

    const fakeIpcRenderer = { invoke: () => Promise.reject(error) }

    return store
      .dispatch(actions.readDatabase(databasePath, fakeIpcRenderer))
      .then(() => {
        expect(store.getActions()).toEqual(expectedActions)
      })
  })
})
