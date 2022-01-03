import React, { Component } from 'react';
import logo from '../logo.png'

class DeviceSelector extends Component {
  constructor(props) {
    super(props)
    this.state = { devices: [] }
    this.selectDevice = this.selectDevice.bind(this)
    this.detectDevices = this.detectDevices.bind(this)
  }
  detectDevices() {
    window.go.main.KoboService.DetectKobos()
      .then(devices => {
        if (devices == null) {
          this.setState({ devices: [] })
        }
      })
      .catch(err => console.log(err))
  }
  selectDevice(name) {
    console.log(name)
  }
  selectLocalDatabase() {
    window.go.main.KoboService.PromptForLocalDBPath()
      .then(result => console.log(result))
      .catch(err => console.log(err))
  }
  render() {
    this.detectDevices()
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-gray-800 py-12 px-24 grid grid-cols-2 gap-14">
        <div className="space-y-2">
          <img
            className="mx-auto h-36 w-auto logo-animation"
            src={logo}
            alt="The Octowise logo, which is a cartoon octopus reading a book."
          />
          <h2 className="text-center text-3xl font-extrabold text-gray-900 dark:text-gray-100">
            Octowise
          </h2>
          <p className="mt-0 text-center text-sm text-gray-600 dark:text-gray-400">
            Easily access your Kobo highlights
          </p>
        </div>
        <div className="space-y-4 text-center">
          <h1 class="text-3xl font-bold">Select your Kobo</h1>
          <button onClick={this.detectDevices}>Don't see your device? Click here to refresh device list.</button>
          <ul>
            {this.state.devices.map(device => (
              <li key={device.MntPath} onClick={() => this.selectDevice(device.Name)}>
                <a className="bg-red-200 hover:bg-red-500 group block rounded-lg p-4 mb-2">
                  <dl>
                    <div>
                      <dt className="sr-only">Title</dt>
                      <dd className="border-gray leading-6 font-medium text-black">
                        {device.Name}
                      </dd>
                      <dt className="sr-only">System Specifications</dt>
                      <dd className="text-sm font-normal">
                        {device.Storage} GB · {device.DisplayPPI} PPI
                      </dd>
                    </div>
                  </dl>
                </a>
              </li>
            ))}
            {/*<li>*/}
            {/*  <a onClick={this.selectLocalDatabase} className="bg-red-200 hover:bg-red-500 group block rounded-lg p-4">*/}
            {/*    <dl>*/}
            {/*      <div>*/}
            {/*        <dt className="sr-only">Title</dt>*/}
            {/*        <dd className="border-gray leading-6 font-medium text-black">*/}
            {/*          Pick a locally available Kobo database (KoboReader.sqlite)*/}
            {/*        </dd>*/}
            {/*      </div>*/}
            {/*    </dl>*/}
            {/*  </a>*/}
            {/*</li>*/}
          </ul>
        </div>
      </div>
    )
  }
}

export default DeviceSelector
