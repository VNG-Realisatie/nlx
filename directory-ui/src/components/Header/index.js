// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { Link } from 'react-router-dom'
import Container from '../Container/Container'
import IconButton from '../IconButton'
import logo from './assets/logo.svg'
import gitlab from './assets/gitlab.svg'
import StyledHeader, { StyledNavigation } from './index.styles'

const Header = () => (
  <StyledHeader>
    <Container>
      <StyledNavigation>
        <a className="navbar-logo" href="https://nlx.io">
          <img src={logo} alt="NLX logo" />
        </a>

        <ul className="navbar-nav">
          <li className="nav-item active">
            <Link className="nav-link" to="/">
              Directory
            </Link>
          </li>
          <li className="nav-item">
            <a className="nav-link" href="https://docs.nlx.io">
              Docs
            </a>
          </li>
        </ul>

        <IconButton
          as="a"
          className="navbar-gitlab"
          href="https://gitlab.com/commonground/nlx"
          target="_blank"
          aria-label="GitLab"
          rel="noopener noreferrer"
        >
          <img src={gitlab} alt="Gitlab logo" />
        </IconButton>
      </StyledNavigation>
    </Container>
  </StyledHeader>
)

export default Header
