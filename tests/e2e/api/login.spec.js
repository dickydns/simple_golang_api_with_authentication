import { test, expect } from '@playwright/test';

test('login testes', async ({ request }) => {
  const response = await request.post('http://127.0.0.1:3100/auth/login', {
    data: {
      email: process.env.TEST_USER_EMAIL,
      password: process.env.TEST_USER_PASSWORD
    }
  });

  


  expect(response.status()).toBe(200);
  const body = await response.json();
  expect(body).toHaveProperty('token');
});