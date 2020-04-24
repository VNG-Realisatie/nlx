// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import UserRepository from './user-repository'
import { preventCaching } from './service-repository'

describe('the UserRepository', () => {
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
        return expect(UserRepository.getAuthenticatedUser()).resolves.toEqual({
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
        return expect(UserRepository.getAuthenticatedUser()).rejects.toEqual(
          new Error('no user is authenticated'),
        )
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
        const user = UserRepository.getAuthenticatedUser()

        await expect(user).rejects.toEqual(
          new Error('unable to handle the request'),
        )

        expect(global.fetch).toHaveBeenCalledWith(
          '/oidc/me',
          expect.objectContaining({
            headers: expect.objectContaining(preventCaching),
          }),
        )
      })
    })
  })
})
