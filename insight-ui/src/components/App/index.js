import React from 'react'
import { NavLink, Route } from 'react-router-dom'

import StyledApp, {StyledNavbar, StyledContent} from './index.styles'
import GlobalStyles from '../../components/GlobalStyles'
import { StyledNavbarNavLinkListItem } from '@commonground/design-system/dist/components/Navbar/index.styles'
import SidebarContainer from '../../containers/SidebarContainer'

import HomePage from '../../components/HomePage'
import OrganizationPageContainer from '../../containers/OrganizationPageContainer'

const App = () =>
  <StyledApp>
    <GlobalStyles/>

    <StyledNavbar homePageURL={process.env.REACT_APP_NAVBAR_HOME_PAGE_URL}
                  aboutPageURL={process.env.REACT_APP_NAVBAR_ABOUT_PAGE_URL}
                  docsPageURL={process.env.REACT_APP_NAVBAR_DOCS_PAGE_URL}>
      <StyledNavbarNavLinkListItem>
        <a href={process.env.REACT_APP_NAVBAR_DIRECTORY_URL}>Directory</a>
      </StyledNavbarNavLinkListItem>
      <StyledNavbarNavLinkListItem>
        <NavLink to="/">Insight</NavLink>
      </StyledNavbarNavLinkListItem>
    </StyledNavbar>

    <SidebarContainer />

    <StyledContent>
      <Route path="/" exact component={HomePage} />
      <Route path="/organization/:organizationName/" component={OrganizationPageContainer} />
    </StyledContent>
  </StyledApp>

export default App
