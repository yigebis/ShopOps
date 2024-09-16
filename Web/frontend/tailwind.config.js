/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
          borderRadius: {
            'custom-shape': '0% 20% 90% 0% ',
          },
          boxShadow: {
            'custom-shadow': 'rgba(99, 99, 99, 0.2) 0px 2px 8px 0px',
          },
    
    },
  },
  plugins: [],
}