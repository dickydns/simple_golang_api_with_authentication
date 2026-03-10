const { setWorldConstructor } = require("@cucumber/cucumber");
const { request } = require("@playwright/test");

class CustomWorld {
  constructor() {
    this.apiContext = null;
    this.payload = null;
    this.response = null;
    this.body = null;
  }

  async init() {
    this.apiContext = await request.newContext({
      baseURL: "http://localhost:3100",
    });
  }
}

setWorldConstructor(CustomWorld);