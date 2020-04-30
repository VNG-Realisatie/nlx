// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { act, render } from '@testing-library/react'
import { PREVENT_CACHING_HEADERS } from '../domain/fetch-utils'
import UserContext, { UserContextProvider } from './index'

describe('UserContext', () => {
  describe('Provider', () => {
    describe('on initialisation', () => {
      beforeEach(() => {
        jest.spyOn(global, 'fetch').mockImplementation(() =>
          Promise.resolve({
            ok: true,
            status: 200,
            json: () => Promise.resolve({ id: '42' }),
          }),
        )
      })

      afterEach(() => global.fetch.mockRestore())

      it('should fetch the current user', async () => {
        await act(async () => render(<UserContextProvider />))
        expect(global.fetch).toHaveBeenCalledWith(
          '/oidc/me',
          expect.objectContaining({
            headers: expect.objectContaining(PREVENT_CACHING_HEADERS),
          }),
        )
      })

      it('should make the user available to the context consumers', () => {
        const { getByTestId } = render(
          <UserContextProvider user={{ id: '43' }}>
            <UserContext.Consumer>
              {({ user }) => (
                <div data-testid="child">{user ? user.id : ''}</div>
              )}
            </UserContext.Consumer>
          </UserContextProvider>,
        )

        expect(getByTestId('child')).toHaveTextContent('43')
      })
    })

    describe('when passing a default user', () => {
      beforeEach(() => {
        jest.spyOn(global, 'fetch')
      })

      afterEach(() => global.fetch.mockRestore())

      it('should not fetch the current user', async () => {
        render(<UserContextProvider user={{ id: '42' }} />)

        expect(global.fetch).not.toHaveBeenCalled()
      })
    })
  })
})
