import * as TYPES from "../types";

const filterOutInvalidOrganizations = organizations =>
  organizations
    .filter(organization =>
      organization.insight_irma_endpoint && organization.insight_log_endpoint
    )

export default (state = [], action) => {
  switch (action.type) {
    case TYPES.FETCH_ORGANIZATIONS_SUCCESS:
      return filterOutInvalidOrganizations(action.data)
    default:
      return state
  }
}
