const { Given, Then, When } = require('@cucumber/cucumber');
const {expect} = require('chai');

Given('payload data user', async function(){
    this.payload = {
        nama :'dickyperdian'+Math.floor(Math.random() * 11),
        email:`perdian${Math.floor(Math.random() * 11)}@gmail.com`,
        password:'asder122'
    }
})

When('membuat data user', async function(){
     if (!this.apiContext) {
        throw new Error("apiContext is null - Before hook tidak berjalan");
    }
    this.response = await this.apiContext.post('/auth/register',this.payload)
    this.body = this.response.data
})

Then('status create harus 201', function(){
    expect(this.response.status).to.equal(201)
})

Then('response data user', function(){
    expect(this.body).to.not.be.empty;
})

