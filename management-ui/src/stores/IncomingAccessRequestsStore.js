// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, observable, flow } from 'mobx'
import IncomingAccessRequestModel, {
  STATES,
} from './models/IncomingAccessRequestModel'

class IncomingAccessRequestsStore {
  incomingAccessRequests = observable.map()

  constructor({ managementApiClient }) {
    makeAutoObservable(this)

    this._managementApiClient = managementApiClient
  }

  updateFromServer = (incomingAccessRequestData) => {
    if (!incomingAccessRequestData) return null

    const cachedIncomingAccessRequest = this.incomingAccessRequests.get(
      incomingAccessRequestData.id,
    )

    if (cachedIncomingAccessRequest) {
      cachedIncomingAccessRequest.update(incomingAccessRequestData)
      return cachedIncomingAccessRequest
    }

    const incomingAccessRequest = new IncomingAccessRequestModel({
      incomingAccessRequestStore: this,
      accessRequestData: incomingAccessRequestData,
    })

    this.incomingAccessRequests.set(
      incomingAccessRequest.id,
      incomingAccessRequest,
    )

    return incomingAccessRequest
  }

  fetchForService = flow(function* fetchForService({ name }) {
    const result = yield this.returnForService({ name })

    // delete access requests which do not exist anymore
    const newIds = result.map((ar) => ar.id)
    this.getForService({ name }).forEach((ar) => {
      if (newIds.includes(ar.id)) {
        return
      }

      this.incomingAccessRequests.delete(ar.id)
    })

    // create and update other existing requests
    result.map((accessRequest) => this.updateFromServer(accessRequest))
  }).bind(this)

  returnForService = async ({ name }) => {
    const result = await this._managementApiClient.managementListIncomingAccessRequest(
      {
        serviceName: name,
      },
    )
    return result.accessRequests
  }

  getForService = ({ name }) => {
    const arrayOfModels = [...this.incomingAccessRequests.values()]

    return arrayOfModels.filter(
      (incomingAccessRequestModel) =>
        incomingAccessRequestModel.serviceName === name,
    )
  }

  approveAccessRequest = async ({ serviceName, id }) => {
    await this._managementApiClient.managementApproveIncomingAccessRequest({
      serviceName,
      accessRequestID: id,
    })
    this.fetchForService({ name: serviceName })
  }

  rejectAccessRequest = async ({ serviceName, id }) => {
    await this._managementApiClient.managementRejectIncomingAccessRequest({
      serviceName,
      accessRequestID: id,
    })
    this.fetchForService({ name: serviceName })
  }

  haveChangedForService = async (service) => {
    let latestAccessRequests = await this.returnForService(service)

    // we are only interested in access request which are not 'resolved'
    latestAccessRequests = latestAccessRequests.filter(
      (ar) => ar.state === STATES.CREATED || ar.state === STATES.RECEIVED,
    )

    if (latestAccessRequests.length !== service.incomingAccessRequests.length) {
      return true
    }

    // we will compare a list of sorted IDs to determine if the
    // latest access requests have changed
    const accessRequestIds = service.incomingAccessRequests.map((ar) => ar.id)
    accessRequestIds.sort()

    const latestAccessRequestIds = latestAccessRequests.map((ar) => ar.id)
    latestAccessRequestIds.sort()

    return (
      JSON.stringify(accessRequestIds) !==
      JSON.stringify(latestAccessRequestIds)
    )
  }
}

export default IncomingAccessRequestsStore
