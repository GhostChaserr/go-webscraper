interface Config {
  ENVIRONMENT: string
  API_URL: string
}

let config: Config = {
  ENVIRONMENT: import.meta.env.VITE_ENV,
  API_URL: 'http://localhost:4000',
}

switch (import.meta.env.VITE_ENV) {
  case 'development':
    config.API_URL = 'http://localhost:4000'
    config.ENVIRONMENT = 'development'
  break

  case 'local':
    config.API_URL = 'http://localhost:4000'
    config.ENVIRONMENT = 'development'
  break

  default:
    config.API_URL = 'http://localhost:4000'
    config.ENVIRONMENT = 'local'

}

export default config