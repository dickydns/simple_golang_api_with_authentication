import { defineConfig } from '@playwright/test';
import dotenv from 'dotenv';
import path from 'path';

// Load .env dari root project, bukan dari folder tests/e2e
dotenv.config({ path: path.resolve(__dirname, '../../.env') });
import { defineConfig } from '@playwright/test';

export default defineConfig({

  testDir: './',

  timeout: 30000,

  use: {
    baseURL: 'http://localhost:3100',
    headless: true
  },


});
