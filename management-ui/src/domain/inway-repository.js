// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import env from '../env'
import { throwOnError } from './fetch-utils'

class InwayRepository {
  static async getAll() {
    const response = await fetch(`${env.managementApiBaseUrl}/v1/inways`)

    throwOnError(response)

    const result = await response.json()
    const inways = result.inways ? result.inways : []
    return inways.map((inway) => {
      inway.name = inway.name || ''
      inway.hostname = inway.hostname || ''
      inway.selfAddress = inway.selfAddress || ''
      inway.services = inway.services || []
      inway.version = inway.version || ''
      return inway
    })
  }

  static async getByName(name) {
    const response = await fetch(
      `${env.managementApiBaseUrl}/v1/inways/${name}`,
    )

    throwOnError(response)

    const inway = await response.json()
    inway.name = inway.name || ''
    inway.hostname = inway.hostname || ''
    inway.selfAddress = inway.selfAddress || ''
    inway.services = inway.services || []
    inway.version = inway.version || ''

    return inway
  }
}

export default InwayRepository
