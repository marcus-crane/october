import configureMockStore from 'redux-mock-store'
import thunk from 'redux-thunk'
import * as actions from './actions'
import expect from 'expect'

const middlewares = [thunk]
const mockStore = configureMockStore(middlewares)

describe("regular actions", () => {
  it('should create READ_KOBO_DB_REQUEST when user selects a volume', () => {
    const expectedAction = { type: actions.READ_KOBO_DB_REQUEST }
    expect(actions.readKoboDbRequest()).toEqual(expectedAction)
  })
})

describe("async actions", () => {

  it('should create READ_KOBO_DB_SUCCESS when fetch database succeeds', () => {

    const expectedActions = [
      { type: actions.READ_KOBO_DB_REQUEST },
      { type: actions.READ_KOBO_DB_SUCCESS, databasePath: '/tmp/.kobo/KoboReader.sqlite' }
    ]

    const store = mockStore({ database: {} })

    const fakeRenderer = {'invoke': () => Promise.resolve(['/tmp']) }

    return store.dispatch(actions.readKoboDb(fakeRenderer)).then(() => {
      expect(store.getActions()).toEqual(expectedActions)
    })
  })

  it('should create READ_KOBO_DB_FAILURE when fetch database fails', () => {

    const expectedActions = [
      { type: actions.READ_KOBO_DB_REQUEST },
      { type: actions.READ_KOBO_DB_FAILURE, error: 'something broke?!' }
    ]

    const store = mockStore({ })

    const fakeRenderer = {'invoke': () => Promise.reject('something broke?!') }

    return store.dispatch(actions.readKoboDb(fakeRenderer)).then(() => {
      expect(store.getActions()).toEqual(expectedActions)
    })
  })
})
