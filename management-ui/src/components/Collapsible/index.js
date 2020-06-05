// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState, useEffect } from 'react'
import { node } from 'prop-types'
import { CSSTransition } from 'react-transition-group'
import {
  CollapsibleBody,
  CollapsibleButton,
  CollapsibleChevron,
  CollapsibleTitle,
  CollapsibleWrapper,
} from './index.styles'

const createRandomId = () => `r${Math.random().toString(36).slice(8)}`

const Collapsible = ({ title, ariaLabel, children }) => {
  const [isOpen, setIsOpen] = useState(false)
  const [id, setId] = useState()
  const toggle = () => setIsOpen(!isOpen)

  useEffect(() => {
    setId(createRandomId())
  }, [])

  return (
    <CollapsibleWrapper onClick={toggle} data-testid="collapsible">
      <CollapsibleButton
        aria-haspopup="true"
        aria-expanded={isOpen}
        aria-controls={id}
        aria-label={ariaLabel || title}
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
        <CollapsibleBody id={id}>{children}</CollapsibleBody>
      </CSSTransition>
    </CollapsibleWrapper>
  )
}

Collapsible.propTypes = {
  title: node.isRequired,
  ariaLabel: (props, propName, componentName) => {
    if (typeof props.title !== 'string' && !props[propName]) {
      return new Error(
        'If Collapsible title is not a string, please provide an ariaLabel',
      )
    }
  },
  children: node.isRequired,
}

export default Collapsible
