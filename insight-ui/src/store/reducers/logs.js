import * as TYPES from '../types'

export default (state = [], action) => {
  switch (action.type) {
    case TYPES.FETCH_ORGANIZATION_LOGS_SUCCESS:
      return action.data.records
    default:
      return state
  }
}
