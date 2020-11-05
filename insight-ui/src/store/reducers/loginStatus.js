// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import * as TYPES from '../types'

const initialState = null

// eslint-disable-next-line default-param-last
const loginStatus = (state = initialState, action) => {
  switch (action.type) {
    case TYPES.IRMA_LOGIN_REQUEST_SUCCESS:
    case TYPES.IRMA_LOGIN_REQUEST_FAILED:
      return action.data
    case TYPES.RESET_LOGIN_INFORMATION:
      return initialState
    default:
      return state
  }
}

export default loginStatus
