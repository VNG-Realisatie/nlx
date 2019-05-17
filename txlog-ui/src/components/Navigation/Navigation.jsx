// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import logo from './assets/logo.svg'
import gitlab from './assets/gitlab.svg'

export default class Navigation extends Component {
    render() {
        return (
            <header
                className="navbar fixed-top bg-white"
                style={{ minHeight: '64px' }}
            >
                <div className="container">
                    <Link className="navbar-logo" to="/">
                        <img src={logo} alt="logo" />
                    </Link>

                    <ul className="navbar-nav flex-row ml-md-auto">
                        <li className="nav-item">
                            <a
                                className="nav-link p-3"
                                href="https://gitlab.com/commonground/nlx"
                                target="_blank"
                                rel="noopener noreferrer"
                                aria-label="gitlab"
                            >
                                <img src={gitlab} alt="logo" />
                            </a>
                        </li>
                    </ul>
                </div>
            </header>
        )
    }
}
