// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { func } from 'prop-types'
import { StyledFilters, StyledInput, StyledSearchIcon, StyledSwitch } from './Filters.styles'

const Filters = ({ onQueryChanged, onStatusFilterChanged, queryValue, ...props }) =>
  <StyledFilters {...props}>
    <StyledSearchIcon />
    <StyledInput dataTest="query" value={queryValue} placeholder="Search for an organization or service…" onChange={event => onQueryChanged(event.target.value)} />
    <StyledSwitch label="Include offline" onChange={event => onStatusFilterChanged(event.target.checked)} />
  </StyledFilters>

Filters.propTypes = {
  onQueryChanged: func,
  onStatusFilterChanged: func
}

Filters.defaultProps = {
  onQueryChanged: () => {},
  onStatusFilterChanged: () => {}
}

export default Filters
