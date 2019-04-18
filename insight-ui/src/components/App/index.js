import React from 'react'
import { NavLink, Route } from 'react-router-dom'

import StyledApp, {StyledNLXNavbar, StyledContent} from './index.styles'
import GlobalStyles from '../../components/GlobalStyles'
import { Navigation } from '@commonground/design-system'
import SidebarContainer from '../../containers/SidebarContainer'

import HomePage from '../../components/HomePage'
import OrganizationPageContainer from '../../containers/OrganizationPageContainer'

const App = () =>
  <StyledApp>
    <GlobalStyles/>

    <StyledNLXNavbar homePageURL={process.env.REACT_APP_NAVBAR_HOME_PAGE_URL || 'https://www.nlx.io'}
                     aboutPageURL={process.env.REACT_APP_NAVBAR_ABOUT_PAGE_URL || 'https://www.nlx.io/about'}
                     docsPageURL={process.env.REACT_APP_NAVBAR_DOCS_PAGE_URL || 'https://docs.nlx.io'}>
      <Navigation.Item>
        <a href={process.env.REACT_APP_NAVBAR_DIRECTORY_URL}>Directory</a>
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
  </StyledApp>

export default App
