import React from 'react';
import { render } from '@testing-library/react';
import App from './App';

test('renders the intro', () => {
  const { getByText } = render(<App />);
  const linkElement = getByText(/nlx management/i);
  expect(linkElement).toBeInTheDocument();
});
