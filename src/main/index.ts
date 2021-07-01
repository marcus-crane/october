import { join } from "path"
import { pathToFileURL } from "url"

import { app, BrowserWindow, ipcMain, dialog } from "electron"

const isDevelopment = process.env.NODE_ENV === "development"

const isSingleInstance = app.requestSingleInstanceLock()

if (!isSingleInstance) {
  app.quit()
  process.exit(0)
}

app.disableHardwareAcceleration()

let mainWindow = null

const createWindow = async () => {
  mainWindow = new BrowserWindow({
    width: 1280,
    height: 720,
    show: false, // Window is shown when 'ready-to-show' is fired
    webPreferences: {
      contextIsolation: isDevelopment, // Spectron tests require this to be disabled
      enableRemoteModule: isDevelopment // Same as above
    },
  })

  mainWindow.on("ready-to-show", () => {
    mainWindow.show()

    if (isDevelopment) {
      mainWindow.webContents.openDevTools()
    }
  })

  const pageUrl = isDevelopment
    ? "http://localhost:3000"
    : pathToFileURL(join(__dirname, "../renderer/index.html")).toString() // TODO: Check new URL alternative

  await mainWindow.loadURL(pageUrl)
}

app.on("second-instance", () => {
  if (mainWindow) {
    if (mainWindow.isMinimized()) mainWindow.restore()
    mainWindow.focus()
  }
})

app.on("window-all-closed", () => {
  if (process.platform !== "darwin") {
    app.quit()
  }
})

app.whenReady()
  .then(createWindow)
  .catch(e => console.error(`Failed to create window: ${e}`))

if (!isDevelopment) {
  app.whenReady()
    .then(() => import("electron-updater"))
    .then(({ autoUpdater }) => autoUpdater.checkForUpdatesAndNotify())
    .catch(e => console.error(`Failed to check for updates: ${e}`))
}

ipcMain.handle("select-mounted-volume", () => {
  dialog.showOpenDialogSync({ properties: ["openDirectory"] })
})
