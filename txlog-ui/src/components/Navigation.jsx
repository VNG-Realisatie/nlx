import React from 'react'
import logo from '../assets/images/logo.svg'
import gitlab from '../assets/icons/gitlab.svg'

export default class Navigation extends React.Component {
    render() {
        return (
            <header
                className="navbar fixed-top bg-white"
                style={{ minHeight: '64px' }}
            >
                <div className="container">
                    <a className="navbar-logo" href="/">
                        <img src={logo} alt="logo" />
                    </a>

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
