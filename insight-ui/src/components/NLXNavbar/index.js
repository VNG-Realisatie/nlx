// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { Children } from 'react'
import { element, string } from 'prop-types'

import { NLXLogo, GitLabLogo } from '@commonground/design-system'
import Header from '../Header'
import Navigation from '../Navigation'
import Container from '../Container'
import IconButton from '../IconButton'

import {
  StyledNavbarNav,
  StyledNavbarLogoLink,
  StyledNavigation,
} from './index.styles'

const NLXNavbar = ({
  children,
  homePageURL,
  aboutPageURL,
  docsPageURL,
  ...props
}) => (
  <Header {...props}>
    <Container>
      <StyledNavbarNav>
        <StyledNavbarLogoLink href={homePageURL}>
          <NLXLogo />
        </StyledNavbarLogoLink>

        <StyledNavigation>
          <Navigation.Item>
            <a href={aboutPageURL}>Over NLX</a>
          </Navigation.Item>
          <Navigation.Item>
            <a href={docsPageURL}>Docs</a>
          </Navigation.Item>
        </StyledNavigation>

        {Children.count(children) > 0 ? (
          <StyledNavigation>{children}</StyledNavigation>
        ) : null}

        <IconButton
          as="a"
          href="https://gitlab.com/commonground/nlx"
          target="_blank"
          aria-label="GitLab"
          rel="noopener noreferrer"
        >
          <GitLabLogo style={{ height: '20px' }} />
        </IconButton>
      </StyledNavbarNav>
    </Container>
  </Header>
)

NLXNavbar.propTypes = {
  homePageURL: string.isRequired,
  aboutPageURL: string.isRequired,
  docsPageURL: string.isRequired,
  children: element,
}

export default NLXNavbar
