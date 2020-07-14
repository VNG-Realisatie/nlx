// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { throwOnError } from './fetch-utils'

class InwayRepository {
  static async getAll() {
    const response = await fetch(`/api/v1/inways`)

    throwOnError(response)

    const result = await response.json()
    return result.inways ? result.inways : []
  }

  static async getByName(name) {
    const response = await fetch(`/api/v1/inways/${name}`)

    throwOnError(response)

    const inway = await response.json()
    inway.services = inway.services || []

    return inway
  }
}

export default InwayRepository
