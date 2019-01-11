import React, { Component } from 'react'
import { Link } from 'react-router-dom'

import logo from './assets/logo.svg'
import menu from './assets/menu.svg'
import gitlab from './assets/gitlab.svg'

export default class Navigation extends Component {
    render() {
        return (
            <header className="navbar fixed-top bg-white navbar-expand-md">
                <div className="container">
                    <Link className="navbar-logo d-md-none" to="/">
                        <img src={logo} alt="logo" />
                    </Link>

                    <button
                        className="navbar-toggler navbar-logo d-md-none"
                        type="button"
                        data-toggle="collapse"
                        data-target="#navbarSupportedContent"
                        aria-controls="navbarSupportedContent"
                        aria-expanded="false"
                        aria-label="Toggle navigation"
                    >
                        <img src={menu} alt="menu" />
                    </button>

                    <nav
                        className="collapse navbar-collapse order-last order-md-first"
                        id="navbarSupportedContent"
                        aria-label="Page navigation"
                    >
                        <ul className="navbar-nav flex-row-md">
                            <li className="nav-item d-none d-md-block">
                                <Link className="navbar-logo" to="/">
                                    <img src={logo} alt="logo" />
                                </Link>
                            </li>
                        </ul>
                    </nav>

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
