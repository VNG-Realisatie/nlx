// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
class AccessRequestRepository {
  static async requestAccess(organizationName, serviceName) {
    const result = await fetch(
      `/api/v1/access-requests/outgoing/organizations/${organizationName}/services/${serviceName}`,
      {
        method: 'POST',
      },
    )

    if (result.status === 400) {
      throw new Error('invalid user input')
    }

    if (result.status === 403) {
      throw new Error('forbidden')
    }

    if (!result.ok) {
      throw new Error('unable to handle the request')
    }

    return await result.json()
  }
}

export default AccessRequestRepository
