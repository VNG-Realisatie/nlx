// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { NavLink, Route } from 'react-router-dom'

import StyledApp, {StyledNLXNavbar, StyledContent} from './index.styles'
import GlobalStyles from '../../components/GlobalStyles'
import Navigation from '../../components/Navigation'
import { VersionLogger } from '@commonground/design-system'
import SidebarContainer from '../../containers/SidebarContainer'

import HomePage from '../../components/HomePage'
import OrganizationPageContainer from '../../containers/OrganizationPageContainer'

window._env = window._env || {}

const App = () =>
  <StyledApp>
    <GlobalStyles/>

    <StyledNLXNavbar homePageURL={window._env.NAVBAR_HOME_PAGE_URL || 'https://www.nlx.io'}
                     aboutPageURL={window._env.REACT_APP_NAVBAR_ABOUT_PAGE_URL || 'https://www.nlx.io/about'}
                     docsPageURL={window._env.REACT_APP_NAVBAR_DOCS_PAGE_URL || 'https://docs.nlx.io'}>
      <Navigation.Item>
        <a href={`${window._env.REACT_APP_NAVBAR_DOCS_PAGE_URL || 'https://docs.nlx.io'}/support`}>Support</a>
      </Navigation.Item>
      <Navigation.Item>
        <a href={window._env.REACT_APP_NAVBAR_DIRECTORY_URL || 'https://directory.nlx.io'}>Directory</a>
      </Navigation.Item>
      <Navigation.Item>
        <NavLink to="/">Insight</NavLink>
      </Navigation.Item>
    </StyledNLXNavbar>

    <SidebarContainer />

    <StyledContent>
      <Route path="/" exact component={HomePage} />
      <Route path="/organization/:organizationName/" component={OrganizationPageContainer} />
    </StyledContent>
    <VersionLogger />
  </StyledApp>

export default App
