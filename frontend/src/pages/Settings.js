import React, { Component } from 'react';

class Settings extends Component {
  constructor(props) {
    super(props)
    this.state = {
      settings: {}
    }
  }
  render() {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-gray-800 py-12 px-24 grid grid-cols-2 gap-14">
        <div className="space-y-2">
          <h2 className="text-center text-3xl font-extrabold text-gray-900 dark:text-gray-100">
            Settings
          </h2>
        </div>
      </div>
    )
  }
}

export default Settings
