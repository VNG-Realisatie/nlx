// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import UserRepositoryBasicAuth, {
  LOCALSTORAGE_KEY_BASIC_AUTH_CREDENTIALS,
} from './user-repository-basic-auth'

describe('the UserRepositoryBasicAuth', () => {
  describe('login', () => {
    it('with valid credentials should return true', () => {
      jest.spyOn(global, 'fetch').mockImplementation(() =>
        Promise.resolve({
          ok: true,
          status: 204,
        }),
      )

      expect(
        UserRepositoryBasicAuth.login('email@example.com', 'password'),
      ).resolves.toEqual(true)

      expect(global.fetch).toHaveBeenCalledWith('/basic-auth/login', {
        body: 'email=email%40example.com&password=password',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8',
        },
        method: 'POST',
      })

      global.fetch.mockRestore()
    })
  })

  it('with invalid credentials should return false', () => {
    jest.spyOn(global, 'fetch').mockImplementation(() =>
      Promise.resolve({
        ok: true,
        status: 401,
      }),
    )

    expect(
      UserRepositoryBasicAuth.login('email@example.com', 'password'),
    ).resolves.toEqual(false)

    global.fetch.mockRestore()
  })

  describe('getting the authenticated user', () => {
    it('should return the user', async () => {
      global.Storage.prototype.getItem = jest.fn(() => 'my-credentials')

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

      await expect(
        UserRepositoryBasicAuth.getAuthenticatedUser(),
      ).resolves.toEqual({
        id: '42',
        fullName: 'full name',
        email: 'email',
        pictureUrl: 'picture url',
      })

      expect(global.fetch).toHaveBeenCalledWith('/basic-auth/me', {
        headers: { Authorization: 'Basic my-credentials' },
      })

      global.Storage.prototype.getItem.mockReset()
      global.fetch.mockRestore()
    })

    it('should not call api when no user is authenticated', async () => {
      global.Storage.prototype.getItem = jest.fn(() => '')

      jest.spyOn(global, 'fetch').mockImplementation(() =>
        Promise.resolve({
          ok: false,
          status: 401,
        }),
      )

      expect(global.fetch).toHaveBeenCalledTimes(0)

      global.Storage.prototype.getItem.mockReset()
      global.fetch.mockRestore()
    })
  })

  describe('storing credentials', () => {
    it('should store user credentials within the local storage', async () => {
      global.Storage.prototype.setItem = jest.fn()

      UserRepositoryBasicAuth.storeCredentials('my-credentials')
      expect(global.Storage.prototype.setItem).toHaveBeenCalledTimes(1)

      global.Storage.prototype.setItem.mockReset()
    })
  })

  describe('retrieving credentials', () => {
    it('should return credentials', async () => {
      global.Storage.prototype.getItem = jest.fn(() => 'my-credentials')

      expect(UserRepositoryBasicAuth.getCredentials()).toEqual('my-credentials')

      global.Storage.prototype.getItem.mockReset()
    })
  })
})
describe('logout', () => {
  it('should logout', async () => {
    delete window.location
    window.location = {}

    global.Storage.prototype.removeItem = jest.fn()

    UserRepositoryBasicAuth.logout()

    expect(global.Storage.prototype.removeItem).toHaveBeenLastCalledWith(
      LOCALSTORAGE_KEY_BASIC_AUTH_CREDENTIALS,
    )

    expect(global.window.location.href).toEqual('/')

    global.Storage.prototype.removeItem.mockReset()
  })
})
