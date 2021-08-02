// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { bool } from 'prop-types'
import { PrimaryNavigation } from '@commonground/design-system'
import { useLocation } from 'react-router-dom'
import { IconHome, IconBox, IconInfo, IconMail } from '../../icons'
import { Container } from '../grid'
import NavLink from '../nav-link'
import {
  StyledIcon,
  LogoWrapper,
  StyledNLXLogo,
  NavigationWrapper,
} from './index.styles'

const Header = ({ homepage }) => {
  const { pathname } = useLocation()

  const HomeIcon = () => <StyledIcon as={IconHome} />
  const FeaturesIcon = () => <StyledIcon as={IconBox} />
  const AboutIcon = () => <StyledIcon as={IconInfo} />
  const ContactIcon = () => <StyledIcon as={IconMail} />

  return (
    <>
      <LogoWrapper homepage={homepage}>
        <Container>
          <StyledNLXLogo />
        </Container>
      </LogoWrapper>

      <NavigationWrapper>
        <PrimaryNavigation
          LinkComponent={NavLink}
          pathname={pathname}
          mobileMoreText="Meer"
          items={[
            {
              name: 'Home',
              to: '/',
              Icon: HomeIcon,
            },
            {
              name: 'Features',
              to: '/features',
              Icon: FeaturesIcon,
            },
            {
              name: 'Over NLX',
              to: '/about',
              Icon: AboutIcon,
            },
            {
              name: 'Contact',
              to: '/contact',
              Icon: ContactIcon,
            },
            {
              name: 'Docs',
              to: 'https://docs.nlx.io/',
              target: '_blank',
            },
          ]}
        />
      </NavigationWrapper>
    </>
  )
}

Header.propTypes = {
  homepage: bool,
}

Header.defaultProps = {
  homepage: false,
}

export default Header
