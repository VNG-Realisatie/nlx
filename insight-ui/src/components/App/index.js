// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { NavLink, Route } from 'react-router-dom'

import GlobalStyles from '../../components/GlobalStyles'
import Navigation from '../../components/Navigation'
import VersionLogger from '../VersionLogger'
import SidebarContainer from '../../containers/SidebarContainer'

import HomePage from '../../components/HomePage'
import OrganizationPageContainer from '../../containers/OrganizationPageContainer'
import StyledApp, { StyledNLXNavbar, StyledContent } from './index.styles'

window._env = window._env || {}

const HOME_PAGE_URL = window._env.NAVBAR_HOME_PAGE_URL || 'https://www.nlx.io'
const ABOUT_PAGE_URL =
  window._env.REACT_APP_NAVBAR_ABOUT_PAGE_URL || 'https://www.nlx.io/about'
const DOCS_PAGE_URL =
  window._env.REACT_APP_NAVBAR_DOCS_PAGE_URL || 'https://docs.nlx.io'
const DIRECTORY_URL =
  window._env.REACT_APP_NAVBAR_DIRECTORY_URL || 'https://directory.nlx.io'
const SUPPORT_PAGE_URL = `${
  window._env.REACT_APP_NAVBAR_DOCS_PAGE_URL || 'https://docs.nlx.io'
}/support`

const App = () => (
  <StyledApp>
    <GlobalStyles />

    <StyledNLXNavbar
      homePageURL={HOME_PAGE_URL}
      aboutPageURL={ABOUT_PAGE_URL}
      docsPageURL={DOCS_PAGE_URL}
    >
      <Navigation.Item>
        <a href={SUPPORT_PAGE_URL}>Support</a>
      </Navigation.Item>
      <Navigation.Item>
        <a href={DIRECTORY_URL}>Directory</a>
      </Navigation.Item>
      <Navigation.Item>
        <NavLink to="/">Insight</NavLink>
      </Navigation.Item>
    </StyledNLXNavbar>

    <SidebarContainer />

    <StyledContent>
      <Route path="/" exact component={HomePage} />
      <Route
        path="/organization/:organizationName/"
        component={OrganizationPageContainer}
      />
    </StyledContent>
    <VersionLogger />
  </StyledApp>
)

export default App
