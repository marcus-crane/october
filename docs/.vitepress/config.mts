import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "October",
  description: "Send Kobo highlights to Readwise with a few clicks",
  lastUpdated: true,
  sitemap: {
    hostname: 'https://october.utf9k.net'
  },
  themeConfig: {
    logo: '/logo.png',
    editLink: {
      pattern: 'https://github.com/marcus-crane/october/edit/main/docs/:path'
    },
    search: {
      provider: 'local'
    },
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Docs', link: '/welcome' },
      { text: 'Changelog', link: '/changelog/index.md' }
    ],

    sidebar: [
      {
        text: 'Getting Started',
        items: [
          { text: 'Welcome', link: '/welcome' },
          { text: 'Prerequisites', link: '/prerequisites' },
          { text: 'Changelog', link: '/changelog/index.md' }
        ]
      },
      {
        text: 'Installation',
        items: [
          { text: 'Linux', link: '/installation/linux' },
          { text: 'macOS', link: '/installation/macos' },
          { text: 'Windows', link: '/installation/windows', }
        ]
      },
      {
        text: 'Contributions',
        items: [
          { text: 'Bug Reports', link: '/contributions/bugs' },
          { text: 'Feature Requests', link: '/contributions/feature-requests' },
          { text: 'Questions', link: '/contributions/questions' }
        ]
      },
      {
        text: 'Technical Details',
        items: [
          { text: 'Reading Formats', link: '/misc/reading-formats' },
          { text: 'Database Analysis', link: '/misc/database-analysis' },
          { text: 'Limitations', link: '/misc/limitations' },
        ]
      },
      {
        text: 'Miscellaneous',
        items: [
          { text: 'Thanks', link: '/misc/thanks' },
          { text: 'LICENSE', link: '/misc/license' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/marcus-crane/october' }
    ],
    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © 2019-present Marcus Crane'
    }
  }
})
