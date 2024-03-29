// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { useContext, useRef, useState } from 'react'
import { CSSTransition } from 'react-transition-group'
import { useTranslation } from 'react-i18next'
import { Icon } from '@commonground/design-system'
import UserContext from '../../user-context'
import useClickOutside from '../../hooks/use-click-outside'
import { IconShutdown } from '../../icons'
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
  const { user, logout } = useContext(UserContext)
  const [menuIsOpen, setMenuIsOpen] = useState(false)

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
        <StyledUserMenu
          id="user-menu-options"
          animationDuration={ANIMATION_DURATION}
          data-testid="user-menu-options"
        >
          <StyledUserMenuItem>
            <button type="button" onClick={logout}>
              <Icon as={IconShutdown} inline />
              {t('Log out')}
            </button>
          </StyledUserMenuItem>
        </StyledUserMenu>
      </CSSTransition>
    </StyledUserNavigation>
  )
}

export default UserNavigation
