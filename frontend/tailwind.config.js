const colors = require('tailwindcss/colors')

module.exports = {
  content: [
    "./dist/index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        cyan: colors.cyan,
      },
    },
  },
  plugins: [
    require("@tailwindcss/forms"),
    require('@tailwindcss/line-clamp')
  ],
}
