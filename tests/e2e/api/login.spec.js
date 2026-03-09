import { test, expect } from '@playwright/test';

test('login testes', async ({ request }) => {
  const response = await request.post('/api/login', {
    data: {
      email: 'dickydraknes@gmail.com',
      password: 'asder122'
    }
  });

  expect(response.status()).toBe(200);
  const body = await response.json();
  expect(body).toHaveProperty('token');
});