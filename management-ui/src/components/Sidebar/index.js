// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import {
    Header,
    Nav,
    NavbarLogoLink,
    NavigationItems,
    NavigationItem,
    Link,
} from './index.styles'
import { NLXLogo } from '@commonground/design-system'

const Navbar = ({ onLinkClickHandler, ...props }) => (
    <Header {...props}>
        <Nav>
            <NavbarLogoLink to="/" exact>
                <NLXLogo width="56px" height="22px" />
            </NavbarLogoLink>

            <NavigationItems>
                <NavigationItem>
                    <Link to="/inways">Inways</Link>
                </NavigationItem>
                <NavigationItem>
                    <Link to="/services">Services</Link>
                </NavigationItem>

                <NavigationItem>
                    <Link to="/logout">Logout</Link>
                </NavigationItem>
            </NavigationItems>
        </Nav>
    </Header>
)

export default Navbar
