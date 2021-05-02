import React, { Component } from "react"
import { connect } from "react-redux"
import { getDbPath, readDb } from "../actions.jsx"

import logo from "../logo.png"

class DeviceSelection extends Component {
  constructor(props) {
    super(props)
    this.readDatabase = this.readDatabase.bind(this)
  }

  readDatabase() {
    this.props.readDb("/Users/marcus/Desktop/KoboReader.sqlite")
  }

  render() {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-gray-800 py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-md w-full space-y-8">
          <div>
            <img
              className="mx-auto h-36 w-auto logo-animation"
              src={logo}
              alt="Octowise logo, which is a cartoon octopus reading a book."
            />
            <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900 dark:text-gray-100">
              Octowise
            </h2>
            <p className="mt-2 text-center text-sm text-gray-600 dark:text-gray-400">
              Sync your Kobo highlights (unofficially) with Readwise
            </p>
          </div>
          <div>
            <button
              onClick={this.readDatabase}
              type="submit"
              className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 dark:bg-indigo-500 hover:bg-indigo-700 dark:bg-indigo-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Select your mounted Kobo device
            </button>
          </div>
          {this.props.errorMessage && (
            <p className="text-center text-sm text-red-600 dark:text-red-400">
              {this.props.errorMessage}
            </p>
          )}
        </div>
      </div>
    )
  }
}

const mapStoreToProps = (store) => {
  return {
    errorMessage: store.errorMessage,
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    getDbPath: () => dispatch(getDbPath()),
    readDb: (path) => dispatch(readDb(path)),
  }
}

export default connect(mapStoreToProps, mapDispatchToProps)(DeviceSelection)
