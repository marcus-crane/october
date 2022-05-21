import { Fragment, useEffect, useState } from 'react';
import Navbar from "../components/Navbar";
import { Menu, Transition } from '@headlessui/react'
import {
  ArchiveIcon as ArchiveIconSolid,
  ChevronDownIcon,
  ChevronUpIcon,
  DotsVerticalIcon,
  FolderDownloadIcon,
  PencilIcon,
  ReplyIcon,
  UserAddIcon,
} from '@heroicons/react/solid'
import { ListBooksOnDevice, ListBookmarksByID } from '../../wailsjs/go/backend/Kobo'
import { backend } from '../../wailsjs/go/models';
import { formatDistanceToNow } from 'date-fns'

function classNames(...classes: Array<string>) {
  return classes.filter(Boolean).join(' ')
}

export default function Library() {
    const [readwiseConfigured, setReadwiseConfigured] = useState(false)
    const [selectedKobo, setSelectedKobo] = useState({})
    const [content, setContent] = useState<Array<backend.Content>>([]);
    const [bookmarks, setBookmarks] = useState<Array<backend.Bookmark>>([])

    useEffect(() => {
      ListBooksOnDevice()
        .then(content => {
          if (content instanceof Error) {
            throw content
          } else {
            setContent(content)
            if (content.length) {
              getBookmarksByID(content[4].content_id)
            }
          }
        })
        .catch(err => console.log(err))
    }, [content.length])

    function getBookmarksByID(contentID: string) {
      ListBookmarksByID(contentID)
        .then(bookmarks => {
          console.log(bookmarks)
          if (bookmarks instanceof Error) {
            throw bookmarks
          } else {
            setBookmarks(bookmarks)
          }
        })
        .catch(err => console.log(err))
    }

    console.log(bookmarks)

    return (
      <>
        <Navbar />
        <main className="min-w-0 flex-1 border-t border-gray-200 xl:flex">
            <section
              aria-labelledby="message-heading"
              className="min-w-0 flex-1 h-full flex flex-col overflow-hidden xl:order-last"
            >
              {/* Top section */}
              <div className="flex-shrink-0 bg-white border-b border-gray-200">
                {/* Toolbar*/}
                <div className="h-16 flex flex-col justify-center">
                  <div className="px-4 sm:px-6 lg:px-8">
                    <div className="py-3 flex justify-between">
                      {/* Left buttons */}
                      <div>
                        <div className="relative z-0 inline-flex shadow-sm rounded-md sm:shadow-none sm:space-x-3">
                          <span className="inline-flex sm:shadow-sm">
                            <button
                              type="button"
                              className="relative inline-flex items-center px-4 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-900 hover:bg-gray-50 focus:z-10 focus:outline-none focus:ring-1 focus:ring-blue-600 focus:border-blue-600"
                            >
                              <ReplyIcon className="mr-2.5 h-5 w-5 text-gray-400" aria-hidden="true" />
                              <span>Reply</span>
                            </button>
                            <button
                              type="button"
                              className="hidden sm:inline-flex -ml-px relative items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-900 hover:bg-gray-50 focus:z-10 focus:outline-none focus:ring-1 focus:ring-blue-600 focus:border-blue-600"
                            >
                              <PencilIcon className="mr-2.5 h-5 w-5 text-gray-400" aria-hidden="true" />
                              <span>Note</span>
                            </button>
                            <button
                              type="button"
                              className="hidden sm:inline-flex -ml-px relative items-center px-4 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-900 hover:bg-gray-50 focus:z-10 focus:outline-none focus:ring-1 focus:ring-blue-600 focus:border-blue-600"
                            >
                              <UserAddIcon className="mr-2.5 h-5 w-5 text-gray-400" aria-hidden="true" />
                              <span>Assign</span>
                            </button>
                          </span>

                          <span className="hidden lg:flex space-x-3">
                            <button
                              type="button"
                              className="hidden sm:inline-flex -ml-px relative items-center px-4 py-2 rounded-md border border-gray-300 bg-white text-sm font-medium text-gray-900 hover:bg-gray-50 focus:z-10 focus:outline-none focus:ring-1 focus:ring-blue-600 focus:border-blue-600"
                            >
                              <ArchiveIconSolid className="mr-2.5 h-5 w-5 text-gray-400" aria-hidden="true" />
                              <span>Archive</span>
                            </button>
                            <button
                              type="button"
                              className="hidden sm:inline-flex -ml-px relative items-center px-4 py-2 rounded-md border border-gray-300 bg-white text-sm font-medium text-gray-900 hover:bg-gray-50 focus:z-10 focus:outline-none focus:ring-1 focus:ring-blue-600 focus:border-blue-600"
                            >
                              <FolderDownloadIcon className="mr-2.5 h-5 w-5 text-gray-400" aria-hidden="true" />
                              <span>Move</span>
                            </button>
                          </span>

                          <Menu as="div" className="-ml-px relative block sm:shadow-sm lg:hidden">
                            <div>
                              <Menu.Button className="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-900 hover:bg-gray-50 focus:z-10 focus:outline-none focus:ring-1 focus:ring-blue-600 focus:border-blue-600 sm:rounded-md sm:px-3">
                                <span className="sr-only sm:hidden">More</span>
                                <span className="hidden sm:inline">More</span>
                                <ChevronDownIcon
                                  className="h-5 w-5 text-gray-400 sm:ml-2 sm:-mr-1"
                                  aria-hidden="true"
                                />
                              </Menu.Button>
                            </div>

                            <Transition
                              as={Fragment}
                              enter="transition ease-out duration-100"
                              enterFrom="transform opacity-0 scale-95"
                              enterTo="transform opacity-100 scale-100"
                              leave="transition ease-in duration-75"
                              leaveFrom="transform opacity-100 scale-100"
                              leaveTo="transform opacity-0 scale-95"
                            >
                              <Menu.Items className="origin-top-right absolute right-0 mt-2 w-36 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 focus:outline-none">
                                <div className="py-1">
                                  <Menu.Item>
                                    {({ active }) => (
                                      <a
                                        href="#"
                                        className={classNames(
                                          active ? 'bg-gray-100 text-gray-900' : 'text-gray-700',
                                          'block sm:hidden px-4 py-2 text-sm'
                                        )}
                                      >
                                        Note
                                      </a>
                                    )}
                                  </Menu.Item>
                                  <Menu.Item>
                                    {({ active }) => (
                                      <a
                                        href="#"
                                        className={classNames(
                                          active ? 'bg-gray-100 text-gray-900' : 'text-gray-700',
                                          'block sm:hidden px-4 py-2 text-sm'
                                        )}
                                      >
                                        Assign
                                      </a>
                                    )}
                                  </Menu.Item>
                                  <Menu.Item>
                                    {({ active }) => (
                                      <a
                                        href="#"
                                        className={classNames(
                                          active ? 'bg-gray-100 text-gray-900' : 'text-gray-700',
                                          'block px-4 py-2 text-sm'
                                        )}
                                      >
                                        Archive
                                      </a>
                                    )}
                                  </Menu.Item>
                                  <Menu.Item>
                                    {({ active }) => (
                                      <a
                                        href="#"
                                        className={classNames(
                                          active ? 'bg-gray-100 text-gray-900' : 'text-gray-700',
                                          'block px-4 py-2 text-sm'
                                        )}
                                      >
                                        Move
                                      </a>
                                    )}
                                  </Menu.Item>
                                </div>
                              </Menu.Items>
                            </Transition>
                          </Menu>
                        </div>
                      </div>

                      {/* Right buttons */}
                      <nav aria-label="Pagination">
                        <span className="relative z-0 inline-flex shadow-sm rounded-md">
                          <a
                            href="#"
                            className="relative inline-flex items-center px-4 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-10 focus:outline-none focus:ring-1 focus:ring-blue-600 focus:border-blue-600"
                          >
                            <span className="sr-only">Next</span>
                            <ChevronUpIcon className="h-5 w-5" aria-hidden="true" />
                          </a>
                          <a
                            href="#"
                            className="-ml-px relative inline-flex items-center px-4 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-10 focus:outline-none focus:ring-1 focus:ring-blue-600 focus:border-blue-600"
                          >
                            <span className="sr-only">Previous</span>
                            <ChevronDownIcon className="h-5 w-5" aria-hidden="true" />
                          </a>
                        </span>
                      </nav>
                    </div>
                  </div>
                </div>
                {/* Message header */}
              </div>
              
              {content.length && (
                <div className="min-h-0 flex-1 overflow-y-auto">
                  <div className="bg-white pt-5 pb-6 shadow">
                    <div className="px-4 sm:flex sm:justify-between sm:items-baseline sm:px-6 lg:px-8">
                      <div className="sm:w-0 sm:flex-1">
                        <h1 id="message-heading" className="text-lg font-medium text-gray-900">
                          {content[4].title}
                        </h1>
                        <p className="mt-1 text-sm text-gray-500 truncate">{content[4].attribution}</p>
                      </div>

                      <div className="mt-4 flex items-center justify-between sm:mt-0 sm:ml-6 sm:flex-shrink-0 sm:justify-start">
                        <span className="inline-flex items-center px-3 py-0.5 rounded-full text-sm font-medium bg-green-100 text-green-800">
                          Synced
                        </span>
                        <Menu as="div" className="ml-3 relative inline-block text-left">
                          <div>
                            <Menu.Button className="-my-2 p-2 rounded-full bg-white flex items-center text-gray-400 hover:text-gray-600 focus:outline-none focus:ring-2 focus:ring-blue-600">
                              <span className="sr-only">Open options</span>
                              <DotsVerticalIcon className="h-5 w-5" aria-hidden="true" />
                            </Menu.Button>
                          </div>

                          <Transition
                            as={Fragment}
                            enter="transition ease-out duration-100"
                            enterFrom="transform opacity-0 scale-95"
                            enterTo="transform opacity-100 scale-100"
                            leave="transition ease-in duration-75"
                            leaveFrom="transform opacity-100 scale-100"
                            leaveTo="transform opacity-0 scale-95"
                          >
                            <Menu.Items className="origin-top-right absolute right-0 mt-2 w-56 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 focus:outline-none">
                              <div className="py-1">
                                <Menu.Item>
                                  {({ active }) => (
                                    <button
                                      type="button"
                                      className={classNames(
                                        active ? 'bg-gray-100 text-gray-900' : 'text-gray-700',
                                        'w-full flex justify-between px-4 py-2 text-sm'
                                      )}
                                    >
                                      <span>Copy email address</span>
                                    </button>
                                  )}
                                </Menu.Item>
                                <Menu.Item>
                                  {({ active }) => (
                                    <a
                                      href="#"
                                      className={classNames(
                                        active ? 'bg-gray-100 text-gray-900' : 'text-gray-700',
                                        'flex justify-between px-4 py-2 text-sm'
                                      )}
                                    >
                                      <span>Previous conversations</span>
                                    </a>
                                  )}
                                </Menu.Item>
                                <Menu.Item>
                                  {({ active }) => (
                                    <a
                                      href="#"
                                      className={classNames(
                                        active ? 'bg-gray-100 text-gray-900' : 'text-gray-700',
                                        'flex justify-between px-4 py-2 text-sm'
                                      )}
                                    >
                                      <span>View original</span>
                                    </a>
                                  )}
                                </Menu.Item>
                              </div>
                            </Menu.Items>
                          </Transition>
                        </Menu>
                      </div>
                    </div>
                  </div>
                  {/* Thread section*/}
                  <ul role="list" className="py-4 space-y-2 sm:px-6 sm:space-y-4 lg:px-8">
                    {bookmarks !== null && bookmarks.map((item) => (
                      <li key={item.content_id} className="bg-white px-4 py-6 shadow sm:rounded-lg sm:px-6">
                        <div className="sm:flex sm:justify-between sm:items-baseline">
                          <h3 className="text-base font-medium">
                            <span className="text-gray-900">Highlight</span>
                          </h3>
                          <p className="mt-1 text-sm text-gray-600 whitespace-nowrap sm:mt-0 sm:ml-3">
                            <time dateTime={item.date_created}>{formatDistanceToNow(new Date(item.date_created), { addSuffix: true })}</time>
                          </p>
                        </div>
                        <div
                          className="mt-4 space-y-6 text-sm text-gray-800"
                          dangerouslySetInnerHTML={{ __html: item.text }}
                        />
                      </li>
                    ))}
                  </ul>
                </div>
              )}
            </section>

            {/* Message list*/}
            <aside className="hidden xl:block xl:flex-shrink-0 xl:order-first">
              <div className="h-full relative flex flex-col w-96 border-r border-gray-200 bg-gray-100">
                <div className="flex-shrink-0">
                  <div className="h-16 bg-white px-6 flex flex-col justify-center">
                    <div className="flex items-baseline space-x-3">
                      <h2 className="text-lg font-medium text-gray-900">Kobo Libra 2</h2>
                      <p className="text-sm font-medium text-gray-500">{content.length} books</p>
                    </div>
                  </div>
                  <div className="border-t border-b border-gray-200 bg-gray-50 px-6 py-2 text-sm font-medium text-gray-500">
                    Sorted by date
                  </div>
                </div>
                <nav aria-label="Message list" className="min-h-0 flex-1 overflow-y-auto">
                  <ul role="list" className="border-b border-gray-200 divide-y divide-gray-200">
                    {content.map((content) => (
                      <li
                        key={content.title}
                        className="relative bg-white py-5 px-6 hover:bg-gray-50 focus-within:ring-2 focus-within:ring-inset focus-within:ring-blue-600"
                      >
                        <div className="flex justify-between space-x-3">
                          <div className="min-w-0 flex-1">
                            <a href="#" className="block focus:outline-none">
                              <span className="absolute inset-0" aria-hidden="true" />
                              <p className="text-sm font-medium text-gray-900 truncate">{content.title}</p>
                              <p className="text-sm text-gray-500 truncate">{content.attribution}</p>
                            </a>
                          </div>
                        </div>
                        <div>
                        <time
                            dateTime={content.date_created}
                            className="flex-shrink-0 whitespace-nowrap text-sm text-gray-500"
                          >
                            {formatDistanceToNow(new Date(content.date_created), { addSuffix: true })}
                          </time>
                        </div>
                      </li>
                    ))}
                  </ul>
                </nav>
              </div>
            </aside>
          </main>
      </>
    )
}