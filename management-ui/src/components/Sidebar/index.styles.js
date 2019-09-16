// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import styled from 'styled-components'
import { NavLink } from 'react-router-dom'

export const Header = styled.header`
    display: block;
    margin: 0;
    padding: 0;
    background-color: ${(p) => p.theme.color.sidebar};
    height: 100%;
    box-shadow: 1px 0 0 #e6eaf5;
    z-index: 1;
`

export const Nav = styled.nav`
    padding-top: 24px;
`

export const NavbarLogoLink = styled(NavLink)`
    display: flex;
    align-items: center;
    padding: 12px 24px;
`

export const NavigationItems = styled.ul`
    list-style-type: none;
    padding: 0;
`

export const NavigationItem = styled.li``

export const Link = styled(NavLink)`
    display: block;
    color: #a3aabf;
    font-weight: ${(p) => p.theme.font.weight.semibold};
    text-decoration: none;
    padding: 12px 24px;

    &:hover {
        background-color: #f7f9fc;
    }

    &:active {
        background-color: #f0f2f7;
    }

    &.active {
        color: #517fff;
        border-left: 2px solid #517fff;
        padding-left: 18px;
    }
`
