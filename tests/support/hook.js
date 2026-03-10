require('dotenv').config({ path: '../.env' });
const { Before, After } = require('@cucumber/cucumber');

Before(async function () {
  await this.init();
});

After(async function () {
  this.apiContext = null;
});