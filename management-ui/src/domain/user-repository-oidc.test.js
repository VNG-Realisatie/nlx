// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import UserRepositoryOIDC from './user-repository-oidc'

describe('the UserRepositoryOIDC', () => {
  describe('getting the authenticated user', () => {
    describe('when the user is authenticated', () => {
      beforeEach(() => {
        jest.spyOn(global, 'fetch').mockImplementation(() =>
          Promise.resolve({
            ok: true,
            status: 200,
            json: () =>
              Promise.resolve({
                id: '42',
                fullName: 'full name',
                email: 'email',
                pictureUrl: 'picture url',
              }),
          }),
        )
      })

      afterEach(() => global.fetch.mockRestore())

      it('should return the user', () => {
        expect.assertions(1)
        return expect(
          UserRepositoryOIDC.getAuthenticatedUser(),
        ).resolves.toEqual({
          id: '42',
          fullName: 'full name',
          email: 'email',
          pictureUrl: 'picture url',
        })
      })
    })

    describe('when no user is authenticated', () => {
      beforeEach(() => {
        jest.spyOn(global, 'fetch').mockImplementation(() =>
          Promise.resolve({
            ok: false,
            status: 401,
          }),
        )
      })

      afterEach(() => global.fetch.mockRestore())

      it('should throw an error', () => {
        expect.assertions(1)
        return expect(
          UserRepositoryOIDC.getAuthenticatedUser(),
        ).rejects.toEqual(new Error('no user is authenticated'))
      })
    })

    describe('when an unexpected error happens', () => {
      beforeEach(() => {
        jest.spyOn(global, 'fetch').mockImplementation(() =>
          Promise.resolve({
            ok: false,
            status: 500,
          }),
        )
      })

      afterEach(() => global.fetch.mockRestore())

      it('should throw an error', async () => {
        const user = UserRepositoryOIDC.getAuthenticatedUser()

        await expect(user).rejects.toEqual(
          new Error('unable to handle the request'),
        )

        expect(global.fetch).toHaveBeenCalledWith('/oidc/me')
      })
    })
  })

  describe('logout', () => {
    beforeEach(() => {
      jest.spyOn(global, 'fetch').mockImplementation(() =>
        Promise.resolve({
          ok: true,
          status: 302,
          url: 'url-to-redirect-to',
        }),
      )
    })

    afterEach(() => global.fetch.mockRestore())

    it('should logout', async () => {
      delete window.location
      window.location = {}

      await UserRepositoryOIDC.logout()

      expect(global.fetch).toHaveBeenCalledWith('/oidc/logout', {
        body: 'csrfmiddlewaretoken=undefined',
        method: 'POST',
      })

      expect(global.window.location.href).toEqual('url-to-redirect-to')
    })
  })
})
