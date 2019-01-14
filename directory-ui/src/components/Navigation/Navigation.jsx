import React, { Component } from 'react'
import { Link } from 'react-router-dom'

import logo from './assets/logo.svg'
import gitlab from './assets/gitlab.svg'

class Navigation extends Component {
    render() {
        return (
            <header className="navbar navbar-expand navbar-sticky flex-column bg-white navbar-expand-md">
                <div className="container">
                    <nav
                        className="collapse navbar-collapse"
                        aria-label="Page navigation"
                    >
                        <a className="navbar-logo" href="https://www.nlx.io">
                            <img src={logo} alt="NLX logo" />
                        </a>
                        <ul className="navbar-nav flex-row">
                            <li className="nav-item">
                                <a
                                    className="nav-link"
                                    href="https://nlx.io/about/"
                                >
                                    Over NLX
                                </a>
                            </li>
                            <li className="nav-item">
                                <a
                                    className="nav-link"
                                    href="https://docs.nlx.io"
                                >
                                    Docs
                                </a>
                            </li>
                            <li className="nav-item active">
                                <Link className="nav-link" to="/">
                                    Directory
                                </Link>
                            </li>
                        </ul>

                        <a
                            className="nav-link"
                            href="https://gitlab.com/commonground/nlx"
                            target="_blank"
                            aria-label="GitLab"
                            rel="noopener noreferrer"
                        >
                            <img src={gitlab} alt="Gitlab logo" />
                        </a>
                    </nav>
                </div>
            </header>
        )
    }
}

export default Navigation
