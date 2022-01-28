// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { createContext } from 'react'
import { configure } from 'mobx'
import { node, object } from 'prop-types'
import ApplicationStore from './ApplicationStore'
import DirectoryServicesStore from './DirectoryServicesStore'
import OutgoingAccessRequestStore from './OutgoingAccessRequestStore'
import AccessGrantStore from './AccessGrantStore'
import AccessProofStore from './AccessProofStore'
import ServiceStore from './ServiceStore'
import IncomingAccessRequestsStore from './IncomingAccessRequestsStore'
import InwayStore from './InwayStore'
import FinanceStore from './FinanceStore'
import AuditLogStore from './AuditLogStore'
import TransactionLogStore from './TransactionLogStore'
import OrderStore from './OrderStore'
import OutwayStore from './OutwayStore'

if (process.env.NODE_ENV !== 'test') {
  // `setupTests` has 'never' set. But some tests include this file,
  // so make sure we don't override the test setup.
  configure({ enforceActions: 'observed' })
}

export const storesContext = createContext(null)

export class RootStore {
  constructor({
    directoryApiClient,
    managementApiClient,
    txLogApiClient,
  } = {}) {
    this.applicationStore = new ApplicationStore({
      rootStore: this,
      managementApiClient,
      directoryApiClient,
    })
    this.directoryServicesStore = new DirectoryServicesStore({
      rootStore: this,
      directoryApiClient,
    })
    this.outgoingAccessRequestStore = new OutgoingAccessRequestStore({
      rootStore: this,
      managementApiClient,
    })
    this.accessGrantStore = new AccessGrantStore({
      managementApiClient,
    })
    this.accessProofStore = new AccessProofStore()
    this.servicesStore = new ServiceStore({
      rootStore: this,
      managementApiClient,
    })
    this.incomingAccessRequestsStore = new IncomingAccessRequestsStore({
      managementApiClient,
    })
    this.inwayStore = new InwayStore({
      rootStore: this,
      managementApiClient,
    })
    this.outwayStore = new OutwayStore({
      rootStore: this,
      managementApiClient,
    })
    this.auditLogStore = new AuditLogStore({
      managementApiClient,
    })
    this.transactionLogStore = new TransactionLogStore({
      txLogApiClient,
      managementApiClient,
    })
    this.financeStore = new FinanceStore({
      managementApiClient,
    })
    this.orderStore = new OrderStore({
      managementApiClient,
    })
  }
}

export const StoreProvider = ({ children, rootStore }) => {
  return (
    <storesContext.Provider value={rootStore}>
      {children}
    </storesContext.Provider>
  )
}

StoreProvider.propTypes = {
  children: node,
  rootStore: object.isRequired,
}
