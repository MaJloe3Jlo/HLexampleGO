const port = 7777
const apiHost = 'http://localhost'

module.exports = {
  devServer: {
    proxy: {
      '/api': {
        target: `${apiHost}:${port}`,
        changeOrigin: true,
      },
    }
  },
  "transpileDependencies": [
    "vuetify"
  ]
}