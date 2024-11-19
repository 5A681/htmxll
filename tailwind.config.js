/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: 'jit',
  content: [
    './views/*.html',
    './**/*.go',
    './**/*.js',
    "./node_modules/flowbite/**/*.js",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('flowbite/plugin')
  ],
}