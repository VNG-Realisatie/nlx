import React from 'react';
import { render } from '@testing-library/react';
import App from './App';

jest.mock('./pages/LoginPage', () => () => (
  <div data-testid="login-page" />
))

test('redirects to the Login page when navigating to /', () => {
  const { getByTestId } = render(<App />);
  expect(getByTestId('login-page')).toBeInTheDocument();
});
