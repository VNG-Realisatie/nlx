import React, {Children} from 'react'
import {string} from 'prop-types'

import {StyledNavbarNav, StyledNavbarLogoLink, StyledNavigation} from './index.styles'
import Header from '../Header'
import Navigation from '../Navigation'
import Container from '../Container'
import IconButton from '../IconButton'
import { NLXLogo } from '@commonground/design-system';
import { GitLabLogo } from '@commonground/design-system';

const NLXNavbar = ({ children, homePageURL, aboutPageURL, docsPageURL, ...props }) =>
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

        {
          Children.count(children) > 0 ?
            <StyledNavigation children={children} /> : null
        }

        <IconButton as="a" href="https://gitlab.com/commonground/nlx" target="_blank" aria-label="GitLab" rel="noopener noreferrer">
          <GitLabLogo style={({ height: '20px' })} />
        </IconButton>
      </StyledNavbarNav>
    </Container>
  </Header>

NLXNavbar.propTypes = {
  homePageURL: string.isRequired,
  aboutPageURL: string.isRequired,
  docsPageURL: string.isRequired
}

export default NLXNavbar
