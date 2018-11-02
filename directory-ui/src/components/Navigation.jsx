import React from 'react'
import logo from '../assets/images/logo.svg'
// import menu from '../assets/icons/menu.svg'
import gitlab from '../assets/icons/gitlab.svg'
// import {Link} from 'react-router-dom'

export default class Navigation extends React.Component {
    render() {
        return (
            <header className="navbar navbar-expand navbar-sticky flex-column bg-white navbar-expand-md">
                <div className="container">
                    <nav
                        className="collapse navbar-collapse "
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
                                <a
                                    className="nav-link"
                                    href="https://directory.nlx.io"
                                >
                                    Directory
                                </a>
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
