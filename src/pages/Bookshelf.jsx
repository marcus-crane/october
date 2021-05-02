import React, { Component } from "react"
import { connect } from "react-redux"

import logo from "../logo.png"

class Bookshelf extends Component {
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
              OctAHHHwise
            </h2>
          </div>
        </div>
      </div>
    )
  }
}

const mapStoreToProps = (store) => {
  return {
    books: store.books,
    errorMessage: store.errorMessage,
  }
}

export default connect(mapStoreToProps)(Bookshelf)
