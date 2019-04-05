// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { Component } from 'react'

import StyledApp, {StyledSidebar, StyledContent, StyledOrganizationList, StyledSidebarHeader} from './App.styles'
import GlobalStyles from './components/GlobalStyles'
import Logo from './components/Logo/index';

const App = () =>
  <StyledApp>
    <GlobalStyles/>
    <StyledSidebar>
      <StyledSidebarHeader href="#">
        <Logo/> Insight
      </StyledSidebarHeader>
      <StyledOrganizationList>
        <li><a href="#">BRP</a></li>
        <li className="active"><a href="#">Haarlem</a></li>
        <li><a href="#">Kadaster</a></li>
      </StyledOrganizationList>
    </StyledSidebar>
    <StyledContent>
      Content
    </StyledContent>
  </StyledApp>

export default App
