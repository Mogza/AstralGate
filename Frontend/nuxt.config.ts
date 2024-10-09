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
      ],
    },
  },
  css: [
    '@/assets/css/main.css'
  ],
})