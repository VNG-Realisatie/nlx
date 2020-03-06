import React from 'react';
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { render } from '@testing-library/react';
import App from './App';

jest.mock('./pages/LoginPage', () => () => (
  <div data-testid="login-page" />
))

test('redirects to /inloggen when navigating to /', async () => {
  const history = createMemoryHistory()
  render(<Router history={history}><App /></Router>)
  expect(history.location.pathname).toEqual('/inloggen')
});

test('the /inloggen route renders the LoginPage', () => {
  const history = createMemoryHistory({ initialEntries: ['/inloggen'] })
  const { getByTestId } = render(<Router history={history}><App /></Router>)
  expect(getByTestId('login-page')).toBeInTheDocument();
})
