const { Application } = require("spectron")
const { assert } = require("assert")

const app = new Application({
  path: require("electron"),
  requireName: "electronRequire",
  args: ["."],
})

app
  .start()
  .then(async () => {
    const isVisible = await app.browserWindow.isVisible()
    assert.ok(isVisible, "Main window is not visible")
  })

  .then(async () => {
    const isDevtoolsOpen = await app.webContents.isDevToolsOpened()
    assert.ok(!isDevtoolsOpen, "Developer tools opened successfully")
  })

  .then(async () => {
    const content = await app.client.$("#app")
    assert.notStrictEqual(
      await content.getHTML(),
      '<div id="app"></div>',
      "Window content is empty"
    )
  })

  .then(() => {
    if (app && app.isRunning()) {
      return app.stop()
    }
  })

  .then(() => process.exit(0))

  .catch((err) => {
    console.error(err)
    if (app && app.isRunning()) {
      app.stop()
    }
    process.exit(1)
  })
