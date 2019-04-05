// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { Component } from 'react'

import StyledApp, {StyledSidebar, StyledContent, StyledOrganizationList, StyledSidebarHeader, StyledCard} from './App.styles'
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
      <StyledCard>
        <p>
          View logs by selecting an organization on the left.
          You can only view logs by disclosing the required IRMA attributes.
        </p>

        <p className="text-muted">
          Read more about IRMA and what it does <a href="https://privacybydesign.foundation/irma/" target="_blank" rel="noopener noreferrer">here</a>.
        </p>
      </StyledCard>
    </StyledContent>
  </StyledApp>

export default App
