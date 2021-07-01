// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { string } from 'prop-types'
import { useTranslation } from 'react-i18next'

import DefaultAvatar from './default-avatar.svg'
import { Figure } from './index.styles'

const Avatar = ({ url, alt, ...props }) => {
  const { t } = useTranslation()
  return (
    <Figure {...props}>
      <img
        className="avatar-image"
        src={url || DefaultAvatar}
        alt={alt || t('Avatar')}
      />
    </Figure>
  )
}

Avatar.propTypes = {
  url: string,
  alt: string,
}

export default Avatar
