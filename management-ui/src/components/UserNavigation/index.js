import React, { useState } from 'react'
import { string } from 'prop-types'
import Cookies from 'js-cookie'
import { useTranslation } from 'react-i18next'
import Avatar from '../Avatar'
import {
  StyledUserMenu,
  StyledMenuToggleButton,
  StyledUserName,
  StyledIconChevron,
} from './index.styles'

const UserNavigation = ({ fullName, pictureUrl }) => {
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
    <StyledUserMenu
      isOpen={menuIsOpen}
      onFocus={onFocusHandler}
      onBlur={onBlurHandler}
      data-testid="user-navigation"
    >
      <StyledMenuToggleButton
        type="button"
        onClick={onClickHandler}
        aria-haspopup="true"
        aria-expanded={menuIsOpen}
        aria-controls="user-menu-options"
        aria-label={t('Account menu')}
      >
        <Avatar data-testid="avatar" alt={t('User avatar')} url={pictureUrl} />
        <StyledUserName data-testid="full-name">{fullName}</StyledUserName>
        <StyledIconChevron flipHorizontal={menuIsOpen} />
      </StyledMenuToggleButton>

      {menuIsOpen && (
        <ul id="user-menu-options" data-testid="user-menu-options">
          <li>
            <form method="POST" action="/oidc/logout/">
              <input
                type="hidden"
                name="csrfmiddlewaretoken"
                value={Cookies.get('csrftoken')}
              />
              <button type="submit">{t('Logout')}</button>
            </form>
          </li>
        </ul>
      )}
    </StyledUserMenu>
  )
}

UserNavigation.propTypes = {
  fullName: string,
  pictureUrl: string,
}

export default UserNavigation
