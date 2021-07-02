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
      databasePath
    }
    expect(actions.selectDevicePathSuccess(databasePath)).toEqual(expectedAction)
  })

  it("should create SELECT_DEVICE_PATH_FAILURE when user back out of volume selection", () => {
    const errorMessage = "something broke?!"
    const expectedAction = {
      type: actionTypes.SELECT_DEVICE_PATH_FAILURE,
      errorMessage
    }
    expect(actions.selectDevicePathFailure(errorMessage)).toEqual(expectedAction)
  })
})

describe("async actions", () => {
  it("should create SELECT_DEVICE_PATH_SUCCESS when a mounted device is selected", () => {
    const databasePath = "/tmp/.kobo/KoboReader.sqlite"
    const expectedActions = [
      { type: actionTypes.SELECT_DEVICE_PATH_REQUEST },
      { 
        type: actionTypes.SELECT_DEVICE_PATH_SUCCESS,
        databasePath
      }
    ]

    const store = mockStore({})

    const fakeIpcRenderer = { invoke: () => Promise.resolve([ "/tmp" ]) }

    return store.dispatch(actions.selectDevicePath(fakeIpcRenderer)).then(() => {
      expect(store.getActions()).toEqual(expectedActions)
    })
  })

  it("should create SELECT_DEVICE_PATH_FAILURE when a user backs out or something goes wrong", () => {
    const error = { message: "something broke?!" }
    const expectedActions = [
      { type: actionTypes.SELECT_DEVICE_PATH_REQUEST },
      { 
        type: actionTypes.SELECT_DEVICE_PATH_FAILURE,
        errorMessage: error.message
      }
    ]

    const store = mockStore({})

    const fakeIpcRenderer = { invoke: () => Promise.reject(error) }

    return store.dispatch(actions.selectDevicePath(fakeIpcRenderer)).then(() => {
      expect(store.getActions()).toEqual(expectedActions)
    })
  })
})