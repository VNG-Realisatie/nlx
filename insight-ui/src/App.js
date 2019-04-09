// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { Component } from 'react'

import StyledApp, {StyledSidebar, StyledContent, StyledOrganizationList, StyledCard} from './App.styles'
import GlobalStyles from './components/GlobalStyles'
import { StyledNavbar, StyledSearch } from "./App.styles";
import { StyledNavbarNavLinkListItem } from '@commonground/design-system/dist/components/Navbar/index.styles'

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

      <StyledSidebar>
        <StyledSearch placeholder="Filter organisations" />
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
