import { test, expect } from '@playwright/test';

test('login testes', async ({ request }) => {
  const email = process.env.TEST_USER_EMAIL;
  const password = process.env.TEST_USER_PASSWORD;

  expect(email).toBeTruthy();
  expect(password).toBeTruthy();

  const response = await request.post('http://127.0.0.1:3100/auth/login', {
    data: {
      email: email,
      password: password
    }
  });


  


  expect(response.status()).toBe(200);
  const body = await response.json();
  expect(body).toHaveProperty('token');
});