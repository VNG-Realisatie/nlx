// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { Component } from 'react'

import StyledApp, {StyledSidebar, StyledContent, StyledOrganizationList, StyledCard} from './App.styles'
import GlobalStyles from './components/GlobalStyles'
import { StyledNavbarNavLinkListItem } from '@commonground/design-system/dist/components/Navbar/index.styles'
import Sidebar from './components/Sidebar'
import Home from "./pages/Home";

const App = () =>
  <StyledApp>
      <GlobalStyles/>

      <StyledNavbar>
        <StyledNavbarNavLinkListItem>
          <a href="#">Directory</a>
        </StyledNavbarNavLinkListItem>
        <StyledNavbarNavLinkListItem>
          <a href="#" className="active">Insight</a>
        </StyledNavbarNavLinkListItem>
      </StyledNavbar>

      <Sidebar />

      <StyledContent>
        <Route path="/" exact component={Home} />
      </StyledContent>
  </StyledApp>

export default App
