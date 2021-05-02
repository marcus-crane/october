import React, { Component } from "react"
import { connect } from "react-redux"
import { Link } from "react-router-dom"

import BookCard from "../components/BookCard.jsx"

import logo from "../logo.png"

class Bookshelf extends Component {
  render() {
    return (
      <div className="min-h-screen bg-gray-100 dark:bg-gray-800">
        <section className="px-4 sm:px-6 lg:px-4 xl:px-6 pt-4 pb-4 space-y-4">
          <header className="flex items-center justify-between">
            <h2 className="text-md leading-6 font-medium text-gray-900 dark:text-gray-100">
              {this.props.books ? this.props.books.length : 0} Books with
              highlights detected
            </h2>
            <button className="hover:bg-light-blue-200 hover:text-light-blue-800 group flex items-center rounded-md bg-light-blue-100 text-gray-900 dark:text-gray-100 text-sm font-medium px-4 py-2">
              <svg
                className="group-hover:text-light-blue-600 text-gray-900 dark:text-gray-100 mr-2"
                width="12"
                height="20"
                fill="currentColor"
              >
                <path
                  fillRule="evenodd"
                  clipRule="evenodd"
                  d="M6 5a1 1 0 011 1v3h3a1 1 0 110 2H7v3a1 1 0 11-2 0v-3H2a1 1 0 110-2h3V6a1 1 0 011-1z"
                />
              </svg>
              Sync with Readwise
            </button>
            <Link className="text-gray-900 dark:text-gray-100" to="/">
              Pick a different Kobo device
            </Link>
          </header>
          <form className="relative">
            <svg
              width="20"
              height="20"
              fill="currentColor"
              className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-900 dark:text-gray-100"
            >
              <path
                fillRule="evenodd"
                d="M18 8a6 6 0 01-7.743 5.743L10 14l-1 1-1 1H6v2H2v-4l4.257-4.257A6 6 0 1118 8zm-6-4a1 1 0 100 2 2 2 0 012 2 1 1 0 102 0 4 4 0 00-4-4z"
                clipRule="evenodd"
              />
            </svg>
            <input
              className="w-full bg-gray-100 dark:bg-gray-800 text-gray-900 dark:text-gray-100 placeholder-gray-500 dark:placeholder-gray-200 border border-gray-200 rounded-md py-2 pl-10 text-sm"
              type="text"
              aria-label="Readwise sync token"
              placeholder="Paste your Readwise sync token here"
            />
          </form>
          <form className="relative">
            <svg
              width="20"
              height="20"
              fill="currentColor"
              className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-900 dark:text-gray-100"
            >
              <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z"
              />
            </svg>
            <input
              disabled={true}
              className="bg-gray-100 dark:bg-gray-800 focus:border-light-blue-500 focus:ring-1 focus:ring-light-blue-500 focus:outline-none w-full text-sm text-gray-900 dark:text-gray-100 placeholder-gray-500 dark:placeholder-gray-200 border border-gray-200 rounded-md py-2 pl-10"
              type="text"
              aria-label="Filter books"
              placeholder="Filter books (disabled)"
            />
          </form>
          <ul className="grid grid-cols-1 gap-4">
            {!this.props.books && (
              <li className="hover:shadow-lg flex rounded-lg">
                <a className="text-gray-900 dark:text-gray-100 hover:shadow-xs w-full flex items-center justify-center rounded-lg border-2 border-dashed border-gray-200 text-sm font-medium py-4">
                  Your first book could be here
                </a>
              </li>
            )}
            {this.props.books &&
              this.props.books.map((book) => <BookCard {...book} />)}
          </ul>
        </section>
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
