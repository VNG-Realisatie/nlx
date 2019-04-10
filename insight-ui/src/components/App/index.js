import React from 'react'
import { Route } from 'react-router-dom'

import StyledApp, {StyledNavbar, StyledContent} from './index.styles'
import GlobalStyles from '../../components/GlobalStyles'
import { StyledNavbarNavLinkListItem } from '@commonground/design-system/dist/components/Navbar/index.styles'
import Sidebar from '../../containers/Sidebar'
import Home from '../../pages/Home'

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