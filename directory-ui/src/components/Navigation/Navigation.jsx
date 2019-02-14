import React from 'react'
import { Link } from 'react-router-dom'

import logo from './assets/logo.svg'
import gitlab from './assets/gitlab.svg'

import StyledNavigation from './Navigation.styles'

const Navigation = () =>
    <StyledNavigation>
        <div className="container">
            <nav>
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

                <a className="nav-link" href="https://gitlab.com/commonground/nlx" target="_blank" aria-label="GitLab" rel="noopener noreferrer">
                    <img src={gitlab} alt="Gitlab logo" />
                </a>
            </nav>
        </div>
    </StyledNavigation>

export default Navigation
