// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import Service from './Service'

interface Order {
  description: string
  delegatee: string
  reference: string
  services: Service[]
  validFrom: Date
  validUntil: Date
  revokedAt?: Date
}

export default Order
