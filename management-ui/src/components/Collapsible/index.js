// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState } from 'react'
import { node } from 'prop-types'
import { CSSTransition } from 'react-transition-group'
import {
  CollapsibleBody,
  CollapsibleButton,
  CollapsibleChevron,
  CollapsibleTitle,
  CollapsibleWrapper,
} from './index.styles'

const Collapsible = ({ title, children }) => {
  const [isOpen, setIsOpen] = useState(false)
  const toggle = () => setIsOpen(!isOpen)
  return children ? (
    <CollapsibleWrapper onClick={toggle} data-testid="collapsible">
      <CollapsibleButton
        aria-haspopup="true"
        aria-expanded={isOpen}
        aria-controls={title}
        aria-label={title}
      >
        <CollapsibleTitle>{title}</CollapsibleTitle>
        <CollapsibleChevron flipHorizontal={isOpen} />
      </CollapsibleButton>

      <CSSTransition
        in={isOpen}
        timeout={300}
        unmountOnExit
        classNames="collapsible"
      >
        <CollapsibleBody>{children}</CollapsibleBody>
      </CSSTransition>
    </CollapsibleWrapper>
  ) : (
    title
  )
}

Collapsible.propTypes = {
  title: node.isRequired,
  children: node,
}

export default Collapsible
