import React from 'react'
import { string } from 'prop-types'

import { StyledAvatar } from '../UserNavigation/index.styles'
import DefaultAvatar from './default-avatar.svg'

const Avatar = ({ url, ...props }) => (
  <StyledAvatar {...props}>
    <img className="avatar-image" src={url || DefaultAvatar} alt="Avatar" />
  </StyledAvatar>
)

Avatar.propTypes = {
  url: string,
}

export default Avatar
