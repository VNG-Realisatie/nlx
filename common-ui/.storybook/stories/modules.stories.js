import React from 'react'
import { storiesOf } from '@storybook/react'
import { BrowserRouter, NavLink } from 'react-router-dom'
import styled from 'styled-components'
import { StateDecorator, Store } from "@sambego/storybook-state";


import { Flex, Box } from '@rebass/grid'
import Hide from '../Hide'

import {Container} from 'src/Grid/Grid'
import Navigation, {NavigationItem} from 'src/Navigation/Navigation'
import Drawer from 'src/Drawer/Drawer'
import List, {ListItem} from 'src/List/List'

import IconButton from 'src/IconButton/IconButton'
import MenuIcon from '@material-ui/icons/Menu'

import {tableStory} from './modules/table.story.js'

const store = new Store({
    showDrawer: false
})

storiesOf('Modules', module)
    .addDecorator(StateDecorator(store))
    .add(
        'Navigation',
        () => {
            const nlxLogo = (
                <svg width="56" height="20" viewBox="0 0 56 22" xmlns="http://www.w3.org/2000/svg"><g id="Page-1" fill="none" fillRule="evenodd"><g id="Directory" transform="translate(-140 -18)" fillRule="nonzero"><g id="Group-9"><g id="logo" transform="translate(140 18)"><path d="M50.423 14.254a3.044 3.044 0 0 0 2.358 1.314 2.425 2.425 0 0 0 1.708-.739 2.434 2.434 0 0 0 .685-1.733 1.844 1.844 0 0 0-.272-1.089l-3.398-3.635 3.514-4.157c.18-.347.269-.733.26-1.123a2.359 2.359 0 0 0-.618-1.754 2.35 2.35 0 0 0-1.694-.76 3.073 3.073 0 0 0-2.618 1.35l-2.381 2.999-2.41-3A3.073 3.073 0 0 0 42.94.58a2.35 2.35 0 0 0-1.694.76 2.359 2.359 0 0 0-.618 1.753c-.01.39.081.778.266 1.123l3.514 4.157-3.404 3.635a1.884 1.884 0 0 0-.26 1.089 2.425 2.425 0 0 0 2.393 2.472 3.056 3.056 0 0 0 2.357-1.314l2.474-3.185 2.456 3.185z" id="Shape" fill="#FEBF24"></path><path d="M36.767 21.861a2.464 2.464 0 0 0 2.468-2.437 2.493 2.493 0 0 0-2.468-2.472h-7.779V3.202A2.62 2.62 0 0 0 26.37.579a2.62 2.62 0 0 0-2.618 2.623v16.036c-.004.697.27 1.367.762 1.86.492.493 1.16.768 1.856.763h10.397zM0 19.238a2.613 2.613 0 0 0 .763 1.858c.492.493 1.16.768 1.855.765a2.604 2.604 0 0 0 1.855-.765 2.613 2.613 0 0 0 .763-1.858V8.783l8.056 6.223v4.232a2.62 2.62 0 0 0 2.618 2.623 2.62 2.62 0 0 0 2.618-2.623V3.202A2.62 2.62 0 0 0 15.91.579a2.62 2.62 0 0 0-2.618 2.623v5.094l-9.218-7.19A2.407 2.407 0 0 0 2.468.579 2.534 2.534 0 0 0 0 3.126v16.112z" fill="#6C757D"></path></g></g></g></g></svg>
            )

            const gitlabLogo = (
                <svg width="26" height="24" viewBox="0 0 26 24" xmlns="http://www.w3.org/2000/svg"><g fill="none" fillRule="evenodd"><g id="Group" fillRule="nonzero"><path d="M25.962 10.291l-1.454 4.474-2.88 8.867a.495.495 0 0 1-.942 0l-2.881-8.867H8.238l-2.88 8.867a.495.495 0 0 1-.943 0l-2.88-8.867L.08 10.291a.99.99 0 0 1 .36-1.107L13.02.044l12.581 9.14a.99.99 0 0 1 .36 1.107" id="path46" fill="#FC6D26" transform="matrix(1 0 0 -1 0 24.018)"></path><path id="path50" fill="#E24329" transform="matrix(1 0 0 -1 0 33.226)" d="M13.021 9.252l4.784 14.722H8.238z"></path><path id="path58" fill="#FC6D26" transform="matrix(1 0 0 -1 0 33.226)" d="M13.021 9.252L8.238 23.974H1.534z"></path><path d="M1.534 23.974L.081 19.5a.99.99 0 0 1 .36-1.107l12.58-9.14-11.487 14.72z" id="path66" fill="#FCA326" transform="matrix(1 0 0 -1 0 33.226)"></path><path d="M1.534.044h6.704L5.358 8.91a.495.495 0 0 1-.943 0L1.535.044z" id="path74" fill="#E24329" transform="matrix(1 0 0 -1 0 9.296)"></path><path id="path78" fill="#FC6D26" transform="matrix(1 0 0 -1 0 33.226)" d="M13.021 9.252l4.784 14.722h6.704z"></path><path d="M24.509 23.974l1.453-4.474a.99.99 0 0 0-.36-1.107l-12.58-9.14 11.487 14.72z" id="path82" fill="#FCA326" transform="matrix(1 0 0 -1 0 33.226)"></path><path d="M24.509.044h-6.704l2.88 8.866a.495.495 0 0 0 .943 0l2.88-8.866z" id="path86" fill="#E24329" transform="matrix(1 0 0 -1 0 9.296)"></path></g></g></svg>
            )

            const StyledDrawer = styled(Drawer)`
                position: fixed;
                top: 64px;
                bottom: 0;
                left: 0;

                width: 100%;
                max-width: 240px;

                transform: ${p => !store.get("showDrawer") && 'translateX(-100%)'};
                box-shadow: ${p => !store.get("showDrawer") && 'none'};

                transition: transform ${p => p.theme.transition.materialNormal}, box-shadow ${p => p.theme.transition.materialNormal};
            `

            const navItems = [
                {
                    label: 'Home',
                    link: '/',
                },
                {
                    label: 'Over NLX',
                    link: '/about',
                },
                {
                    label: 'Docs',
                    link: '/docs',
                },
                {
                    label: 'Directory',
                    link: '/directory',
                },
            ]

            return (
                <BrowserRouter>
                    <React.Fragment>
                        <Navigation style={{position: 'fixed', top: 0, left: 0, right: 0, zIndex: 1}}>
                            <Container>
                                <Flex justifyContent="space-between" alignItems="center">
                                    <Box>
                                        <Hide breakpoints={[1,2,3]}>
                                            <IconButton size="large" variant="tertiary" onClick={() => store.set({ showDrawer: !store.get("showDrawer") })}>
                                                <MenuIcon/>
                                            </IconButton>
                                        </Hide>
                                        <Hide breakpoints={[0]}>
                                            <Flex>
                                                <NavigationItem as={NavLink} to="/">{nlxLogo}</NavigationItem>
                                                <NavigationItem as={NavLink} to="/about">Over NLX</NavigationItem>
                                                <NavigationItem as={NavLink} to="/docs">Docs</NavigationItem>
                                                <NavigationItem as={NavLink} to="/directory">Directory</NavigationItem>
                                            </Flex>
                                        </Hide>
                                    </Box>
                                    <Box>
                                        <IconButton size="large" variant="tertiary" as={'a'} href="https://gitlab.com/commonground/nlx" target="_blank">
                                            {gitlabLogo}
                                        </IconButton>
                                    </Box>
                                </Flex>
                            </Container>
                        </Navigation>
                        <Hide breakpoints={[1,2,3]}>
                            <StyledDrawer>
                                <Flex flexDirection="column" justifyContent="space-between" style={{height: '100%'}}>
                                    <Box>
                                        <List>
                                            {navItems.map((item, index) => (
                                                <ListItem key={index} as={NavLink} to={item.link} size="normal">
                                                    {item.label}
                                                </ListItem>
                                            ))}
                                        </List>
                                    </Box>
                                    <Box p={3}>
                                        {nlxLogo}
                                    </Box>
                                </Flex>
                            </StyledDrawer>
                        </Hide>
                    </React.Fragment>
                </BrowserRouter>
            )
        }
    )
    .add(
        'Table',
        () => tableStory
        )