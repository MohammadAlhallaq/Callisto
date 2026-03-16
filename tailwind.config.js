/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./index.html', './src/**/*.{ts,tsx}'],
  theme: {
    extend: {
      colors: {
        base: '#09090b',
        mantle: '#0a0a0c',
        surface0: '#18181b',
        surface1: '#27272a',
        overlay0: '#71717a',
        text: '#fafafa',
        accent: '#6366f1',
        'accent-hover': '#818cf8',
        red: '#ef4444',
        green: '#22c55e',
        yellow: '#eab308',
      },
    },
  },
  plugins: [],
};
