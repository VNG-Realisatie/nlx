// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import * as TYPES from '../types'

const defaultState = {
  loaded: false,
  error: false,
  value: null,
  response: null,
}

// eslint-disable-next-line default-param-last
const proofReducer = (state = defaultState, action) => {
  switch (action.type) {
    case TYPES.FETCH_PROOF_SUCCESS:
      return {
        loaded: true,
        error: false,
        value: action.data,
        message: null,
      }
    case TYPES.FETCH_PROOF_FAILED:
      return {
        loaded: true,
        error: true,
        value: null,
        message: action.data.response,
      }
    case TYPES.RESET_LOGIN_INFORMATION:
      return defaultState
    default:
      return state
  }
}

export default proofReducer
