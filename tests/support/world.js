const { setWorldConstructor } = require("@cucumber/cucumber");
const axios = require("axios");

class CustomWorld {
  constructor() {
    this.apiContext = null;
    this.payload = null;
    this.response = null;
    this.body = null;
    this.token=null;
  }

  async init() {
    this.apiContext = await axios.create({
      baseURL: "http://localhost:3100",
    });
  }
}

setWorldConstructor(CustomWorld);