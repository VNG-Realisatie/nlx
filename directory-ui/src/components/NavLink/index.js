// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { string, node, oneOfType, element, func } from 'prop-types'
import { NavLink as RouterNavLink } from 'react-router-dom'
import { StyledIcon, IconExternalLink } from './index.styles'

const NavLink = ({ to, children, target, ...props }) => {
  const isExternal = to.substring(0, 4) === 'http'
  const rel = isExternal ? { rel: 'noreferrer' } : {}

  return (
    <RouterNavLink to={{ pathname: to }} target={target} {...rel} {...props}>
      {children}
      {isExternal && <StyledIcon as={IconExternalLink} inline />}
    </RouterNavLink>
  )
}

NavLink.propTypes = {
  to: string,
  className: string,
  children: node,
  target: string,
  Icon: oneOfType([element, func]),
}

export default NavLink
