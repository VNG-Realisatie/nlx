// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

class UserRepository {
  static async getAuthenticatedUser() {
    const result = await fetch('/oidc/me')

    if (result.status === 401) {
      throw new Error('no user is authenticated')
    }

    if (!result.ok) {
      throw new Error('unable to handle the request')
    }

    return await result.json()
  }
}

export default UserRepository
