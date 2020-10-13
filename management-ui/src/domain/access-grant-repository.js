// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { throwOnError } from './fetch-utils'

class AccessGrantRepository {
  static async getByServiceName(serviceName) {
    const response = await fetch(
      `/api/v1/access-grants/services/${serviceName}`,
    )

    throwOnError(response)

    const result = await response.json()
    return result.accessGrants || []
  }

  static async revokeAccessGrant({
    organizationName,
    serviceName,
    accessGrantId,
  }) {
    const response = await fetch(
      `/api/v1/access-grants/service/${serviceName}/organizations/${organizationName}/${accessGrantId}/revoke`,
      {
        method: 'POST',
        body: JSON.stringify({
          organizationName: organizationName,
          serviceName: serviceName,
          accessGrantID: accessGrantId,
        }),
      },
    )

    throwOnError(response, {
      409: 'Access has already been revoked.',
    })

    return null
  }
}

export default AccessGrantRepository
