// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import Cookies from 'js-cookie'
import { throwOnError } from './fetch-utils'

class UserRepositoryOIDC {
  static async getAuthenticatedUser() {
    const response = await fetch('/oidc/me')

    throwOnError(response)

    return await response.json()
  }

  static async logout() {
    const csrftoken = Cookies.get('csrftoken')

    const response = await fetch('/oidc/logout', {
      method: 'POST',
      body: `csrfmiddlewaretoken=${csrftoken}`,
    })

    window.location.href = response.url
  }
}

export default UserRepositoryOIDC
