import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'

import { renderWithProviders } from '../../test-utils'
import UserNavigation from './index'

describe('the UserNavigation', () => {
  let result

  beforeEach(() => {
    result = renderWithProviders(
      <Router>
        <UserNavigation fullName="John Doe" />
        <div data-testid="outside-user-menu" />
      </Router>,
    )
  })

  it('should display the the full name and avatar', () => {
    const { getByTestId } = result

    expect(getByTestId('full-name').textContent).toEqual('John Doe')
    expect(getByTestId('avatar')).toBeTruthy()
  })

  it('hides the user menu by default', () => {
    const { queryByTestId } = result

    expect(queryByTestId('user-menu-options')).toBeNull()
  })

  describe('and toggled the menu', () => {
    beforeEach(() => {
      const { queryByTestId } = result
      queryByTestId('user-menu-toggle').click()
    })

    it('should display the user menu', async () => {
      const { queryByTestId } = result
      expect(queryByTestId('user-menu-options')).toBeTruthy()
    })

    describe('on blur', () => {
      beforeEach(() => {
        const { queryByTestId } = result
        queryByTestId('user-menu-toggle').click()
        queryByTestId('outside-user-menu').click()
      })

      it('should hide the user menu', async () => {
        const { queryByTestId } = result
        expect(queryByTestId('user-menu-options')).toBeNull()
      })
    })
  })
})
