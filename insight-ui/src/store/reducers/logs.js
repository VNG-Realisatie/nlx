import * as TYPES from '../types'

export const modelFromAPIResponse = logFromAPIResponse => ({
    id: logFromAPIResponse['logrecord-id'],
    subjects: logFromAPIResponse.data['doelbinding-data-elements'].split(','),
    requestedBy: logFromAPIResponse['source_organization'],
    requestedAt: logFromAPIResponse['destination_organization'],
    application: logFromAPIResponse.data['doelbinding-application-id'],
    reason: logFromAPIResponse.data['doelbinding-process-id'],
    date: new Date(logFromAPIResponse['created'])
})

const defaultState = {
  records: [],
  pageCount: 0
}

export default (state = defaultState, action) => {
  switch (action.type) {
    case TYPES.FETCH_ORGANIZATION_LOGS_SUCCESS:
      return {
        pageCount: Math.ceil(action.data.rowCount / action.data.rowsPerPage),
        records: action.data.records.map(record => modelFromAPIResponse(record))
      }
    default:
      return state
  }
}
