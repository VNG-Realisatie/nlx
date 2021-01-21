// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { renderWithProviders, waitFor } from '../../test-utils'
import { UserContextProvider } from '../../user-context'
import UserNavigation from './index'

describe('the UserNavigation', () => {
  describe('when not authenticated', () => {
    it('should not render', () => {
      const authenticationHandler = () => {
        throw new Error('not authenticated')
      }
      const { getByTestId } = renderWithProviders(
        <Router>
          <UserContextProvider
            user={null}
            fetchAuthenticatedUser={authenticationHandler}
          >
            <UserNavigation />
          </UserContextProvider>
        </Router>,
      )

      expect(() => getByTestId('user-navigation')).toThrow()
    })
  })

  describe('when authenticated', () => {
    let result

    beforeEach(() => {
      result = renderWithProviders(
        <Router>
          <UserContextProvider
            user={{
              fullName: 'John Doe',
              pictureUrl: 'https://my-pictures.com/nlx.jpg',
            }}
          >
            <UserNavigation />
            <div data-testid="outside-user-menu" />
          </UserContextProvider>
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

      expect(queryByTestId('user-menu-options')).not.toBeInTheDocument()
    })

    describe('and toggled the menu', () => {
      beforeEach(() => {
        const { queryByLabelText } = result
        queryByLabelText('Account menu').click()
      })

      it('should display the user menu', async () => {
        const { queryByTestId } = result
        expect(queryByTestId('user-menu-options')).toBeTruthy()
      })

      describe('on blur', () => {
        beforeEach(() => {
          const { queryByTestId, queryByLabelText } = result
          queryByLabelText('Account menu').click()
          queryByTestId('outside-user-menu').click()
        })

        it('should hide the user menu', async () => {
          const { queryByTestId } = result
          await waitFor(() => {
            expect(queryByTestId('user-menu-options')).toBeNull()
          })
        })
      })
    })
  })
})
