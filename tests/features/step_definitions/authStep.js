const { Given, Then, When } = require('@cucumber/cucumber');
const {expect} = require('chai');

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
    this.response = await this.apiContext.post("auth/login", this.payload);
    this.body = this.response.data;
})

Then('response status harus {int}', function (status){
    expect(this.response.status).to.equal(200)
})

Then('response memiliki token', function(){
    expect(this.response.data.token).to.not.be.empty;
})


Then('response memiliki isi token', function(){
    const token = this.response.data.token
    console.log("Token", token)
    expect(this.response.data.token).to.not.be.empty;
})
