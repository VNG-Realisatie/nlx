// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { bool } from 'prop-types'
import { PrimaryNavigation } from '@commonground/design-system'
import { useLocation } from 'react-router-dom'
import { IconHome } from '../../icons'
import { Container } from '../Grid'
import NavLink from '../NavLink'

import {
  StyledIcon,
  LogoWrapper,
  StyledNLXLogo,
  NavigationWrapper,
} from './index.styles'

const Header = ({ homepage }) => {
  const { pathname } = useLocation()

  const HomeIcon = () => <StyledIcon as={IconHome} />

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
              name: 'Docs',
              to: 'https://docs.nlx.io/',
              target: '_blank',
            },
            {
              name: 'NLX.io',
              to: 'https://nlx.io',
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
