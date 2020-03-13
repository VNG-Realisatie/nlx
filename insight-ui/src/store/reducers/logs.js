// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import * as TYPES from '../types'

export const modelFromAPIResponse = (logFromAPIResponse) => ({
  id: logFromAPIResponse['logrecord-id'],
  subjects: logFromAPIResponse.data['doelbinding-data-elements'].split(','),
  requestedBy: logFromAPIResponse.source_organization,
  requestedAt: logFromAPIResponse.destination_organization,
  application: logFromAPIResponse.data['doelbinding-application-id'],
  reason: logFromAPIResponse.data['doelbinding-process-id'],
  date: new Date(logFromAPIResponse.created),
})

const defaultState = {
  records: [],
  rowsPerPage: 0,
  rowCount: 0,
}

export default (state = defaultState, action) => {
  switch (action.type) {
    case TYPES.FETCH_ORGANIZATION_LOGS_SUCCESS:
      const { rowCount, rowsPerPage, records } = action.data
      return {
        rowCount,
        rowsPerPage,
        records: records.map(modelFromAPIResponse),
      }
    case TYPES.RESET_LOGIN_INFORMATION:
      return defaultState
    default:
      return state
  }
}
