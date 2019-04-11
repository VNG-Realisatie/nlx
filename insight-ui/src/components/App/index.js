import React from 'react'
import { NavLink, Route } from 'react-router-dom'

import StyledApp, {StyledNavbar, StyledContent} from './index.styles'
import GlobalStyles from '../../components/GlobalStyles'
import { StyledNavbarNavLinkListItem } from '@commonground/design-system/dist/components/Navbar/index.styles'
import SidebarContainer from '../../containers/SidebarContainer'

import HomePage from '../../components/HomePage'
import LoginPageContainer from '../../containers/LoginPageContainer'

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

    <SidebarContainer />

    <StyledContent>
      <Route path="/" exact component={HomePage} />
      <Route path="/organization/:organizationName/login" component={LoginPageContainer} />
    </StyledContent>
  </StyledApp>

export default App