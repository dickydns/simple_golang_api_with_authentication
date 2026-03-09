import { defineConfig } from '@playwright/test';

export default defineConfig({

  testDir: './',

  timeout: 30000,

  use: {
    baseURL: 'http://localhost:3100',
    headless: true
  },

  webServer: {
    command: 'go run main.go',
    port: 3100,
    timeout: 20000
  }

});
