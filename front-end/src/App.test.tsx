import { render, screen } from '@testing-library/react';
import App from './App';

test('initial render', () => {
  render(<App />);
  const countElement = screen.getByText(/Your count/i);
  const btnElement = screen.getByText(/Push me!/i);
  const messageElement = screen.queryByText(/Fizz|Buzz|FizzBuzz/i);
  expect(countElement).toBeInTheDocument();
  expect(btnElement).toBeInTheDocument();
  expect(messageElement).not.toBeInTheDocument();
});
