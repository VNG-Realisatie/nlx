// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import * as TYPES from '../types'

const initialState = {}

// eslint-disable-next-line default-param-last
const loginRequestInfo = (state = initialState, action) => {
  switch (action.type) {
    case TYPES.FETCH_IRMA_LOGIN_INFORMATION_SUCCESS:
      return action.data
    case TYPES.RESET_LOGIN_INFORMATION:
      return initialState
    default:
      return state
  }
}

export default loginRequestInfo
