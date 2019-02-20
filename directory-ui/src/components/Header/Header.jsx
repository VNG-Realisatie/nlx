import React from 'react'
import { Link } from 'react-router-dom'

import logo from './assets/logo.svg'
import gitlab from './assets/gitlab.svg'

import StyledHeader, {StyledNavigation} from './Header.styles'

const Header = () =>
    <StyledHeader>
        <StyledNavigation>
            <a className="navbar-logo" href="https://www.nlx.io">
                <img src={logo} alt="NLX logo" />
            </a>
            <ul className="navbar-nav">
                <li className="nav-item">
                    <a className="nav-link" href="https://nlx.io/about/">Over NLX</a>
                </li>
                <li className="nav-item">
                    <a className="nav-link" href="https://docs.nlx.io">Docs</a>
                </li>
            </ul>

            <ul className="navbar-nav">
                <li className="nav-item active">
                    <Link className="nav-link" to="/">Directory</Link>
                </li>
            </ul>

            <a className="navbar-gitlab" href="https://gitlab.com/commonground/nlx" target="_blank" aria-label="GitLab" rel="noopener noreferrer">
                <img src={gitlab} alt="Gitlab logo" />
            </a>
        </StyledNavigation>
    </StyledHeader>

export default Header
