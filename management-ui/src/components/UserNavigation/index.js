// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { useContext, useState } from 'react'
import Cookies from 'js-cookie'
import { useTranslation } from 'react-i18next'
import Avatar from '../Avatar'
import UserContext from '../../user-context'
import {
  StyledUserNavigation,
  StyledToggleButton,
  StyledUsername,
  UserNavigationChevron,
} from './index.styles'

const UserNavigation = ({ ...props }) => {
  const { t } = useTranslation()
  const [menuIsOpen, setMenuIsOpen] = useState(false)
  const { user } = useContext(UserContext)

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

  return !user ? null : (
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
        <Avatar
          data-testid="avatar"
          alt={t('User avatar')}
          url={user.pictureUrl}
        />
        <StyledUsername data-testid="full-name" title={user.fullName}>
          {user.fullName}
        </StyledUsername>
        <UserNavigationChevron flipHorizontal={menuIsOpen} />
      </StyledToggleButton>

      {menuIsOpen && (
        <ul id="user-menu-options" data-testid="user-menu-options">
          <li>
            <form method="POST" action="/oidc/logout">
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

export default UserNavigation
