// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { string, bool } from 'prop-types'
import { useTranslation } from 'react-i18next'

import DefaultAvatar from './default-avatar.svg'
import DefaultAvatar2 from './default-avatar-p.svg'
import { Figure } from './index.styles'

const notProd = process.env.NODE_ENV !== 'production'

const Avatar = ({ url, alt, menuIsOpen, ...props }) => {
  const { t } = useTranslation()
  const TheAvatar = notProd && menuIsOpen ? DefaultAvatar2 : DefaultAvatar
  return (
    <Figure {...props}>
      <img
        className="avatar-image"
        src={url || TheAvatar}
        alt={alt || t('Avatar')}
      />
    </Figure>
  )
}

Avatar.propTypes = {
  url: string,
  alt: string,
  menuIsOpen: bool,
}

export default Avatar
