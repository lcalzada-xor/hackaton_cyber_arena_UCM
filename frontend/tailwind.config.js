/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{html,ts}",
  ],
  theme: {
    extend: {
      colors: {
        'cyber-black': '#0f172a', // Deep blue-black
        'cyber-dark': '#1e293b', // Slate 800
        'cyber-primary': '#a855f7', // Purple 500
        'cyber-secondary': '#ec4899', // Pink 500
        'cyber-accent': '#3b82f6', // Blue 500
        'cyber-text': '#f8fafc', // Slate 50
        'cyber-muted': '#94a3b8', // Slate 400
      },
      fontFamily: {
        sans: ['Outfit', 'Inter', 'sans-serif'],
      },
      boxShadow: {
        'glow': '0 0 20px rgba(168, 85, 247, 0.5)',
        'glow-sm': '0 0 10px rgba(168, 85, 247, 0.3)',
      },
    },
  },
  plugins: [],
}
