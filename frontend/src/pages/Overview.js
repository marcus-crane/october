import React, { Component } from 'react';

class Overview extends Component {
  constructor(props) {
    super(props)
    console.log(props)
    this.state = {
      readwiseConfigured: true,
      selectedKobo: []
    }
  }
  render() {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-gray-800 py-12 px-24 grid grid-cols-2 gap-14">
        <div className="space-y-2">
          <h2 className="text-center text-3xl font-extrabold text-gray-900 dark:text-gray-100">
            Overview
          </h2>
        </div>
      </div>
    )
  }
}

export default Overview
