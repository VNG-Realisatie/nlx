import React from 'react'
import { NavLink } from 'react-router-dom'
import { arrayOf, string } from 'prop-types'
import { StyledOrganizationList, StyledSearch, StyledSidebar } from './index.styles'

const Sidebar = ({ organizations, ...props }) =>
  <StyledSidebar {...props}>
    <StyledSearch placeholder="Filter organisations" />

    <StyledOrganizationList>
      {
        organizations
          .map((organization, i) =>
            <li key={i}>
              <NavLink to={`/organization/${organization}/login`}>{organization}</NavLink>
            </li>
          )
      }
    </StyledOrganizationList>
  </StyledSidebar>

Sidebar.propTypes = {
  organizations: arrayOf(string)
}

Sidebar.defaultProps = {
  organizations: []
}

export default Sidebar
