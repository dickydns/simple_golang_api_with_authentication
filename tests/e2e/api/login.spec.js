import { test, expect } from '@playwright/test';

test('login testes', async ({ request }) => {
  const response = await request.post('http://127.0.0.1:3100/auth/login', {
    data: {
      email: 'dickydraknes@gmail.com',
      password: 'asder122s'
    }
  });

  expect(response.status()).toBe(200);
  const body = await response.json();
  expect(body).toHaveProperty('token');
});