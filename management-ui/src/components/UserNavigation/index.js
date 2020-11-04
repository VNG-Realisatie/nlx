// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { useContext, useRef, useState } from 'react'
import Cookies from 'js-cookie'
import { CSSTransition } from 'react-transition-group'
import { useTranslation } from 'react-i18next'

import UserContext from '../../user-context'
import useClickOutside from '../../hooks/use-click-outside'
import {
  StyledAvatar,
  StyledToggleButton,
  StyledUserMenu,
  StyledUserMenuItem,
  StyledUsername,
  StyledUserNavigation,
  UserNavigationChevron,
} from './index.styles'

const ANIMATION_DURATION = 150

const UserNavigation = ({ ...props }) => {
  const { t } = useTranslation()
  const [menuIsOpen, setMenuIsOpen] = useState(false)
  const { user } = useContext(UserContext)

  const onClickHandler = ({ currentTarget }) => {
    setMenuIsOpen(!menuIsOpen)
    currentTarget.focus()
  }

  const onClickOutside = () => {
    setMenuIsOpen(false)
  }

  const wrapperRef = useRef(null)
  useClickOutside(wrapperRef, onClickOutside)

  return !user ? null : (
    <StyledUserNavigation
      animationDuration={ANIMATION_DURATION}
      isOpen={menuIsOpen}
      {...props}
      ref={wrapperRef}
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
        <StyledAvatar
          data-testid="avatar"
          alt={t('User avatar')}
          url={user.pictureUrl}
        />
        <StyledUsername data-testid="full-name" title={user.fullName}>
          {user.fullName}
        </StyledUsername>
        <UserNavigationChevron
          animationDuration={ANIMATION_DURATION}
          flipHorizontal={menuIsOpen}
        />
      </StyledToggleButton>

      <CSSTransition
        in={menuIsOpen}
        timeout={ANIMATION_DURATION}
        classNames="user-menu-slide"
        unmountOnExit
      >
        <StyledUserMenu id="user-menu-options" data-testid="user-menu-options">
          <StyledUserMenuItem>
            <form method="POST" action="/oidc/logout">
              <input
                type="hidden"
                name="csrfmiddlewaretoken"
                value={Cookies.get('csrftoken')}
              />
              <button type="submit">{t('Log out')}</button>
            </form>
          </StyledUserMenuItem>
        </StyledUserMenu>
      </CSSTransition>
    </StyledUserNavigation>
  )
}

export default UserNavigation
