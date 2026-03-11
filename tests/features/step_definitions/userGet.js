const { Then, When } = require('@cucumber/cucumber');
const {expect} = require('chai');

When('ambil data user', async function(){
    if (!this.apiContext) {
        throw new Error("apiContext is null - Before hook tidak berjalan");
    }

    this.response = await this.apiContext.get("user")
    this.body = this.response.data
})

Then('status harus 200',  function(){
    expect(this.response.status).to.equal(200)
})

Then('response memiliki user list',  function(){
    expect(this.response.data).to.be.an('array');
})