// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import {
  node,
  oneOf,
  bool,
  string,
  oneOfType,
  arrayOf,
  objectOf,
} from 'prop-types'
import { List, Item } from './index.styles'

/**
 * Not supporting multiple titles or values
 */
const SpecItem = ({ title, value, alignValue }) => (
  <Item alignValue={alignValue}>
    <dt>{title}</dt>
    <dd>{value}</dd>
  </Item>
)

SpecItem.propTypes = {
  title: node,
  value: node,
  alignValue: oneOf(['left', 'right']),
}

/**
 * Displays key/value pairs with the key in a deviating style.
 * Intended to be used with two td's per row only.
 */
const SpecList = ({ children, ...props }) => {
  return <List {...props}>{children}</List>
}

SpecList.propTypes = {
  alignValuesRight: bool,
  className: string,
  children: oneOfType([objectOf(SpecItem), arrayOf(objectOf(SpecItem))]),
}

SpecList.defaultProps = {
  className: '',
}

SpecList.Item = SpecItem

export default SpecList
