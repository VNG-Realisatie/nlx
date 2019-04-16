import * as TYPES from '../types'

export default (state = {}, action) => {
  switch (action.type) {
    case TYPES.FETCH_IRMA_LOGIN_INFORMATION_SUCCESS:
      return action.data
    default:
      return state
  }
}
