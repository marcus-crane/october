import React from "react"

const BookCard = (book) => (
  <li key={book.VolumeID}>
    <a className="hover:bg-light-blue-500 hover:border-transparent hover:shadow-lg group block rounded-lg p-4 border border-gray-200">
      <div>
        <p className="font-semibold text-gray-900 dark:text-gray-100">
          {book.Attribution}
        </p>
        <p className="text-gray-900 dark:text-gray-100">{book.Title}</p>
        Reading progress:{" "}
        <progress value={book.___PercentRead} max="100">
          {book.___PercentRead}%
        </progress>
      </div>
    </a>
  </li>
)

export default BookCard
