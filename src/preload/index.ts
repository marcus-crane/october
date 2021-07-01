import { contextBridge } from "electron"

const refObjectName = "electron"

const api = {
  versions: process.versions,
  environment: process.env.NODE_ENV
}

if (process.env.NODE_ENV === "development") {
  contextBridge.exposeInMainWorld(refObjectName, api)
} else {
  const deepFreeze = (obj) => {
    if (typeof obj === "object" && obj !== null) {
      Object.keys(obj).forEach(prop => {
        const val = obj[prop]
        if ((typeof val === "object" || typeof val === "function") && !Object.isFrozen(val)) {
          deepFreeze(val)
        }
      })
    }

    return Object.freeze(obj)
  }

  deepFreeze(api)

  window[refObjectName] = api

  window.electronRequire = require
}
