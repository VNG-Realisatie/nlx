// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { string } from 'prop-types'
import { Button } from '@commonground/design-system'
import { Icon, IconExternalLink } from './index.styles'

const LinkButton = ({ href, text, ...props }) => {
  const isExternal = href && href.substring(0, 4) === 'http'
  const rel = isExternal ? { rel: 'noreferrer' } : {}

  return text && href ? (
    <Button as="a" variant="link" {...rel} {...props} href={href}>
      {text}
      {isExternal && <Icon as={IconExternalLink} inline />}
    </Button>
  ) : null
}

LinkButton.propTypes = {
  href: string,
  text: string,
}

export default LinkButton
