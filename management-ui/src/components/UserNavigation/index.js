// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { useState } from 'react'
import { string } from 'prop-types'
import Cookies from 'js-cookie'
import { useTranslation } from 'react-i18next'
import Avatar from '../Avatar'
import {
  StyledUserNavigation,
  StyledToggleButton,
  StyledUsername,
  UserNavigationChevron,
} from './index.styles'

const UserNavigation = ({ fullName, pictureUrl, ...props }) => {
  const { t } = useTranslation()
  const [menuIsOpen, setMenuIsOpen] = useState(false)

  const onClickHandler = (event) => {
    setMenuIsOpen(!menuIsOpen)
    event.currentTarget.focus()
  }

  let timeoutId
  const onBlurHandler = () => {
    timeoutId = setTimeout(() => {
      setMenuIsOpen(false)
    })
  }

  const onFocusHandler = () => {
    clearTimeout(timeoutId)
  }

  return (
    <StyledUserNavigation
      isOpen={menuIsOpen}
      onFocus={onFocusHandler}
      onBlur={onBlurHandler}
      {...props}
      data-testid="user-navigation"
    >
      <StyledToggleButton
        type="button"
        onClick={onClickHandler}
        aria-haspopup="true"
        aria-expanded={menuIsOpen}
        aria-controls="user-menu-options"
        aria-label={t('Account menu')}
      >
        <Avatar data-testid="avatar" alt={t('User avatar')} url={pictureUrl} />
        <StyledUsername data-testid="full-name">{fullName}</StyledUsername>
        <UserNavigationChevron flipHorizontal={menuIsOpen} />
      </StyledToggleButton>

      {menuIsOpen && (
        <ul id="user-menu-options" data-testid="user-menu-options">
          <li>
            <form method="POST" action="/oidc/logout/">
              <input
                type="hidden"
                name="csrfmiddlewaretoken"
                value={Cookies.get('csrftoken')}
              />
              <button type="submit">{t('Log out')}</button>
            </form>
          </li>
        </ul>
      )}
    </StyledUserNavigation>
  )
}

UserNavigation.propTypes = {
  fullName: string,
  pictureUrl: string,
}

export default UserNavigation
