import React from 'react'
import { StyledOrganizationList, StyledSearch, StyledSidebar } from './index.styles'

const Sidebar = () =>
  <StyledSidebar>
    <StyledSearch placeholder="Filter organisations" />

    <StyledOrganizationList>
      <li><a href="#">BRP</a></li>
      <li className="active"><a href="#">Haarlem</a></li>
      <li><a href="#">Kadaster</a></li>
    </StyledOrganizationList>
  </StyledSidebar>

export default Sidebar
