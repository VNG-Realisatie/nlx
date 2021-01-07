// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { throwOnError } from './fetch-utils'

class UserRepository {
  static async getAuthenticatedUser() {
    const response = await fetch('/oidc/me')

    throwOnError(response)

    return await response.json()
  }
}

export default UserRepository
