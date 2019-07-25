// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { NavLink } from 'react-router-dom'
import { arrayOf, string, func } from 'prop-types'
import { StyledOrganizationList, StyledSearch, StyledSidebar } from './index.styles'

const Sidebar = ({ organizations, onSearchQueryChanged, ...props }) =>
  <StyledSidebar {...props}>
    <StyledSearch inputProps={({
      placeholder: 'Filter organizations'
    })} onQueryChanged={onSearchQueryChanged} />

    <StyledOrganizationList>
      {
        organizations
          .map((organization, i) =>
            <li key={i}>
              <NavLink to={`/organization/${organization}`}>{organization}</NavLink>
            </li>
          )
      }
    </StyledOrganizationList>
  </StyledSidebar>

Sidebar.propTypes = {
  organizations: arrayOf(string),
  onSearchQueryChanged: func,
}

Sidebar.defaultProps = {
  organizations: [],
  onSearchQueryChanged: () => {}
}

export default Sidebar
