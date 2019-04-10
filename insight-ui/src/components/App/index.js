import React from 'react'
import { NavLink, Route } from 'react-router-dom'

import StyledApp, {StyledNavbar, StyledContent} from './index.styles'
import GlobalStyles from '../../components/GlobalStyles'
import { StyledNavbarNavLinkListItem } from '@commonground/design-system/dist/components/Navbar/index.styles'
import Sidebar from '../../containers/SidebarContainer'
import Home from '../../pages/Home'

const App = () =>
  <StyledApp>
    <GlobalStyles/>

    <StyledNavbar>
      <StyledNavbarNavLinkListItem>
        <a href="#">Directory</a>
      </StyledNavbarNavLinkListItem>
      <StyledNavbarNavLinkListItem>
        <NavLink to="/">Insight</NavLink>
      </StyledNavbarNavLinkListItem>
    </StyledNavbar>

    <Sidebar />

    <StyledContent>
      <Route path="/" exact component={Home} />
    </StyledContent>
  </StyledApp>

export default App