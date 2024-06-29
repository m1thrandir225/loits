/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
  safelist: [],
  theme: {
    extend: {
      fontFamily: {
        yung: ['"Yeon Sung"', "system-ui"],
        nanum: ['"Nanum Myeongjo"', "serif"],
      },
    },
  },
};
