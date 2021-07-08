const now = new Date()
const buildVersion = `${now.getFullYear() - 2000}.${
  now.getMonth() + 1
}.${now.getDate()}`

const config = {
  appId: "net.utf9k.octowise",
  productName: "Octowise",
  copyright: "Copyright © 2021 Marcus Crane",
  directories: {
    output: "dist",
    buildResources: "build",
    app: "app",
  },
  mac: {
    category: "public.app-category.utilities",
    icon: "assets/icon.icns",
  },
  extraMetadata: {
    version: buildVersion,
  },
}

module.exports = config
