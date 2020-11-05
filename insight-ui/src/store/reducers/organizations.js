// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import * as TYPES from '../types'

const filterOutInvalidOrganizations = (organizations) =>
  organizations.filter(
    (organization) =>
      organization.insight_irma_endpoint && organization.insight_log_endpoint,
  )

// eslint-disable-next-line default-param-last
const organizationsReducer = (state = [], action) => {
  switch (action.type) {
    case TYPES.FETCH_ORGANIZATIONS_SUCCESS:
      return filterOutInvalidOrganizations(action.data)
    default:
      return state
  }
}

export default organizationsReducer
