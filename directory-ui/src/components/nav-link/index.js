// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { string, node, oneOfType, element, func } from 'prop-types'
import { useLocation } from 'react-router-dom'
import { StyledIcon, IconExternalLink } from './index.styles'

const NavLink = ({ to, className, children, Icon, ...props }) => {
  const { basePath, pathname } = useLocation()
  const isExternal = to.substring(0, 4) === 'http'
  const href = isExternal ? to : basePath + to
  const rel = isExternal ? { rel: 'noreferrer' } : {}
  const finalClassName = pathname === to ? `${className} active` : className

  return (
    <a href={href} className={finalClassName} {...rel} {...props}>
      {children}
      {isExternal && <StyledIcon as={IconExternalLink} inline />}
    </a>
  )
}

NavLink.propTypes = {
  to: string,
  className: string,
  children: node,
  Icon: oneOfType([element, func]),
}

export default NavLink
