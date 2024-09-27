/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: 'jit',
  content: [
    './**/*.html',
    './**/*.go',
    "./node_modules/flowbite/**/*.js",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('flowbite/plugin')
  ],
}