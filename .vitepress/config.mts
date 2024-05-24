import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "October",
  description: "Getting highlights off of your Kobo is very fiddly on a technical level. October is a community-driven desktop application that makes it really simple to send them to Readwise. 100% open source with support for Windows, macOS and Linux!",
  srcDir: "./docs",
  srcExclude: [
    "node_modules",
    "contributions/*.md",
    "miscellaneous/limitations.md",
    "miscellaneous/database-analysis.md"
  ],
  cleanUrls: true,
  lastUpdated: true,
  head: [
    ['link', { rel: "apple-touch-icon", sizes: "180x180", href: "/apple-touch-icon.png"}],
    ['link', { rel: "icon", type: "image/png", sizes: "32x32", href: "/favicon-32x32.png"}],
    ['link', { rel: "icon", type: "image/png", sizes: "16x16", href: "/favicon-16x16.png"}],
    ['link', { rel: "manifest", href: "/site.webmanifest"}],
    ['link', { rel: "shortcut icon", href: "/favicon.ico"}],
    ['link', { rel: "mask-icon", href: "/safari-pinned-tab.svg", color: "#5bbad5"}],
    ['meta', { name: "msapplication-TileColor", content: "#00a300"}],
    ['meta', { name: "theme-color", content: "#ffffff"}]
  ],
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Documentation', link: '/overview' },
      { text: 'Download', link: 'https://github.com/marcus-crane/october/releases' },
      { text: 'Changelog', link: '/changelog' }
    ],
    logo: '/android-chrome-192x192.png',
    editLink: {
      pattern: 'https://github.com/marcus-crane/october/edit/main/docs/:path'
    },
    search: {
      provider: 'local'
    },
    sidebar: [
      {
        text: 'Getting Started',
        items: [
          { text: 'Overview', link: '/overview' },
          { text: 'Prerequisites', link: '/prerequisites' }
        ]
      },
      {
        text: 'Installation',
        items: [
          { text: 'Windows', link: '/installation/windows' },
          { text: 'macOS', link: '/installation/macos' },
          { text: 'Linux', link: '/installation/linux' }
        ]
      },
      {
        text: 'Extras',
        items: [
          { text: 'Changelog', link: '/changelog' },
          { text: 'LICENSE', link: '/miscellaneous/license' },
          { text: 'Reading Formats', link: '/miscellaneous/reading-formats' },
          { text: 'Thanks', link: '/thanks' }
        ]
      }
    ],
    footer: {
      message: 'Source code released under the <a href="https://github.com/marcus-crane/october/blob/main/LICENSE">MIT License</a>.',
      copyright: '<a href="https://soco-st.com/18147">Hero image</a> by <a href="https://soco-st.com">Soco St</a> used under the <a href="https://creativecommons.org/licenses/by/4.0/">CC Attribution License</a>.'
    },
    socialLinks: [
      { icon: 'github', link: 'https://github.com/marcus-crane/october' }
    ],
  },
  sitemap: {
    hostname: "https://october.utf9k.net"
  }
})
