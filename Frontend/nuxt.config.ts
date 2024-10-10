export default defineNuxtConfig({
  app: {
    head: {
      title: 'AstralGate',
      meta: [
        { name: 'description', content: 'Welcome to AstralGate!' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
        { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Lato:ital,wght@0,100;0,300;0,400;0,700;0,900;1,100;1,300;1,400;1,700;1,900&display=swap" rel="stylesheet' }
      ],
    },
  },

  css: [
    '@/assets/css/main.css'
  ],

  compatibilityDate: '2024-10-10',
})