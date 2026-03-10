Feature: Login API

Scenario: User berhasil login
  Given API Data Payload
  When user login dengan email user tester
  Then response status harus 200
  And response memiliki token
  And response memiliki isi token 