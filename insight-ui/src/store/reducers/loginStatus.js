import * as TYPES from "../types";

export default (state = null, action) => {
  switch (action.type) {
    case TYPES.IRMA_LOGIN_REQUEST_SUCCESS:
    case TYPES.IRMA_LOGIN_REQUEST_FAILED:
      return action.data
    default:
      return state
  }
}
