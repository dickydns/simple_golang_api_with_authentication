const { Given, Then, When, Before  } = require('@cucumber/cucumber');
const assert = require("assert");


Given('API Data Payload', async function(){
    this.payload = {
        email:process.env.TEST_USER_EMAIL,
        password:process.env.TEST_USER_PASSWORD
    };
})

When('user login dengan email user tester', async function(){
    if (!this.apiContext) {
        throw new Error("apiContext is null - Before hook tidak berjalan");
    }

    this.response = await this.apiContext.post('auth/login', {
        data: this.payload
    });

    this.body = await this.response.json()
})

Then('response status harus {int}', function (status){
    assert.strictEqual(this.response.status(), status);
})

Then('response memiliki token', function(){
    assert.ok(this.body.token)
})