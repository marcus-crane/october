const { app, BrowserWindow, ipcMain, dialog } = require("electron")
const fs = require("fs")
const Database = require("better-sqlite3")

// Handle creating/removing shortcuts on Windows when installing/uninstalling.
if (require("electron-squirrel-startup")) {
  app.quit()
}

const createWindow = () => {
  // Create the browser window.
  const mainWindow = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      nodeIntegration: true,
      contextIsolation: false,
    },
  })

  // and load the index.html of the app.
  mainWindow.loadURL(MAIN_WINDOW_WEBPACK_ENTRY) // eslint-disable-line

  // Open the DevTools.
  mainWindow.webContents.openDevTools()
}

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on("ready", createWindow)

// Quit when all windows are closed, except on macOS. There, it's common
// for applications and their menu bar to stay active until the user quits
// explicitly with Cmd + Q.
app.on("window-all-closed", () => {
  if (process.platform !== "darwin") {
    app.quit()
  }
})

app.on("activate", () => {
  // On OS X it's common to re-create a window in the app when the
  // dock icon is clicked and there are no other windows open.
  if (BrowserWindow.getAllWindows().length === 0) {
    createWindow()
  }
})

// In this file you can include the rest of your app's specific main process
// code. You can also put them in separate files and import them here.

ipcMain.handle("select-mounted-volume", () =>
  dialog.showOpenDialogSync({ properties: ["openDirectory"] })
)

ipcMain.handle("read-database", (_, { path }) => {
  if (!fs.existsSync(path)) {
    throw Error(
      "Couldn't find a Kobo database. Are you sure this is the correct path?"
    )
  }
  const db = new Database(path, {
    readonly: true,
    fileMustExist: true,
    verbose: console.error,
  })

  const bookQuery = db.prepare(
    "SElECT DISTINCT b.VolumeID, c.Title, c.Attribution, c.___PercentRead FROM Content c INNER JOIN Bookmark b on c.ContentID = b.VolumeID WHERE c.ContentType = 6 AND c.MimeType = 'application/x-kobo-epub+zip'"
  )

  const bookResult = bookQuery.all()

  // const query = db.prepare(
  //   "SELECT b.VolumeID, c.Title, c.Attribution, c.___PercentRead, b.Text, b.Annotation FROM Content c INNER JOIN Bookmark b ON c.ContentID = b.VolumeID WHERE c.ContentType = 6 AND c.MimeType = 'application/x-kobo-epub+zip'"
  // )
  // const result = query.all()

  return { books: bookResult }
})
