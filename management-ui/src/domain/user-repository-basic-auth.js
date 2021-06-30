// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { throwOnError } from './fetch-utils'

export const LOCALSTORAGE_KEY_BASIC_AUTH_CREDENTIALS = 'basic_auth_credentials'

const getCredentials = () => {
  return localStorage.getItem(LOCALSTORAGE_KEY_BASIC_AUTH_CREDENTIALS)
}

class UserRepositoryBasicAuth {
  static async login(email, password) {
    const emailParam = `email=${encodeURIComponent(email)}`
    const passwordParam = `password=${encodeURIComponent(password)}`

    const response = await fetch('/basic-auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8',
      },
      body: `${emailParam}&${passwordParam}`,
    })

    throwOnError(response)

    return response.status === 204
  }

  static async getAuthenticatedUser() {
    const credentials = getCredentials()
    if (!credentials) {
      return null
    }

    const response = await fetch('/basic-auth/me', {
      headers: {
        Authorization: `Basic ${credentials}`,
      },
    })

    throwOnError(response)

    return await response.json()
  }

  static storeCredentials(email, password) {
    const credentials = window.btoa(`${email}:${password}`)
    localStorage.setItem(LOCALSTORAGE_KEY_BASIC_AUTH_CREDENTIALS, credentials)
  }

  static getCredentials() {
    return getCredentials()
  }

  static logout() {
    localStorage.removeItem(LOCALSTORAGE_KEY_BASIC_AUTH_CREDENTIALS)
    window.location.href = '/'
  }
}

export default UserRepositoryBasicAuth
